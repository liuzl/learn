# SQL Server on Linux

* client tool: `mssql-cli`

## 常用命令

### 列出数据库里所有的表名
```sql
select name from sysobjects where type='U'
```
### 列出表里的所有的列
```sql
select name from syscolumns where id=object_id('TableName')
```
