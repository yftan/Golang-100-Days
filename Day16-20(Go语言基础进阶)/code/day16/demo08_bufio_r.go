package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	/*
		bufio:高效io读写
			buffer缓存
			io：input/output

		将io包下的Reader，Write对象进行包装，带缓存的包装，提高读写的效率

			ReadBytes()
			ReadString()
			ReadLine()

	*/

	fileName := "/Users/tyf/Project/go/Golang-100-Days/bb.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	//bytes := make([]byte, 1024)
	//ret1, err1 := reader.Read(bytes)
	//if err1 != nil {
	//	log.Fatal(err1)
	//}
	//fmt.Println(ret1)
	//
LOOP:
	data, flag, err2 := reader.ReadLine()
	fmt.Println(flag)
	fmt.Println(err2)
	fmt.Println(data)
	fmt.Println(string(data))
	if err2 != io.EOF {
		goto LOOP
	}
	//创建Reader对象
	//b1 := bufio.NewReader(file)
	//1.Read()，高效读取
	//p := make([]byte,1024)
	//n1,err := b1.Read(p)
	//fmt.Println(n1)
	//fmt.Println(string(p[:n1]))

	//2.ReadLine()
	//flag,flag,err := b1.ReadLine()
	//fmt.Println(flag)
	//fmt.Println(err)
	//fmt.Println(data)
	//fmt.Println(string(data))

	//3.ReadString()
	// s1,err :=b1.ReadString('\n')
	// fmt.Println(err)
	// fmt.Println(s1)
	//
	// s1,err = b1.ReadString('\n')
	// fmt.Println(err)
	// fmt.Println(s1)
	//
	//s1,err = b1.ReadString('\n')
	//fmt.Println(err)
	//fmt.Println(s1)
	//
	//for{
	//	s1,err := b1.ReadString('\n')
	//	if err == io.EOF{
	//		fmt.Println("读取完毕。。")
	//		break
	//	}
	//	fmt.Println(s1)
	//}

	//4.ReadBytes()
	//data,err :=b1.ReadBytes('\n')
	//fmt.Println(err)
	//fmt.Println(string(data))

	//Scanner
	//s2 := ""
	//fmt.Scanln(&s2)
	//fmt.Println(s2)

	//b2 := bufio.NewReader(os.Stdin)
	//s2, _ := b2.ReadString('\n')
	//fmt.Println(s2)

}
