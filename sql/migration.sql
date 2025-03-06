CREATE DATABASE IF NOT EXISTS devbook;
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

DROP TABLE IF EXISTS followers;
CREATE TABLE followers(
  user_id int not null,
  FOREIGN KEY (user_id)
  REFERENCES users(id)
  ON DELETE CASCADE,

  follower_id int not null,
  FOREIGN KEY (follower_id)
  REFERENCES users(id)
  ON DELETE CASCADE,

  primary key(user_id, follower_id)
) ENGINE=INNODB;

DROP TABLE IF EXISTS posts;
CREATE TABLE posts(
  id int auto_increment primary key,
  user_id int not null,
  content varchar(200) not null,
  likes int not null default 0,
  created_at timestamp not null default current_timestamp(),
  
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=INNODB;