package staff

import (
	"signin-go/common/types"
	"signin-go/internal/core"
	"signin-go/repository/staff"
)

func GetConsultantRenewData(ctx core.Context, studioID, staffUserID uint32) (datas []*types.ConsultantRenewData, err error) {
	datas = []*types.ConsultantRenewData{}
	if staffUserID != 0 {
		return
	}

	_, err = staff.StudioConsultantOnlyID(ctx, studioID)
	if err != nil {
		return
	}

	// var wg sync.WaitGroup
	// mu := sync.Mutex{}
	// appendDatas := func(data *types.ConsultantRenewData) {
	// 	mu.Lock()
	// 	defer func() {
	// 		mu.Unlock()
	// 	}()
	// 	datas = append(datas, data)
	// }

	// for _, idsStudioConsultant := range idsStudioConsultants {
	// 	wg.Add(1)
	// 	go func(c core.Context) {
	// 		renewTargeValueRedisKey := redis.GetRenewTargeValueRedisKey(studioID, staffUserID)
	// 		renewTargeValue, _ := redis.GetRenewTargeValue(c, renewTargeValueRedisKey)

	// 	}(ctx)
	// }

	return
}
