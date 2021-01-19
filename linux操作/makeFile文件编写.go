package linux操作

// makefile .PHONY用法
// 声明为伪目标后：在执行对应的命令时，make 就不会去检查是否存在 build / clean / tool / lint / help 其对应的文件，而是每次都会运行标签对应的命令
// 若不声明：恰好存在对应的文件，则 make 将会认为 xx 文件已存在，没有重新构建的必要了

// 赋值语句
// = 是最基本的赋值
//:= 是覆盖之前的值
//?= 是如果没有被赋值过就赋予等号后面的值
//+= 是添加等号后面的值

// @ 符号
// 后面会知道是注释还是代码

// go 编译设置git 版本
// s3port 为 包名
// CGO_ENABLED=0 GOOS= GOARCH= go build "-v" -ldflags '-X "s3porter/internal/utils.BuildTime=2020/07/23/02:43:46UTC" -X "s3porter/internal/utils.BuildVersion=75bc01e" -X "s3porter/internal/utils.BuildAppName=s3ctl"' -tags '""' -o s3ctl