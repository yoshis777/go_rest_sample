# GoによるRest Apiサンプル
### パッケージ
go-json-rest
gorm

### DB作成
```bash
docker-compose up
mysql -h 127.0.0.1 -u root # ローカルPCのmysqlクライアントから

CREATE DATABASE go_rest_sample_development;
use go_rest_sample_development;

# goファイルから、
# db.Table("users").CreateTable(&User{}) でもよい
CREATE TABLE IF NOT EXISTS users (id bigint(20) not null primary key auto_increment, username varchar(255));
```

### gore導入（repl）
```bash
go get -u github.com/motemen/gore/cmd/gore
```
### 参考
* [golangでREST APIをやってみた①](https://qiita.com/katekichi/items/d94e078b376151858ca4)
* [Goでこれどうやるの？ 入門](https://www.slideshare.net/zaruhiroyukisakuraba/go-80884259)
    * 但し、goreに関しては、https://qiita.com/nagata03/items/cec2587140a376e345a9
* [[Go] Realizeが便利なので、もう少し仲良くなってみる](https://qiita.com/enta0701/items/9f60ad18600acab8c93d)

* [gorm リファレンス](https://gorm.io/ja_JP/docs/conventions.html)