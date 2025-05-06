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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddApplication(reader *bufio.Reader) {
	fmt.Print("ID абитуриента: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	ctx, cancel := context.WithTimeout(context.Background(), 24*time.Second)
	defer cancel()

	// Найдём абитуриента
	var applicant models.Applicant
	err := utils.ApplicantCollection.FindOne(ctx, bson.M{"_id": applicant.ID}).Decode(&applicant)
	if err != nil {
		fmt.Println("Абитуриент не найден.")
		return
	}

	fmt.Print("ID образовательной программы: ")
	progIDStr, _ := reader.ReadString('\n')
	progIDStr = strings.TrimSpace(progIDStr)
	progID, err := primitive.ObjectIDFromHex(progIDStr)
	if err != nil {
		fmt.Println("Неверный ID программы.")
		return
	}

	fmt.Print("Статус заявления (submitted/approved/rejected): ")
	status, _ := reader.ReadString('\n')

	app := models.Application{
		ApplicantID: applicant.ID,
		ProgramID:   progID,
		Status:      strings.TrimSpace(status),
		SubmittedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	_, err = utils.ApplicationCollection.InsertOne(ctx, app)
	if err != nil {
		fmt.Println("Ошибка при добавлении заявления:", err)
		return
	}

	fmt.Println("Заявление добавлено.")
}

func ListApplications() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := utils.ApplicationCollection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Ошибка получения заявлений:", err)
		return
	}
	defer cursor.Close(ctx)

	fmt.Println("Список заявлений:")
	for cursor.Next(ctx) {
		var a models.Application
		if err := cursor.Decode(&a); err != nil {
			continue
		}

		fmt.Printf("- Заявление от %s на программу %s | Статус: %s | Подано: %s\n",
			a.ApplicantID.Hex(), a.ProgramID.Hex(), a.Status, a.SubmittedAt)
	}
}
