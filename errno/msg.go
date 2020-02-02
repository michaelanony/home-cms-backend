package errno

var MsgFlags = map[int64]string{
	SUCCESS:          "ok",
	SYS_ERROR:        "sys err",
	REQ_TIME_OUT:     "time out",
	REDIS_MANAGE_ERR: "redis get data err",
	MYSQL_MANAGE_ERR: "mysql get data err",
	NOT_LOGIN:        "not login",
	HTTP_REQ_ERR:     "http request err",
	PARAM_ERR:        "param err",
	USER_EXIST:       "userHandle existed",
	ACCESS_FORBIDDEN: "access forbidden",
	PASSWORD_WRONG:   "wrong password",
	WRONG_FORMAT:"wrong format",
}

func GetMsg(code int64) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[SYS_ERROR]
}
