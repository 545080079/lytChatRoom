package utils

import (
	"bou.ke/monkey"
	"encoding/json"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func mockRead(b []byte) (n int, err error) {
	return 4, nil
}

func mockMarshal(v interface{}) ([]byte, error) {
	return []byte{'a', 'b'}, nil
}

func mockUnmarshal(data []byte, v interface{}) error {
	return nil
}

func TestTransfer_ReadPkg(t *testing.T) {

	//var conn *net.Conn
	//monkey.PatchInstanceMethod(reflect.TypeOf(conn),"Read", mockRead)
	monkey.Patch(json.Marshal, mockMarshal)
	monkey.Patch(json.Unmarshal, mockUnmarshal)
	//transfer := &Transfer{
	//	Conn: nil,
	//	Buf : [8096]byte{},
	//}
	convey.Convey("test ReadPkg", t, func() {
		a := 50
		convey.So(a, convey.ShouldEqual, 50)
	})

}

//
//func TestTransfer_WritePkg(t *testing.T) {
//
//}
