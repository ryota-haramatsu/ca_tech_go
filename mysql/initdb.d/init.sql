DROP DATABASE IF EXISTS ca_tech;
CREATE DATABASE ca_tech;
USE ca_tech;
DROP TABLE IF EXISTS users;

CREATE TABLE users ( 
  id       INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name     VARCHAR(255),
  token    VARCHAR(255)
)
-- なぜかSQLエラーが発生するのでコメント
-- INSERT INTO users (name, token) VALUES ('Nagaoka', 'asdjflkasdjflkkljj');
-- INSERT INTO users (name, token) VALUES ('tanaka', 'lkfjaskldjfljasdlkf');