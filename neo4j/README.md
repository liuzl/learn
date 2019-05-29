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

## 批量插入数据 neo4j-import
```sh
time ./bin/neo4j-import  --into ./data/graph.db/ \
--array-delimiter TAB
--nodes:Company "/data/graph/company-header.csv,/data/graph/company.csv" \
--nodes:Person "/data/graph/person-header.csv,/data/graph/person.csv" \
--relationships:Invest "/data/graph/company_invest.csv" \
--relationships:Invest "/data/graph/person_invest-header.csv,/data/graph/person_invest.csv" \
--relationships:LegalPerson "/data/graph/legal-header.csv,/data/graph/legal.csv" \
--relationships:Staff "/data/graph/staff-header.csv,/data/graph/staff.csv"
```
