package entity

/**
 * 用户登陆接口实体
 */
type AdminLoginEntity struct {
	User_name string `json:"user_name"`
	Password  string `json:"password"`
}

/**
 * 某一日期统计的结果
 */
type StatisEntity struct {
	StatisDate  string `json:"statis_date"`
	StatisCount int    `json:"statis_count"`
}
