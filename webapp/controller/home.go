package controller

import (
	"time"

	"github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"
)

type Home struct {
}

func (h *Home) Default(k *knot.WebContext) interface{} {
	k.Config.OutputType = knot.OutputTemplate
	return ""
}

func (h *Home) Login(k *knot.WebContext) interface{} {
	k.Config.OutputType = knot.OutputTemplate
	k.Config.LayoutTemplate = ""
	return ""
}

func (h *Home) LoginAuth(k *knot.WebContext) interface{} {
	var loginRequest struct {
		UserID   string `json:"userid"`
		Password string `json:"password"`
	}

	k.Config.OutputType = knot.OutputJson
	ret := toolkit.NewResult()

	if err := k.GetPayload(&loginRequest); err != nil {
		return ret.SetErrorTxt("unable to extract payload: " + err.Error())
	}

	loginStatus := struct {
		LoginID, SessionID string
		LoginTime          time.Time
	}{}

	if !(loginRequest.UserID == "user01" && loginRequest.Password == "Password@1234") {
		k.Server.Log().Errorf("login failed. user %s IP: %s", loginRequest.UserID, k.Request.RemoteAddr)
		return ret.SetErrorTxt("Invalid credential")
	}

	loginStatus.LoginID = "user01"
	loginStatus.LoginTime = time.Now()
	loginStatus.SessionID = toolkit.RandomString(32)

	return ret.SetData(loginStatus)
}
