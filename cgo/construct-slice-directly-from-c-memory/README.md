# 获取c数组中的元素

- 一个` unsafe.Pointer `可以转化成任何类型的指针
- 一个 `uintptr` 可以转化成一个 `unsafe.Pointer`
- 一个 `unsafe.Pointer` 可以转化成一个 `uintptr`

通过` unsafe.Pointer `指向c数组中的内存地址，`reflect.SliceHeader` 提取出数组中的值

