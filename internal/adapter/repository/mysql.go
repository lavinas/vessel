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

	"github.com/go-sql-driver/mysql"
	gssh "golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

const (
	ConnectionTimeout = 5 * time.Second

	ErrRepoSshInvalid          = "invalid ssh dns"
	ErrRepoPassNotImplemented  = "password authentication is not implemented"
	ErrRepoSshTimeout          = "ssh connection timeout"
	ErrRepoProtoNotImplemented = "protocol is not implemented"
	ErrTransactionIsNil        = "transaction is nil"
	ErrInsertAutoParams        = "mysql insert auto invalid params"
	ErrGetParams               = "mysql get invalid params"

	QUse          = "USE %s;"
	QSimpleInsert = "INSERT INTO %s.%s (%s) VALUES (%s);"
	QGet          = "SELECT * FROM %s.%s WHERE %s;"
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

// InsertAuto is a method that inserts an object into the database and return the i d
func (m *MySql) Insert(tx interface{}, base, object string, vals *map[string]interface{}) (int64, error) {
	if tx == nil || base == "" || object == "" || vals == nil || len(*vals) == 0 {
		return 0, errors.New(ErrInsertAutoParams)
	}
	txi := tx.(*sql.Tx)
	keys, vls, objs := m.getKeysValues(vals)
	sql := fmt.Sprintf(QSimpleInsert, base, object, keys, vls)
	result, err := txi.Exec(sql, *objs...)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Get is a method that gets all columns of a list of objects by a field
func (m *MySql) Get(tx interface{}, base, object string, where *map[string]interface{}) (*[]map[string]interface{}, error) {
	if tx == nil || base == "" || object == "" || where == nil || len(*where) == 0 {
		return nil, errors.New(ErrGetParams)
	}
	txi := tx.(*sql.Tx)
	whr, objs := m.getFormatWhere(where)
	sql := fmt.Sprintf(QGet, base, object, whr)
	rows, err := txi.Query(sql, objs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return m.formatRows(rows)
}

// getFormatVals is a method that gets keys and values of map of vals
func (m *MySql) getKeysValues(vals *map[string]interface{}) (string, string, *[]interface{}) {
	var keys, vls string
	var objs []interface{}
	for key, val := range *vals {
		if val == nil {
			continue
		}
		keys += key + ", "
		vls += "?, "
		objs = append(objs, val)
	}
	return keys[:len(keys)-2], vls[:len(vls)-2], &objs
}

// getFormatKeys is a method that gets the keys of the object formatted
func (m *MySql) getFormatWhere(where *map[string]interface{}) (string, []interface{}) {
	var whr string
	var objs []interface{}
	for key, val := range *where {
		whr += fmt.Sprintf("%s = ? AND ", key)
		objs = append(objs, val)
	}
	return whr[:len(whr)-5], objs
}

// queyMountMap is a method that mounts the slice of maps os result
func (r *MySql) formatRows(rows *sql.Rows) (*[]map[string]interface{}, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([]interface{}, len(cols))
	valuePtrs := make([]interface{}, len(cols))
	for i := range cols {
		valuePtrs[i] = &values[i]
	}
	result := make([]map[string]interface{}, 0)
	for rows.Next() {
		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}
		row, err := r.formatRow(cols, values)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return &result, nil
}

// queryFormatRow is a method that formats the row
func (r *MySql) formatRow(cols []string, values []interface{}) (map[string]interface{}, error) {
	row := make(map[string]interface{}, len(values))
	for i, val := range values {
		if obj, ok := val.([]byte); ok {
			row[cols[i]] = string(obj)
			continue
		}
		row[cols[i]] = val
	}
	return row, nil
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
