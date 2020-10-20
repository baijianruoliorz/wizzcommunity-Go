package models

/*
*  @author liqiqiorz
*  @data 2020/10/20 16:57
 */
//这里的tag就变成 db了 因为要和数据库交互
type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Token    string
}
