package 一致性算法

NWR理论:


cap理论:
	分区容错性 可用性 一致性


一致性模型:
	弱一致性:
		最终一致性:
			DNS
			Gossip(Cassandra的通讯协议)
	强一致性:
		同步
		Paxos
		Raft(multi-paxos)
		ZAB(multi-paxos)


主从同步:
	1. Master接受写请求
	2. Master复制日志到从节库
	3. Master等待, 直到所有从库返回结果
	问题:
		可用性很低

多数派:


Paxos:
	client:
		请求发起者, 用户

	Proposer:
		像集群提出议案

	Acceptor:
		投票者

	Learner:


Paxos流程:

	1. 客户端请求
	2. 提案
	3. 议员通过
	4. 告诉议员数据
	5. 议员写数据, 并rpc到leaner 和 提议案的
	6. leaner 返回成功给客户端


Basic Paxos问题:
	难实现 效率低(两轮RPC) 活锁


Multi Paxos:
	减少提案的rpc过程, 使用一个proposer, 避免活锁

	减少角色, 进一步简化:


Raft:
	划分为三个子问题:
		Leader选举
		日志复制
		安全
	角色:
		leader
		候选者
		追随者


平票问题:
	两个候选者获得平票是, 开始随机等候, 吸纳到先得
分区问题:
	单数节点
	多数的节点集群可以写入数据, 少数的只写log
	恢复后, 追随新任期的leader


问题:
	访问时怎么访问主节点的 -> 路由

Gossip 协议:
	基于gossip协议的系统:
		Cassandra, Redis(cluster模式), Consol

