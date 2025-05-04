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

func AddDocument(reader *bufio.Reader) {
	fmt.Print("Введите email абитуриента: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Найдём абитуриента по email
	var applicant models.Applicant
	err := utils.ApplicantCollection.FindOne(ctx, bson.M{"email": email}).Decode(&applicant)
	if err != nil {
		fmt.Println("Абитуриент не найден")
		return
	}

	// Проверим, есть ли уже документ для этого абитуриента
	count, _ := utils.DocumentCollection.CountDocuments(ctx, bson.M{"applicant_id": applicant.ID})
	if count > 0 {
		fmt.Println("Документы уже существуют для этого абитуриента")
		return
	}

	fmt.Print("Введите паспорт: ")
	passport, _ := reader.ReadString('\n')
	fmt.Print("Введите СНИЛС: ")
	snils, _ := reader.ReadString('\n')

	doc := models.Document{
		ApplicantID: applicant.ID,
		Passport:    strings.TrimSpace(passport),
		SNILS:       strings.TrimSpace(snils),
	}

	_, err = utils.DocumentCollection.InsertOne(ctx, doc)
	if err != nil {
		fmt.Println("Ошибка при добавлении документа:", err)
		return
	}

	fmt.Println("Документ добавлен.")
}

func ShowDocuments() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := utils.DocumentCollection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Ошибка получения документов:", err)
		return
	}
	defer cursor.Close(ctx)

	fmt.Println("Документы абитуриентов:")
	for cursor.Next(ctx) {
		var doc models.Document
		if err := cursor.Decode(&doc); err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		fmt.Printf("- ApplicantID: %s | Паспорт: %s | СНИЛС: %s\n",
			doc.ApplicantID.Hex(), doc.Passport, doc.SNILS)
	}
}

func DeleteDocument(reader *bufio.Reader) {
	fmt.Print("Введите email абитуриента: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var applicant models.Applicant
	err := utils.ApplicantCollection.FindOne(ctx, bson.M{"email": email}).Decode(&applicant)
	if err != nil {
		fmt.Println("Абитуриент не найден")
		return
	}

	res, err := utils.DocumentCollection.DeleteOne(ctx, bson.M{"applicant_id": applicant.ID})
	if err != nil {
		fmt.Println("Ошибка при удалении:", err)
		return
	}

	if res.DeletedCount == 0 {
		fmt.Println("Документ не найден.")
	} else {
		fmt.Println("Документ удалён.")
	}
}
