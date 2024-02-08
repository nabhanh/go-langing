package handlers

import (
	"context"
	"nabhanh/simple-rest-api/internal/db"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// this is like an interface in Go
type Product struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"` // bson is the format that MongoDB uses to store data
	Name      string             `json:"name"`
	Price     float64            `json:"price"`
	CreatedAt time.Time          `json:"createdAt" bson:"created_at"`
	UpdateAt  time.Time          `json:"updatedAt" bson:"updated_at"`
}

func validateProduct(product Product) validator.ValidationErrors {

	validate := validator.New()

	err := validate.Struct(product)

	if err != nil {
		return err.(validator.ValidationErrors)
	}

	return nil
}

func CreateProduct(c *fiber.Ctx) error {
	product := Product{
		ID:        primitive.NewObjectID(),
		Name:      "Product 1",
		Price:     100,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	// c.BodyParser is a middleware that parses the body of the request if there is an error it will return an error
	if err := c.BodyParser(&product); err != nil {
		return err
	}

	errors := validateProduct(product)

	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	client, err := db.GetClient()

	if err != nil {
		return err
	}

	collection := client.Database(db.Db).Collection("products")

	_, err = collection.InsertOne(context.TODO(), product)

	if err != nil {
		return err
	}

	return c.JSON(product)
}

func GetProducts(c *fiber.Ctx) error {
	client, err := db.GetClient()

	if err != nil {
		return err
	}

	collection := client.Database(db.Db).Collection("products")

	cursor, err := collection.Find(context.TODO(), nil)

	if err != nil {
		return err
	}

	var products []Product

	err = cursor.All(context.TODO(), &products)

	if err != nil {
		return err
	}

	return c.JSON(products)
}
