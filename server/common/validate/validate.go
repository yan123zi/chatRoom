package validate

import (
	"errors"
	"regexp"
)

//判断密码是否正确
func IsPassword(password, rePassword string) error {
	if len(password) == 0 {
		return errors.New("密码不能为空！")
	}
	if len(password) < 6 {
		return errors.New("密码不能过短！")
	}
	if password != rePassword {
		return errors.New("两次密码输入不一致！")
	}
	return nil
}

//判断用户名是否正确
func IsUsername(username string) error {
	if len(username) == 0 {
		return errors.New("请输入用户名")
	}
	matched, err := regexp.MatchString("^[0-9a-zA-Z_-]{5,12}$", username)
	if err != nil || !matched {
		return errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	}
	matched, err = regexp.MatchString("^[a-zA-Z]", username)
	if err != nil || !matched {
		return errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	}
	return nil
}
