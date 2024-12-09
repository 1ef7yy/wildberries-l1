package main

func createHugeString(size int) string {
	return string(make([]byte, size))
}

var justString string

// Проблема состоит в том, что срез - ссылочный тип
// Поэтому в изначальном коде в justString мы хранили
// ссылку на срез, а не сам срез, что будет проблематично из-за GC, который не удалит v
// Из-за этого надо делать копию среза

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string(v[:100])
}
func main() {
	someFunc()
}
