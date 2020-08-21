package linux操作

// 30天 找文件夹
find / -name "*.txt" -type d -mtime +30

// 当天 找文件
find / -name "*.txt" -type f -mtime -1 -size +4m -perm |xargs rm -rf {} \;

-exec 参数可以操作找出来的文件

// 查看nginx文件
grep -v "#" nginx.conf | grep -v "^$"


// 过滤go 项目中的一些其他文件
find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./tests/*" -not -path "./assets/*"

// awk 切割字符串, 取值


sed:

// sed 替换文件内容
sed -E 's/'