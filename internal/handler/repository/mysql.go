package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"os"
	"regexp"
	"time"
	"strings"

	"github.com/go-sql-driver/mysql"
	gssh "golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

const (
	ConnectionTimeout          = 5 * time.Second
	ErrRepoSshInvalid          = "invalid ssh dns"
	ErrRepoPassNotImplemented  = "password authentication is not implemented"
	ErrRepoSshTimeout          = "ssh connection timeout"
	ErrRepoProtoNotImplemented = "protocol is not implemented"
	ErrTransactionIsNil		   = "transaction is nil"
	ErrNotFound  			   = "not found"

	QUse = "USE %s;"
	QSimpleInsert = "INSERT INTO %s.%s (%s) VALUES (%s) RETURNING id;"
)

// MySql represents the mysql repository
type MySql struct {
	Conn *sql.DB
	Ssh  *gssh.Client
	Sdns string
	Sssh string
}

// NewMySql is a function that creates a new mysql repository
func NewMySql(dns, ssh string) (*MySql, error) {
	db, sshConn, err := connect(dns, ssh)
	if err != nil {
		return nil, err
	}
	return &MySql{Conn: db, Ssh: sshConn, Sdns: dns, Sssh: ssh}, nil
}

// Close is a method that closes the mysql repository
func (m *MySql) Close() error {
	if m.Ssh != nil {
		m.Ssh.Close()
	}
	return m.Conn.Close()
}

// Ping is a method that pings the mysql repository
func (m *MySql) Ping() error {
	return m.Conn.Ping()
}

// Reconnect is a method that reconnects to the mysql repository
func (m *MySql) Reconnect() error {
	if err := m.Close(); err != nil {
		return err
	}
	db, sshConn, err := connect(m.Sdns, m.Sssh)
	if err != nil {
		return err
	}
	m.Conn = db
	m.Ssh = sshConn
	return nil
}

// Check is a method that checks the mysql repository and reconnects if necessary
func (m *MySql) Check() error {
	if err := m.Ping(); err != nil {
		if err := m.Reconnect(); err != nil {
			return err
		}
	}
	return nil
}

// Begin is a method that begins a transaction
func (m *MySql) Begin(base string) (interface{}, error) {
	if err := m.Check(); err != nil {
		return nil, err
	}
	if base != "" {
		if _, err := m.Conn.Exec(fmt.Sprintf(QUse, base)); err != nil {
			return nil, err
		}
	}
	return m.Conn.Begin()
}

// Commit is a method that commits a transaction
func (m *MySql) Commit(tx interface{}) error {
	return tx.(*sql.Tx).Commit()
}

// Rollback is a method that rolls back a transaction
func (m *MySql) Rollback(tx interface{}) error {
	return tx.(*sql.Tx).Rollback()
}

// InsertAuto is a method that inserts an object into the database and return the id
func (m *MySql) InsertAuto(tx interface{}, base, object string, fields *[]string, vals *[]string) (int64, error) {
	if tx == nil {
		return 0, errors.New(ErrTransactionIsNil)
	}
	txi := tx.(*sql.Tx)
	svals := ""
	for _, val := range *vals {
		svals += fmt.Sprintf("'%s', ", val)
	}
	svals = strings.TrimSuffix(svals, ", ")
	sql := fmt.Sprintf(QSimpleInsert, base, object, strings.Join(*fields, ", "), svals)
	var id int64
	if err := txi.QueryRow(sql).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// GetId is a method that gets an object by id
func (m *MySql) GetId(tx interface{}, base, object string, id int64, fields *[]string) (*[]string, error) {
	if tx == nil {
		return nil, errors.New(ErrTransactionIsNil)
	}
	txi := tx.(*sql.Tx)
	sql := fmt.Sprintf("SELECT * FROM %s.%s WHERE id = %d", base, object, id)
	rows, err := txi.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, errors.New(ErrNotFound)
	}
	vals := make([]string, len(*fields))
	err = rows.Scan(&vals)
	if err != nil {
		return nil, err
	}
	return &vals, nil
}

