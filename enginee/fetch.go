package enginee

import (
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const Qps = 5

var (
	rateLimiter = time.Tick(time.Second / Qps)
)

func Fetch(uri string) ([]byte, error) {
	log.Printf("start fetch with http:%s", uri)
	<-rateLimiter //限速
	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bodyReader := bufio.NewReader(res.Body)
	decodeReader := transform.NewReader(bodyReader, Determineencoding(bodyReader).NewDecoder())
	all, err := ioutil.ReadAll(decodeReader)
	if err != nil {
		return nil, err
	}
	return all, nil

}

func Determineencoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
