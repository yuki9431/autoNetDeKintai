package component

import (
	"fmt"

	"github.com/sclevine/agouti"
)

type User struct {
	Id       string
	Password string
}

type WebInfo struct {
	User User
	Page *agouti.Page
}

// come to work: true
// leave to work: false
func Punch(user User, isCome bool) (err error) {
	web, err := new(user)
	if err != nil {
		return
	}

	if err := web.accessToNetdeKomon(); err != nil {
		return fmt.Errorf("Error: Faild to access netDeKomon: %v", err)
	}

	if err := web.login(); err != nil {
		return fmt.Errorf("Error: Faild to login %v", err)
	}

	if err := web.punch(isCome); err != nil {
		return fmt.Errorf("Error: Faild to punch %v", err)
	}

	web.Page.AllByClass("timeBtnR").Click()

	if err := web.Page.Screenshot("dakoku.png"); err != nil {
		return fmt.Errorf("Error: Failed to screenshot:%v", err.Error())
	}

	if err := web.logout(); err != nil {
		return fmt.Errorf("Error: Faild to logout %v", err)
	}

	return nil
}

func new(user User) (web *WebInfo, err error) {
	// start driver
	//driver := agouti.PhantomJS()
	driver := agouti.ChromeDriver()
	if err = driver.Start(); err != nil {
		return nil, fmt.Errorf("Error: Failed to start driver:%v", err.Error())
	}
	// defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("phantomjs"))
	if err != nil {
		return nil, fmt.Errorf("Error: Failed to open page:%v", err.Error())
	}

	return &WebInfo{
		User: User{
			Id:       user.Id,
			Password: user.Password,
		},
		Page: page,
	}, nil
}

func (web *WebInfo) accessToNetdeKomon() (err error) {
	if err := web.Page.Navigate("https://www1.shalom-house.jp/komon/"); err != nil {
		return fmt.Errorf("Error: Failed to navigate:%v", err.Error())
	}

	err = web.Page.ClearCookies()
	if err != nil {
		return
	}

	return nil
}

func (web *WebInfo) login() (err error) {

	err = web.Page.AllByID("txtID").Fill(web.User.Id)
	if err != nil {
		return
	}

	err = web.Page.AllByID("txtPsw").Fill(web.User.Password)
	if err != nil {
		return
	}

	err = web.Page.AllByID("btnLogin").Click()
	if err != nil {
		return
	}

	err = web.Page.AllByID("ctl00_ContentPlaceHolder1_imgBtnSyuugyou").Click()
	if err != nil {
		return
	}

	return nil
}

func (web *WebInfo) logout() (err error) {
	err = web.Page.AllByID("ctl00_LnkbtnLogOut").Click()
	if err != nil {
		return
	}

	err = web.Page.CloseWindow()
	if err != nil {
		return
	}

	return nil
}

func (web *WebInfo) punch(isCome bool) (err error) {
	if isCome == true {
		// come to work ctl00_ContentPlaceHolder1_imgIn1
		return web.Page.AllByID("ctl00_ContentPlaceHolder1_ibtnIn3").Click()
	} else {
		// leave to work ctl00_ContentPlaceHolder1_imgnOut1
		return web.Page.AllByID("ctl00_ContentPlaceHolder1_ibtnOut3").Click()
	}
}
