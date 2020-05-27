package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func fetch(url string, ch chan <- string){
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil{
		ch <- fmt.Sprint(err)
		return
	}
	nbs, _ := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	sec := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", sec, nbs, url)
}

func rotate(data []int) []int{
	i := 0
	j := len(data ) - 1

	for i < j{
		x := data[i]
		data[i] = data[j]
		data[j] = x

		i++
		j--
	}
	return data
}

func change(x int) *int {
	return &x
}

func helper(str []string) []string{
	i := 1
	for i < len(str) {
		if strings.EqualFold(str[i - 1], str[i]){
			str = append(str[:i - 1], str[i:]...)
		}
		i++
	}
	return str
}

func main(){
	data := []int{1,2,3,4,5}
	data = rotate(data)
	fmt.Println(data)

	str := []string{"qwert", "abc", "abc", "lyt", "lyy", "jfafj"}
	str = helper(str)
	fmt.Println(str)
	bytes := []byte{'f'}
	bytes = append(bytes[:])
}
