package note

import (
	"fmt"
	"note-app/pkg/io"
	"time"
)

type Note struct {
	Id          int
	Username    string
	Title       string
	Description string
	CreatedAt   string
}

var DB []Note

func Notes() {
	username := *io.UserInput("what is your name : ")
	fmt.Println("Hey ", username)

	title := *io.UserInput(fmt.Sprintf("%v, what is your note called : ", username))

	if io.FsExist(fmt.Sprintf("../asset/%v.txt", title)) {
		fmt.Printf("Ohh my bad, there's an already that %v..", title)
	}

	fmt.Println("Okay, that's a good one...")
	desc := *io.UserInput(fmt.Sprintf("%v, what kind of description you want : ", username))

	user := Note{
		Id:          len(DB) + 1,
		Username:    username,
		Title:       title,
		Description: desc,
		CreatedAt:   fmt.Sprintf("%v : %v", time.Now().Hour(), time.Now().Minute()),
	}

	result := io.Format("Note App !\n", user.Username, user.Title, user.Description, user.CreatedAt)

	if io.CreateFile(fmt.Sprintf("../asset/%v.txt", title), result) {
		DB = append(DB, user)
		fmt.Println("Database's : ", DB)
		fmt.Println("Note successfully created..")
	}
}
