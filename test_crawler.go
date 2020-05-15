package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	//"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"regexp"
)

func main() {
	name:="雨伞"
	encodeName := url.QueryEscape(name)
	println(encodeName)
	//data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(name)), simplifiedchinese.GBK.NewEncoder()))
	//println(data)
	//return
	//返送请求获取返回结果
	head:="https://search.jd.com/Search?keyword=" + encodeName + "&enc=utf-8"

	//third:="&click_id="
	url:=head+encodeName
	log.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		panic(fmt.Errorf("Error: http Get, err is %v\n", err))
	}

	//关闭response body
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: statuscode is ", resp.StatusCode)
		return
	}

	//utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	utf8Reader := transform.NewReader(resp.Body, determinEncoding(resp.Body).NewDecoder())
	body, err := ioutil.ReadAll(utf8Reader)

	if err != nil {
		fmt.Println("Error read body, error is ", err)
	}
	log.Printf("%s",body)

	//printAllCityInfo(body)
	//打印返回值
	//fmt.Println("body is ", string(body))
}

func determinEncoding(r io.Reader) encoding.Encoding {

	//这里的r读取完得保证resp.Body还可读
	body, err := bufio.NewReader(r).Peek(1024)

	if err != nil {
		fmt.Println("Error: peek 1024 byte of body err is ", err)
	}

	//这里简化,不取是否确认
	e, _, _ := charset.DetermineEncoding(body, "")
	return e
}

func printAllCityInfo(body []byte){

	//href的值都类似http://www.zhenai.com/zhenghun/XX
	//XX可以是数字和小写字母,所以[0-9a-z],+表示至少有一个
	//[^>]*表示匹配不是>的其他字符任意次
	//[^<]+表示匹配不是<的其他字符至少一次
	compile := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

	/*matches := compile.FindAll(body, -1)

	//matches是二维数组[][]byte
	for _, m := range matches {
	  fmt.Printf("%s\n", m)
	}
	*/

	submatch := compile.FindAllSubmatch(body, -1)

	//submatch是三维数组[][][]byte
	/*  for _, matches := range submatch {

	   //[][]byte
	   for _, m := range matches {
		 fmt.Printf("%s ", m)
	   }

	   fmt.Println()
	 }*/

	for _, matches := range submatch {

		//打印
		fmt.Printf("City:%s URL:%s\n", matches[2], matches[1])

	}

	//可以看到匹配个数为470个
	fmt.Printf("Matches count: %d\n", len(submatch))

	//打印abc
	//fmt.Printf("%s\n", []byte{97,98,99})
}