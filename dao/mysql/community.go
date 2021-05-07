package mysql

import (
	"database/sql"
	"wizzcommunity/models"

	"go.uber.org/zap"
)

/*
*  @author liqiqiorz
*  @data 2020/10/20 20:10
 */
// GetCommunityDetailByID 根据ID查询社区详情

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select 
			community_id, community_name, introduction, create_time
			from community 
			where community_id = ?
	`

	/*	sqlStr2:=`select community_id,community_name,introduction,create_time from community where community_id=?`
		if err:=db.Get(community,sqlStr2,id);err!=nil{
			if err ==sql.ErrNoRows{
				err=ErrorInvalidID
			}
		}*/

	if err := db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return community, err
}
