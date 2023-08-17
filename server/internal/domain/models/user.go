package models

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	RoleID    int64  `json:"roleID"`
	Role      Role   `json:"role"`
}

type Statistics struct {
	Plan                  Plan   `json:"plan"`
	TotalConsumptionKb    uint64 `json:"totalConsumptionKb"`
	TotalDailyUploadCount int    `json:"dailyUploadCount"`
	TotalConsumptionCount int    `json:"totalConsumptionCount"`
}

type ExpandedStatistics struct {
	User                  User   `json:"user"`
	Plan                  Plan   `json:"plan"`
	TotalConsumptionKb    uint64 `json:"totalConsumptionKb"`
	TotalDailyUploadCount int    `json:"dailyUploadCount"`
	TotalConsumptionCount int    `json:"totalConsumptionCount"`
}

type UserStatsReq struct {
	UserID int `json:"id"`
}
