package main

import (
	"os"
	"bufio"
	"strings"
	"time"
	"strconv"
	"fmt"
	"flag"
)

const cmdg = "GET"
const cmdpo = "POST"
const cmdpu = "PUT"
const cmdd = "DELETE"
const longForm = "02/01/2006/15:04:05/MST"

func main()  {
	datetimefilterflag := false
	var unixformatst int64
	var unixformatet int64
	user := flag.String("u","","username")
	starttimestr := flag.String("s","","start time like: 3/06/2018/12:47:00")
	logfilename := flag.String("f","","log file name")
	endtime := flag.String("e","","end time like: 3/06/2018/12:47:00")
	urlstringpart := flag.String("a","","address (URL) filter")
	ip := flag.String("i","","ip address")
	flag.Parse()
	if len(*starttimestr) > 0 && (len(*endtime) >0) {
		datetimefilterflag = true
		st,_ := time.Parse(longForm,*starttimestr)
		et,_ := time.Parse(longForm,*endtime)
		fmt.Print("Start date and time: "+st.String()+"\r\n")
		fmt.Print("End date and time: "+et.String()+"\r\n")
		unixformatst = st.Unix()
		unixformatet = et.Unix()
	}
	fmt.Print("Username: "+*user+"\r\n")
	fmt.Print("IP: "+*ip+"\r\n")
	fmt.Print("URL contain: "+*urlstringpart+"\r\n")
	fmt.Println("=========================================\r\n")

	fd,err := os.Open(*logfilename)
	if err != nil{
		fmt.Println("Can not open fiile!")
		os.Exit(1)
	}
	scan1 := bufio.NewScanner(fd)

	for scan1.Scan(){
		scl := scan1.Text()
		//fmt.Println(scl)
		allargs := strings.Split(scl," ")
		//fmt.Println(len(allargs))
		if scl[0] == '#'{
			//fmt.Println("Comments")
			continue
		}

		if allargs[5] == cmdg || allargs[5] == cmdpo || allargs[5] == cmdpu || allargs[5] == cmdd {
			if len(*user) > 0{
				if strings.Contains(allargs[7],*user) == false{
					continue
				}
			}
			if len(*ip) >1{
				if strings.Contains(allargs[2],*ip) == false{
					continue
				}
			}
			itsd := strings.Split(allargs[0], ".")
			its, err := strconv.Atoi(itsd[0])
			itsus, err := strconv.Atoi(itsd[1])
			if err != nil {
				fmt.Print("Invalid Unix date format")
				os.Exit(1)
			}
			uits := time.Unix(int64(its), int64(itsus))
			if err != nil {
				fmt.Print("Invalid Unix date format")
				os.Exit(1)
			}

			if datetimefilterflag{
				if uits.Unix() < unixformatst || uits.Unix() > unixformatet {
					continue
				}
			}

			if len(*urlstringpart) > 0{
				if strings.Contains(allargs[6],*urlstringpart) == false{
					continue
				}
			}
			for allargsi := range allargs {
				if allargsi == 0 {
					datetime := time.Unix(int64(its), int64(itsus))
					fmt.Println(datetime.String())
				} else if allargsi ==  2 || allargsi ==  5 || allargsi ==  6 || allargsi ==  7 || allargsi ==  10{
					fmt.Println(allargs[allargsi])
				}
			}
			fmt.Println("=========================================")
		}
	}
}
