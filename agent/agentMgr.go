package agent

import (
	"gameserver/tcpListener"
	"fmt"
)

var agents map[int]Agent
var curId int


func Start(){
	agents = make(map[int]Agent)
	curId = 0

	go waitForCreateAgent()
}


func waitForCreateAgent(){
	for {
		conn := <- tcpListener.AcceptWiat
		fmt.Printf("start create agent\n")
		agent := CreateAgent(conn, curId)
		curId++
		agents[agent.id] = agent
		fmt.Printf("create agent:%d\n", agent.id)
	}
}


func DeleteAgent(id int){
	delete(agents, id)
}

func GetAgent(id int)(Agent){
	agent, _ := agents[id]
	return agent
}

