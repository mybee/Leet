package linux操作

// 30天 找文件夹
find / -name "*.txt" -type d -mtime +30

// 当天 找文件
find / -name "*.txt" -type f -mtime -1 -size +4m -perm |xargs rm -rf {} \;

-exec 参数可以操作找出来的文件

// 查看nginx文件
grep -v "#" nginx.conf | grep -v "^$"


// awk 切割字符串, 取值


// sed 替换文件内容
