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

func ListApplicants() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := utils.ApplicantCollection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Ошибка получения данных:", err)
		return
	}
	defer cursor.Close(ctx)

	fmt.Println("Абитуриенты:")
	for cursor.Next(ctx) {
		var a models.Applicant
		if err := cursor.Decode(&a); err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}
		fmt.Printf("- %s (%s)\n", a.Name, a.Email)
	}
}

func AddApplicant(reader *bufio.Reader) {
	fmt.Print("Имя: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')
	fmt.Print("Телефон: ")
	phone, _ := reader.ReadString('\n')
	fmt.Print("Дата рождения (YYYY-MM-DD): ")
	birthDate, _ := reader.ReadString('\n')

	applicant := models.Applicant{
		Name:      strings.TrimSpace(name),
		Email:     strings.TrimSpace(email),
		Phone:     strings.TrimSpace(phone),
		BirthDate: strings.TrimSpace(birthDate),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := utils.ApplicantCollection.InsertOne(ctx, applicant)
	if err != nil {
		fmt.Println("Ошибка добавления:", err)
		return
	}

	fmt.Println("Абитуриент добавлен.")
}

func DeleteApplicant(reader *bufio.Reader) {
	fmt.Print("Введите email абитуриента для удаления: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := utils.ApplicantCollection.DeleteOne(ctx, bson.M{"email": email})
	if err != nil {
		fmt.Println("Ошибка удаления:", err)
		return
	}

	if res.DeletedCount == 0 {
		fmt.Println("Абитуриент не найден.")
	} else {
		fmt.Println("Абитуриент удалён.")
	}
}
