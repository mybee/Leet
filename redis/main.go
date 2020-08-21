package redis

//
//

有哪些数据类型 使用场景:


	GEO: 计算距离
	bitmap: 可以做布隆过滤器
	hyperloglog: 统计多少个不同的值, 所需的空间很少, 只要12k

基本使用:
	计数:
		incr key1
		desr key1
	list: 消息列表 消息队列
		lpush list
		lrange list 0 -1
		rpush list 1 2 4 6
		lpop list 22
	hash: 用户属性
		hset key field value
		hkeys key
		hvals key
		hget key field
	set:  唯一性 交集 共同好友
		sadd set 1 3 4 5 5
		smembers set
		sdiff seta setb  // 取seta中特有的
		sdiff setb seta  // 取setb中特有的
		sinter seta setb // 交集
		sunion seta setb  // 并集

	zset: 排行榜 带权重的消息
		zadd zset 1 a 2 b 4 c
 		zrange zset 0 -1
		zrevrange zset 0 -1
		zrem zset a
		zrange zset 0 -1 withscores




设置key的过期时间:
	expire key second

查看key的过期时间:
	ttl key

清除key的过期时间:
	persist key

redis的 pub/sub 有什么优缺点:
	消费者挂掉时, 生产的消息会丢失, 所以要用专业的消息队列

如果大量的key设置同一时间过期, 需要注意什么:
	到过期的时间点, redis 会出现短暂的卡顿现象

如何生产一次, 消费多次:
	使用pub/sub 主题订阅模式, 实现 1:N的消息队列

管道有什么好处 为啥用pipeline:
	可以将多次io往返的时间变为一次, 前提是没有因果相关性的指令

redis分布式锁实现:
	setnx争抢锁, 抢到后, 加过期时间 -> set 和expire 合成一个指令以防中间down机

keys 会导致线程阻塞一段时间..
	scan可以无阻塞取出keys, 但是有重复, 需要去重, 时间比keys至指令长
	scan 游标 match xxx  count 100   // 游标从0开始, 下一次用上一次获取的
redis队列:
	blpop可以阻塞知道消息到来

redis持久化:
	bgsave: 镜像全量持久化
	aof: 增量持久化

bgsave:
	fork: 创建一个相同的子进程
	cow: copy on write 父进程共享数据段给子进程, 父进程继续提供读写,

redis同步机制:
	可以主从同步, 从从同步.
	第一次: 主节点做一次bgsave, 同时操作记录加到内存中, rdb文件全量到复制节点
	复制节点将rdb加载到内存中, 通知主节点同步操作记录, 然后复制节点重放


读写分离

容灾恢复
	配从不配主

	一主二仆


集群:

哨兵机制(高可用):
	当主节点挂掉, 用来选主, 不过有几秒不可用, 所以有cluster模式

cluster模式(扩展性):
	数据分片, 高可用

BASE:
	基本可用 (响应时间的损失, 功能上的损失)
	软状态  存在中间状态
	最终一致性

CAP原理: (布鲁尔定理)
	传统数据库: ACID 原子性 一致性 独立性  持久性
	noSql的: CAP  C: 强一致性 A: 高可用性 P: 分区容错性
	CAP 只能 三选二:

		CA: RDBMS
		AP:
		CP:


持久化:
	RDB:
		fork一个与当前一模一样的子进程来持久化,(为啥: 因为是单线程)
		进程的所有的数据跟原进程一样, 会将数据先写道一个临时文件中
	AOF:


缓存的四大问题:
	1. 缓存穿透:
		数据库 和 缓存中 都没有的数据
		解决: 缓存空对象, 布隆过滤器

	2. 缓存击穿:
		数据库中有, 缓存没有, 并发导致
		解决: 分布式锁, 预热数据

	3. 缓存雪崩:
		大量数据失效
		解决: 搭建高可用集群
             失效时间错开

	4. 缓存与数据库一致性:

		解决: 延时双删 串行化


问题:
	1. 怎么快速删除10w个key?
		管道
	2. redis单线程为什么快?


io多路复用:
	epoll select
	拿到数据给redis
	通俗点: 老师检查课堂作业, 做完的举手

NIO 设计:
	accept() 方法 设置成非阻塞

select:


	hash算法 -> 一致性hash算法 -> hash slot算法

	老hash:
		hash取模算法
	一致性hash算法:
		hash环, hash后不取模, 而是在hash环上找, 顺时针最近一个取数据

redis集群原理:
	cluster:
		hash环, 根据key计算crc16的值, 然后对16484取模, 拿到对应的槽位

		虚拟节点:
			给每个节点做虚拟节点, 均匀分布到环上, 解决数据偏移问题


		gossip:
			每个节点维护集群的元数据
			优点: 元数据分布, 压力小
			缺点: 元数据更新有延时

			包含多种消息: meet ping pong fail 等
			meet: 某个节点发送meet消息给新加入的节点, 让新节点加入集群中

			redis-trib add-node xxx 时 内部发meet 给新节点

			ping:

				?每个节点每秒都会执行10 次ping, 包含5个最久没有通讯的其他节点
				如果发现某个节点延时达到了cluster_node_timeout / 2 ,
				那么立即发送ping, 避免数据延时交换过长
				ping的信息:
					自己节点的信息, 和 1/10其他节点的信息, 发送出去, 进行数据交换
					至少包含三个其他节点的信息, 最多总节点减2的信息



			客户端:
				本地维护 slot -> node 的表
				当返回moved, 更新 slot -> node 表


			高可用原理:
				1. 判断节点down机:
				如果一个节点认为另一个节点宕机(根据超时时间), 那么久pfail, 主观宕机
				如果多个节点都认为另一个节点宕机, 那么久fail, 跟哨兵原理几乎一样, sdown odown
				2. 从节点过滤:

				3. 从节点选举:

				4. 与哨兵比较:
