自動ネットde勤怠
====

## Overview

ネットde勤怠の出社と退社を自動化するためのパッケージ.

## Description
CUIのコマンドから出社と退社を実施できる.

## Install
```bash:#
brew cask install phantomjs
go get github.com/yuki9431/autoNetDeKintai/component
```

## Requirement
- Go 1.10 or later
- github.com/sclevine/agouti

## Configuration
```go:main.go
package main

import (
	"github.com/yuki9431/autoNetDeKintai/component"
)

func main() {
	userId := "Anzu0902"
	password := "KirariSuki"

	userInfo := component.User{
		Id:       *userId,
		Password: *password,
	}

	// come in work
	component.Punch(userInfo, true)

	// leave in work
	component.Punch(userInfo, false)
}
```

## Contribution
1. Fork ([https://github.com/yuki9431/autoNetDeKintai](https://github.com/yuki9431/autoNetDeKintai))
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Create new Pull Request


## Author
[Dillen H. Tomida](https://twitter.com/t0mihir0)
