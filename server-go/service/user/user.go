package user

import (
	"github.com/yuanji6666/gopherAI/common/code"
	myemail "github.com/yuanji6666/gopherAI/common/email"
	"github.com/yuanji6666/gopherAI/common/redis"
	"github.com/yuanji6666/gopherAI/dao/user"
	"github.com/yuanji6666/gopherAI/schema"
	"github.com/yuanji6666/gopherAI/utils"
	"github.com/yuanji6666/gopherAI/utils/myjwt"
)

func Register(email, password, captcha string) (string, code.Code) {

	var ok bool
	var userInformation *schema.User

	//判断用户是否存在
	if exist, _ := user.IsExistUser(email); exist {
		return "", code.CodeUserExist
	}

	//检查验证码
	if ok, _ := redis.CheckCaptchaForEmail(email, captcha); !ok {
		return "", code.CodeInvalidCaptcha
	}

	//初始用户名11位随机数
	username := utils.GetRandomNumbers(11)

	//调dao层register加入数据库
	if userInformation, ok = user.Register(username, email, password); !ok {
		return "", code.CodeServerBusy
	}

	//把用户名发给用户邮箱
	if err := myemail.SendCaptcha(email, username, myemail.UserNameMsg); err != nil {
		return "", code.CodeServerBusy
	}

	//根据id和用户名生成token
	token, err := myjwt.GenerateJwt(int64(userInformation.ID), userInformation.Username)

	if err != nil {
		return "", code.CodeServerBusy
	}

	return token, code.CodeSuccess

}

func Login(username, password string) (string, code.Code) {
	userInformation := new(schema.User)
	var ok bool

	if ok, userInformation = user.IsExistUser(username); !ok {
		return "", code.CodeUserNotExist
	}

	if userInformation.Password != utils.MD5(password) {
		return "", code.CodeIllegalPassword
	}

	token, err := myjwt.GenerateJwt(int64(userInformation.ID), username)

	if err != nil {
		return "", code.CodeServerBusy
	}

	return token, code.CodeSuccess
}

func SendCaptcha(email string) code.Code {
	sendCode := utils.GetRandomNumbers(6)

	if err := redis.SetCaptchaForEmail(email, sendCode); err != nil {
		return code.CodeServerBusy
	}

	if err := myemail.SendCaptcha(email, sendCode, myemail.CodeMsg); err != nil {
		return code.CodeServerBusy
	}

	return code.CodeSuccess
}
