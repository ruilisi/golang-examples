### 并发的选择

[在两个channel并发的条件下，通过switch选择当前channel中有值的那一个,或者输出无，执行9秒。](main.go)

[两个channel并发，在两个channel同时不存在的时候终止，不然输出其中有值的那一个。](remember_to_close_channel.go)

[每0.1给tick传入一个参数，使得tick的函数每过0。1秒执行一次，zoom等待10秒，10秒钟后switch选取到zoom终止掉程序]()
