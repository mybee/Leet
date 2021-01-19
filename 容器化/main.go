package 容器化

docker只是辅助启动和管理容器

容器实现:
	Namespace:
		Linux内核用来隔离内核资源的方式
		1. 创建namespace
		2. 将进程加入到namespace, 只能看到namespace下的资源

		创建namespace:
			int clone (...)
		  		flag: NewNet NewPid
		加入namespace:
			setnx(namespace)

	Cgroup:
		约束进程使用的资源:
			//	[root@i1-dev ~]# mount -c cgroup
			//mount: cgroup is already mounted or /sys/fs/cgroup/perf_event busy
			//       cgroup is already mounted on /sys/fs/cgroup/systemd
			//       cgroup is already mounted on /sys/fs/cgroup/net_cls,net_prio
			//       cgroup is already mounted on /sys/fs/cgroup/cpuset
			//       cgroup is already mounted on /sys/fs/cgroup/freezer
			//       cgroup is already mounted on /sys/fs/cgroup/pids
			//       cgroup is already mounted on /sys/fs/cgroup/hugetlb
			//       cgroup is already mounted on /sys/fs/cgroup/blkio
			//       cgroup is already mounted on /sys/fs/cgroup/cpu,cpuacct
			//       cgroup is already mounted on /sys/fs/cgroup/memory
			//       cgroup is already mounted on /sys/fs/cgroup/devices
			//       cgroup is already mounted on /sys/fs/cgroup/perf_event
			进程cpu目录下 的 cpu.cfs_quots_us 文件的值改为10000, 则只会使用10%的cpu资源

	Rootfs: (容器文件分层设计)
		删除也是写入


	容器的文件系统:
		mount namespace 将文件 mount /a b (将a目录挂载为b)

	镜像文件分层: (原理: 联合文件系统)
		docker image inspect $image  (查看镜像分层)

总结:
	容器即进程

执行: entrypoint + cmd
      cmd可以覆盖 entrypoint不可以


