# GC

![](https://tva1.sinaimg.cn/large/008eGmZEgy1gmrod3db2xj30ju0gwaal.jpg)

## 为什么
 消耗系统内容

## 手动垃圾回收
    问题: 
        - 悬挂指针
        - 内存溢出

## 三色抽象(标记清扫算法)
    问题:
        - 碎片化

## 分代回收
    原理:
        大部分对象在年轻时死亡

## 引用计数垃圾回收
    优点: 及时回收无用内存
    缺点:
        - 引用计数开销
        - 循环引用问题

## 增量式垃圾回收
![](https://tva1.sinaimg.cn/large/008eGmZEgy1gmrop6fne2j315i0l6dio.jpg)
    强三色不变性:(不容许白色写到黑色)
        插入写屏障
    弱三色不变性:(容许白色写到黑色, 但有灰色可达白色)

## 主体增量并发垃圾回收
图

## GO の GC

### 准备阶段
    为每个 `P` 创建一个 `mark worker` 协程, 并把该协程的指针存到 `P` 中, 开始是休眠状态, 等待调度执行


### 
![](https://tva1.sinaimg.cn/large/008eGmZEgy1gmrtt8jjiaj30t00mg0uz.jpg)
    gcphrase = _GCMARK -> _GCMARKTermination -> GCOff
    writeBarrier.enable = true
    gcBlackenEnable = true

### mheap


### 工作队列

wbuf1: 2为空, 且 1> 4 , 放到全局 
wbuf2:  
wbuf: 写屏障缓冲区

### CPU 使用率




### 混合写屏障

- 栈区:
STW全部扫描一遍, 标记为黑色
新建的元素为黑色

- 堆区:
被删除的标记为 灰色
被添加的标记为 灰色


### 触发方式
    - 手动触发
    - 分配内存
    - systom
  