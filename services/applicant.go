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
	"go.mongodb.org/mongo-driver/mongo"
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
		fmt.Printf("ID: %s", a.ID)
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

	// Проверка валидности email
	email = strings.TrimSpace(email)
	if !isValidEmail(email) {
		fmt.Println("Ошибка: некорректный формат email")
		return
	}

	// Проверка валидности даты
	birthDate = strings.TrimSpace(birthDate)
	if _, err := time.Parse("2006-01-02", birthDate); err != nil {
		fmt.Println("Ошибка: некорректный формат даты. Используйте YYYY-MM-DD")
		return
	}

	applicant := models.Applicant{
		Name:      strings.TrimSpace(name),
		Email:     email,
		Phone:     strings.TrimSpace(phone),
		BirthDate: birthDate,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Проверка уникальности email перед вставкой
	count, err := utils.ApplicantCollection.CountDocuments(ctx, bson.M{"email": email})
	if err != nil {
		fmt.Println("Ошибка проверки email:", err)
		return
	}
	if count > 0 {
		fmt.Println("Ошибка: абитуриент с таким email уже существует")
		return
	}

	_, err = utils.ApplicantCollection.InsertOne(ctx, applicant)
	if err != nil {
		fmt.Println("Ошибка добавления:", err)
		return
	}

	fmt.Println("Абитуриент успешно добавлен.")
}

func isValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func DeleteApplicant(reader *bufio.Reader) error {
	fmt.Print("Введите email абитуриента для удаления: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var applicant models.Applicant
	err := utils.ApplicantCollection.FindOne(ctx, bson.M{"email": email}).Decode(&applicant)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("абитуриент с email %s не найден", email)
		}
		return fmt.Errorf("ошибка поиска абитуриента: %v", err)
	}

	// 2. Удаляем все заявления этого абитуриента
	filter := bson.M{"applicant_id": applicant.ID}
	deleteResult, err := utils.ApplicationCollection.DeleteMany(ctx, filter)
	if err != nil {
		return fmt.Errorf("ошибка удаления заявлений: %v", err)
	}
	fmt.Printf("Удалено %d заявлений абитуриента\n", deleteResult.DeletedCount)

	// 3. Удаляем самого абитуриента
	_, err = utils.ApplicantCollection.DeleteOne(ctx, bson.M{"_id": applicant.ID})
	if err != nil {
		return fmt.Errorf("ошибка удаления абитуриента: %v", err)
	}
	_, err = utils.DocumentCollection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("ошибка удаления документа: %v", err)
	}

	return nil
}
