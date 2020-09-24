DROP TABLE IF EXISTS articles;

CREATE TABLE articles
(
  id MEDIUMINT NOT NULL AUTO_INCREMENT,
  users_id VARCHAR(30),
  title VARCHAR(240),
  main_text VARCHAR(3000),
  tag VARCHAR(30),
  release_date DATE,
  PRIMARY KEY (id)
);
INSERT INTO articles (users_id,title,main_text,tag,release_date) VALUES ("ramo123", "タイトル","メインテキスト","タグタグ","2019-09-07");