# Array 
### Khai báo mảng 1 chiều
#### Khai báo biết trước độ dài mảng 
```
var s [5]int {2, 3, 5, 7, 11}
```
#### Khai báo chưa biết trước độ dài mảng
```
var s []int
```
### Khai báo mảng 2 chiều 
khai cáo cấu trúc
``` 
cấu trúc <ten_mang> := [][]<kieu_du_lieu>  
{  
    []<kieu_du_lieu{data}  
}
```

```
board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
```

### append
Để thêm dữ liệu vào mảng ta sử dụng hàm append
#### Cấu trúc
```
append(<Mảng dữ liệu>, data)
```
### Chỉ mục 
Bạn có thể bỏ qua chỉ mục hoặc giá trị bằng cách gán cho nó  `_`

```
for i, _ := range pow
for _, value := range pow
```

```
for i := range pow
```