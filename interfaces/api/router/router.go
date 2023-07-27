package router

import (
	"clockwork-server/application"
	"clockwork-server/config"
	"clockwork-server/domain/repository"
	"clockwork-server/helper"
	"clockwork-server/interfaces/api/auth"
	"clockwork-server/interfaces/api/handler"
	"clockwork-server/interfaces/api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterInterface interface {
	RegisterAPI() *gin.Engine
}

type router struct {
	db *gorm.DB
}

func NewRouter(db *gorm.DB) RouterInterface {
	return &router{
		db,
	}
}

func (r router) RegisterAPI() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	authService := auth.NewService(*config.GetConfig())
	authMiddleware := middleware.NewAuthMiddleware()

	// exampleTrue := new(bool)
	// *exampleTrue = true

	// exampleFalse := new(bool)
	// *exampleFalse = false

	// productData := model.Product{
	// 	Title:        "test",
	// 	Description:  "test",
	// 	SerialNumber: "aaaxxwww",
	// 	UnitPrice:    20000,
	// 	UserID:       1,
	// 	InventoryID:  1,
	// 	CategoryID:   2,
	// 	Attributes:   [{"id" : 1}]
	// }

	// err := r.db.Create(&productData).Error

	// fmt.Println("RabbitMQ in Golang: Getting started")

	// connection, err := amqp.Dial("amqp://user:rabbitmq@localhost:5672/")
	// if err != nil {
	// 	panic(err)
	// }
	// defer connection.Close()

	// fmt.Println("Successfully connected to RabbitMQ instance")

	// // opening a channel over the connection established to interact with RabbitMQ
	// channel, err := connection.Channel()
	// if err != nil {
	// 	panic(err)
	// }
	// defer channel.Close()

	// // declaring queue with its properties over the the channel opened
	// queue, err := channel.QueueDeclare(
	// 	"queue_testing", // name
	// 	false,           // durable
	// 	false,           // auto delete
	// 	false,           // exclusive
	// 	false,           // no wait
	// 	nil,             // args
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// body, err := json.Marshal(map[string]string{
	// 	"Name":  "dika",
	// 	"Hobby": "football",
	// })

	// // publishing a message
	// err = channel.Publish(
	// 	"",              // exchange
	// 	"queue_testing", // key
	// 	false,           // mandatory
	// 	false,           // immediate
	// 	amqp.Publishing{
	// 		ContentType: "application/json",
	// 		Body:        []byte(body),
	// 	},
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Queue status:", queue)
	// fmt.Println("Successfully published message")

	// // declaring consumer with its properties over channel opened
	// msgs, err := channel.Consume(
	// 	"queue_testing", // queue
	// 	"",              // consumer
	// 	true,            // auto ack
	// 	false,           // exclusive
	// 	false,           // no local
	// 	false,           // no wait
	// 	nil,             //args
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// var result struct {
	// 	Name  string
	// 	Hobby string
	// }

	// // print consumed messages from queue
	// forever := make(chan bool)
	// go func() {
	// 	for msg := range msgs {
	// 		err = json.Unmarshal(msg.Body, &result)
	// 		fmt.Printf("Received Message: %s\n", result.Hobby)
	// 	}
	// }()

	// fmt.Println("Waiting for messages...")
	// <-forever

	cfig := config.GetConfig()

	// cacheConnection := database.NewDBRedis(cfig)

	cartItemHelper := helper.NewCartItemHelper()
	globalHelper := helper.NewGlobalHelper()
	orderHelper := helper.NewOrderHelper()

	userRepository := repository.NewRepository(r.db)
	customerRepository := repository.NewCustomerRepository(r.db)
	productAttributeRepository := repository.NewProductAttributeRepository(r.db)
	orderRepository := repository.NewOrderRepository(r.db)
	inventoryRepository := repository.NewInventoryRepository(r.db)
	cartRepository := repository.NewCartRepository(r.db)
	cartItemRepository := repository.NewCartItemRepository(r.db)
	categoryRepository := repository.NewCategoryRepository(r.db)
	attributeRepository := repository.NewAttributeRepository(r.db)
	productRepository := repository.NewProductRepository(r.db)
	attributeItemRepository := repository.NewAttributeItemRepository(r.db)
	cartItemAttributeItemRepository := repository.NewCartItemAttributeItemRepository(r.db)
	paymentRepository := repository.NewPaymentRepository(r.db)
	imageRepository := repository.NewImageRepository(r.db)
	organizationRepository := repository.NewOrganizationRepository(r.db)
	locationRepository := repository.NewLocationRepository(r.db)

	userService := application.NewUserService(userRepository)
	customerService := application.NewCustomerService(customerRepository)
	productService := application.NewProductService(productRepository, categoryRepository,
		inventoryRepository,
		productAttributeRepository,
		attributeRepository,
		globalHelper,
	)
	midtransService := application.NewMidtransService(cfig, orderRepository)
	orderService := application.NewOrderService(
		orderRepository,
		midtransService,
		cartRepository,
		paymentRepository,
		orderHelper,
	)
	inventoryService := application.NewInventoryService(inventoryRepository, cartItemRepository)
	cartItemAttributeItemService := application.NewCartItemAttributeItemService(cartItemAttributeItemRepository)
	cartService := application.NewCartService(cartRepository)
	cartItemService := application.NewCartItemService(
		inventoryService,
		cartRepository,
		cartItemRepository,
		inventoryRepository,
		cartItemAttributeItemRepository,
		cartItemAttributeItemService,
		cartService,
		productRepository,
		cartItemHelper,
	)
	attributeService := application.NewAttributeService(attributeRepository)
	attributeItemService := application.NewAttributeItemService(attributeItemRepository)
	imageService := application.NewImageService(imageRepository)
	organizationService := application.NewOrganizationService(organizationRepository)
	locationService := application.NewLocationService(locationRepository)
	categoryService := application.NewCategoryService(categoryRepository)

	userHandler := handler.NewUserHandler(userService, authService)
	customerHandler := handler.NewCustomerHandler(customerService, authService)
	productHandler := handler.NewProductHandler(productService)
	orderHandler := handler.NewOrderHandler(orderService)
	cartHandler := handler.NewCartHandler(cartService)
	cartItemHandler := handler.NewCartItemHandler(cartItemService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	attributeHandler := handler.NewAttributeHandler(attributeService)
	attributeItemHandler := handler.NewAttributeItemHandler(attributeItemService)
	imageHandler := handler.NewImageHandler(imageService)
	organizationHandler := handler.NewOrganizationHandler(organizationService)
	locationHandler := handler.NewLocationHandler(locationService)

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers, Authorization"},
	}))

	router.Static("/images", "./images")

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/users/login", userHandler.Login)
	api.GET("/users/:id", userHandler.UserDetails)
	api.GET("/users", userHandler.UserFindAll)
	api.PUT("/users/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), userHandler.UpdateUser)
	api.DELETE("/users/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), userHandler.DeleteUser)

	api.POST("/customers", customerHandler.RegisterCustomer)
	api.POST("/customers/login", customerHandler.Login)
	api.GET("/customers/:id", customerHandler.CustomerDetails)
	api.GET("/customers", customerHandler.CustomerFindAll)
	api.PUT("/customers/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), customerHandler.UpdateCustomer)
	api.DELETE("/customers/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), customerHandler.DeleteCustomer)

	api.POST("/products", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), productHandler.Create)
	api.PUT("/products/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), productHandler.Update)
	api.DELETE("/products/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), productHandler.Delete)
	api.GET("/products/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), productHandler.FindById)
	api.GET("/products/code/:code", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), productHandler.FindByCode)
	api.GET("/products", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), productHandler.FindAll)

	api.POST("/images", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), imageHandler.Create)
	api.DELETE("/images/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), imageHandler.Delete)

	api.POST("/categories", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), categoryHandler.Create)
	api.PUT("/categories/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), categoryHandler.Update)
	api.DELETE("/categories/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), categoryHandler.Delete)
	api.GET("/categories/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), categoryHandler.FindById)
	api.GET("/categories", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), categoryHandler.FindAll)

	api.POST("/organizations", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), organizationHandler.Create)
	api.PUT("/organizations/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), organizationHandler.Update)
	api.DELETE("/organizations/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), organizationHandler.Delete)
	api.GET("/organizations/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), organizationHandler.FindById)
	api.GET("/organizations", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), organizationHandler.FindAll)

	api.POST("/locations", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), locationHandler.Create)
	api.PUT("/locations/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), locationHandler.Update)
	api.DELETE("/locations/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), locationHandler.Delete)
	api.GET("/locations/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), locationHandler.FindById)
	api.GET("/locations", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), locationHandler.FindAll)

	api.POST("/orders", authMiddleware.AuthMiddleware(authService, userService, customerService, "customer"), orderHandler.Create)
	api.POST("/place-order", authMiddleware.AuthMiddleware(authService, userService, customerService, "customer"), orderHandler.PlaceOrder)
	api.GET("/orders", orderHandler.FindAll)
	api.GET("/orders/:id", orderHandler.FindById)
	api.PUT("/orders/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "customer"), orderHandler.Update)

	api.POST("/order-items", authMiddleware.AuthMiddleware(authService, userService, customerService, "customer"), cartItemHandler.Create)
	api.PUT("/order-items/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "customer"), cartItemHandler.Update)
	api.DELETE("/order-items/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "customer"), cartItemHandler.Delete)

	api.POST("/attributes", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), attributeHandler.Create)
	api.PUT("/attributes/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), attributeHandler.Update)
	api.DELETE("/attributes/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), attributeHandler.Delete)
	api.GET("/attributes/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), attributeHandler.FindById)
	api.GET("/attributes", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), attributeHandler.FindAll)

	api.POST("/attribute-items", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), attributeItemHandler.Create)
	api.PUT("/attribute-items/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), attributeItemHandler.Update)
	api.DELETE("/attribute-items/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), attributeItemHandler.Delete)
	api.GET("/attribute-items/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), attributeItemHandler.FindById)
	api.GET("/attribute-items", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), attributeItemHandler.FindAll)

	api.GET("/carts/active", authMiddleware.AuthMiddleware(authService, userService, customerService, "customer"), cartHandler.CheckActiveCart)
	api.POST("/carts", authMiddleware.AuthMiddleware(authService, userService, customerService, "customer"), cartHandler.Create)
	api.GET("/carts/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "customer"), cartHandler.FindById)

	api.POST("/cart-items", authMiddleware.AuthMiddleware(authService, userService, customerService, "customer"), cartItemHandler.Create)
	api.PUT("/cart-items/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "customer"), cartItemHandler.Update)

	return router
}
