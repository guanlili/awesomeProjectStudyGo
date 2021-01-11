package main


import "fmt"

func findTheDifference(s, t string) byte {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i := 0; ; i++ {
		ch := t[i]
		cnt[ch-'a']--
		if cnt[ch-'a'] < 0 {
			return ch
		}
	}
}


func main() {
	fmt.Println("Hello, World!")

	s:= "abcd"
	t:= "abcde"
	ret:= string(findTheDifference(s, t))
	fmt.Printf( "不同值 : %v\n", ret )

}


