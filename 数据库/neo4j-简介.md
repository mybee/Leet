# 图数据库之 Neo4j

### 为什么要有图数据库

大多数社交网络应用程序和视频托管应用程序都在使用更多连接的数据。这些应用程序包含大量的结构化，半结构化和非结构化的连接数据。 在RDBMS数据库中表示这种非结构化连接数据并不容易。

图分析可以深入探索各种实体（组织，人员，交易）之间复杂的相互关系，几乎任何可以建模和查询的数据。

如下面的示例：

![Google+应用个人资料](https://tva1.sinaimg.cn/large/007S8ZIlly1gga2pfiisjj309d0730so.jpg)

![连接的数据](https://tva1.sinaimg.cn/large/007S8ZIlly1gga2pxt2yrj309d0730so.jpg)

图数据库以图形结构（即网络状）的形式存储数据，他以“节点”为基本的存储单位，“节点”之间则以“关系”作为关联，“节点”与“关系”中可以包含许多“属性”，正如RDS里的“行”、“列”，图数据库以图形结构进行存储。

![从网络上找到的一个对比图](https://tva1.sinaimg.cn/large/007S8ZIlgy1gga2q8m2cuj30jp0jpjrv.jpg)

### 常见的图数据库
- 	Neo4j
-  Microsoft Azure Cosmos DB
-  ArangoDB
-  OrientDB

### Neo4j 安装
- 可以直接安装 desktop 版本，或者使用 docker 安装；
```bash
docker pull neo4j
docker run -it -d -p 7474:7474 -p 7687:7687 neo4j:latest
```

说明：7474 for http, 7473 for https, *7687* for bolt

##  基本概念

Neo4j是一个世界领先的图数据库，具有本机图形存储与处理功能，具有展示方便、支持面广、性能优异等优点。

包含有灵活的架构、完整的事务管理、集群以提升高可用与性能、强大的Cypher（CQL）查询语言、可视化Web界面 Neo4j Browser、驱动支持时下流行的语言与框架、方便的数据导入、成熟的云端服务等特性。

#### 支持

- 它遵循属性图数据模型

- 它通过使用Apache Lucence支持索引

- 它支持UNIQUE约束

- 它包含一个用于执行CQL命令的UI：Neo4j数据浏览器

- 它支持完整的ACID（原子性，一致性，隔离性和持久性）规则

- 它采用原生图形库与本地GPE（图形处理引擎）

- 它支持查询的数据导出到JSON和XLS格式

- 它提供了REST API，可以被任何编程语言（如Java，Spring，Scala等）访问

- 它提供了可以通过任何UI MVC框架（如Node JS）访问的Java脚本

- 它支持两种Java API：Cypher API和Native Java API来开发Java应用程序

#### 优点
- 它很容易表示连接的数据

- 检索/遍历/导航更多的连接数据是非常容易和快速的

- 它非常容易地表示半结构化数据

- Neo4j CQL查询语言命令是人性化的可读格式，非常容易学习

- 使用简单而强大的数据模型

- 它不需要复杂的连接来检索连接的/相关的数据，因为它很容易检索它的相邻节点或关系细节没有连接或索引
#### 缺点

- 不支持分片，没有水平扩展性能。



#### 数据模型

- 节点（使用圆圈表示，包含零到多个属性）
- 关系（使用带箭头的线表示，连接节点，包含零到多个属性）
- 属性（命名的数据值，键值对表示）

示例：

<img src="https://tva1.sinaimg.cn/large/007S8ZIlgy1gga2qg5blaj30ks0ccaaj.jpg" alt="图的表示" style="zoom:57%;" />

这里使用圆圈作为节点，箭头表示关系，关系是有方向的。这里节点中的值为 ID 属性。

单个节点内部：

![img](https://tva1.sinaimg.cn/large/007S8ZIlgy1gga2r0a0ksj303g03gq2q.jpg)

**注：**

- 关系具有方向：单向和双向。
- 每个关系包含“开始节点”或“从节点”和“到节点”或“结束节点”

在属性图数据模型中，关系应该是定向的。

##### 属性中支持的数据类型

| 序号 | CQL数据类型 | 用法                            |
| ---- | ----------- | ------------------------------- |
| 1.   | boolean     | 用于表示布尔文字：true，false。 |
| 2.   | byte        | 用于表示8位整数。               |
| 3.   | short       | 用于表示16位整数。              |
| 4.   | int         | 用于表示32位整数。              |
| 5.   | long        | 用于表示64位整数。              |
| 6.   | float       | 用于表示32位浮点数。            |
| 7.   | double      | 用于表示64位浮点数。            |
| 8.   | char        | 用于表示16位字符。              |
| 9.   | string      | 用于表示字符串。                |

#### 标签

关联的一组节点。一个节点可以有零个或者多个标签，标签不具有属性。



## CQL 使用

CQL 代表 Cypher 查询语言，相当于 MySQL 中的 SQL。

准备测试数据，直接使用官方指南的例子（Movie-Graph）。

### 创建

其中包括，创建演员，创建电影，创建关系

```CQL
CREATE (TomH:Person {name:'Tom Hanks', born:1956})
CREATE (HelenH:Person {name:'Helen Hunt', born:1963})
CREATE (RobertZ:Person {name:'Robert Zemeckis', born:1951})
CREATE (CastAway:Movie {title:'Cast Away', released:2000, tagline:'At the edge of the world, his journey begins.'})
CREATE
(TomH)-[:ACTED_IN {roles:['Chuck Noland']}]->(CastAway),
(HelenH)-[:ACTED_IN {roles:['Kelly Frears']}]->(CastAway),
(RobertZ)-[:DIRECTED]->(CastAway)
```

### 独立节点查找

```CQL
// 按名字查演员
MATCH (tom {name: "Tom Hanks"}) RETURN tom
// 按标题查电影
MATCH (cloudAtlas {title: "Cloud Atlas"}) RETURN cloudAtlas
// 列表显示演员，指定显示字段
MATCH (people:Person) RETURN people.name LIMIT 10
// 查找发行于九十年代的电影
MATCH (nineties:Movie) WHERE nineties.released >= 1990 AND nineties.released < 2000 RETURN nineties.title
```

### 模式查找

```CQL
// 查询 Tom Hanks 出演的所有电影
MATCH (tom:Person {name: "Tom Hanks"})-[:ACTED_IN]->(tomHanksMovies) RETURN tom,tomHanksMovies

// 查询电影的导演名称
MATCH (cloudAtlas {title: "Cloud Atlas"})<-[:DIRECTED]-(directors) RETURN directors.name

// 查找与 Tom Hanks 出演过同样电影的合作者
MATCH (tom:Person {name:"Tom Hanks"})-[:ACTED_IN]->(m)<-[:ACTED_IN]-(coActors) RETURN coActors.name

// 查找与电影云图相关的人，展示这些人名、与电影的关系类型、关系详情
MATCH (people:Person)-[relatedTo]-(:Movie {title: "Cloud Atlas"}) RETURN people.name, Type(relatedTo), relatedTo
```

### 解决路径问题

培根定律”（Bacon's Law）是基于“六度分隔”概念的客厅游戏，该概念假定地球上的任何两个人相隔六个或六个以下熟人。电影爱好者互相挑战，以找到任意演员和多产演员凯文·培根之间的最短路径。

1. 可变长度模式
2. 内置的`shortestPath()`算法

```CQL
// 距离 Kevin Bacon 4 度的所有人和电影
MATCH (bacon:Person {name:"Kevin Bacon"})-[*1..4]-(hollywood)
RETURN DISTINCT hollywood

// 两个人之间的最短距离
MATCH p=shortestPath((bacon:Person {name:"Kevin Bacon"})-[*]-(meg:Person {name:"Meg Ryan"}))
RETURN p
```

4 度节点结果：

![image-20200702113050537](https://tva1.sinaimg.cn/large/007S8ZIlly1ggcgxym1t0j318g0u04cr.jpg)

最短路径结果：

<img src="https://tva1.sinaimg.cn/large/007S8ZIlly1ggch1nw8p8j30os0o0tb2.jpg" alt="image-20200702113424098" style="zoom:50%;" />

### 推荐问题

给一个演员推荐新的合作者，基本的方案就是基于与他的合作者的其他合作者，可以演变出两种查询：

- 给定一个演员，推荐新的合作者；

  ![neo4j](https://tva1.sinaimg.cn/large/007S8ZIlly1ggacbr1c0uj309807h0sw.jpg)

- 查找中间人（可以将一个演员介绍给潜在合作人的人）

```CQL
// 推荐的合作者，按出现次数排序
MATCH (tom:Person {name:"Tom Hanks"})-[:ACTED_IN]->(m)<-[:ACTED_IN]-(coActors),
  (coActors)-[:ACTED_IN]->(m2)<-[:ACTED_IN]-(cocoActors)
WHERE NOT (tom)-[:ACTED_IN]->()<-[:ACTED_IN]-(cocoActors) AND tom <> cocoActors
RETURN cocoActors.name AS Recommended, count(*) AS Strength ORDER BY Strength DESC
// 注：CQL 中没有 group by，聚合函数 count 的结果取决于 return 返回的其他非聚合字段，会使用该值作为聚合的 key，本例中使用 Recommended 分组。

// 可以把 Tom Hanks 介绍给 Tom Cruise 的人
MATCH (tom:Person {name:"Tom Hanks"})-[:ACTED_IN]->(m)<-[:ACTED_IN]-(coActors),
  (coActors)-[:ACTED_IN]->(m2)<-[:ACTED_IN]-(cruise:Person {name:"Tom Cruise"})
RETURN tom, m, coActors, m2, cruise
```

查找中间人结果：

![image-20200702142130132](https://tva1.sinaimg.cn/large/007S8ZIlgy1ggclvnwg17j31jk0ooagu.jpg)



## CQL 语法详细说明

neo4j 的 CQL 语言是专门用来处理图形数据。

特点：

- 使用模式来描述图形语言；
- 类似于 SQL 语法；
- 说明性的语法，描述要查找的内容，而不是如何查找；



#### 常用命令与关键字

| 序号 | CQL命令/条       | 用法                         |
| ---- | ---------------- | :--------------------------- |
| 1    | CREATE 创建      | 创建节点，关系和属性         |
| 2    | MATCH 匹配       | 检索有关节点，关系和属性数据 |
| 3    | RETURN 返回      | 返回查询结果                 |
| 4    | WHERE 条件       | 提供条件过滤检索数据         |
| 5    | DELETE 删除      | 删除节点和关系               |
| 6    | REMOVE 移除      | 删除节点和关系的属性         |
| 7    | ORDER BY 以…排序 | 排序检索数据                 |
| 8    | SET 组           | 添加或更新标签               |

##### Create 命令

示例：

```CQL
CREATE (ee:Person { name: "Emil", from: "Sweden", klout: 99 })
```

说明：

- `()`中包含的是节点；
- `ee:Person`，其中`ee` 是变量，如果没有用到可以省略，`Person` 是标签；
- `{}`中的是节点的属性；

可以创建一个没有属性的人：

`CREATE (ee:Person)`

##### Match 命令

该命令不可以单独使用，必须包含 `RETURN` 或者是一个更新语句。

示例：

```CQL
MATCH (ee:Person) WHERE ee.name = "Emil" RETURN ee;
```

说明：

- `(ee:Person)`必须指定标签和变量
- 类似于 `SQL` 的 `WHERE` 条件，数据中的变量使用变量加点的方式；
- `RETURN` 用于指定返回的数据，可以是具体字段；

##### 创建多个节点和关系

示例：

```CQL
MATCH (ee:Person) WHERE ee.name = "Emil"
CREATE (js:Person { name: "Johan", from: "Sweden", learn: "surfing" }),
(ir:Person { name: "Ian", from: "England", title: "author" }),
(rvb:Person { name: "Rik", from: "Belgium", pet: "Orval" }),
(ally:Person { name: "Allison", from: "California", hobby: "surfing" }),
(ee)-[:KNOWS {since: 2001}]->(js),(ee)-[:KNOWS {rating: 5}]->(ir),
(js)-[:KNOWS]->(ir),(js)-[:KNOWS]->(rvb),
(ir)-[:KNOWS]->(js),(ir)-[:KNOWS]->(ally),
(rvb)-[:KNOWS]->(ally)
```

说明：

- 基于`MATCH`查到的节点赋值到变量 ee 中；
- `CREATE`创建节点，关系
- `(node1)-[:KNOWS]{}]->(node2)` 为创建关系的语法，`{}`中指定属性，没有属性可以省略；

##### 模式查询

示例，查找 Email 的朋友：

```CQL
MATCH (ee:Person)-[:KNOWS]-(friends)
WHERE ee.name = "Emil" RETURN ee, friends
```

示例，查找会冲浪的新朋友：

```CQL
MATCH (js:Person)-[:KNOWS]-()-[:KNOWS]-(surfer)
WHERE js.name = "Johan" AND surfer.hobby = "surfing"
RETURN DISTINCT surfer
```

- `()`表示忽略这些节点；
- ``DISTINCT`多条路径会匹配到相同的人；



##### `EXPLAIN` or `PROFILE` 分析查询执行过程

```CQL
PROFILE MATCH (js:Person)-[:KNOWS]-()-[:KNOWS]-(surfer)
WHERE js.name = "Johan" AND surfer.hobby = "surfing"
RETURN DISTINCT surfer
```



##### Where 字句

语法：其中布尔运算符号包括，与、或、非、异或

```CQL
// 简单语法
WHERE <condition>
// 复杂语法
WHERE <condition> <boolean-operator> <condition>
#condition 语法
<property-name> <comparison-operator> <value>
```

查找名字为 Emil 的人：

```CQL
MATCH (p:Person) 
WHERE p.name = 'Emil'
RETURN p
```

##### Delete 

语法：删除节点或关系

```CQL
// 删除关系
match(p1:Person)-[ref:KNOWS]-(p2:Person)
DELETE ref

// 删除标签为 People 的所有人
MATCH (p: Person) DELETE p

// 删除所有的数据（包括关系和节点）
MATCH (n) DETACH DELETE n
```

**注：** 带关系的节点无法直接删除，必须先删除关系。

##### Set 字句

向现有节点或关系添加新属性。

```CQL
MATCH (rp:Person)
SET rp.kind = "humanity"
RETURN rp
```

##### Remove 字句

删除节点或关系的属性，删除节点的标签：

```CQL
MATCH (rp:Person)
REMOVE rp.kind
RETURN rp

// 删除标签
MATCH (m:Movie) 
REMOVE m:Picture
```

##### Order by 字句

```CQL
MATCH (p:Peopel)
RETURN p.name,p.from
ORDER BY p.name DESC
```

##### UNION 字句

它将两组结果中的公共行组合并返回到一组结果中。 

Union 不从两个节点返回重复的行；

Union all 子句不过滤它们的重复行；

```CQL
Create(cc:CreditCard{id:1, number:111, name:"Spike",cc:"cc"});
Create(cc:CreditCard{id:2, number:222, name:"Tom",cc:"cc"});
Create(cc:CreditCard{id:3, number:333, name:"Jerry",cc:"cc"});
Create(dc:DebitCard{id:4, number:444, name:"Spike",dc:"dc"});
Create(dc:DebitCard{id:5, number:555, name:"Tom",dc:"dc"})
```

必须使用相同的列才能正确返回，因此需要提供别名语法 `AS`：

```CQL
MATCH (cc:CreditCard)
RETURN cc.id as id,cc.number as number,cc.name as name
UNION
MATCH (dc:DebitCard)
RETURN dc.id as id,dc.number as number,dc.name as name
```

##### LIMIT 和 SKIP 子句

```CQL
MATCH (cc:CreditCard) 
RETURN cc
LIMIT 1

// skip 必须在 limit 前面
MATCH (cc:CreditCard) 
RETURN cc
SKIP 1
LIMIT 1
```

##### MERGE 命令（CREATE + MATCH）

在图中搜索给定的模式，存在则返回结果；不存在就创建新的节点/关系，然后返回结果。

##### NULL

 CQL将空值视为对节点或关系的属性的缺失值或未定义值。

```CQL
MATCH (cc:CreditCard) 
WHERE cc.id IS NOT NULL
RETURN cc
```

##### IN 操作符

```CQL
MATCH (cc:CreditCard) 
WHERE cc.id IN [1,124]
RETURN cc
```

##### id

`id`节点和关系的默认内部属性。这意味着，当我们创建一个新的节点或关系时，Neo4j数据库服务器将为内部使用分配一个数字。 它会自动递增。节点的`id`属性的最大值约为35亿。

##### 聚合

|  1   | COUNT |      它返回由MATCH命令返回的行数。      |
| :--: | :---: | :-------------------------------------: |
|  2   |  MAX  |  它从MATCH命令返回的一组行返回最大值。  |
|  3   |  MIN  | 它返回由MATCH命令返回的一组行的最小值。 |
|  4   |  SUM  | 它返回由MATCH命令返回的所有行的求和值。 |
|  5   |  AVG  | 它返回由MATCH命令返回的所有行的平均值。 |

**注：**CQL 的分组由返回的其他非聚合字段决定，不需要显示指定。

示例：

```CQL
match(p:Person)-[]-(m:Movie)
return p.name, max(m.released)
```

##### 索引

```CQL
// 创建
CREATE INDEX ON :<label_name> (<property_name>)

// 删除
DROP INDEX ON :<label_name> (<property_name>)

//创建唯一索引
CREATE CONSTRAINT ON (<label_name>)
ASSERT <property_name> IS UNIQUE

// 删除唯一索引
DROP CONSTRAINT ON (<label_name>)
ASSERT <property_name> IS UNIQUE
```



## 驱动模式

neo4j 提供了两种驱动模式，一种是 Bolt，二进制协议，更加紧凑，吞吐量更高；另一种是 http，使用比较方便。

!(https://tva1.sinaimg.cn/large/007S8ZIlly1ggbwedhey9j30u00wlwi4.jpg)

<img src="https://tva1.sinaimg.cn/large/007S8ZIlly1ggbwedhey9j30u00wlwi4.jpg" style="zoom:40%;" />




### http api

#### 事务流程

事务是通过几个不同的 URI 管理的，这些 URI 被设计为以规定的模式使用。提供了在单个 HTTP 请求，或者多个 HTTP 请求上执行整个事务周期的功能。

总体流程说明如下，每一个盒子代表一个单独的 HTTP 请求：

![http cypher transaction api flow](https://tva1.sinaimg.cn/large/007S8ZIlly1ggagfcszxdj30m80dwtap.jpg)

#### 事务有效期

事务开始后，它的状态会在服务器上保存。一段时间不活跃，事务就会自动过期。默认为 60 秒。

可以在/ tx / {n} URI 中提交空的声明，来达到不提交新的查询，也能保持事务的存活。

示例：

```bash
// 使用一个请求完成事务
curl --location --request POST 'http://localhost:7474/db/neo4j/tx/commit' \
--header 'Authorization: Basic bmVvNGo6cm9vdA==' \
--header 'Content-Type: application/json' \
--data-raw '{
    "statements": [
        {
            "statement": "CREATE (n:test $props) RETURN n",
            "parameters": {
                "props": {
                    "name": "Tom",
                    "age":3
                }
            }
        }
    ]
}'
```



### 参考：

1. https://neo4j.com/doc
2. https://www.cnblogs.com/tinging/p/12088259.html
3. https://www.tony-bro.com/posts/2624331944/
4. https://www.w3cschool.cn/neo4j/neo4j_need_for_graph_databses.html

