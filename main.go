package main

import (
	"fmt"
	"net"
	"os"

	"github.com/maxmind/mmdbwriter"
	"github.com/maxmind/mmdbwriter/inserter"
	"github.com/maxmind/mmdbwriter/mmdbtype"
)

func main() {
	//加载当前IP库
	writer, err := mmdbwriter.Load("country.mmdb", mmdbwriter.Options{})
	if err != nil {
		fmt.Println(err)
	}
        //解析IP段 自行修改
	_, subnet, err := net.ParseCIDR("46.xx.xx.0/24")
	if err != nil {
		fmt.Println(err)
	}
	//构建需要修改mmdbtype数据 自行修改
	subData := mmdbtype.Map{
		mmdbtype.String("country"):        mmdbtype.String("HK"),
		mmdbtype.String("continent"):      mmdbtype.String("AS"),
		mmdbtype.String("continent_name"): mmdbtype.String("Asia"),
		mmdbtype.String("country_name"):   mmdbtype.String("Hong Kong"),
	}
	//使用InsertFunc的inserter.ReplaceWith直接替换原有数据
	if err := writer.InsertFunc(subnet, inserter.ReplaceWith(subData)); err != nil {
		fmt.Println(err)
	}
	// 创建新的mmdb库文件
	file, err := os.Create("country-new.mmdb")
	if err != nil {
		panic(err)
	}
	// 用WriteTo方法写回到新的mmdb文件中
	_, err = writer.WriteTo(file)
	if err != nil {
		fmt.Println(err)
	}

}

