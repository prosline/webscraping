package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"io/ioutil"
	"github.com/prosline/webscraping/link"
)

var exampleHtml=`
<html>
<body>
    <h1>Hello Marcio</h1>
    <a ref="/root-page1">Link to Brazil</a ref>
    <a ref="/root-page2">Link to Spain</a ref>
    <a ref="/root-page3">Link to USA</a ref>
</body>
</html>
`
	//Example to run main.go  https://google.com
func main() {
	url := os.Args[1]	
	fmt.Println("URL parameter ->",url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	rs, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		panic(er)
	}
	r := strings.NewReader(string(rs))
	links, err := link.Parse(r)
	if err != nil{
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
