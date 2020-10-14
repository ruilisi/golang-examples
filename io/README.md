## io

[copy.go](copy.go) 将输入流和输出流进行同步

[multi-reader.go](multi-reader.go) io.MultiReader读取一个io.Reader数组

[multi-reader.go](multi-writer.go) io.MultiReader同时向文件和stdout写入

[pipe.go](pipe.go) io.Pipe中包含一个write和一个read，在两方中有一方有值的时候，另一方才能工作

[relay.go](relay.go) 双向通信

[tee-reader.go](tee-reader.go) 分段读取，借助 `io.teeReader(r, w)` 中的w作为过渡，从输入流读取写到另一输出流中



