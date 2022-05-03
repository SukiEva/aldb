package e

type code struct {
	Success            int
	CaptchaError       int
	CaptchaMismatch    int
	AuthBindError      int
	AuthEmpty          int
	AuthFormatError    int
	AuthTokenError     int
	AuthTokenTimeOut   int
	AuthError          int
	UserBindError      int
	UserAlreadyExists  int
	RiverBindError     int
	RiverAlreadyExists int
}

var CODE code

func init() {
	CODE = code{
		Success:            200,
		CaptchaError:       301,
		CaptchaMismatch:    302,
		AuthBindError:      400,
		AuthEmpty:          401,
		AuthFormatError:    402,
		AuthTokenError:     403,
		AuthTokenTimeOut:   404,
		AuthError:          405,
		UserBindError:      500,
		UserAlreadyExists:  501,
		RiverBindError:     601,
		RiverAlreadyExists: 602,
	}
}

func ParseCode(num int) string {
	var msg = ""
	switch num {
	case CODE.Success:
		msg = "OK"
	case CODE.CaptchaError:
		msg = "无法获取验证码"
	case CODE.CaptchaMismatch:
		msg = "验证码错误"
	case CODE.AuthBindError:
		msg = "授权格式错误"
	case CODE.AuthEmpty:
		msg = "未授权"
	case CODE.AuthFormatError:
		msg = "授权Token格式错误"
	case CODE.AuthTokenError:
		msg = "授权错误"
	case CODE.AuthTokenTimeOut:
		msg = "授权超时，请重新登录"
	case CODE.AuthError:
		msg = "密码错误"
	case CODE.UserBindError:
		msg = "用户格式错误"
	case CODE.UserAlreadyExists:
		msg = "用户已被注册"
	case CODE.RiverBindError:
		msg = "河流格式错误"
	case CODE.RiverAlreadyExists:
		msg = "河流已经存在"
	default:
		msg = "未知错误"
	}
	return msg
}
