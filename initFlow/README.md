
Go语言程序的初始化和执行总是从main.main函数开始的。但是如果main包导入了其它的包，则会按照顺序将它们包含进main包里（这里的导入顺序依赖具体实现，一般可能是以文件名或包路径名的字符串顺序导入）。如果某个包被多次导入的话，在执行的时候只会导入一次。当一个包被导入时，如果它还导入了其它的包，则先将其它的包包含进来，然后创建和初始化这个包的常量和变量,再调用包里的init函数，如果一个包有多个init函数的话，调用顺序未定义(实现可能是以文件名的顺序调用)，同一个文件内的多个init则是以出现的顺序依次调用（init不是普通函数，可以定义有多个，所以也不能被其它函数调用）。最后，当main包的所有包级常量、变量被创建和初始化完成，并且init函数被执行后，才会进入main.main函数，程序开始正常执行。下图是Go程序函数启动顺序的示意图：


![Init Flow](https://chai2010.cn/advanced-go-programming-book/images/ch1-11-init.ditaa.png)

*图 1-11 包初始化流程*

要注意的是，在`main.main`函数执行之前所有代码都运行在同一个goroutine，也就是程序的主系统线程中。因此，如果某个`init`函数内部用go关键字启动了新的goroutine的话，新的goroutine只有在进入`main.main`函数之后才可能被执行到。
