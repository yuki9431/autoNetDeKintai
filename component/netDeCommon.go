package netDeKomon

import (
	"fmt"

	"github.com/sclevine/agouti"
)

type user struct {
	Id       string
	password string
}

type webInfo struct {
	user user
	page *agouti.Page
}

// come to work
func Punch(userId, password string, isCome bool) (err error) {
	web, err := new(userId, password)
	if err != nil {
		return
	}

	if err := web.login(); err != nil {
		return fmt.Errorf("Error: Faild to login %v", err)
	}

	if err := web.punch(isCome); err != nil {
		return fmt.Errorf("Error: Faild to punch %v", err)
	}

	if err := web.logout(); err != nil {
		return fmt.Errorf("Error: Faild to logout %v", err)
	}

	return nil
}

func TestLogin(userId, password string) (err error) {
	web, err := new(userId, password)
	if err != nil {
		return
	}

	if err := web.login(); err != nil {
		return fmt.Errorf("Error: Faild to login %v", err)
	}

	if err := web.logout(); err != nil {
		return fmt.Errorf("Error: Faild to logout %v", err)
	}

	return nil
}

func new(userId, password string) (web *webInfo, err error) {
	// start driver
	driver := agouti.ChromeDriver()
	if err = driver.Start(); err != nil {
		return nil, fmt.Errorf("Error: Failed to start driver:%v", err.Error())
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		return nil, fmt.Errorf("Error: Failed to open page:%v", err.Error())
	}

	if err := page.Navigate("https://www1.shalom-house.jp/komon/"); err != nil {
		return nil, fmt.Errorf("Error: Failed to navigate:%v", err.Error())
	}

	return &webInfo{
		user: user{
			Id:       userId,
			password: password,
		},
		page: page,
	}, nil
}

func (web *webInfo) login() (err error) {
	err = web.page.ClearCookies()
	if err != nil {
		return
	}

	err = web.page.AllByID("txtID").Fill(web.user.Id)
	if err != nil {
		return
	}

	err = web.page.AllByID("txtPsw").Fill(web.user.password)
	if err != nil {
		return
	}

	err = web.page.AllByID("btnLogin").Click()
	if err != nil {
		return
	}

	err = web.page.AllByID("ctl00_ContentPlaceHolder1_imgBtnSyuugyou").Click()
	if err != nil {
		return
	}

	return nil
}

func (web *webInfo) logout() (err error) {
	err = web.page.AllByID("ctl00_LnkbtnLogOut").Click()
	if err != nil {
		return
	}

	return nil
}

func (web *webInfo) punch(isCome bool) (err error) {
	if isCome == true {
		// come to work
		return web.page.AllByID("ctl00_ContentPlaceHolder1_ibtnIn3").Click()
	} else {
		// leave to work
		return web.page.AllByID("ctl00_ContentPlaceHolder1_imgOut3").Click()
	}
}
