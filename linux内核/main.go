package linux内核


// cpu 总线  内存

// cpu通过总线从内存中拿指令, 执行下一条执行

// ABI 二进制指令接口

// 内核的工作:
//  -进程管理 进程调度 进程间通讯机制 -内存管理
//  中断异常处理 -文件系统 IO系统 -设备驱动 -网络部分

// 系统调用: 系统的提供的接口, 应用程序通过 它 调用 内核功能
//  open recv send write close

// 用户态:


// 内核态:


// 进程通讯:

// 中断机制:

// 时间系统:
	// 时钟中断
	// 通过脉冲信号
	// 相关硬件:
			// 实时时钟RTC 可编程间隔器 时间戳计数器
			// 高精度计数器 CPU 本地定时器 公安经服定时器


//  进程切换
//    进程上下文切换 :
//



// 红黑树 -- 在内核中的用途
//   红黑树(平衡二叉树的一种):
//     平衡黑色节点的高度
//
//   用途:(a. kv类型快速查找, b.根据红黑树的中序顺序, 查找一段)
//      epoll  a
//      socket管理
//      定时器  b
//      cfs(进程的公平调度)
//      内存管理 a
//      map   a

//   内核用处:
//		1. 进程管理:
//       每一个进程, task_struct, 写时复制
