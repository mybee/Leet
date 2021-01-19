## Rust

- 特点:
    - 安全
    - 无需 GC
    - 性能好
  
- 擅长:
    - WebAssembly
    - 高性能

- 数据类型
    - bool
    - str
    - char  32位 
       可以是字符, 也可以是汉字
    - i8 i16 i32 u8 u32 u64 f32 f64
    - 自适应类型 isize usize
    - 数组
      - [u32; 5]
    - 元组
      - let (x, y, z) = tup // 元组拆解
   
-  函数
   - 蛇形命名
   - fn name(a: i32, b: i64) -> i64 {}
   - 不加分号, 为返回值
 
- 注释

- 控制流

- 所有权
  -   通过所有权管理内存
  - 堆和zhan
    - 编译的时候数据的类型和大小是固定的, 分配到栈上
    - 编译的时候数据的类型和大小是不固定的, 分配到堆上  string 类型
  - 作用域 
    - string 类型离开作用域 会调用 drop 方法
  - string 内存回收
  - clone
  - 移动
    - 所有权移动, move
    - 浅拷贝后, 原来的引用失效(引用类型)
  - 深拷贝
        - let s1 = s3.clone()  
  - 栈上数据拷贝
        
  - 函数和作用域
          
       

- copy trait
 - 整形
 - 浮点型
 - char
 - 上面类型的元组
 
- 引用
  -  让我们创建一个指向值的引用, 但是并不拥有它, 因为不拥有这个值, 所以, 当引用离开其值指向的作用域后, 也不会被丢弃
  - 创建一个指向地址的引用 ptr -> ptr -> value
  - 悬垂引用 野指针
  - 引用必须有效
  - 在任意给定时间, 有了可变引用之后不能再有不可变引用
  
  
- slice
  -  let h = &s[0..5]  或者 &[0..=4] 

- 结构体
  - 定义结构体
  - 创建结构体实例
  - 修改结构体字段
    - mut
  - 参数名字和字段名同名 简写
  - 从其他结构体创建实例
  - let user2 = User{
      name: String::from("user2")
      ..user1
      }
  - 元组结构体
     - 字段没有名字
     - 圆括号
     - Struct Point(i32, i32)
  - 没有任何字段的结构体
  - 打印结构体
    - 加 #[derive(Debug)]
    - {:#?}
           
- 方法
  - 构造函数
  - impl Dog {
      fn get_name(&self) -> &str {
          &(self.name[..])
      }
    }             
    
- 枚举


- 所有权: (https://www.bilibili.com/video/BV1hp4y1k7SV?p=16)
    - 原因:
      - 清理 Heap 上未使用的数据
      - 最小化 Heap 上的重复数据

    - Stack vs heap:
      - 所有存在 Stack 上的 Size 必须是固定的
      - 访问 Heap 比访问 Stack 慢
    
    - 规则:
      - scope 作用域 (drop 函数)
      - double free 问题
      - clone 会深拷贝
      - Copy trait
  
- 引用:
  - 把引用作为函数参数叫为 借用
  - 数据竞争:
    - 多个指针方位同一个数据
    - 至少有一个指针用于写数据
    - 没有使用任何机制来同步对数据的访问
  - 悬空引用
  
- Option
  - Null 的问题, billion mistakes
  - 
