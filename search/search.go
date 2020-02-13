package search

import (
	"fmt"
	"sync"
)

type hitokotov1 struct {
	Id			int			`json:"id"`
	Hitokoto	string		`json:"hitokoto"`
	Type		string		`json:"type"`
	From		string		`json:"from"`
	Fromwho		string		`json:"from_who"`
	Creator		string		`json:"creator"`
	Creatoruid	string		`json:"creator_uid"`
	Reviewer	string		`json:"reviewer"`
	UUID		string		`json:"uuid"`
	Createdat	string		`json:"created_at"`
}

// API
var URI string = "https://v1.hitokoto.cn"

var wg sync.WaitGroup

func Run() {

	var a int

	fmt.Println("输入一个数字，这个数字是获取一言的数量:")
	fmt.Scan(&a)

	// 存储响应数据的通道
	yiyan := make(chan hitokotov1)

	wg.Add(a)

	for b := 0; b < a; b++{
		// 狗肉听
		go func(){
			retrieve(URI, yiyan)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		//关闭通道，通知Dispaly退出
		close(yiyan)
	}()

	Display(yiyan)

	fmt.Println("over==========")

}
