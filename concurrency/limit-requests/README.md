### requests限流

[通过在每个channel中的添加`time.Sleep`,来使得一些线程进入睡眠状态，从而达到控制requests执行的间隔，限制requests的速度](main.go)