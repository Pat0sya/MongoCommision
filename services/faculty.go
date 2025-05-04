package services

import (
	"bufio"
	"context"
	"fmt"
	"strings"
	"time"

	"RBD_dev/models"
	"RBD_dev/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func AddFaculty(reader *bufio.Reader) {
	fmt.Print("Название факультета: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Корпус (здание): ")
	building, _ := reader.ReadString('\n')

	faculty := models.Faculty{
		Name:     strings.TrimSpace(name),
		Building: strings.TrimSpace(building),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := utils.FacultyCollection.InsertOne(ctx, faculty)
	if err != nil {
		fmt.Println("Ошибка при добавлении факультета:", err)
		return
	}

	fmt.Println("Факультет добавлен.")
}

func ListFaculties() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := utils.FacultyCollection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Ошибка получения факультетов:", err)
		return
	}
	defer cursor.Close(ctx)

	fmt.Println("Список факультетов:")
	for cursor.Next(ctx) {
		var f models.Faculty
		if err := cursor.Decode(&f); err != nil {
			continue
		}
		fmt.Printf("- %s (корпус: %s)\n", f.Name, f.Building)
		fmt.Printf("ID: %s", f.ID)
	}
}
