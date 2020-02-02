package errno

type ErrorCode int64

const (
	SUCCESS          int64 = 200
	SYS_ERROR        int64 = 65521
	REQ_TIME_OUT     int64 = 65522
	REDIS_MANAGE_ERR int64 = 65523
	MYSQL_MANAGE_ERR int64 = 65524
	NOT_LOGIN        int64 = 65525
	HTTP_REQ_ERR     int64 = 65526
	PARAM_ERR        int64 = 65527
	USER_EXIST       int64 = 401
	PASSWORD_WRONG   int64 = 402
	ACCESS_FORBIDDEN int64 = 403
	WRONG_FORMAT	int64=600
)
