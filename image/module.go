package image

import (
	"encoding/json"
	"fmt"
	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"messiahJKBot/util"
	"strconv"
	"strings"
	"sync"
)

const ModuleName = "us.messiahjk.image"

var instance *image

func init() {
	instance = &image{}
	bot.RegisterModule(instance)
}

func (img *image) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       ModuleName,
		Instance: instance,
	}
}

type image struct{}

var imageMap map[string]int

var groupSet util.Int64Set

func (img *image) Init() {
	imageMap = make(map[string]int)
	groupSet = make(util.Int64Set)
	groupSet.Add(741804645)
	groupSet.Add(294422676)
	groupSet.Add(1050516360)
}

func (img *image) PostInit() {

}

func (img *image) Serve(bot *bot.Bot) {
	bot.OnGroupMessage(func(client *client.QQClient, msg *message.GroupMessage) {
		if groupSet.Has(msg.GroupCode) {
			if strings.ToLower(msg.ToString()) == "/jk" {
				groupCode := msg.GroupCode
				sendingMsg := message.NewSendingMessage()
				sendingMsg.Append(message.NewReply(msg))
				sendingMsg.Append(message.NewText("jk真可爱"))
				client.SendGroupMessage(groupCode, sendingMsg)
			} else if msg.Sender.Uin != bot.Uin && len(msg.Elements) <= 2 {
				marshal, _ := json.Marshal(msg)
				fmt.Println("0:" + string(marshal))
				marshal1, _ := json.Marshal(msg.OriginalObject.Body.RichText.Elems[0])
				fmt.Println("1:" + string(marshal1))
				marshal2, _ := json.Marshal(msg.OriginalObject.Body.RichText.Elems[1])
				fmt.Println("2:" + string(marshal2))
				fmt.Println(1)
				for _, e := range msg.Elements {
					fmt.Println(2)
					if e.Type() == message.Image {
						fmt.Println(3)
						if imageElement, ok := e.(*message.ImageElement); ok {
							fmt.Println(4)
							//这一长串的傻逼玩意是用来判断是否是表情包的
							//图片：CAAQADIAUAB4Ag==   PbReserve:"CAAQADIAUAB4CA=="
							//表情包：CAEQADIASg5b5Yqo55S76KGo5oOFXVAAeAY=
							marshal3, _ := json.Marshal(msg.OriginalObject.Body.RichText.Elems[0].CustomFace.PbReserve)
							if len(marshal3) < 32 {
								fmt.Println(5)
								var key = string(imageElement.Md5) + strconv.FormatInt(msg.GroupCode, 10)
								size, has := imageMap[key]
								if has {
									size++
									imageMap[key] = size
									groupCode := msg.GroupCode
									sendingMsg := message.NewSendingMessage()
									sendingMsg.Append(message.NewReply(msg))
									sendingMsg.Append(message.NewText("该图片已在群里发过，你是第" + strconv.Itoa(size) + "个发的"))
									client.SendGroupMessage(groupCode, sendingMsg)
								} else {
									imageMap[key] = 1
								}
							}
						}
					}
				}
			}
		}
	})
}

func (img *image) Start(bot *bot.Bot) {

}

func (img *image) Stop(bot *bot.Bot, wg *sync.WaitGroup) {
	defer wg.Done()
}
