package main

// 获得城市空气信息

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type RespJson struct {
	SKInfo struct {
		// 城市名
		CityName string `json:"city"`
		// 城市ID
		CityID string `json:"cityid"`
		// 温度
		Temp string `json:"temp"`
		// 风向
		WD string `json:"WD"`
		// 风级
		WS string `json:"WS"`
		// 湿度
		SD string `json:"SD"`
		// 压强
		AP string `json:"AP"`
	} `json:"weatherinfo"`
}

func main() {
	fmt.Println("请输入城市代码, 比如北京为101010100:")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	if sc.Err() != nil {
		log.Fatal(sc.Err())
	}

	url := fmt.Sprintf("http://www.weather.com.cn/data/sk/%v.html", strings.Trim(sc.Text(), " \t"))

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	v := RespJson{}
	if json.Unmarshal(b, &v) != nil {
		log.Fatal(err)
	}

	fmt.Printf("城市: %v\t城市ID: %v\t风向: %v\t风级: %v\t湿度: %v\t压强: %v\n", v.SKInfo.CityName, v.SKInfo.CityID, v.SKInfo.WD, v.SKInfo.WS, v.SKInfo.SD, v.SKInfo.AP)
}
