package gamelogic

import (
	"gameserver/module"
	"fmt"
	"encoding/json"
	//"gameserver/event"
	"gameserver/agent"
)

type Player struct{
	x float32
	y float32
	id int
	agentId int
}


type SyncDataC_S struct{
	PosX float32
	PosY float32
}


type NewPlayerSync struct{
	PlayerId int
}

func Create(id int, agentId int)(Player){
	var player Player
	player.id = id
	player.agentId = agentId

	var newPlayerSyncData NewPlayerSync
	newPlayerSyncData.PlayerId = id
	fmt.Println("create player and send message to client")
	agent.GetAgent(player.agentId).SendMessage(0, newPlayerSyncData)

	thisModule.AddNetEventHandler(2, func(data module.PackageData){
		player.onSync(data)
	})


	return player
}


func (player Player)Remove(){
	thisModule.RemoveNetEventHandler(player.id)
}



func (player Player)SetPosition(x float32, y float32){
	player.x = x
	player.y = y
}


func (player Player)Sync(syncData SyncDataS_C){
	temp := fmt.Sprintf("%d", syncData.PlayerId)
	fmt.Println("syncdata:"+temp)
	agent.GetAgent(player.agentId).SendMessage(1, syncData)
}


func (player Player)onSync(data module.PackageData){
	var syncData SyncDataC_S
	pkgData :=  data.([]byte)
	err := json.Unmarshal(pkgData, syncData)
	if err != nil {
		player.SetPosition(syncData.PosX, syncData.PosY)
	}
}