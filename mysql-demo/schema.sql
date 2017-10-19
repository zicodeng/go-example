create table if not exists contacts
(
    id int not null auto_increment primary key,
    email varchar (128) not null,
    first_name varchar (64) not null,
    last_name varchar (128) not null
)