package linux内核

程序所处的状态:
	用户态或内核态

用户态到内核态:
	申请外部资源:
		1. 系统调用(man syscalls):(进程, 文件, 设备, 信息, 通信)
           读写文件 (open/write/read)
           申请内存 (malloc)
		2. 中断
		3. 异常

内核态到用户态:
	内核态执行完成, 切换到用户态
