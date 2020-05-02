package engine

import "reptile/porsist"

//定义一个Request，里面含有Url，和它对应的解析器
type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult
}

//定义一个ParseResult，解释器的返回，返回一个Request，和解析出来的值
type ParseResult struct {
	Requests []Request
	Items [] porsist.Profile
}

//零时使用 返回一个空ParseResult
func NilParser([]byte)ParseResult  {
	return ParseResult{}
}