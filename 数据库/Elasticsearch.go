package 数据库

// 开源高扩展的分布式全文搜索引擎, 可以近乎实时的存储, 检索数据

// 7.x

//  功能: 全文搜索 结构化搜索 搜索纠错

//  es 和 solr 基于 lucene

//  es 性能不受数据量影响, solr 会变慢

// 核心概念:
//  1. 面向文档, 一切都是json
//  2. 一个索引类型, 包含多个文档.
//  3. 最小的单位是文档
//  4.

//  类型:

//  索引:
//		就是数据库

     //倒排索引:
     // 完全过滤掉无关的数据, 提高效率


// IK 分词器: 使用中文的话用这个
//   ik_smart 和 ik_max_word
//   ik_smart:
//
//   ik_max_word:
//
///  可以自己配置分词

// 查询:
//  参数体:
//    {
//		query: match 匹配词
//		sort: 排序
//		from: 从哪个数据开始
//      size: 返回多少数据
//    }