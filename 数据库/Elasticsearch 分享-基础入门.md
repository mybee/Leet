

## Elasticsearch 分享

本次分享主要分为从以下几点展开：

基础理论、查询、聚合、文本分析和映射、节点与分片、监控与诊断

### 1. 基础理论

Elasticsearch 是一个开源的搜索引擎，建立在一个全文搜索引擎库 [Apache Lucene™](https://lucene.apache.org/core/) 基础之上。 Lucene 仅仅只是一个 Java 库，Elasticsearch 内部使用 Lucene，屏蔽了它的复杂性，提供一套简单一致的 RESTful API（es 也是 Java 写的）。

主要特点为**分布式**、**近实时**、**全文搜索引擎**。

使用上来说，因为预设了一些适当的默认值，隐藏了大部分复杂性，可以做到开箱即用。

#### 一些基本术语

##### NRT 近实时（Near Realtime NRT）

当文档索引后，当变为能够搜索可见，会有一个轻微的延迟，通常来说是 1s。

##### cluster

集群,一个 ES 集群由一个或多个节点（Node）组成，每个集群都有一个 cluster name 作为标识。

##### node

节点，一个 ES 实例就是一个节点，一个机器可以有多个实例，所以并不能说一台机器就是一个节点，大多数情况下每个节点运行在一个独立的环境或虚拟机上。

##### index

索引，即一系列文档的集合，一个集群中可以有多个。

##### type

类型，索引下的分类，可以在一个索引下区分类型存储文档。低版本（6.x 之前）中可以存在多个类型，因为引入这个概念并不是很合理，所以高版本中逐步淡化这个概念，最终的目标是完全移除这个概念。

##### document

文档，是可以被索引的基本信息单元。

##### shard & replicas

分片和副本。为了应对大规模数据下，单点性能不足、无法并行操作，提出了数据分片，即将数据切割为多份，分布于不同节点。在创建索引可以指定分片数，默认 5。副本是为了提供服务的高可用，而提供的额外数据复制集，创建索引时也可以指定，默认 1。

#### 安装

安装部分直接使用 docker 方式，其他方式大家可以自行 google

```bash
docker run -d --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:6.3.2
```

注：9200 是提供服务的端口，9300 是集群节点内部通讯使用的端口；

可以配合 kibana 使用更加方便：

```bash
docker run --name kibana --link=elasticsearch -p 5601:5601 -d docker.elastic.co/kibana/kibana:6.3.2
```

注：--link 中的值要与 elasticsearch 中的name 对应

在 kibana 的 Console 中可以省略 curl 请求中的 IP、端口、Header 等信息，下面示例中都会使用这种简化的操作。

#### 创建一个索引

```bash
# 完整请求
curl -X PUT "localhost:9200/customer?pretty"

# 简略请求
PUT /customer?pretty

# 查看所有索引
GET /_cat/indices?v
```

该操作可以省略，为某个索引增加一个文档也会自动创建索引

#### Elasticsearch 中的增删改查

在 elasticsearch 中使用 Json 作为文档序列化格式，一个文档或者说一条记录就是一个 Json，当然这个 Json 可以很复杂。

- es 中增加数据一般称为索引

  ```bash
  PUT /customer/_doc/1?pretty
  {
    "name": "John Doe"
  }
  ```
  
- 修改这个文档（一般来说可以全部覆盖的方式修改，或者指定字段修改）

  ```bash
  PUT /customer/_doc/1
  {
    "name": "Jerry"
  }
  
  # 指定字段修改
  POST /customer/_doc/1/_update?pretty
  {
    "doc": { "name": "Jerry" }
  }
  
  # 修改的同时也可以添加新的字段
  POST /customer/_doc/1/_update?pretty
  {
    "doc": { "name": "Jerry", "age": 2 }
  }
  ```

- 获取这个文档

  ```bash
  GET /customer/_doc/1?pretty
  ```

- 删除这个文档

  ```bash
  DELETE /customer/_doc/1
  ```

#### 批量操作

批量操作是通过一个请求，执行多个操作，可以极大的提升效率，目前支持 4 中操作，在同一个请求中并不局限与单种类型的操作；其每一个操作都是通过`\n`进行分隔。

- create 创建文档
- index 添加新文档或者覆盖已有文档
- update 仅更新已有文档
- delete 已有文档

```bash
POST /customer/_doc/_bulk?pretty
{"create":{"_id":"2"}}
{"name": "tom" }
{"index":{"_id":"3"}}
{"name": "spike" }
{"update":{"_id":"3"}}
{"doc":{"age": "3" }}
{"delete":{"_id":"4"}}
```

### 文本分析与映射
索引其实是有表结构，这里称为映射，不过默认情况下会根据创建的文档进行自动生成，该配置可以在主动创建索引时更改。

#### 数据类型

- 数字类型，具体按存储的长度又细分为 long，integer，short 等；
- 布尔类型；
- 日期类型，可以指定时间戳格式，也可以指定字符串格式，可选格式很多；
- Keyword 类型，一个不分词的字符串，有存储长度限制；
- Text 类型，分词类型，可以主动指定分词器；
- 对象类型，复合数据类型；

因为不同的字段类型存储方式不同，一旦创建好，即不可修改，不可删除。

可以通过如下方式查看索引映射：

```bash
GET customer/_mapping
```

映射的创建

```bash
PUT twitter/_mapping/_doc
{
    "properties": {
        "line_id": {
            "type": "long"
        },
        "line_number": {
            "type": "keyword"
        },
        "play_name": {
            "type": "keyword"
        },
        "speaker": {
            "type": "keyword"
        },
        "speech_number": {
            "type": "long"
        },
        "text_entry": {
            "type": "text",
            "analyzer":"ik_max_word"
        }
    }
}
```
#### 文本分析（Analysis）

它是将文本块转换为有区别的、规范化的 token 的一个过程，并最终创建底层存储的倒排索引。


> 安装中文分词器，elasticsearch-analysis-ik-6.3.2.zip，过程略。

这个插件带了 2 种解析器，`ik_max_word`，`ik_smart`。

- ik_max_word，会将文本做最细粒度的拆分；
- ik_smart，会做较粗粒度的拆分；

```bash
# 示范
POST _analyze
{
  "analyzer": "ik_max_word", 
  "text":"中华人民共和国国歌"
}
```

### 查询（搜索）

搜索是 es 中一个最基本的工具，主要的搜索方式有空搜索、范围搜索、精确搜索、全文搜索等，且可以指定返回字段、排序、分页等。格式为使用 `_search`端点，GET 或者 POST 请求，加上查询 DSL（JSON-style domain-specific language）；

导入一些测试数据：

```bash
wget https://raw.githubusercontent.com/zq2599/blog_demos/master/files/create_shakespeare_index.sh \

&& chmod a+x create_shakespeare_index.sh \

&& ./create_shakespeare_index.sh localhost 9200
```

#### 空搜索

```bash
# 就是没有 dsl
GET shakespeare/_search?pretty

# 等价于
GET /_search?pretty
{
  "query": {"match_all": {}}
}
```

#### 精确搜索

 `term` 关键字，最通用的精确匹配；

```bash
GET shakespeare/_search?pretty
{
  "query": {
    "term": {
      "line_id": {
        "value": "20"
      }
    }
  }
}
# terms 支持多个，相当于 MySQL 中的 in 操作
GET shakespeare/_search?pretty
{
  "query": {
    "terms": {
      "line_id": [
        20,
        30
      ]
    }
  }
}
```

#### 范围搜索

使用 `range` 关键字，可以查询数字、日期、字符串类型

```bash
# 范围 [10, 12]
GET shakespeare/_search?pretty
{
  "query": {
    "range": {
      "line_id": {
        "gte": 10,
        "lte": 12
      }
    }
  }
}
```

#### 全文搜索

**相关性（Relevance）**

它是评价查询与其结果间的相关程度，并根据这种相关程度对结果排名的一种能力，这种计算方式可以是 TF/IDF（检索词频率/反向文档频率） 方法（参见 [相关性的介绍](https://www.elastic.co/guide/cn/elasticsearch/guide/current/relevance-intro.html)）、地理位置邻近、模糊相似，或其他的某些算法。


这里有基于词项的查询`term`，匹配查询的`match`，基于短语的`match_phrase`。其中 `match` 是很重要的一个关键字，它会根据不同的数据类型做出相应的变化。对于文本来说，会分析查询文本，生成词项列表，然后会对每个词项逐一执行底层的查询，再将结果合并，然后为每个文档生成一个最终的相关度评分。

```bash
# term
GET shakespeare/_search?pretty
{
  "query": {
    "term": {
      "text_entry": "shaken"
    }
  }
}

# match
GET shakespeare/_search?pretty
{
  "query": {
    "match": {
      "text_entry": "So shake as we are"
    }
  }
}
# match_phrase
GET shakespeare/_search?pretty
{
  "query": {
    "match_phrase": {
      "text_entry": "oak not to be"
    }
  }
}
```



#### 组合查询（AND OR NOT）

组合查询主要由布尔过滤器实现，一个布尔过滤器可以有 4 部分组成（旧版本无 filter）：

```json
{
   "bool" : {
      "must" :     [],
      "should" :   [],
      "must_not" : [],
      "filter" :   []
   }
}
```

其中，`must` 相当于与操作，`should` 相当于或操作，`must_not`相当于非。`filter` 比较特殊，是一个过滤操作，跟 `must` 类似，但是它不会计算评分。这些组合词可以跟上面的查询相互组合，完成更复杂的查询。

```bash
# where speech_number <> 1 and line_id in (15000, 20000)
GET shakespeare/_search?pretty
{
  "query": {
    "bool": {
      "must_not": [
        {
          "term": {
            "speech_number": 1
          }
        }
      ],
      "must": [
        {
          "terms": {
            "line_id": [
              20000,
              15000
            ]
          }
        }
      ]
    }
  }
}
```

```bash
# where line_id between 1 and 1000 and (speech_number = 1 or speaker = "POMPEY" or line_number <> "2.1.1")
GET shakespeare/_search?pretty
{
  "query": {
    "bool": {
      "must": [
        {
          "range": {
            "line_id": {
              "gte": 1,
              "lte": 20000
            }
          }
        }
      ],
      "should": [
        {
          "term": {
            "speech_number": 1
          }
        },
        {
          "term": {
            "speaker": "POMPEY"
          }
        },
        {
          "bool": {
            "must_not": [
              {
                "term": {
                  "line_number": "2.1.1"
                }
              }
            ]
          }
        }
      ]
    }
  }
}
```

```bash
#filter 不会计算评分，所有的评分都是 0，会缓存结果集
GET shakespeare/_search?pretty
{
  "query": {
    "bool": {
      "filter": [
        {
          "range": {
            "line_id": {
              "gte": 1,
              "lte": 20000
            }
          }
        }
      ]
    }
  }
}
```

#### 通用功能

- 使用`source` 进行过滤返回的字段；
- 使用`from`，`size`进行分页； 
- 使用`sort` 进行排序；

### 聚合

ES 除了全文搜索引擎，另一个亮点就是聚合。特点是实时性非常高，所有的计算结果都是即时返回。在 ES 的聚合中一共分为四大类：

- `指标聚合（Metric）` 可以对文档字段进行统计分析，比如计算最大值、最小值、平均值等；
- `桶聚合（Bucket）` 就是满足特定条件的文档的集合；
- `管道聚合（Pipeline）` 对其他聚合结果进行二次聚合；
- `矩阵聚合（Metrix）` 支持对多个字段的操作并提供相应统计量；


#### 指标聚合
常用的指标有最小值、最大值、平均值、求和、文档数、数据百分数；
平均值聚合：

```bash
GET shakespeare/_search?size=0
{
  "aggs": {
    "avg_line_id": {
      "avg": {
        "field": "line_id"
      }
    }
  }
}
```
百分数聚合：

```bash
# line_id 特定百分位的值
GET shakespeare/_search?size=0
{
  "aggs": {
    "percent_line_id": {
      "percentiles": {
        "field": "line_id"
      }
    }
  }
}
# 给定某些 line_id 的值，查看所处的百分位
GET shakespeare/_search?size=0
{
  "aggs": {
    "percent_line_id": {
      "percentile_ranks": {
        "field": "line_id",
        "values": [100, 10000,20000]
      }
    }
  }
}
```

#### 桶聚合
terms 分桶

```bash
GET shakespeare/_search?size=0
{
  "aggs": {
    "group_by_play_name": {
      "terms": {
        "field": "play_name"
      }
    }
  }
}
```

柱状图分桶，范围分桶

```bash
# 固定区间分桶（柱状图）
GET shakespeare/_search?size=0
{
  "aggs": {
    "line_ids": {
      "histogram": {
        "field": "line_id",
        "interval": 50000
      }
    }
  }
}

# 返回指定分桶规则（左闭右开）
GET shakespeare/_search?size=0
{
  "aggs": {
    "range_line_id": {
      "range": {
        "field": "line_id",
        "ranges": [
          {
            "to": 100
          },
          {
            "from": 100,
            "to": 1000
          }
        ]
      }
    }
  }
}
```

桶聚合与指标聚合嵌套

```bash
# 这里是按 play_name 分桶，然后计算每个 play_name 桶中 line_id 的 平均值。
GET shakespeare/_search?size=0
{
  "aggs": {
    "group_by_play_name":{
      "terms": {
        "field": "play_name",
        "size": 10
      },
      "aggs": {
        "avg_id": {
          "avg": {
            "field": "line_id"
          }
        }
      }
    }
  }
}
```
#### 管道聚合
```bash
GET shakespeare/_search?size=0
{
  "aggs": {
    "play_names":{
      "terms": {
        "field": "play_name",
        "size": 10
      },
      "aggs": {
        "avg_id": {
          "avg": {
            "field": "line_id"
          }
        }
      }
    },
     "avg_play_name": {
      "avg_bucket": {
        "buckets_path": "play_names>avg_id" 
      }
    }
  }
}
```

### 节点与分片
##### shard
1. 分片，ES是分布式搜索引擎，每个索引有一个或多个分片，索引的数据被分配到各个分片上，相当于一桶水用了 N 个杯子装。
2. 分片有助于横向扩展，N 个分片会被尽可能平均地（rebalance）分配在不同的节点上（例如你有2个节点,4个主分片（不考虑备份），那么每个节点会分到 2 个分片，后来增加了 2 个节点，那么这 4 个节点上都会有 1 个分片，这个过程叫 Relocation，ES感知后自动完成）。
3. 每个分片都是一个 Lucene Index，所以一个分片存放的文档数量是有限的。

##### replica
1. 复制，可以理解为备份分片或者从分片，相应地有 primary shard(主分片)。
2. 主分片和从分片不会出现在同一个节点上（防止单点故障）,默认情况下一个索引创建 5 个分片，1 
个备份，即 5 * (primary + replica) = 10个分片
3. 如果你只有一个节点,那么 5 个 replica 都无法分配（unassigned）。

用一个图来表示：

<img src="es1.png" alt="c184b2e6030d9ef05d33bcc08309a031" style="zoom:80%;" />

##### 节点
ES 作为一个分布式应用，为了更好的管理集群，维护集群的稳定，区分节点功能。
严格来说节点可以分为 4 种，协调（coordinating）、摄取（ingest）、主（master）、数据（data）。

- 候选主节点：可以成为整个集群主节点的节点，当主节点故障时，其他候选主节点会通过 raft 协议选择新主节点。
- 协调节点：每个节点是协调节点，不可能更改，只会用到机器的内存与 cpu。当收到请求会转发给数据节点，并合并数据。
- 数据节点：存储数据的节点，磁盘性能会影响机器的性能。
- 摄取节点：可以在文档索引前对文档进行预处理操作。
![node](./node.jpeg)

**节点发现**

Elasticsearch 默认使用单播发现，以防止节点无意中加入集群。
不建议使用组播可能会连接到错误的节点。

### 监控与诊断

#### 监控
借助一些可视化管理工具，可以很直观的看到集群各方面的状态。目前 ES 可用的监控工具或插件很多，主要为以下几种：

1. X-Pack+kibana 提供索引信息、集群整体信息，及各索引的索引、搜索速率、索引延迟数据等；其中，X-Pack 是官方给出的插件，各方面完善，推荐使用，缺点是只能试用一个月；
2. ElasticHQ 支持监控，并提供了折线图，方便查看变化，支持实时搜索，支持索引列表信息查看等功能；
3. cerebro 支持查看分片分配、索引操作、监控等功能；
4. elasticsearch-head 谷歌浏览器插件，支持查看分片分配、索引操作等功能，简单方便，数据量大界面会有卡顿，显示不友好；
5. _cat es 自带，可以在命令行使用，提供一些常用的监控项输出；

#### 诊断
1. 集群健康度
    - Status：状态群集的状态。红色：部分主分片未分配。黄色：部分副本分片未分配。绿色：所有分片分配 ok。
    - Unassigned Shards。未分配的分片。尚未创建或分配副本的分片计数。

2. 查询性能
    - 查询队列长度（队列里堆积的查询任务太多）
3. 索引性能
    - 刷新数，可以配置 `refresh_interval`
4. JVM 健康度
    - 堆使用情况
    - gc 情况
5. 系统健康度
    - 内存
    - 磁盘
    - cpu

### 补充
#### percolate 查询
`percolate` 查询可以用匹配存储于索引中的查询。`percolate`查询本身包含文档，用来作为查询去匹配存储的「查询」。具体文档见：
https://www.elastic.co/guide/en/elasticsearch/reference/6.3/query-dsl-percolate-query.html

### 参考
1. Elasticsearch: 权威指南（2.x 版本）https://www.elastic.co/guide/cn/elasticsearch/guide/current/foreword_id.html
2. 官方 6.3 版本文档 https://www.elastic.co/guide/en/elasticsearch/reference/6.3/getting-started.html
3. 中文社区 https://elasticsearch.cn/
4. ik 分词插件地址 https://github.com/medcl/elasticsearch-analysis-ik/releases/tag/v6.3.2
5. 分析 GC 文件 https://gceasy.io

