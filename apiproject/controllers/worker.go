package controllers

import (
	"fmt"
	"apiproject/models"
//	"encoding/json"
	"strconv"
	"github.com/astaxie/beego"
)

// Operations about Users
type WorkerController struct {
	beego.Controller
}



// @Title Get
// @Description get user by uid
// @Param	workerid		path 	string	true		"The key for staticblock"
// @Param	token		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :workerid is empty
// @router /querybagreq/:workerid/:token [get]
func (u *WorkerController) Querybagreq() {
	workerId := u.GetString(":workerid")
	token := u.GetString(":token")
	//index := u.GetString(":index")
	
	var worker models.Worker
	worker.WorkerId = workerId
	worker.Token = token
	//worker.Index = index
	
	u.Data["json"] = worker
	u.ServeJson()
}

//http://localhost:8080/v1/worker/getworkerid?workerid=1&token=2
func (u *WorkerController) GetWorkerId() {
	workerId := u.GetString("workerid")
	token := u.GetString("token")
	//index := u.GetString(":index")
	
	var worker models.Worker
	worker.WorkerId = workerId
	worker.Token = token
	//worker.Index = index
	
	u.Data["json"] = worker
	u.ServeJson()
}

func (u *WorkerController) Getbagreq() {
	workerId := u.GetString(":workerid")
	token := u.GetString(":token")
	
	var reworker models.ReWorker
	var worker models.Worker
	worker.WorkerId = workerId + "aaa"
	worker.Token = token
	reworker.Code = "000000"
	reworker.Result = worker
	
	u.Data["json"] = reworker
	u.ServeJson()
}

func (u *WorkerController) GetbagreqList() {
	//workerId := u.GetString(":workerid")////----string方式
	workerId, _ := u.GetInt(":workerid")//----int方式
	token := u.GetString(":token")

	var reworker models.ReWorker
	var worker models.Worker
	workers := make([]models.Worker, 2)
	
	worker.WorkerId =strconv.Itoa(workerId) + "aaaa"//----int方式
	
	worker.Token = token
	reworker.Code = "000001"
	workers[0] = worker
	workers[1] = worker
	reworker.Result = workers
	
	u.Data["json"] = reworker
	u.ServeJson()
}


//******************************************************************
//********post 测试**************************************************
//******************************************************************
func (u *WorkerController) PostWorkerInfo() {
	
	fmt.Println("11postpostpost:", u.Ctx.Input.RequestBody)
	
	var worker models.Worker
	worker.WorkerId = "wdkId"
	worker.Token = "wdkToken"
	//worker.Index = index
	
	u.Data["json"] = worker
	u.ServeJson()
}

func (u *WorkerController) PostWorkerInfo1() {
	
	workerId := u.GetString(":workerid")////----string方式
	token := u.GetString(":token")
	fmt.Println("11postpostpost:", u.Ctx.Input.RequestBody)
	
	var worker models.Worker
	worker.WorkerId = workerId
	worker.Token = token
	//worker.Index = index
	
	u.Data["json"] = worker
	u.ServeJson()
}




