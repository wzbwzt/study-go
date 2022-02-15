package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/wzbwzt/studyGo/Excelize/instore"
)

func main() {
	instore.Init()
	return

	compile := regexp.MustCompile("\\d+\\.?\\d*|\\D*")
	stuffone := "20140923001 | 深沟球轴承 | 6004 | 6004 | 10"
	details := strings.Split(stuffone, "|")
	for k, v := range details {
		space := strings.TrimSpace(v)
		details[k] = space
	}
	fmt.Println(details)
	if len(details) == 4 {
		all := compile.FindAllStringSubmatch(details[3], 2)
		fmt.Println(all)
		//atoi, err := strconv.Atoi(all[0][0])
		float, err := strconv.ParseFloat(all[0][0], 32)
		if err != nil {
			println(err)
		}
		fmt.Println(int64(float))
		return
		if all[0] != nil {
			count, err := strconv.Atoi(all[0][0])
			if err != nil {
				panic(err)
			}
			fmt.Println(count)
		}

	}
	if len(details) == 5 {
		all := compile.FindAllStringSubmatch(details[4], 2)
		//count, err := strconv.ParseFloat(all[0][0], 32)
		count, err := strconv.Atoi(all[0][0])
		fmt.Println(all, len(all))
		if err != nil {
			all = compile.FindAllStringSubmatch(details[3], 2)
			//atoi, err := strconv.ParseFloat(all[0][0], 32)
			atoi, err := strconv.Atoi(all[0][0])
			if err != nil {
				println(err)
			}
			fmt.Println(atoi)
		} else {
			if len(all) > 1 {
				unit := all[1][0]
				fmt.Println(unit, count)
			}
			fmt.Println(123)

		}
	}
}
