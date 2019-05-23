Package math has  

- rand  
    usage
```
    rand.Intn(10)
```
- Sqrt  
usage
```
    math.Sqrt(7)
```
- Pi  
    usage
```
    math.Pi
```
Lưu ý : Cách khai báo biến : Loại biến đi sau tên biến  
    example:    

    
        func add(x int, y int) int {
        return x + y
        }

        var x int32
khi 2 biến cùng loại ta có thể để biến ở kiểu biến ở cuối

Các kiểu dữ liệu cơ bản 
```
    bool

    string

    int  int8  int16  int32  int64
    uint uint8 uint16 uint32 uint64 uintptr

    byte // alias for uint8

    rune // alias for int32
        // represents a Unicode code point

    float32 float64

    complex64 complex128
```
Các hằng số const thì không thể khai báo bằng cú pháp :=  
