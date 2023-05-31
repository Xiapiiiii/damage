package logic

import (
	"damage/model"
	"damage/svc"
	"damage/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

func GetRankingList(c *gin.Context, req types.GetRankingReq) (rankingResp types.GetRankingResp, err error) {
	ctx := c.MustGet("ctx").(*svc.ServiceContext)
	rankingRecordList, err := ctx.RankingRecord.GetRankingRecordList(c.Request.Context(), &req)

	if err != nil {
		if err != mongo.ErrNoDocuments {
			return types.GetRankingResp{}, err
		}
		return types.GetRankingResp{}, nil
	}

	data := make([]types.RankingRespData, 0)
	for _, rankingRecord := range rankingRecordList {
		rankingRespData := types.RankingRespData{
			ID:          rankingRecord.ID.Hex(),
			RoleUID:     rankingRecord.RoleUID,
			RoleName:    rankingRecord.RoleName,
			Occupation:  rankingRecord.Occupation,
			AreaService: rankingRecord.AreaService,
			Damage:      rankingRecord.Damage,
		}
		data = append(data, rankingRespData)
	}

	rankingResp.Data = data
	rankingResp.Code = 200
	rankingResp.Message = "success"

	return
}

func UpdateRankingRecord(c *gin.Context, req types.UpdateRankingReq) (rankingResp types.UpdateRankingResp, err error) {
	ctx := c.MustGet("ctx").(*svc.ServiceContext)

	log.Println("one ", req)
	if req.ID == "" {
		rankingRecord := model.RankingRecord{
			RoleUID:     req.RoleUID,
			RoleName:    req.RoleName,
			Occupation:  req.Occupation,
			AreaService: req.AreaService,
			Damage:      req.Damage,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		log.Println("two", rankingRecord)
		_, err := ctx.RankingRecord.CreateRankingRecord(c.Request.Context(), &rankingRecord)
		if err != nil {
			return types.UpdateRankingResp{}, err
		}

		return types.UpdateRankingResp{
			Code:    200,
			Message: "success",
		}, nil
	} else {
		objectId, _ := primitive.ObjectIDFromHex(req.ID)
		record, err := ctx.RankingRecord.GetRankingRecord(c.Request.Context(), objectId)
		if err != nil {
			return types.UpdateRankingResp{}, err
		}
		record.RoleName = req.RoleName
		record.RoleUID = req.RoleUID
		record.Occupation = req.Occupation
		record.AreaService = req.AreaService
		record.Damage = req.Damage
		record.UpdatedAt = time.Now()
		err = ctx.RankingRecord.UpdateRankingRecord(c.Request.Context(), record)
		if err != nil {
			return types.UpdateRankingResp{}, err
		}

		return types.UpdateRankingResp{
			Code:    200,
			Message: "success",
		}, nil
	}

}
