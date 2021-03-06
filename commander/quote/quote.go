package quote

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/chinglinwen/wechat-commander/commander"
	"gopkg.in/resty.v1"
)

type QuoteResult []struct {
	Quote    string `json:"quote"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

var (
	url = "https://andruxnet-random-famous-quotes.p.mashape.com/?cat=movies&count=10"
	key = "dKHmQTWJM0mshA6NKKPhOt9zboacp1AABODjsnTbTPMv5KhafO"
)

func GetQuote() (QuoteResult, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("X-Mashape-Key", key).
		SetHeader("Accept", "application/json").
		Post(url)
	if err != nil {
		return nil, err
	}
	var q QuoteResult
	err = json.Unmarshal(resp.Body(), &q)
	return q, nil

}

type Quote struct {
	quotes QuoteResult
	i      int
}

func (b *Quote) Command(cmd string) (data string, err error) {
	//log.Printf("got cmd %v from quote", cmd)
	if b.i == 0 || b.i > 10 || len(b.quotes) == 0 {
		//get another 10 quotes
		log.Printf("start requesting quote...")
		b.quotes, err = GetQuote()
		b.i = 1
	}
	log.Printf("got %v quotes, repsonse with %v", len(b.quotes), b.i)
	data = fmt.Sprintf("%v\n  --%v\n", b.quotes[b.i].Quote, b.quotes[b.i].Author)
	b.i += 1
	return
}

func init() {
	commander.Register("quote", &Quote{})
}
