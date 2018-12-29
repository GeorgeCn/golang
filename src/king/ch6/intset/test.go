package main

type Buffer struct {
	buf    []byte
	initial    [64]byte
}

//Grow 方法按需扩展缓冲区的大小
//保证 n 个字节的空间
func (b *Buffer) Grow(n int) {
	if b.buf == nil {
		b.buf = b.initial[:0]// 最初使用预分配的空间
	}
	if len(b.buf)+n > cap(b.buf) {
		buf := make([]byte, b.Len(), 2*cap(b.buf)+n)
		copy(buf, b.buf)
		b.buf = buf

	}
}
