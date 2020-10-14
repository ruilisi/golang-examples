## channel

通过[go](./pass-ball.go)并发实现，channel中的数据决定channel中线程运行的顺序，在ch->ch1之后执行ch1->ch2,
最后执行ch2->ch3，所以最后的myball停留在ch3，详见字符串输出顺序

