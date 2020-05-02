package parser

import (
	"log"
	"regexp"
	"reptile/engine"
	"reptile/porsist"
	"strings"
)
//获取诗人的正则
const poetRe = `<a href="(/authorv_[a-z 0-9]+.aspx)">([^<]+)</a>`

//这是获取诗人的解析器
//<a href="/shiren/libai/shiju/" title="李白的名言名句和诗句">李白(491句)</a>
func ParsePoetList(contents []byte) engine.ParseResult {
	//利用正则的方式来获取想要的数据
	re := regexp.MustCompile(poetRe)
	matches := re.FindAllSubmatch(contents, -1)
	if len(matches) <= 0 {
		panic("获取到零个诗人")
	}
	//创建一个ParseResult用来返回
	result := engine.ParseResult{}
	for _, val := range matches {
		//result.Items = append(result.Items, string(val[2])) //作者
		result.Requests = append(result.Requests, engine.Request{
			Url:        "https://so.gushiwen.org" + string(val[1]), //下一个页面的URL
			ParserFunc: ParseIntroList,                                 //下一个页面的URL所对应的解析器
		})
	}

	return result
}

const introRe = `<a href="(/authors/authorvsw_[a-z A-Z 0-9]+.aspx)">[^<]*</a></p>`
//const Author = `height:22px;"><b>([^<]+)</b></span>`
//这是诗人作品页面的解析器
//"<a href="/gushi/shi/7.html">将进酒</a>"
func ParseIntroList(contents []byte) engine.ParseResult {
	//利用正则的方式来获取想要的数据
	re := regexp.MustCompile(introRe)
	matches := re.FindAllSubmatch(contents, -1)

	//re := regexp.MustCompile(Author)
	//author := re.FindAllSubmatch(contents, -1)

	//创建一个ParseResult用来返回
	result := engine.ParseResult{}
	for _, val := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        "https://so.gushiwen.org" + string(val[1]), //下一个页面的URL
			ParserFunc: ParseProductList,                                 //下一个页面的URL所对应的解析器
		})
	}
	return result
}


var	productRe=regexp.MustCompile(`id="txtare[0-9]+">([^-]+)——([^·]+).([^《]+)《([^》]+)》`)
var	nextRe = regexp.MustCompile(`href="(/authors/authorvsw_[a-z A-Z 0-9]+.aspx)">下一页</a>`)



func ParseProductList(contents []byte) engine.ParseResult {
	//创建一个ParseResult用来返回
	result := engine.ParseResult{}
	index:= strings.Index(string(contents) ,"因服务器开支对本站造成巨大")
	if index != -1 {
		log.Printf("已经第十页了")
		return result
	}
	//利用正则的方式来获取想要的数据
	next:=nextRe.FindAllSubmatch(contents,-1)
	if len(next)>=1 && len(next[0])==2 {
		result.Requests=append(result.Requests,engine.Request{
			Url:        "https://so.gushiwen.org"+string(next[0][1]),
			ParserFunc: ParseProductList,
		})
	}

	product := productRe.FindAllSubmatch(contents, -1)

	//Items:=make([]model.Profile,0)
	for _, val := range product {
		profile:=porsist.Profile{
			Title:   string(val[4]),
			Dynasty: string(val[2]),
			Author:  string(val[3]),
			Content: string(val[1]),
		}
		//Items = append(Items, profile) //作者
		result.Items = append(result.Items, profile)
	}

	//result.Items=[]interface{}{Items}
	return result
}

