package db

import (
	"encoding/json"
	"log"
	"os"
)

func InsertData() error {
	user := Users{Username: "Rayan Gosling", Password: "robert"}
	result := DB.Create(&user)
	if result.Error != nil {
		log.Fatal("Ошибка при добавлении пользователя:", result.Error)
	}
	loadPosts()
	return nil
}

func loadPosts() {
	data, err := os.ReadFile("./db/post.json")
	if err != nil {
		log.Fatal("Ошибка чтения файла posts.json:", err)
	}

	var posts []Posts
	err = json.Unmarshal(data, &posts)
	if err != nil {
		log.Fatal("Ошибка разбора JSON:", err)
	}

	for _, post := range posts {
		result := DB.Create(&post)
		if result.Error != nil {
			log.Printf("Ошибка при добавлении поста %s: %v", post.Name, result.Error)
		}
	}
}
