package role

import "log"

// 某个用户是否有某个端口的权限
var UserPort = make(map[string][]int)

func SetUserPortRole(user string, port int) {
	ints := append(UserPort[user], port)
	print(ints)
}

func CheckUserPortRole(user string, port int) bool {
	if len(UserPort[user]) == 0 {
		return false
	}
	for _, v := range UserPort[user] {
		if v == port {
			return true
		}
	}
	return false

}

func DelUserPortRole(user string, port int) {
	for i, v := range UserPort[user] {
		if v == port {
			_ := append(UserPort[user][:i], UserPort[user][i+1:]...)

		}
	}

}

func init() {
	initUserPort := append(UserPort["admin"], 9200, 3306, 9999, 10121, 10221, 8761)
	log.Println("初始化角色端口权限数据 ： ", initUserPort)
}
