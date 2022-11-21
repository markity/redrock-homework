package main

// sentence to create user table
var migrateTableUser = `CREATE TABLE IF NOT EXISTS user(
	id INT PRIMARY KEY AUTO_INCREMENT,
	username VARCHAR(32) NOT NULL UNIQUE		COMMENT 'unique user name',
	password_md5 TINYBLOB NOT NULL				COMMENT "user's md5 password digest",
	created_at DATETIME NOT NULL				COMMENT "user's entry datatime"
) COMMENT 'a table to store user login data'
`

// sentence to create a middle table security_user to implement many-to-many structure
var migrateTableSecurityUser = `CREATE TABLE IF NOT EXISTS security_user(
	id INT PRIMARY KEY AUTO_INCREMENT,
	security_id INT NOT NULL					COMMENT "associate security's id field",
	user_id INT NOT NULL						COMMENT "associate user's id field",
	answer VARCHAR(50) NOT NULL					COMMENT "user's security answer"
) COMMENT 'a middle table to implment many-to-many relationship between user and security'
`

// sentence to create security question table named security
var migrateTableSecurity = `CREATE TABLE IF NOT EXISTS security(
	id INT PRIMARY KEY AUTO_INCREMENT,
	question VARCHAR(50) NOT NULL
) COMMENT "a table to store security questions"
`

// sentence to create table comment
var migreateTableComment = `CREATE TABLE IF NOT EXISTS comment(
	id INT PRIMARY KEY AUTO_INCREMENT,
	user_id INT NOT NULL						COMMENT "the sender's user id",
	message TEXT(512) NOT NULL					COMMENT 'comment content',
	parent INT DEFAULT NULL						COMMENT "specify a comment's parent, null when it is the top comment",
	created_at DATETIME NOT NULL				COMMENT 'comment created at'
) COMMENT 'a table to stroe comments'
`
