package 锁

互斥同步锁(悲观锁):
	适合并发写入多的操作

乐观锁:
	适合并发写入少, 大多数是读的操作
	常用是CAS算法, 原子类
	使用: git的push


数据库:
	悲观锁: select xxx for update;
	乐观锁: 添加一个字段 version=1

