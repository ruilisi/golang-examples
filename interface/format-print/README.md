## interface

#### 运行方式

`go run interface.go`

定义一个输出流

```
type UpperWriter struct {
	io.Writer
}
```

自定义格式化输出

```
func (p *UpperWriter) Write(data []byte) (n int, err error) {
	return p.Writer.Write(bytes.ToUpper(data))
}
```

输出时使用自定义输出格式

```
fmt.Fprintln(&UpperWriter{os.Stdout}, "hello, world")
```

输出结果

```
fmt.Fprintln(&UpperWriter{os.Stdout}, "hello, world")

输出：HELLO, WORLD
```


