package app

import (
	"clockwork-server/auth"
	"clockwork-server/config"
	"clockwork-server/controller"
	"clockwork-server/middleware"
	"clockwork-server/repository"
	"clockwork-server/service"

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

	userRepository := repository.NewRepository(r.db)
	customerRepository := repository.NewCustomerRepository(r.db)
	addressRepository := repository.NewAddressRepository(r.db)
	productRepository := repository.NewProductRepository(r.db)
	orderRepository := repository.NewOrderRepository(r.db)
	inventoryRepository := repository.NewInventoryRepository(r.db)
	orderItemRepository := repository.NewOrderItemRepository(r.db)
	categoryRepository := repository.NewCategoryRepository(r.db)

	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService, authService)

	customerService := service.NewCustomerService(customerRepository)
	customerController := controller.NewCustomerController(customerService, authService)

	addressService := service.NewAddressService(addressRepository, customerRepository)
	addressController := controller.NewAddressController(addressService)

	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	midtransService := service.NewMidtransService(config.GetConfig())

	orderService := service.NewOrderService(orderRepository, midtransService)
	orderController := controller.NewOrderController(orderService)

	inventoryService := service.NewInventoryService(inventoryRepository, orderItemRepository)
	inventoryController := controller.NewInventoryController(inventoryService)

	orderItemService := service.NewOrderItemService(inventoryService, orderRepository, orderItemRepository, inventoryRepository)
	orderItemController := controller.NewOrderItemController(orderItemService)

	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers, Authorization"},
	}))

	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	api.POST("/users", userController.RegisterUser)
	api.POST("/users/login", userController.Login)
	api.GET("/users/:id", userController.UserDetails)
	api.GET("/users", userController.UserFindAll)
	api.PUT("/users/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), userController.UpdateUser)
	api.DELETE("/users/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), userController.DeleteUser)

	api.POST("/customers", customerController.RegisterCustomer)
	api.POST("/customers/login", customerController.Login)
	api.GET("/customers/:id", customerController.CustomerDetails)
	api.GET("/customers", customerController.CustomerFindAll)
	api.PUT("/customers/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), customerController.UpdateCustomer)
	api.DELETE("/customers/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), customerController.DeleteCustomer)

	api.POST("/adress", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), addressController.Create)
	api.PUT("/adress/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), addressController.Update)
	api.DELETE("/adress/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), addressController.Delete)
	api.GET("/adress/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), addressController.FindById)
	api.GET("/adress", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), addressController.FindAll)

	api.POST("/products", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), productController.Create)
	api.PUT("/products/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), productController.Update)
	api.DELETE("/products/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), productController.Delete)
	api.GET("/products/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), productController.FindById)
	api.GET("/products/code/:code", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), productController.FindByCode)
	api.GET("/products", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), productController.FindAll)

	api.POST("/categories", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), categoryController.Create)
	api.PUT("/categories/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), categoryController.Update)
	api.DELETE("/categories/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), categoryController.Delete)
	api.GET("/categories/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), categoryController.FindById)
	api.GET("/categories", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), categoryController.FindAll)

	api.POST("/orders", authMiddleware.AuthMiddleware(authService, userService, customerService, "customer"), orderController.Create)
	api.GET("/orders", orderController.FindAll)
	api.PUT("/orders/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "customer"), orderController.Update)

	api.POST("/inventories", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), inventoryController.Create)
	api.PUT("/inventories/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), inventoryController.Update)
	api.DELETE("/inventories/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), inventoryController.Delete)
	api.GET("/inventories/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), inventoryController.FindById)
	api.GET("/inventories/product-id/:code", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), inventoryController.FindByProductId)
	api.GET("/inventories", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), inventoryController.FindAll)

	api.POST("/order-items", authMiddleware.AuthMiddleware(authService, userService, customerService, "customer"), orderItemController.Create)
	api.PUT("/order-items/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "customer"), orderItemController.Update)
	api.DELETE("/order-items/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "customer"), orderItemController.Delete)

	return router
}
