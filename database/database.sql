-- Script for creating database 

CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id int auto_increment primary key,
    name varchar(50) not null,
    nickname varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(50) not null unique,
    created_at timestamp default CURRENT_TIMESTAMP
) ENGINE=InnoDB;
