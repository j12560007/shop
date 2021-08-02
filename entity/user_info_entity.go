package entity

type TUserRegisterRequest struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Email    string `json:"email"`
}

type TUserRegisterResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type TUserLoginRequest struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type TUserLoginResponse struct {
	Status   int    `json:"status"`
	Message  string `json:"message"`
	UserName string `json:"username"`
}

type TUserInfoRequest struct {
	UserName string `json:"username"`
}

type TUserInfoResponse struct {
	Status   int    `json:"status"`
	Message  string `json:"message"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type TUserEditInfoRequest struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type TUserEditInfoResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type TUserEditPassWordRequest struct {
	UserName    string `json:"username"`
	PassWord    string `json:"password"`
	NewPassWord string `json:"newpassword"`
}

type TUserEditPassWordResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
