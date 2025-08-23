package main

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {}

func (f *File) Read() {}

func ReadFile(reader Reader) {
	c := reader.(Closer)	// file은 Close 메소드를 가지고 있지 않기 때문에 Closer로 사용 불가능, 런타임 에러 발생
	c.Close()

	// 실제 사용 시에는 타입 변환 성공 여부 확인
	// if c, ok := reader.(Closer); ok {
	// 	...
	// }
}

func main() {
	file := &File{}
	ReadFile(file)	// file은 Read 메소드를 가지고 있기에 Reader로 사용 가능
}