create database PICKUP;

use PICKUP;

--用户信息表--
drop table if exists `users`;
create table if not exists  `users`(
   `id` int unsigned auto_increment,
   `password` varchar(100) not null,
   `number` varchar null,
   `gender` tinyint not null,
   `username` varchar(100) not null,
   `qq` varchar(100) null,
   `phone` varchar(100) null,
   `picture` varchar(100) not null,
   `student_number` varchar(100) null,
   primary key(`id`)
)engine=innodb default charset=utf8;

--索引--
create index idx_users_id on users(id);

--匹配成功表--
drop table if exists `match`;
create table if not exists `match`(
   `id` int unsigned auto_increment,
   `diver_id` int not null,
   `passenger_id` int not null,
   `start_time` varchar(100) not null,
   `end_time` varchar(100) not null,
   `start_spot` varchar(100) null,
   `diver_phone` varchar(100) null,
   primary key(`id`)
)engine=innodb default charset=utf8;

--司机需求--
drop table if exists `require_diver`;
create table if not exists `require_diver`(
   `id` int unsigned auto_increment,
   `diver_id` int not null,
   `start_spot` varchar(100) not null,
   `start_time` varchar(100) not null,
   `end_time` varchar(100) not null,
   `passing_spots` varchar(100) null,
   `notes` text null,
   `status`tinyint ,
   primary key(`id`)
)engine=innodb default charset=utf8;

--索引--
create index idx_start_spot on require_diver(start_spot);
create index idx_passing_spots on require_diver(passing_spots);
create index idx_d_status on require_diver(status);

--乘客需求--
drop table if exists `require_passenger`;
create table if not exists `require_passenger`(
   `id` int unsigned auto_increment,
   `passenger_id` int not null,
   `start_spot` varchar(100) not null,
   `end_spot` varchar(100) not null,
   `start_time` varchar(100) not null,
   `end_time` varchar(100) not null,
   `urgent` tinyint not null,
   `notes` text null,
   `status`tinyint ,
   primary key(`id`)
)engine=innodb default charset=utf8;

--索引--
create index idx_p_status on require_passenger(status);

--评价司机--
drop table if exists `comment_diver`;
create table if not exists `comment_diver`(
   `id` int unsigned auto_increment,
   `diver_id` int not null,
   `diver_score` float null,
   `words` varchar(100) null,
   primary key(`id`)
)engine=innodb default charset=utf8;

--评价乘客--
drop table if exists `comment_passenger`;
create table if not exists `comment_passenger`(
   `id` int unsigned auto_increment,
   `passenger_id` int not null,
   `passenger_score` float null,
   `words` varchar(100) null,
   primary key(`id`)
)engine=innodb default charset=utf8;

--确认匹配--
drop table if exists `confirm_passenger`;
create table if not exists `confirm_passenger`(
   `id` int unsigned auto_increment,
   `require_diver_id` int not null,
   `passenger_attitude` tinyint null,
   primary key(`id`)
)engine=innodb default charset=utf8;

--索引--
create index idx_p_attitude on confirm_passenger(passenger_attitude);

--确认匹配--
drop table if exists `confirm_diver`;
create table if not exists `confirm_diver`(
   `id` int unsigned auto_increment,
   `confirm_passenger_id` int not null,
   `diver_attitude` tinyint null,
   primary key(`id`)
)engine=innodb default charset=utf8;