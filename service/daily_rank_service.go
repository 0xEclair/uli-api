package service

import (
	"fmt"
	"go-crud/cache"
	"go-crud/model"
	"go-crud/serializer"
	"strings"
)

type DailyRankService struct{
	
}


// Get
func (service *DailyRankService)Get()serializer.Response{
	var videos []model.Video
	
	// 从redis获取top10
	vids,_:=cache.RedisClient.ZRevRange(cache.DailyRankKey,0,9).Result()
	if len(vids)>1{
		order:=fmt.Sprintf("FIELD(id,%s)",strings.Join(vids,","))
		err:=model.DB.Where("id in (?)",vids).Order(order).Find(&videos).Error
		if err!=nil{
			return serializer.Response{
				Status: 50000,
				Msg:    "数据库连接错误daily_rank",
				Error:  err.Error(),
			}
		}
	}

	return serializer.Response{
		Data:serializer.BuildVideos(videos),
	}
}