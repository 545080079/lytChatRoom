package processor



var (
	userMgr *UserMgr
)

type UserMgr struct{
	onlineUsers map[int]*UserProcess
	index int
}

func init(){
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}



func(t *UserMgr) addOnlineUser(userProcess *UserProcess){
	t.onlineUsers[t.index] = userProcess
	t.index++
}

func(t *UserMgr) removeOnlineUser(userId int){

}