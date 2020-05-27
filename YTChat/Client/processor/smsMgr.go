package processor

import (
	"awesomeProject/YTChat/Common"
	"encoding/json"
	"fmt"
)

//管理Message
func outputGroupMes(mes *Common.Message){//Type一定=SmsMessageType

	var smsMes Common.SmsMessage
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil{
		fmt.Println("outputGroupMes unmarshal error:", err)
		return
	}

	//Display
	content := fmt.Sprintf("用户id:\t%d昵称:\t%s 说:\t%s", smsMes.UserId, smsMes.UserName, smsMes.Content)
	fmt.Println(content)
	fmt.Println()
	fmt.Println()

}
