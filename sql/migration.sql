CREATE DATABASE IF NOT EXISTS;
USE devbook;

DROP TABLE IF EXISTS users;
CREATE TABLE users(
  id int auto_increment primary key,
  name varchar(50) not null,
  nick varchar(50) not null unique,
  email varchar(50) not null unique,
  password text not null,
  created_at timestamp not null default current_timestamp()
  ) ENGINE=INNODB;
