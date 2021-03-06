package main

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"strings"
	
	"github.com/urfave/cli"
)

func main(){
	var rate string
	app := cli.NewApp()
	app.Name = "sameura"
	app.Usage = "get Sameura Dam status"
	app.Action = func(c *cli.Context) error {
		domain := "http://www1.river.go.jp"
		resp, _ := http.Get(domain + "/cgi-bin/DspDamData.exe?ID=1368080700010&KIND=3")
		defer resp.Body.Close()
		b1, _ := ioutil.ReadAll(resp.Body)
		b2 := GetObsPage(domain, string(b1))

		tr := "<TR>"
		trs := strings.Split(b2, tr)

		for i := 1; i < len(trs); i++ {
			td := "</TD>"
			tds := strings.Split(trs[i], td)
			tag := ">"
			rates := strings.Split(tds[len(tds)-2], tag)
		  if string(rates[1]) != "-" {
				rate = rates[2]
				break
			}
		}

		un := "<"
		value := strings.Split(rate, un)
		
		fmt.Printf("早明浦ダムの現在の貯水率は %s %%です.\n", string(value[0]))
		return nil
	}

	app.Run(os.Args)
}
