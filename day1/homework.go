package main

import "fmt"

func main() {
	//1.提示用户输入麻花藤，判断用户输入的对不对，如果对，提示真聪明，如果不对，提示不对，再试试
	// var username string

	// fmt.Print("请输入麻花藤：")
	// fmt.Scanln(&username)
	// if username == "麻花藤" {
	// 	fmt.Println("真聪明")
	// } else {
	// 	fmt.Println("不对，再试试")
	// }

	//2. 提示用户输入两个数字，计算两个数的和并输出。
	// var num1, num2 int
	// fmt.Print("请输入第一个数：")
	// fmt.Scanln(&num1)
	// fmt.Print("请输入第二个数：")
	// fmt.Scanln(&num2)

	// fmt.Println("两个数的和为：", num1+num2)

	//3. 提示用户输入姓名、位置、行为三个值，然后做字符串的拼接 并输出，例如：xx在xx做xx

	// var (
	// 	name     string
	// 	site     string
	// 	behavior string
	// )
	// fmt.Print("请输入姓名、位置、行为：")
	// fmt.Scanln(&name, &site, &behavior)
	// fmt.Printf("%v在%v做%v", name, site, behavior)

	//4.设定一个理想数字比如：66，让用户输入数字，比如比66大，则显示猜测的结果大了；
	//如果比66小，则显示猜测的结果小了;只有等于66，显示猜测结果正确

	// var num int
	// idealNum := 66
	// fmt.Print("请输入一个数字：")
	// fmt.Scanln(&num)
	// if num > idealNum {
	// 	fmt.Println("猜测的结果大了")
	// } else if num < idealNum {
	// 	fmt.Println("猜测的结果小了")
	// } else if num == idealNum {
	// 	fmt.Println("猜测结果正确")
	// }

	// 5. 写程序，输出成绩等级。成绩有ABCDE5个等级，与分数的对应关系如下
	// A 90-100  B 80-89  C 60-79  D 40-59  E 0-39
	// 要求用户输入0-100的数字后，正确输出他的对应成绩等级。

	var score int
	fmt.Print("请输入数字：")
	fmt.Scanln(&score)
	if score >= 90 && score <= 100 {
		fmt.Println("A")
	} else if score >= 80 && score <= 89 {
		fmt.Println("B")
	} else if score >= 60 && score <= 79 {
		fmt.Println("C")
	} else if score >= 40 && score <= 59 {
		fmt.Println("D")
	} else if score >= 0 && score <= 39 {
		fmt.Println("E")
	}
}
