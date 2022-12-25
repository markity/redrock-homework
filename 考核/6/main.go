package main

/*
CREATE TABLE user(
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(32) NOT NULL,
    password_md5 TINYBLOB   NOT NULL
)

CREATE TABLE book(
    id INT PRIMARY KEY AUTO_INCREMENT,
    book_name VARCHAR(64) NOT NULL,
    isbn varchar(64) NOT NULL UNIQUE,
    author VARCHAR(32) NOT NULL,
    pub VARCHAR(64) NOT NULL
)

CREATE TABLE borrowing_records(
    id INT PRIMARY KEY AUTO_INCREMENT,
    borrow_userid INT NOT NULL,
    borrow_bookid INT NOT NULL,
    book_status TINYINT NOT NULL,
    expire_at DATETIME NOT NULL
)

book_status 0 代表未归还但未逾期
			1 代表已归还
			2 代表未归还已经逾期


预处理语句可以减少处理时间，提交效率
*/

func main() {

}
