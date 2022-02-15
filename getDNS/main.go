package main

import (
	"fmt"
	"github.com/miekg/dns"
	"time"
)

func CNAME2(src string, dnsService string) (dst []string, err error) {
	c := dns.Client{
		Timeout: 5 * time.Second,
	}

	var lastErr error
	// retry 3 times
	for i := 0; i < 3; i++ {
		m := dns.Msg{}
		// 最终都会指向一个ip 也就是typeA, 这样就可以返回所有层的cname.
		m.SetQuestion(src+".", dns.TypeA)
		r, _, err := c.Exchange(&m, dnsService+":53")
		if err != nil {
			lastErr = err
			time.Sleep(1 * time.Second * time.Duration(i+1))
			continue
		}

		dst = []string{}
		for _, ans := range r.Answer {
			record, isType := ans.(*dns.CNAME)
			if isType {
				dst = append(dst, record.Target)
			}
		}
		lastErr = nil
		break
	}

	err = lastErr

	return
}

func main(){
	cname2, _ := CNAME2("www.baidu.com", "8.8.8.8")
	fmt.Println(cname2)
}

