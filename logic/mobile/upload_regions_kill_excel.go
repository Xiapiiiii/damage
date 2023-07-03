package mobile

import (
	"damage/model"
	"damage/svc"
	"damage/types"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

func UploadRegionsKillExcel(c *gin.Context) (common types.CommonResp, err error) {
	ctx := c.MustGet("ctx").(*svc.ServiceContext)
	file, err := c.FormFile("file")
	if err != nil {
		return types.CommonResp{
			Code:    501,
			Message: "打开文件异常",
		}, err
	}
	fmt.Println(file.Size)

	err = c.SaveUploadedFile(file, file.Filename)
	if err != nil {
		c.String(500, "保存文件失败")
		return types.CommonResp{
			Code:    502,
			Message: "导入数据异常",
		}, err
	}

	// 删除临时文件
	defer func() {
		err = ioutil.WriteFile(file.Filename, []byte(""), 0644)
		if err != nil {
			c.String(503, "删除文件失败")
			return
		}
	}()

	f, err := excelize.OpenFile(file.Filename)
	if err != nil {
		fmt.Println(err)
		return types.CommonResp{
			Code:    504,
			Message: "导入数据异常",
		}, err
	}

	rows, err := f.GetRows("Sheet1")

	for i, row := range rows {
		if i == 0 { // Skip the header
			continue
		}
		adventure := &model.RegionsKill{
			Name:             row[0],
			TriggerCondition: row[4],
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		}
		adventure.CoordX, err = strconv.ParseInt(row[1], 10, 64)
		if err != nil {
			fmt.Println(err)
			return types.CommonResp{
				Code:    505,
				Message: "导入数据异常",
			}, err
		}
		adventure.CoordY, err = strconv.ParseInt(row[2], 10, 64)
		if err != nil {
			fmt.Println(err)
			return types.CommonResp{
				Code:    506,
				Message: "导入数据异常",
			}, err
		}
		adventure.Location, err = strconv.ParseInt(row[3], 10, 64)
		if err != nil {
			fmt.Println(err)
			return types.CommonResp{
				Code:    507,
				Message: "导入数据异常",
			}, err
		}
		adv, err := ctx.RegionsKillModel.GetRegionsKillByName(c.Request.Context(), adventure.Name)
		if err != nil {
			if err != mongo.ErrNoDocuments {
				return types.CommonResp{
					Code:    509,
					Message: "导入数据异常",
				}, err
			}
		}
		if adv != nil {
			adv.CoordX = adventure.CoordX
			adv.CoordY = adventure.CoordY
			adv.Location = adventure.Location
			adv.TriggerCondition = adventure.TriggerCondition
			adv.UpdatedAt = time.Now()

			err := ctx.RegionsKillModel.UpdateRegionsKill(c.Request.Context(), adv)
			if err != nil {
				return types.CommonResp{
					Code:    510,
					Message: "导入数据异常",
				}, err
			}
		} else {
			_, err = ctx.RegionsKillModel.CreateRegionsKill(c.Request.Context(), adventure)
			if err != nil {
				return types.CommonResp{
					Code:    511,
					Message: "导入数据异常",
				}, err
			}
		}
		log.Println(adventure)

	}

	return types.CommonResp{
		Code:    200,
		Message: "success",
	}, nil
}
