package logic

import (
	"wizzcommunity/dao/mysql"
	"wizzcommunity/models"
)

/*
*  @author liqiqiorz
*  @data 2020/10/22 10:03
 */

//GetCommunityList 查询所有的community并返回
func GetCommunityList() ([]*models.Community, error) {
	//	查询数据库 查找所有的community并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
