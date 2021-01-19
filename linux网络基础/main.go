package linux网络基础

TCP / IP 四层模型
应用层:     应用层
内核层:     传输层
           网络层
           链路层


IP协议:
	版本 IPv4 IPv6
	TTL time to live  在路由器中跳的次数, 每次减一, 到0抛弃
	源IP: 32位 4个字节

UDP:
	16位 源端口号
	16位 目的端口号

网络套接字:
	套接字, 在通讯过程中一定是成对出现的.
	一个文件描述符指向一个套接字, 建两个缓冲区

网络字节序:
	
	小端法 (计算机用的): 高位存高地址
	大端法 (网络用的): 高位存低地址

socket 通讯流程分析:

	server:
		socket() 创建socket
		bind()   绑定服务器地址结构
		listen() 设置监听
		accept() 阻塞监听客户端连接
		read()   读socket获取的内容
		write()  写
		close()	 关闭连接

	client:
		socket()	创建socket
		connect()	与服务端建议连接
		write() 	写数据到socket
		read()		读数据到socket
		close()		关闭



三次握手:



四次挥手:
	client: FIN 标志位
	server: ACK 标志位
	client: 半关闭, 关闭的是内部的缓冲区
	server: FIN 标志位
	client: ACK 标志位

		举例子: 打电话 挂电话 的 例子


滑动窗口:
	tcp 流量控制, 防止数据丢失
	client 和 server 沟通缓冲空间剩余多少


TCP 状态:

	client                  server

	CLOSED--   				CLOSE
			--
    SYN_SEND   -- syn
					  ------

	SYN_SEND     syn ack--

					ack		SYN_RCVD

	ESTABLISHED				ESTABLISHED

	主动                     被动

	FIN_WAIT_1		FIN

					ACK
	FIN_WAIT_2(半关闭)		CLOSE_WAIT

					FIN
	FIN_WAIT_2
							LAST_ACK
	TIME_WAIT		ACK

	2msl时长 40s

	CLOSE


	TIME_AWIT 以及 2MSL 只有主动关闭方会经历该状态

2MSL 时长:
	保证最后一个ACK能够被对端接收, 等待期间, 如果对端没收到,
	会给我发FIN, 我再给他发SCK

半关闭:
	通信双方中, 只有一端关闭通讯
	shutdown 在关闭多个文件描述符应用的文件时, 采用全关闭方法, close 只关闭一个


Select:

	内核的select

多路IO转接:

	select:
	poll:
	epoll:
				事件模型:

			     	ET: 边缘触发 (只有往缓冲区写数据才会触发) 剩余数据不会导致epoll_wait触发

				    LT: 水平触发(默认) (将缓冲区中的数据一下都取出来)  剩余数据会导致epoll_wait触发

