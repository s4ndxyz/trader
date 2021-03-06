package mock

import (
	"github.com/long2ice/trader/db"
	"github.com/long2ice/trader/exchange"
	"time"
)

type KLineService struct {
	exchange.KLineService
}

func (service *KLineService) Do() ([]exchange.KLine, error) {
	var kLines []db.KLine
	startTime := time.Unix(int64(*service.StartTime), 0)
	db.Client.Where("symbol = ?", service.Symbol).Where("close_time > ?", startTime).Limit(*service.Limit).Order("close_time").Find(&kLines)
	var ret []exchange.KLine
	for _, line := range kLines {
		ret = append(ret, exchange.KLine{
			Open:      line.Open,
			Close:     line.Close,
			High:      line.High,
			Low:       line.Low,
			Amount:    line.Amount,
			Volume:    line.Vol,
			Finish:    true,
			CloseTime: line.CloseTime,
		})
	}
	return ret, nil
}
