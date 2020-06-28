package 信号

进程间通讯:
	1. 管道:
		ls | grep ma
		缺点: 不适合大量的数据

	2. 共享内存:
		不同进程虚拟内存映射相同的物理内存
		缺点: 没有触发机制

	3. 套接字:
		mysql(本机访问不用tcp)
		传输效率低

	4. 信号:
		kill -l
		ctrl + c 是发送的2的信号(sigint)
		不带参数的kill 是发送terminated信号(sigterm)
		9是killed 信号(sigkill)

	5. 消息队列:
		系统内核创建的队列, 不同程序可以读写
		缺点: send recv 两次copy 会消耗