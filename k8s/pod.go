package k8s


// 容器技术原理

// 其实就是进程, 使用namespace隔离
// 进程可以创建子进程, 文件做namespace
// cgroup 做容器的cpu限制

// infra 容器 (共享网络, volume, namespace)

// namespace 是容器的那个namespace


// 容器设计模式

//   挎斗模式	一个读一个写日志

//   外交官模式   配置中心策略

//

//   选举模式     像etcd那样

//   工作队列模式

//   分散收集模式
