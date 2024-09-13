package main

// thêm vào 2 module chúng ta cần
import (
    "fmt"
    "io/ioutil"
)

func main() {
    // Đọc nội dung trong localfile.data
    data, err := ioutil.ReadFile("localfile.data")
    // Nếu chương trình không thể đọc file
    // in ra nguyên nhân tại sao
    if err != nil {
        fmt.Println(err)
    }

    // Nếu đọc file thành công thì
    // in ra nội dung như một string
    fmt.Println(string(data))

}