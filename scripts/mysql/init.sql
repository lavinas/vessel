-- Date: 2021-07-07
-- Author: P Lavinas
-- Description: Create the database and tables for the assets service

-- Create the database and tables for the assets service
create database if not exists `assets` default character set utf8 collate utf8_general_ci;
use `assets`;

-- Create the tables for the assets service
create table if not exists `class` (
    `id` bigint not null auto_increment,
    `name` varchar(100) not null,
    `description` varchar(500),
    `created_at` datetime not null,
    primary key (`id`),
    unique key `name` (`name`)
) engine=InnoDB default charset=utf8;

create table if not exists `asset` (
    `id` bigint not null auto_increment,
    `class_id` bigint not null,
    `name` varchar(100) not null,
    `description` varchar(500),
    `created_at` datetime not null,
    primary key (`id`),
    unique key `name` (`name`),
    index `class_id` (`class_id`),
    foreign key (`class_id`) references `class` (`id`)
) engine=InnoDB default charset=utf8;

create table if not exists `event` (
    `id` bigint not null auto_increment,
    `name` varchar(100) not null,
    `description` varchar(500),
    `created_at` datetime not null,
    primary key (`id`),
    unique key `name` (`name`)
) engine=InnoDB default charset=utf8;	

create table if not exists `history` (
    `id` bigint not null auto_increment,
    `at` datetime not null,
    `asset_id` bigint not null,
    `event_id` bigint not null,
    `value` decimal(10, 2) not null,
    primary key (`id`),
    index `asset_id` (`asset_id`),
    index `event_id` (`event_id`),
    index `at` (`at`),
    foreign key (`asset_id`) references `asset` (`id`)
    foreign key (`event_id`) references `event` (`id`)
) engine=InnoDB default charset=utf8;


create table if not exists `test` (
    `id` bigint not null auto_increment,
    `name` varchar(100) not null,
    `created_at` datetime not null,
    `value` decimal(10, 2) not null,
    primary key (`id`),
    unique key `name` (`name`)
) engine=InnoDB default charset=utf8;