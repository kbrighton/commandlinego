package pkg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"time"
)

type Configuration struct {
	Password string
	User     string
	Host     string
	Port     uint16
	Database string
}

type Task struct {
	TaskName    string
	Description string
}

func GetTasks(conn *pgx.Conn) error {
	rows, _ := conn.Query(context.Background(), "select id,taskname,description,date_updated from todos where is_complete = false order by date_updated")

	fmt.Println("The currently active tasks are:")

	for rows.Next() {
		var id int
		var taskName string
		var description string
		var date_updated time.Time
		err := rows.Scan(&id, &taskName, &description, &date_updated)
		if err != nil {
			return err
		}
		fmt.Printf("%d\t%s\t%s\t%d-%02d-%02d\n", id, taskName, description, date_updated.Year(), date_updated.Month(), date_updated.Day())

	}
	return rows.Err()

}

func CreateTask(conn *pgx.Conn, task Task) error {
	_, err := conn.Exec(context.Background(), "insert into todos(taskName,description,date_created,date_updated) values($1,$2,$3,$4)", task.TaskName, task.Description, time.Now(), time.Now())

	return err

}

func CompleteTask(conn *pgx.Conn, taskId int) error {
	_, err := conn.Exec(context.Background(), "update todos set is_complete = true where id=$1", taskId)

	return err

}

func UpdateTask(conn *pgx.Conn, taskId int, description string) error {
	_, err := conn.Exec(context.Background(), "update todos set description = $2 where id=$1", taskId, description)

	return err

}

func DisplayHelp() {

	fmt.Println("The options are:")
	fmt.Printf("list:\tThis will list active tasks\n")
	fmt.Printf("complete <id>\tThis will complete the task\n")
	fmt.Printf("add \"<title>\" \"<description>\"\tThis will create a new task\n")
	fmt.Printf("update <id> \"<description>\"\tThis will update a task's description\n")
}
