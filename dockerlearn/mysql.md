# MySQL on docker
### Some usefull commands
```
docker pull mysql
docker run --name first -p 3306:3306 -e MYSQL\_ROOT\_PASSWORD=123456 -d mysql
mysql -h127.0.0.1 -P3306 -uroot -p123456
docker exec -it first bash
```
```
# using dir on host system
docker run --name first -p 3306:3306 -v /Users/baidu/go_misc/mysql/data:/var/lib/mysql -e MYSQL\_ROOT\_PASSWORD=123456 -d mysql
```

```SQL
CREATE TABLE `revival_card` (
  `uid` INT NOT NULL,
  `code` VARCHAR(45) NOT NULL,
  `count` INT(10) NOT NULL DEFAULT 0,
  `could_apply_by_sub` INT(1) NOT NULL DEFAULT 1,
  `could_apply_by_invite` INT(1) NOT NULL DEFAULT 1,
  `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`),
  INDEX `code_index` USING BTREE (`code` ASC),
  INDEX `uid_count_update_time_index` (`count` ASC, `uid` ASC, `update_time` ASC)
);
```
