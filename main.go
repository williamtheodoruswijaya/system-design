package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Sekarang buat API dengan Database MongoDB

// 1. Buat struct untuk data yang akan diambil, dsbnya di Database.
type Todo struct {
	ID			primitive.ObjectID		`json:"_id,omitempty" bson:"_id,omitempty"` // tambahin `bson:"_id"` basically ini data format di MongoDB
	Completed 	bool					`json:"completed"`
	Body		string					`json:"body"`
} // ID disini jadi primitive.ObjectID karena bawaan dari MongoDB-nya.
// Tambahin juga omitempty agar ga jadi 000000...0000 (bawaan dari MongoDB)

// 2. Buat collection untuk menghubungkan ke MongoDB
var collection *mongo.Collection

func main() {
	// 3. Load .env file-nya
	// Notes: kalau di production, kita akan skip env file ini.
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// 4. Ambil connection string-nya
	MONGODB_URI := os.Getenv("MONGODB_URI")

	// 5. Establish connection ke MongoDB
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client,err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	// Pinging Test
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB:", err)
	}

	// Appendix 3: Kita juga pengen agar client-nya disconnect ketika dia udah selesai (dia akan disconnect setelah func Main selesai)
	defer client.Disconnect(context.Background())

	// Test Connection
	fmt.Println("Connected to MongoDB!")

	// 6. Kita connect ke database (collection) dari MongoDB
	collection = client.Database("golang_db").Collection("todos")

	// 7. Lakukan CRUD API disini
	app := fiber.New()

	// CORS Middleware adjustment (ini biar bisa diakses dari localhost:5173)
	// CORS itu Cross-Origin Resource Sharing, jadi kita bisa akses API dari domain yang berbeda (localhost:5173)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173/",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))
	// CORS dipakai kalau kita mau jalanin dalam local devices, tapi kalau dalam devvelopment, kita bisa comment ni code.

	// 7.1. Get All Todos
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		var todos []Todo

		cursor, err := collection.Find(context.Background(), bson.M{}) // bayangin aja kita mau fetch semua collection-nya (kalau ada filter kita taro ke dalam bson.M{})
		// Kenapa ada cursor dan err? cursor ibaratnya return-value dari query di MongoDB, err ma ya error ya
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Error fetching todos",
			})
		}

		// Appendix 2: ketika Get request ini dijalankan, kita ingin logic-nya dijalankan secara synchronous, untuk memastikan kita jalankan ini:
		defer cursor.Close(context.Background())

		// ingat nih, collection itu ga cuman 1 data tapi bisa banyak dan dia return dalam bentuk cursor
		// Nah kita looping cursor-nya untuk ambil datanya satu-satu
		for cursor.Next(context.Background()) {
			var todo Todo // kita buat variable untuk menampung cursor yang lagi iterate si collections-nya
			if err := cursor.Decode(&todo); err != nil {
				return c.Status(500).JSON(fiber.Map{ // ini basically try-catch-nya
					"error": "Error decoding todo",
				})
			} // Basically kita coba decode cursor-nya ke dalam variable yang kita buat diatas
			todos = append(todos, todo) // Nah ini kita append ke dalam array todos tadi
		}

		// Appendix 1: Decode() itu function buat mindahin data dari collection ke dalam bentuk struct

		return c.JSON(todos)
	})

	// 7.2. Create a Todo
	app.Post("/api/todos", func(c * fiber.Ctx) error {
		// Buat variable berdasarkan struct diatas
		todo := &Todo{}

		// Ambil data dari request body-nya
		if err := c.BodyParser(todo); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Error parsing request body",
			})
		}

		// Basic validation incase body-nya kosong
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Body cannot be empty",
			})
		}

		// Masukin ke dalam collection-nya
		insertResult, err := collection.InsertOne(context.Background(), todo)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Error inserting todo",
			})
		}

		// Masukin ke dalam struct diatas
		todo.ID = insertResult.InsertedID.(primitive.ObjectID)
		todo.Completed = false
		// todo.Body kita skip karena udah diambil diatas

		return c.Status(201).JSON(todo)
	})

	// 7.3. Update a Todo
	app.Put("/api/todos/:id", func(c * fiber.Ctx) error {
		id := c.Params("id") // ambil id dari URL-nya

		// convert id-nya ke dalam bentuk ObjectID (biar bisa kita lakuin comparison "==")
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid ID",
			})
		}

		// Update data-nya pake UpdateOne(context.Background(), filter, update)
		_, err = collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, bson.M{"$set":bson.M{"completed":true}})
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Error updating todo",
			})
		}

		// Kalau berhasil ywd kita return 200 OK
		return c.Status(200).JSON(fiber.Map{
			"message": "Todo updated",
		})
	})

	// 7.4. Delete a Todo
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id") // ambil id dari URL

		// convert id ke ObjectID
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid ID",
			})
		}

		// Delete data-nya pake DeleteOne(context.Background(), filter)
		_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Error deleting todo",
			})
		}

		// Kalau berhasil ywd kita return 200 OK
		return c.Status(200).JSON(fiber.Map{
			"message": "Todo deleted",
		})
	})

	// Run server di sini
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// Sekarang kita akan set agar client folder-nya jadi static
	if os.Getenv("ENV") == "production" {
		app.Static("/", "./client/dist")
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}