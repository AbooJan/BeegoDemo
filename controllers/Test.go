package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"BeegoDemo1/models"
)

// TestController is the test operation
type TestController struct {
	beego.Controller
}

// Get all user
func (ctr *TestController) Get() {

	userID, err := ctr.GetInt64("userID")

	if err != nil {
		ctr.failResponse("can't get param userID")

	} else {

		userInfo := models.User{}
		userInfo.ID = userID
		userInfo.Age = 20
		userInfo.Gender = 1
		userInfo.Height = 165.0
		userInfo.Name = "Aboo"

		res := models.BaseResponse{Code: "111", Msg:"success", Data:userInfo}

		ctr.Data["json"] = res
		ctr.ServeJSON()
	}
}

func (ctr *TestController) Friends() {
	userID, err := ctr.GetInt64("userID")

	if err != nil {
		ctr.failResponse("invalid param")

	}else {
		logs.Info(userID)

		var friendsList = make([]models.User, 10)

		i := 0
		for i < 10 {

			userInfo := models.User{}
			userInfo.ID = int64(i)
			userInfo.Age = 20+i
			userInfo.Gender = 1
			userInfo.Height = 160.0 + float64(i)
			userInfo.Name = "Aboo"+string(i)

			friendsList[i] = userInfo

			i++
		}

		res := models.BaseResponse{}
		res.Code = "111"
		res.Msg = "success"
		res.Data = friendsList

		ctr.Data["json"] = res
		ctr.ServeJSON()
	}
}

// Post to login user
func (ctr *TestController) Post() {

	phone := ctr.GetString("phone")
	pw := ctr.GetString("pw")

	if phone == "" || pw == "" {
		ctr.failResponse("invalid param")

	}else {

		userInfo := models.User{}
		userInfo.ID = 123456789
		userInfo.Age = 22
		userInfo.Gender = 0
		userInfo.Height = 185.0
		userInfo.Name = "Aboo"

		res := models.BaseResponse{Code: "111", Msg:"success", Data:userInfo}

		ctr.Data["json"] = res
		ctr.ServeJSON()
	}
}

// UploadAvatar is Operation for avatar
func (ctr *TestController) UploadAvatar() {

	f, h, err := ctr.GetFile("avatar")
	phone := ctr.GetString("phone")

	logs.Info(h.Filename)
	logs.Info(phone)

	if f != nil {
		defer  f.Close()
	}

	if err != nil {

		ctr.failResponse("read file err:"+err.Error())

	} else {

		ctr.SaveToFile("avatar", "static/upload/"+h.Filename)

		res := models.BaseResponse{}
		res.Code = "111"
		res.Msg = "保存成功：" + h.Filename
		ctr.Data["json"] = res
		ctr.ServeJSON()
	}
}

/// Avatar request is download user avatar
func (ctr *TestController) Avatar() {

	userId, err := ctr.GetInt64("userId")

	if err == nil {
		logs.Info("GET AVATAR" + string(userId))

		// 返回 头像
		ctr.Ctx.Output.Download("static/upload/avatar.jpeg", "avatar.jpeg")

	}else {

		ctr.failResponse("Request Fail: no param userId")
	}
}

// Version will return news api version
func (ctr *TestController) Version() {

	res := models.BaseResponse{}
	res.Code = "111"
	res.Msg = "v1.0.0"
	ctr.Data["json"] = res
	ctr.ServeJSON()
}

func (ctr *TestController) failResponse(msg string) {

	res := models.BaseResponse{}
	res.Code = "000"
	res.Msg = msg
	ctr.Data["json"] = res
	ctr.ServeJSON()
}
