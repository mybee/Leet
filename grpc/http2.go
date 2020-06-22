package grpc

HTTP 1.0 协议缺点:
	1. 连接无法复用, tcp握手次数很多, 占用资源, 可以keep-alive避免
	2. 排队处理

HTTP 2.0 协议:
	1. 连接复用, 资源复用率搞
	2. 二进制分帧


