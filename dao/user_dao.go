package dao

import (
	"database/sql"
	"fmt"
	"shop/config"
	"shop/tool"

	_ "github.com/go-sql-driver/mysql"
)

func SetUserRegisterInfo(userName string, passWord string, email string) (status int) {
	serverConf := config.ServerConf
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		serverConf.SqlUser, serverConf.SqlPassword, serverConf.SqlSchema))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(
		"INSERT INTO user (username, password, email, authority) VALUES(?, ?, ?, 0)",
		userName,
		tool.EncryptPassword(passWord),
		email,
	)

	if err != nil {
		status = tool.ResponseFailed
	} else {
		status = tool.ResponseOK
	}

	return
}

func GetUserInfoByUserNameAndPassword(userName string, passWord string) (status int, resultName string) {
	serverConf := config.ServerConf
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		serverConf.SqlUser, serverConf.SqlPassword, serverConf.SqlSchema))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT username FROM user WHERE username=? AND password=?", userName, passWord)
	if err != nil {
		return tool.ResponseFailed, ""
	}

	for rows.Next() {
		err = rows.Scan(&resultName)
		if err != nil {
			return tool.ResponseFailed, ""
		}
	}

	status = tool.ResponseOK
	return
}

func GetUserInfoByUserName(userName string) (status int, resultName string, resultEmail string) {
	serverConf := config.ServerConf
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		serverConf.SqlUser, serverConf.SqlPassword, serverConf.SqlSchema))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT username, email FROM user WHERE username=?", userName)
	if err != nil {
		status = tool.ResponseFailed
		return
	}

	for rows.Next() {
		err = rows.Scan(&resultName, &resultEmail)
		if err != nil {
			status = tool.ResponseFailed
			return
		}
	}

	status = tool.ResponseOK
	return
}

func SetUserEditInfo(userName string, newEmail string) (status int, resultEmail string) {
	serverConf := config.ServerConf
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		serverConf.SqlUser, serverConf.SqlPassword, serverConf.SqlSchema))
	if err != nil {
		status = tool.ResponseFailed
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE user SET email=? WHERE username=?")
	if err != nil {
		status = tool.ResponseFailed
		return
	}

	_, err = stmt.Exec(newEmail, userName)
	if err != nil {
		status = tool.ResponseFailed
		return
	}

	status = tool.ResponseOK
	resultEmail = newEmail
	return
}

func SetUserEditPassWord(userName string, oldPassword string, newPassword string) (status int) {
	serverConf := config.ServerConf
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		serverConf.SqlUser, serverConf.SqlPassword, serverConf.SqlSchema))
	if err != nil {
		status = tool.ResponseFailed
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE user SET password=? WHERE username=? AND password=?")
	if err != nil {
		status = tool.ResponseFailed
		return
	}

	_, err = stmt.Exec(newPassword, userName, oldPassword)
	if err != nil {
		status = tool.ResponseFailed
		return
	}

	status = tool.ResponseOK
	return
}

// test
type TUserInfo struct {
	Account string
	Name    string
	Address string
	Phone   string
}

func GetUserInfoById(userId int) (result TUserInfo) {
	serverConf := config.ServerConf
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		serverConf.SqlUser, serverConf.SqlPassword, serverConf.SqlSchema))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user WHERE id=?", userId)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var pw string
		err = rows.Scan(&id, &result.Name, &result.Address, &result.Account, &pw, &result.Phone)
		if err != nil {
			panic(err)
		}
	}

	return
}
