package 并发模式


CPU:
	南桥:

	北桥:

	cache:


命令查看CPU:

root@VM-0-8-ubuntu:~# vmstat
procs -----------memory---------- ---swap-- -----io---- -system-- ------cpu-----
r  b   swpd   free   buff  cache   si   so    bi    bo   in   cs us sy id wa st
1  0      0  92976  43568 285596    0    0    30    16    1    0  1  1 99  0  0
root@VM-0-8-ubuntu:~# uptime
16:25:54 up 375 days, 23:28,  1 user,  load average: 0.01, 0.01, 0.00
root@VM-0-8-ubuntu:~# cat /proc/cpuinfo
processor       : 0
vendor_id       : AuthenticAMD
cpu family      : 23
model           : 1
model name      : AMD EPYC Processor
stepping        : 2
microcode       : 0x1000065
cpu MHz         : 1999.995
cache size      : 512 KB
physical id     : 0
siblings        : 1
core id         : 0
cpu cores       : 1
apicid          : 0
initial apicid  : 0
fpu             : yes
fpu_exception   : yes
cpuid level     : 13
wp              : yes
flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm rep_good nopl extd_apicid eagerfpu pni pclmulqdq ssse3 fma cx16 sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand hypervisor lahf_lm cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw vmmcall fsgsbase bmi1 avx2 smep bmi2 rdseed adx smap clflushopt sha_ni xsaveopt xsavec xgetbv1 arat
bugs            : fxsave_leak sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass
bogomips        : 3999.99
TLB size        : 1024 4K pages
clflush size    : 64
cache_alignment : 64
address sizes   : 48 bits physical, 48 bits virtual
power management:






