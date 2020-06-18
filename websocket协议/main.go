package websocket协议


websocket 是基于tcp, 不是http的

流式的socket

粘包的原因:
	没有对包进行一个分割
	 1. 协议头加长度
	 2. 每个包结束符, 分割符

websocket 是什么:

	别的缺点:
	1. http协议 头信息多, 信息利用率没有websocket多
	2. tcp协议, 需要包头比较麻烦

	1. 握手
	2. websocket协议头 (如何分包粘包)
	3. 明文 到 密文
	4. 关闭

