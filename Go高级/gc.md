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

### 
![](https://tva1.sinaimg.cn/large/008eGmZEgy1gmrtt8jjiaj30t00mg0uz.jpg)
    gcphrase = _GCMARK -> _GCMARKTermination -> GCOff
    writeBarrier.enable = true
    gcBlackenEnable = true

### mheap


### 触发方式
    - 手动触发
    - 分配内存
    - systom
  