package processor

import (
	"awesomeProject/YTChat/Client/utils"
	"awesomeProject/YTChat/Common"
	"encoding/json"
	"fmt"
)

type SmsProcessor struct{

}
func (t *SmsProcessor) SendGroupMessage(content string) (err error) {

	//创建mes实例
	var mes Common.Message
	mes.Type = Common.SmsMessageType

	//创建SmsMessage实例
	var smsMes Common.SmsMessage
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserName = CurUser.UserName

	//序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("Send Group Message's Json Marshal Error:", err)
		return
	}
	mes.Data = string(data)

	//对mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("Send Group Message's Json Marshal Error:", err)
		return
	}

	//发送mes给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	tf.WritePkg(data)
	if err != nil {
		fmt.Println("Send Group Message Error:", err)
		return err
	}

	return err
}