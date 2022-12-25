package dao

var SentenceCreateDepository = `
CREATE TABLE IF NOT EXISTS depository(
	id INT PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(32) NOT NULL
) DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_bin COMMENT '仓库表'
`

var SentenceCreateCargo = `
CREATE TABLE IF NOT EXISTS cargo(
	id INT PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(64) NOT NULL,
	amount INT NOT NULL DEFAULT 0
) DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_bin COMMENT '货物表'
`

var SentenceCreateCargoDepository = `
CREATE TABLE IF NOT EXISTS cargo_depository(
	id INT PRIMARY KEY AUTO_INCREMENT,
	cargo_id INT NOT NULL,
	depository_id INT NOT NULL
) DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_bin COMMENT '货物和仓库对应表'
`
