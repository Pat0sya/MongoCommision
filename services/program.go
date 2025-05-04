package services

import (
	"bufio"
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"RBD_dev/models"
	"RBD_dev/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddProgram(reader *bufio.Reader) {
	fmt.Print("Название программы: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("ID факультета (получить через просмотр факультетов): ")
	facultyIDStr, _ := reader.ReadString('\n')
	facultyID, err := primitive.ObjectIDFromHex(strings.TrimSpace(facultyIDStr))
	if err != nil {
		fmt.Println("Неверный ID факультета")
		return
	}

	fmt.Print("Длительность обучения (в годах): ")
	durStr, _ := reader.ReadString('\n')
	duration, err := strconv.Atoi(strings.TrimSpace(durStr))
	if err != nil {
		fmt.Println("Неверный формат длительности")
		return
	}

	program := models.Program{
		Name:      strings.TrimSpace(name),
		FacultyID: facultyID,
		Duration:  duration,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = utils.ProgramCollection.InsertOne(ctx, program)
	if err != nil {
		fmt.Println("Ошибка при добавлении программы:", err)
		return
	}

	fmt.Println("Программа добавлена.")
}

func ListPrograms() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := utils.ProgramCollection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Ошибка получения программ:", err)
		return
	}
	defer cursor.Close(ctx)

	fmt.Println("Список образовательных программ:")
	for cursor.Next(ctx) {
		var p models.Program
		if err := cursor.Decode(&p); err != nil {
			continue
		}
		fmt.Printf("- %s (ID факультета: %s, %d лет)\n", p.Name, p.FacultyID.Hex(), p.Duration)
		fmt.Printf("ID: %s", p.ID)
	}
}
