package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//这里做一个延迟
var rateLimiter = time.Tick(10*time.Microsecond)
//这是获取网页数据
func Fetch(url string)([]byte,error)  {
	<-rateLimiter
	resp, err := http.Get(url)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()

	//检查下网页返回是否正常
	if resp.StatusCode != http.StatusOK{
		return nil,fmt.Errorf("wrong status code: %d",resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
	//fmt.Printf(string(dataByte))
}