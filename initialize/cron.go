package initialize

import (
	"github.com/robfig/cron/v3"
	"rmall/dao"
)

func timedTask() {
	c := cron.New()
	c.AddFunc("@every 1m", func() {
		err := dao.UpdateOrderTimeout()
		if err != nil {
			return
		}
	})
	c.Start()
}
