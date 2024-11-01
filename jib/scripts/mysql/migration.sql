use jib;
create table user (
    id int primary key auto_increment,
    name varchar(255) not null,
    email varchar(255) not null,
    age int not null,
    is_employee boolean not null 
)