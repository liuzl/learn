# neo4j misc

## 添加索引

[参考资料](https://neo4j.com/docs/getting-started/current/cypher-intro/schema/#cypher-intro-schema-indexes)
```cql
CREATE CONSTRAINT ON (movie:Movie) ASSERT movie.title IS UNIQUE
```

## 根据ID插入关系
```cql
MATCH(i) where ID(i)=9117940 match(d) where ID(d)=9117941  MERGE (i)-[:FEATURES]->(d)
```
