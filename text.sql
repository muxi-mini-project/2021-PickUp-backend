drop database mytext;
create database mytext;

use mytext;

drop table if exists `users`;
create table if not exists  `users`(
    `sid` varchar(100) not null,
    `name` varchar(100) not null,
   primary key(`sid`)
)engine=innodb default charset=utf8;