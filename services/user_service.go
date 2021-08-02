package services

import (
	"fmt"
	"net/http"
	"shop/dao"
	"shop/entity"
	"shop/tool"

	"github.com/gin-gonic/gin"
)

func UserRegisterService(ctx *gin.Context) {
	var request entity.TUserRegisterRequest
	if err := ctx.BindJSON(&request); err != nil {
		panic(err)
	}

	status := dao.SetUserRegisterInfo(
		request.UserName,
		request.PassWord,
		request.Email,
	)

	var response entity.TUserRegisterResponse
	response.Status = status
	if status == tool.ResponseOK {
		response.Message = "註冊成功！"
	} else {
		response.Message = "註冊失敗！使用者名稱已存在！"
	}

	ctx.JSON(http.StatusOK, response)
}

func UserLoginService(ctx *gin.Context) {
	var request entity.TUserLoginRequest
	if err := ctx.BindJSON(&request); err != nil {
		panic(err)
	}

	status, userName := dao.GetUserInfoByUserNameAndPassword(request.UserName, request.PassWord)

	var response entity.TUserLoginResponse
	response.Status = status
	response.UserName = userName

	if response.Status != tool.ResponseOK {
		response.Message = "登入失敗！帳號密碼錯誤或使用者不存在！"
	} else {
		response.Message = fmt.Sprintf("登入成功！歡迎 %s ！", userName)
	}

	ctx.JSON(http.StatusOK, response)
}

func UserInfoService(ctx *gin.Context) {
	var request entity.TUserInfoRequest
	if err := ctx.BindJSON(&request); err != nil {
		panic(err)
	}

	status, userName, userEmail := dao.GetUserInfoByUserName(request.UserName)

	var response = entity.TUserInfoResponse{
		Status:   status,
		UserName: userName,
		Email:    userEmail,
	}

	if response.Status != tool.ResponseOK {
		response.Message = fmt.Sprintf("使用者 %s 資訊載入失敗！", userName)
	} else {
		response.Message = fmt.Sprintf("使用者 %s 資訊載入成功！", userName)
	}

	ctx.JSON(http.StatusOK, response)
}

func UserEditInfoService(ctx *gin.Context) {
	var request entity.TUserEditInfoRequest
	if err := ctx.BindJSON(&request); err != nil {
		panic(err)
	}

	status, _ := dao.SetUserEditInfo(request.UserName, request.Email)

	var response entity.TUserEditInfoResponse
	response.Status = status

	if response.Status != tool.ResponseOK {
		response.Message = "編輯使用者資訊失敗！"
	} else {
		response.Message = "編輯使用者資訊成功！"
	}

	ctx.JSON(http.StatusOK, response)
}

func UserEditPasswordService(ctx *gin.Context) {
	var request entity.TUserEditPassWordRequest
	if err := ctx.BindJSON(&request); err != nil {
		panic(err)
	}

	status := dao.SetUserEditPassWord(request.UserName, request.PassWord, request.NewPassWord)

	var response entity.TUserEditInfoResponse
	response.Status = status

	if response.Status != tool.ResponseOK {
		response.Message = "變更使用者密碼失敗！"
	} else {
		response.Message = "變更使用者密碼成功！請使用新密碼進行登入！"
	}

	ctx.JSON(http.StatusOK, response)
}
