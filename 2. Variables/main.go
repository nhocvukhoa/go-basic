package main

import "fmt"

func main() {
	// Vd1
	var number int
	number = 10
	fmt.Println(number)

	// Vd2
	var number2 int = 11
	fmt.Println(number2)

	// Vd3: nếu biến chưa được gán giá trị, go sẽ tự động khởi tạo giá trị zero tương ứng với kiểu của biến
	var age int
	fmt.Println("my age is", age) // 0

	// Vd4: Type inference => tự động suy luận kiểu vủa biến dựa trên giá trị khai báo
	var email = "admin@gmail.com"
	fmt.Println("email is", email)

	// Vd5: Khai báo nhiều biến với cùng kiểu dữ liệu
	var a, b int
	a = 1
	b = 2
	fmt.Println(a, b)

	var a1, b1 int = 3, 4
	fmt.Println(a1, b1)

	var a2, b2 = 4, 5
	fmt.Println(a2, b2)

	// Vd6: Khai báo nhiều biến khác kiểu dữ liệu
	var (
		name    string
		address string
	)

	name = "Khoa"
	address = "Hội An"

	fmt.Println(name, address)

	var (
		name1    string = "Khoa1"
		address1 string = "Hội An 1"
	)

	fmt.Println(name1, address1)

	var name2, address2 = "Khoa2", "Hội An2"
	fmt.Println(name2, address2)

	// Vd7: Short hand declaration => Khai báo ngắn gọn
	language := "Go"
	fmt.Println(language)
}
