package processor

import (
	"awesomeProject/YTChat/Common"
	"awesomeProject/YTChat/Server/model"
	"awesomeProject/YTChat/Server/utils"
	"bou.ke/monkey"
	"encoding/json"
	"github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func mockUnmarshal(b []byte, v interface{}) error {
	v = &Common.LoginMessage{
		UserId:   1,
		UserName: "admin",
		UserPwd:  "admin",
	}
	return nil
}

func mockMarshal(v interface{}) ([]byte, error) {
	var rer = []byte{
		'a', 'd', 'm', 'i', 'n',
	}
	return rer, nil
}

func TestServerProcessLogin(t *testing.T) {
	mess := &Common.Message{
		Type: Common.LoginMesType,
		Data: "default",
	}
	user := &UserProcess{
		Conn: nil,
	}

	//对涉及到的单元以外系统函数打Patch
	monkey.Patch(json.Unmarshal, mockUnmarshal)
	monkey.Patch(json.Marshal, mockMarshal)

	//对实例函数打Patch
	var udao *model.UserDao
	monkey.PatchInstanceMethod(reflect.TypeOf(udao), "Login", func(_ *model.UserDao, _ int, _ string) (*Common.User, error) {
		return &Common.User{
			UserId:   1,
			UserName: "admin",
			UserPwd:  "admin",
		}, nil
	})

	var ts *utils.Transfer
	monkey.PatchInstanceMethod(reflect.TypeOf(ts), "WritePkg", func(_ *utils.Transfer, _ []byte) error {
		return nil
	})
	//执行测试
	convey.Convey("Test Server Login.", t, func() {
		err := user.ServerProcessLogin(mess)
		convey.So(err, convey.ShouldBeNil)
	})
	monkey.UnpatchAll()
	return
}
