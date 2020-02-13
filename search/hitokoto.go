package search

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func retrieve(URI string, yiyan chan hitokotov1)  {

	if URI == "" {
		log.Fatalln("no nrl")
	}

	resp, err := http.Get(URI)
	if err != nil {
		log.Fatalln("nothing...")
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalln("error...")
	}

	var hitokoto hitokotov1
	err = json.NewDecoder(resp.Body).Decode(&hitokoto)

	yiyan <- hitokoto
}

func Display(yiyan chan hitokotov1) {

	// 遍历通道，这里可以自定义输出的数据
	for yy := range yiyan {
		fmt.Printf("\n类型：%s\nID:%d\nHitokoto:%s\n来自：%s\n提交者：%s\nUUID:%s\n===============", yy.Type, yy.Id, yy.Hitokoto, yy.From, yy.Creator,yy.UUID)
	}

}
