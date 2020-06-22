package grpc


// https://www.cnblogs.com/sunsky303/p/11119300.html


grpc:
	基于 http2 标准设计:
		Magic 帧:
			建立 Http2连接 请求的前言


//gRPC 基于 HTTP/2 + Protobuf。
//gRPC 有四种调用方式，分别是一元、服务端/客户端流式、双向流式。
//gRPC 的附加信息都会体现在 HEADERS 帧，数据在 DATA 帧上。
//Client 请求若使用 grpc.Dial 默认是异步建立连接，当时状态为 Connecting。
//Client 请求若需要同步则调用 WithBlock()，完成状态为 Ready。
//Server 监听是循环等待连接，若没有则休眠，最大休眠时间 1s；若接收到新请求则起一个新的 goroutine 去处理。
//grpc.ClientConn 不关闭连接，会导致 goroutine 和 Memory 等泄露。
//任何内/外调用如果不加超时控制，会出现泄漏和客户端不断重试。
//特定场景下，如果不对 grpc.ClientConn 加以调控，会影响调用。
//拦截器如果不用 go-grpc-middleware 链式处理，会覆盖。
//在选择 gRPC 的负载均衡模式时，需要谨慎。


// 注意:
//  grpc 会复用之前的连接, 所以做负载均衡的时候要考虑下




grpc-gateway:

gateway 接收 HTTP 请求
然后把请求再转换成 gRPC 的请求发给 gRPC Server

从 Server 收到响应之后
再将 gRPC 的响应转换为 HTTP 响应返回给客户端
