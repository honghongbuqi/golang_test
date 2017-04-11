package main

//M P G
//
//goroutine 与 thread 的区别

//并发: 程序逻辑结构, 调用多线程可实现并发
//并行: 需多核或多机支持, 即使多核程序不是并发逻辑也不会并行

func add(x ...int) int {
	total := 0
	for _, num := range x {
		total += num
	}
	return total
}

func main() {
	x, y := 0x100, 0x200
	go add(x, y)
}
