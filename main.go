package main

import "fmt"

func main() {
	//-15:向右填充,填充满15个空格,包括字符在内
	//$4v : 向左填充,不满4 则用空格填充到4个
	fmt.Printf("%-15v $%4v\n", "SpaceX", "94")
	//文本右对齐
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}
