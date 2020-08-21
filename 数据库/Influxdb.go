package 数据库

// influxdb

// 模型:
//   时间点
//   tags 属性
//   fields  值, 可以多个
//   时间戳

以时间做主键和索引
measurement 表


// 数据保留策略
//   最短一个小时
//   分片(shard)时长
//   定义在 database空间下
//   数据清理以分片为单位
//   分片中所有数据过期才会删除
//   过期数据处理:
//      丢弃
//      降采样



// schema设计
//   tagk/tagv简短
//   tag有索引 field没有
//   tag 是 字符串 field 支持 int, float等数据类型

// 查询优化