
CREATE DATABASE `libaray` CHARACTER SET 'utf8';

create table borrower
(
id int not null primary key auto_increment,
name varchar(20) not null,
card_num varchar(50) not null unique,
type int not null
);
create table undergraduate
(
id int not null primary key references borrower(id),
major varchar(50)
);

create table graduate
(
id int not null primary key references borrower(id),
major varchar(50),
director varchar(50)
);

create table doctor
(
id int not null primary key references borrower(id),
major varchar(50)
);

create table college_student
(
id int not null primary key references borrower(id),
major varchar(50)
);

create table teacher
(
id int not null primary key references borrower(id),
department varchar(50)
);

