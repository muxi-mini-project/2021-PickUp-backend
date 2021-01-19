create database PICKUP;

use PICKUP;

--用户信息表--
drop table if exists `users`;
create table if not exists  `users`(
   `ID` int unsigned auto_increment,
   `password` varchar(100) not null,
   `gender` tinyint not null,
   `name` varchar(100) not null,
   `major` varchar(100) null,
   `QQ` varchar(100) null,
   `phone` varchar(100) null,
   `picture` varchar(100) not null,
   `car_information` varchar(100) null,
   `age` int null,
   primary key(`ID`)
)engine=innodb default charset=utf8;

--配对成功表--
drop table if exists `match`;
create table if not exists `match`(
   `ID` int unsigned auto_increment,
   `diver_ID` int not null,
   `passenger_ID` int not null,
   `start_time` varchar(100) not null,
   `start_spot` varchar(100) null,
   `diver_phone` varchar(100) null,
   primary key(`ID`)
)engine=innodb default charset=utf8;

--司机订单--
drop table if exists `require_diver`;
create table if not exists `require_diver`(
   `ID` int unsigned auto_increment,
   `diver_ID` int not null,
   `start_spot` varchar(100) not null,
   `time range` varchar(100) not null,
   `passing_spots` varchar(100) null,
   `notes` text null,
   `status`tinyint ,
   primary key(`ID`)
)engine=innodb default charset=utf8;

--乘客订单--
drop table if exists `require_passenger`;
create table if not exists `require_passenger`(
   `ID` int unsigned auto_increment,
   `passenger_ID` int not null,
   `start_spot` varchar(100) not null,
   `end_spot` varchar(100) not null,
   `start_time` varchar(100) not null,
   `urgent` tinyint not null,
   `notes` text null,
   `status`tinyint ,
   primary key(`ID`)
)engine=innodb default charset=utf8;

--互评之司机--
drop table if exists `comment_diver`;
create table if not exists `comment_diver`(
   `ID` int unsigned auto_increment,
   `diver_ID` int not null,
   `diver_score` float null,
   `words` varchar(100) null,
   primary key(`ID`)
)engine=innodb default charset=utf8;

--互评之乘客--
drop table if exists `comment_passenger`;
create table if not exists `comment_passenger`(
   `ID` int unsigned auto_increment,
   `passenger_ID` int not null,
   `passenger_score` float null,
   `words` varchar(100) null,
   primary key(`ID`)
)engine=innodb default charset=utf8;

--双向选择之乘客--
drop table if exists `confirm_passenger`;
create table if not exists `confirm_passenger`(
   `ID` int unsigned auto_increment,
   `require_diver_ID` int not null,
   `passenger_attitude` tinyint null,
   primary key(`ID`)
)engine=innodb default charset=utf8;

--双向选择之司机--
drop table if exists `confirm_diver`;
create table if not exists `confirm_diver`(
   `ID` int unsigned auto_increment,
   `confirm_passenger_ID` int not null,
   `diver_attitude` tinyint null,
   primary key(`ID`)
)engine=innodb default charset=utf8;