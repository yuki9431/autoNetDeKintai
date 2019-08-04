// main.go
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sclevine/agouti"
)

var usage = `Usage:
  netDeCommon [options] -usr [<userId>] -pwd [<password>]
  netDeCommon -come -usr UserId -pwd Password

Options:
  -h                  ヘルプを表示.
  -come               出社に打刻.
  -leave              退社に打刻.
  -usr <userId>       ユーザIDを入力できる.
  -pwd <password>     パスワードを入力できる.`

func main() {
	// help
	help := flag.Bool("h", true, "ヘルプを表示")
	if *help == true {
		fmt.Println(usage)
	}

	// get userId and password from option
	userId := flag.String("usr", "", "enter your userId")
	password := flag.String("pwd", "", "enter your password")

	// get punch type
	isCome := flag.Bool("come", false, "use when come to work")
	isLeave := flag.Bool("leave", false, "use when leave to work")
	flag.Parse()

	if (*isCome == true && *isLeave == true) || (*isCome == false && *isLeave == false) {
		log.Fatalf("choose \"come\" or \"leave\"")
	}

	if *userId == "" || *password == "" {
		log.Fatalf("enter userId and password")
	}

	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	if err := page.Navigate("https://www1.shalom-house.jp/komon/"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	// login
	page.ClearCookies()
	page.AllByID("txtID").Fill(*userId)
	page.AllByID("txtPsw").Fill(*password)
	page.AllByID("btnLogin").Click()

	page.AllByID("ctl00_ContentPlaceHolder1_imgBtnSyuugyou").Click()

	if *isCome == true {
		// come to work
		page.AllByID("ctl00_ContentPlaceHolder1_ibtnIn3").Click()
	} else if *isLeave == true {
		// leave to work
		page.AllByID("ctl00_ContentPlaceHolder1_imgOut3").Click()
	}

}
