package cron

import (
	"EpidemicAssistant/config"
	"EpidemicAssistant/service"
	"flag"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

var configFile = flag.String("config", "app.yaml", "配置文件路径")

var conf = config.Init(*configFile)

func StartSchedule() {
	log.Info("Epidemic assistant server started successfully~")
	crontab := cron.New(cron.WithSeconds())
	// 添加定时任务
	_, _ = crontab.AddFunc(conf.Cron, Run)
	// 启动定时器
	crontab.Start()
	select {}
}

func Run() {

	//推送消息到微信
	filePath := service.Screenshot(conf.Screenshot.Token)
	md5 := service.GetFileMd5(filePath)
	base64 := service.GetPicBase64(filePath)
	log.Info(md5)
	service.SendImageMessage(base64, md5, conf.QyWechat.Key)
}
