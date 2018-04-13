package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//返回200 继续执行
	if resp.StatusCode != http.StatusOK {
		panic("error start code")
		return
	}
	//determineEncoding 获得读取网页的编码(checkset)
	e := determineEncoding(resp.Body)

	//转码
	utf8Reader := transform.NewReader(
		resp.Body, e.NewDecoder())
	//ioutil readAll 读去所有数据 返回[]byte数据
	all, err := ioutil.ReadAll(utf8Reader)

	if err != nil {
		panic(err)
	}
	printCityList(all)

}
func printCityList(content []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(content,-1)
	for _, v := range matches {
		fmt.Printf("city :%s --> URL :%s \n",v[2],v[1])
	}
	fmt.Printf("number : %d",len(matches))
}

//获得网页编码
func determineEncoding(r io.Reader) encoding.Encoding {
	//bufiio.NewReader() 相当于 bufio.NewReaderSize(rd,4096)
	// NewReaderSize 将 rd 封装成一个带缓存的 bufio.Reader 对象，
	// 缓存大小由 size 指定（如果小于 16 则会被设置为 16）。
	// 如果 rd 的基类型就是有足够缓存的 bufio.Reader 类型，则直接将
	// rd 转换为基类型返回。
	// Peek 返回缓存的一个切片，该切片引用缓存中前 n 个字节的数据，
	// 该操作不会将数据读出，只是引用，引用的数据在下一次读取操作之
	// 前是有效的。如果切片长度小于 n，则返回一个错误信息说明原因。
	// 如果 n 大于缓存的总大小，则返回 ErrBufferFull。
	// n不能大于4096
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	fmt.Println("------------------->", e)
	return e
}
