## interface

`io.Writer`, `io.Reader` 都是比较常用的接口类型

### 接口优势

可以实现各种不同的数据类型的相互转化，实现各种 `io`, 之间的过渡

## `io.Reader/Writer` 几种常用的实现

'net.Conn, os.Stdin, os.File': 网络、标准输入输出、文件的流读取

'strings.Reader': 把字符串抽象成Reader

'bytes.Reader': 把[]byte抽象成Reader

'bufio.Reader/Writer': 抽象成带缓冲的流读取（比如按行读写）

## [wrap-interface.go](wrap-interface.go)

通过结构体包装接口，`io.Writer` 使用 `os.Stdout`，输出到控制台

## [replace-method-of-interfacewrap-interface.go](replace-method-of-interfacewrap-interface.go)

`Fatal()` 中可以存放多种任何接口，无论是什么类型都可以调用
