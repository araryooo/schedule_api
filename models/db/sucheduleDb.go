package db

import (
	// フォーマットI/O
	"fmt"

	// Go言語のORM
	"github.com/jinzhu/gorm"

	// エンティティ(データベースのテーブルの行に対応)
	entity "SCH/models/entity"
)

// DB接続する
func open() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "Schedule"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		// 処理停止とランタイムエラー出力
		panic(err.Error())
	}

	// DBエンジンを「InnoDB」に設定
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	// 詳細なログを表示
	db.LogMode(true)

	// 登録するテーブル名を単数形にする（デフォルトは複数形）
	db.SingularTable(true)

	// マイグレーション（テーブルが無い時は自動生成）
	db.AutoMigrate(&entity.Schedule{})

	fmt.Println("db connected: ", &db)
	return db
}

// FindAllschedules は スケジュールテーブルのレコードを全件取得する
func FindAllSchedules() []entity.Schedule {
	schedules := []entity.Schedule{}

	db := open()
	// select
	db.Order("ID asc").Find(&schedules)

	// defer 関数がreturnする時に実行される
	defer db.Close()

	return schedules
}

// FindSchedule は スケジュールテーブルのレコードを１件取得する
func FindSchedule(scheduleID int) []entity.Schedule {
	schedule := []entity.Schedule{}

	db := open()
	// select
	db.First(&schedule, scheduleID)
	defer db.Close()

	return schedule
}

// Insertschedule は スケジュールテーブルにレコードを追加する
func InsertSchedule(registerschedule *entity.Schedule) {
	db := open()
	// insert
	db.Create(&registerschedule)
	defer db.Close()
}

// UpdateStateschedule は スケジュールテーブルの指定したレコードの状態を変更する
func UpdateStateschedule(scheduleID int, registerschedule *entity.Schedule) {
	schedule := []entity.Schedule{}

	db := open()
	// update
	db.Model(&schedule).Where("ID = ?", scheduleID).Update(schedule)
	defer db.Close()
}

// DeleteSchedule は スケジュールテーブルの指定したレコードを削除する
func DeleteSchedule(scheduleID int) {
	schedule := []entity.Schedule{}

	db := open()
	// delete
	db.Delete(&schedule, scheduleID)
	defer db.Close()
}