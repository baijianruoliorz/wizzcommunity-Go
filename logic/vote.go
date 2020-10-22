package logic

import (
	"go.uber.org/zap"
	"strconv"
	"wizzcommunity/dao/redis"
	"wizzcommunity/models"
)

/*
*  @author liqiqiorz
*  @data 2020/10/22 10:17
 */
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))

}
