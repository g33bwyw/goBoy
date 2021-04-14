package main

import (
	"fmt"
	"strings"
)

//定义金币分配规则
type CoinRuleType map[byte]int
//初始化金币规则
var coinRule = CoinRuleType{
	'e' : 1,
	'i' : 2,
	'o' : 3,
	'u' : 4,
}
var coin int = 50
func getCoinByWord(name string) (userCoin int){
	username := strings.ToLower(name)
	for _,v := range username{
		coinValue,ok := coinRule[byte(v)]
		if ok {
			userCoin += coinValue
			coin -= coinValue
		}
	}
	return
}

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
func main() {
	allUser := []string{"Matthew", "Sarah", "Augustuuuus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	userStat := make(map[string]int)
	for _,user := range allUser{
		userCoin := getCoinByWord(user)
		userStat[user] = userCoin
		fmt.Printf("用户名：%s,金币数：%d\n", user, userCoin)
	}
	fmt.Printf("剩余金币数：%d\n", coin)
}