// GetField is a method that gets an object by field
func (m *MySql) GetField(tx interface{}, base, object, field, value string, fields *[]string) (*[]string, error) {
	if tx == nil {
		return nil, errors.New(ErrTransactionIsNil)
	}
	txi := tx.(*sql.Tx)
	sql := fmt.Sprintf("SELECT * FROM %s.%s WHERE %s = '%s'", base, object, field, value)
	rows, err := txi.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, errors.New(ErrNotFound)
	}
	vals := make([]string, len(*fields))
	err = rows.Scan(&vals)
	if err != nil {
		return nil, err
	}
	return &vals, nil
}


// DeleteId is a method that deletes an object by id
func (m *MySql) DeleteId(tx interface{}, base, object string, id int64) error {
	if tx == nil {
		return errors.New(ErrTransactionIsNil)
	}
	txi := tx.(*sql.Tx)
	sql := fmt.Sprintf("DELETE FROM %s.%s WHERE id = %d", base, object, id)
	_, err := txi.Exec(sql)
	return err
}

// connect is a function that connects to the database
func connect(dns string, ssh string) (*sql.DB, *gssh.Client, error) {
	sshConn, err := sshConnect(ssh)
	if err != nil {
		return nil, nil, err
	}
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, nil, err
	}
	return db, sshConn, nil
}

// sshConnect is a method that connects to the ssh server
func sshConnect(ssh string) (*gssh.Client, error) {
	if ssh == "" {
		return nil, nil
	}
	var sshUser, sshKfile, sshHost, sshPort string
	sshUser, sshKfile, sshHost, sshPort, err := parseSshDns(ssh)
	if err != nil {
		return nil, err
	}
	var agentClient agent.Agent
	if conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		defer conn.Close()
		agentClient = agent.NewClient(conn)
	}
	pemBytes, err := os.ReadFile(sshKfile)
	if err != nil {
		return nil, err
	}
	signer, err := gssh.ParsePrivateKey(pemBytes)
	if err != nil {
		return nil, err
	}
	sshConfig := &gssh.ClientConfig{
		User:            sshUser,
		Auth:            []gssh.AuthMethod{gssh.PublicKeys(signer)},
		HostKeyCallback: gssh.InsecureIgnoreHostKey(),
	}
	if agentClient != nil {
		sshConfig.Auth = append(sshConfig.Auth, gssh.PublicKeysCallback(agentClient.Signers))
	}
	// sshConn, err := gssh.Dial("tcp", fmt.Sprintf("%s:%s", sshHost, sshPort), sshConfig)
	sshConn, err := dialTimeout(sshHost, sshPort, sshConfig)
	if err != nil {
		return nil, err
	}
	mysql.RegisterDialContext("mysql+tcp", func(_ context.Context, addr string) (net.Conn, error) {
		return sshConn.Dial("tcp", addr)
	})
	return sshConn, nil
}

// dialTimeout is a method that dials the ssh server with timeout
func dialTimeout(sshHost, sshPort string, sshConfig *gssh.ClientConfig) (*gssh.Client, error) {
	ch := make(chan *gssh.Client, 1)
	ech := make(chan error, 1)
	ctx, cancel := context.WithTimeout(context.Background(), ConnectionTimeout)
	defer cancel()
	go func() {
		conn, err := gssh.Dial("tcp", fmt.Sprintf("%s:%s", sshHost, sshPort), sshConfig)
		if err != nil {
			ech <- err
			return
		}
		ch <- conn
	}()
	select {
	case conn := <-ch:
		return conn, nil
	case err := <-ech:
		return nil, err
	case <-ctx.Done():
		return nil, errors.New(ErrRepoSshTimeout)
	}
}

// parseSshDns is a method that parses the ssh dns
func parseSshDns(ssh string) (string, string, string, string, error) {
	pat := `^(\w+):(\w+)\(([^)]+)\)@(\w+)\(([^:]+):(\d+)\)$`
	re := regexp.MustCompile(pat)
	if !re.MatchString(ssh) {
		return "", "", "", "", errors.New(ErrRepoSshInvalid)
	}
	m := re.FindStringSubmatch(ssh)
	if len(m) != 7 {
		return "", "", "", "", errors.New(ErrRepoSshInvalid)
	}
	if m[2] != "file" {
		return "", "", "", "", errors.New(ErrRepoPassNotImplemented)
	}
	if m[4] != "tcp" {
		return "", "", "", "", errors.New(ErrRepoProtoNotImplemented)
	}
	return m[1], m[3], m[5], m[6], nil
}
