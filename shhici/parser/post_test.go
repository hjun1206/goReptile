package parser

import (
	"io/ioutil"
	"testing"
)

//测试
func TestParsePoetList(t *testing.T) {
	//contents, err := fetcher.Fetch("http://www.gushicimingju.com/shiju/zuozhe/")
	file, err := ioutil.ReadFile("./post_test_data.txt")
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(file))
	result := ParsePoetList(file)
	const resultSize  = 99
	if len(result.Requests) != resultSize{
		t.Errorf("result should have %drequests; but had %d",resultSize,len(result.Requests))
	}

	if len(result.Items) != resultSize{
		t.Errorf("Items should have %drequests; but had %d",resultSize,len(result.Items))
	}
}
