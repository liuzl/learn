# 常用SQL

* 列出`DATA_PUMP_DIR`目录里面的文件
```sql
select * from table(RDSADMIN.RDS_FILE_UTIL.LISTDIR('DATA_PUMP_DIR'));
```
