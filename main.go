package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	. "github.com/kbrighton/commandlinego/pkg"
	"github.com/tkanos/gonfig"
	"os"
	"strconv"
)

//Separating this from main to keep it small

func Initialize() *pgx.Conn {
	configuration := Configuration{}
	err := gonfig.GetConf("config.json", &configuration)
	if err != nil {
		fmt.Println(err)
	}

	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		configuration.User,
		configuration.Password,
		configuration.Host,
		configuration.Port,
		configuration.Database)

	config, _ := pgx.ParseConfig(url)

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to establish connection to database: %v\n", err)
		os.Exit(1)
	}

	return conn

}

//Main should just be the command line args
func main() {

	conn := Initialize()

	if len(os.Args) == 1 {
		DisplayHelp()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "add":
		tempTask := Task{Description: os.Args[3], TaskName: os.Args[2]}
		err := CreateTask(conn, tempTask)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not create task: %v\n", err)
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	case "update":
		tempTask := os.Args[3]
		tempId, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "That is not a valid value: %v\n", err)
			os.Exit(1)
		}
		err = UpdateTask(conn, tempId, tempTask)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not update task: %v\n", err)
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	case "list":
		GetTasks(conn)
		os.Exit(0)
	case "complete":
		tempInt, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "That is not a valid value: %v\n", err)
			os.Exit(1)
		}
		err = CompleteTask(conn, tempInt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not complete task: %v\n", err)
			os.Exit(1)
		} else {
			os.Exit(0)
		}

	default:
		fmt.Println("This is not a valid option")
		DisplayHelp()
		os.Exit(0)
	}

}
