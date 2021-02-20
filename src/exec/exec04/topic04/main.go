package main

import "fmt"

//根据用户输入显示对应的星期时间（string），如果星期一，显示干煸豆角，如果星期二，显示醋溜土豆，如果星期三，显示红烧狮子头，
//如果星期四，显示油炸花生米，如果星期五，显示蒜蓉扇贝，如果星期六， 显示东北乱炖，如果星期日，显示大盘鸡。
func main() {
	fmt.Println("请输入星期时间")
	var weekday string

	fmt.Scanf("%v",&weekday)
	switch weekday {
	case "星期一":
		fmt.Println("干煸豆角")
	case "星期二":
		fmt.Println("醋溜土豆")
	case "星期三":
		fmt.Println("红烧狮子头")
	case "星期四":
		fmt.Println("油炸花生米")
	case "星期五":
		fmt.Println("蒜蓉扇贝")
	case "星期六":
		fmt.Println("东北乱炖")
	case "星期日":
		fmt.Println("大盘鸡")
	default:
		fmt.Println("输入不合法")




	}

}