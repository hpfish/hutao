package geecache

// ByteView 缓存值的抽象与封装，提供一个只读数据结构
type ByteView struct {
	b []byte // 存储真实的缓存值
}

func (v ByteView) Len() int {
	return len(v.b)
}

func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
