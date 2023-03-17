package validation

type TagErrorCode map[string]ErrorCode

const (
	EmailIsExist = "email_is_exist"
	PhoneIsExist = "phonenumber_is_exist"
)

type ErrorCode string

const (
	IsExistErrorCode ErrorCode = "isExist"
)

var TagMap TagErrorCode

func init() {
	TagMap = map[string]ErrorCode{
		EmailIsExist: IsExistErrorCode,
		PhoneIsExist: IsExistErrorCode,
	}
}
