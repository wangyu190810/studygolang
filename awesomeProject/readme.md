### gin 框架

[文档信息](https://gin-gonic.com/zh-cn/docs/#gin-v1-%E7%A8%B3%E5%AE%9A%E7%89%88) 

[本程序参考文章](https://www.jianshu.com/p/a3f63b5da74c)

程序编译注意事项：

    此时运行项目，不能像之前简单的使用go run main.go，
    因为包main包含main.go和router.go的文件，
    因此需要运行go run *.go命令编译运行。
    如果是最终编译二进制项目，则运行go build -o app
