package main

import (
	"autoNetDeKintai/component"
	"flag"
	"fmt"
	"log"
	"os"
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
	help := flag.Bool("h", false, "ヘルプを表示")

	// get userId and password from option
	userId := flag.String("usr", "", "enter your userId")
	password := flag.String("pwd", "", "enter your password")

	flag.Parse()

	if *help == true {
		fmt.Println(usage)
		os.Exit(0)
	}

	if *userId == "" || *password == "" {
		log.Fatalf("Error: enter userId and password")
	}

	userInfo := component.User{
		Id:       *userId,
		Password: *password,
	}

	component.Punch(userInfo, true)

}
