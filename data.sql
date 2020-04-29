create database if not exists mydb;
use mydb;
drop table if exists users;

create table users(
	Id int primary key,
	Name char(25) not null,
	Age int not null,
	Dept char(20),
	Subject char(10)
);
rakesh singh
INSERT INTO users VALUES(123,'',23,'MBA','Commerce');
INSERT INTO users VALUES(134,'bhavesh nadurdikar',21,'LAW','Social Science');
INSERT INTO users VALUES(135,'payal jain',22,'MATH','AM-I');
INSERT INTO users VALUES(136,'esha varma',22,'STAT','Statistics-II');
INSERT INTO users VALUES(141,'rohit jatt',23,'ELECTRONICS','Digital ELE');
INSERT INTO users VALUES(133,'shubham natekar',23,'CS','OS');
