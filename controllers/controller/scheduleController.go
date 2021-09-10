package controller

import (
	// 文字列と基本データ型の変換パッケージ
	strconv "strconv"

	// Gin
	"github.com/gin-gonic/gin"

	// エンティティ(データベースのテーブルの行に対応)
	entity "SCH/models/entity"
	// DBアクセス用モジュール
	db "SCH/models/db"
)

// FetchAllschedules は 全てのスケジュール情報を取得する
func FetchAllSchedules(c *gin.Context) {
	resultschedules := db.FindAllSchedules()

	// URLへのアクセスに対してJSONを返す
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.JSON(200, resultschedules)
}

// Findschedule は 指定したIDのスケジュール情報を取得する
func FindSchedule(c *gin.Context) {
	scheduleIDStr := c.Query("scheduleId")

	scheduleID, _ := strconv.Atoi(scheduleIDStr)

	resultschedule := db.FindSchedule(scheduleID)

	// URLへのアクセスに対してJSONを返す
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.JSON(200, resultschedule)
}

// Addschedule は スケジュールをDBへ登録する
func AddSchedule(c *gin.Context) {
	scheduleName := c.PostForm("scheduleName")
	scheduleMemo := c.PostForm("scheduleMemo")
	scheduleStartDate := c.PostForm("scheduleStartDate")
	scheduleEndDate := c.PostForm("scheduleEndDate")

	var schedule = entity.Schedule{
		Company:    scheduleName,
		Task:       scheduleMemo,
		Start_Date: scheduleStartDate,
		End_Date:   scheduleEndDate,
	}

	db.InsertSchedule(&schedule)
	// URLへのアクセスに対してJSONを返す
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.JSON(200, "")

}

// スケジュールを変更する
func ChangeSchedule(c *gin.Context) {
	scheduleID, _ := strconv.Atoi(c.PostForm("scheduleId"))

	scheduleName := c.PostForm("scheduleName")
	scheduleMemo := c.PostForm("scheduleMemo")
	scheduleStartDate := c.PostForm("scheduleStartDate")
	scheduleEndDate := c.PostForm("scheduleEndDate")

	var schedule = entity.Schedule{
		Company:    scheduleName,
		Task:       scheduleMemo,
		Start_Date: scheduleStartDate,
		End_Date:   scheduleEndDate,
	}

	db.UpdateStateschedule(scheduleID, &schedule)
	// URLへのアクセスに対してJSONを返す
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.JSON(200, "")
}

// スケジュールをDBから削除する
func DeleteSchedule(c *gin.Context) {
	scheduleID, _ := strconv.Atoi(c.PostForm("scheduleId"))

	db.DeleteSchedule(scheduleID)

	// URLへのアクセスに対してJSONを返す
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.JSON(200, "")
}
