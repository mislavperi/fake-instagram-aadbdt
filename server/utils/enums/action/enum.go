package enums

type Action int64

const (
	UNDEFINED Action = iota
	UPLOAD_IMAGE
	DELETE_IMAGE
	MODIFY_IMAGE

	CREATE_USER
	MODIFY_USER
	DELETE_USER
	LOGIN_USER
	GET_USER_INFO
	GET_USER_LOGS
	GET_USERS

	CHANGE_PLAN
	CHOOSE_PLAN

	GET_CONSUMPTION
)

func (a Action) String() string {
	switch a {
	case UPLOAD_IMAGE:
		return "upload_image"
	case DELETE_IMAGE:
		return "delete_image"
	case MODIFY_IMAGE:
		return "modify_image"

	case CREATE_USER:
		return "create_user"
	case MODIFY_USER:
		return "modify_user"
	case DELETE_USER:
		return "delete_user"
	case LOGIN_USER:
		return "login_user"
	case GET_USER_INFO:
		return "get_user_info"
	case GET_USER_LOGS:
		return "get_user_logs"
	case GET_USERS:
		return "get_users"

	case CHANGE_PLAN:
		return "change_plan"
	case CHOOSE_PLAN:
		return "choose_plan"

	case GET_CONSUMPTION:
		return "get_consumption"
	}
	return "unknown"
}
