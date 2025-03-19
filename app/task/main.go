package main

import (
	"china-russia/dao"
	"china-russia/global"
	"china-russia/model"
	"github.com/gocolly/colly"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	//初始化viper
	global.Viper()
	//初始化log
	global.Log()
	//dao连接
	global.DB = dao.Gorm()
	global.REDIS = dao.Redis()

	ticker := time.NewTicker(time.Minute * 10)
	for {
		<-ticker.C
		c := colly.NewCollector()
		c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

		c.OnHTML("#blk_hdline_01 > h3 > a", func(e *colly.HTMLElement) {
			url := e.Attr("href")
			intro := e.Text
			if url != "" && intro != "" {
				var title, content, dateTime, source string
				c.OnHTML("body > div.main-content.w1240", func(h *colly.HTMLElement) {
					title = h.ChildText("body > div.main-content.w1240 > h1")
					content = h.ChildText("#artibody")
					dateTime = h.ChildText("#top_bar > div > div.date-source > span.date")
					source = h.ChildText("#top_bar > div > div.date-source > span.source.ent-source")
					if source == "" {
						source = h.ChildText("#top_bar > div > div.date-source > a")
					}
				})
				c.Visit(url)
				log.Println("新闻链接：", url)
				//检查是否重复
				value := global.REDIS.HGet("sina_news", url).Val()
				if value == "" {
					//入库
					news := model.News{
						Title:    title,
						Content:  content,
						Status:   1,
						Sort:     0,
						Intro:    intro,
						Cover:    "",
						Category: 0,
						Source:   source,
						DateTime: dateTime,
						Url:      url,
					}
					news.Insert()
					global.REDIS.HSet("sina_news", url, "1")
				}
			}

		})
		c.Visit("https://finance.sina.com.cn/")
	}
	//ticker.Run()
}
