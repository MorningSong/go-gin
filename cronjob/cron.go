package cronjob

import (
	"github.com/robfig/cron"
	"log"
	"github.com/MorningSong/go-gin/models"
	"time"
)

func Setup() {
	go Run()
}

func Run()  {
	log.Println("Run Cronjob.cron.Setup...")

	c := cron.New()
	c.AddFunc("*/30 * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("*/30 * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
