package main

import (
	"fmt"
	"strings"
)

func main(){
	str := "http://192.168.1.1:8080"
	str2 := "http://192.168.1.1/wecd"
	str3 := "http://192.168.1.1/"
	str4 := "http://192.168.1.1/erreg/tth"
	fmt.Println(strings.Split(str, "/")[2])
	fmt.Println(strings.Split(str2, "/")[2])
	fmt.Println(strings.Split(str3, "/")[2])
	fmt.Println(strings.Split(str4, "/")[2])
	//fmt.Println(strings.Split(str, "/")[3])
	fmt.Println(strings.Split(str2, "/")[3])
	fmt.Println(strings.Split(str3, "/")[3])
	fmt.Println(strings.Split(str4, "/")[3:])
	fmt.Println(strings.Join(strings.Split(str4, "/")[3:], "/"))

	fmt.Println(len(strings.Split(str, "/")))
	fmt.Println(len(strings.Split(str2, "/")))
	fmt.Println(len(strings.Split(str3, "/")))
	fmt.Println(len(strings.Split(str4, "/")))


}
