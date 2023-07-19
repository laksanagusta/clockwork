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
	cartRepository := repository.NewCartRepository(r.db)
	cartItemRepository := repository.NewCartItemRepository(r.db)
	categoryRepository := repository.NewCategoryRepository(r.db)
	attributeRepository := repository.NewAttributeRepository(r.db)
	attributeItemRepository := repository.NewAttributeItemRepository(r.db)
	cartItemAttributeItemRepository := repository.NewCartItemAttributeItemRepository(r.db)

	userService := service.NewUserService(userRepository)
	customerService := service.NewCustomerService(customerRepository)
	addressService := service.NewAddressService(addressRepository, customerRepository)
	productService := service.NewProductService(productRepository, categoryRepository, inventoryRepository)
	midtransService := service.NewMidtransService(config.GetConfig(), orderRepository)
	orderService := service.NewOrderService(orderRepository, midtransService)
	inventoryService := service.NewInventoryService(inventoryRepository, cartItemRepository)
	cartItemAttributeItemService := service.NewCartItemAttributeItemService(cartItemAttributeItemRepository)
	cartService := service.NewCartService(cartRepository)
	cartItemService := service.NewCartItemService(
		inventoryService,
		cartRepository,
		cartItemRepository,
		inventoryRepository,
		cartItemAttributeItemRepository,
		cartItemAttributeItemService,
		cartService,
	)
	attributeService := service.NewAttributeService(attributeRepository)
	attributeItemService := service.NewAttributeItemService(attributeItemRepository)

	userController := controller.NewUserController(userService, authService)
	customerController := controller.NewCustomerController(customerService, authService)
	addressController := controller.NewAddressController(addressService)
	productController := controller.NewProductController(productService)
	orderController := controller.NewOrderController(orderService)
	inventoryController := controller.NewInventoryController(inventoryService)
	cartController := controller.NewCartController(cartService)
	cartItemController := controller.NewCartItemController(cartItemService)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)
	attributeController := controller.NewAttributeController(attributeService)
	attributeItemController := controller.NewAttributeItemController(attributeItemService)

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

	api.POST("/order-items", authMiddleware.AuthMiddleware(authService, userService, customerService, "customer"), cartItemController.Create)
	api.PUT("/order-items/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "customer"), cartItemController.Update)
	api.DELETE("/order-items/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "customer"), cartItemController.Delete)

	api.POST("/attributes", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), attributeController.Create)
	api.PUT("/attributes/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), attributeController.Update)
	api.DELETE("/attributes/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), attributeController.Delete)
	api.GET("/attributes/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), attributeController.FindById)
	api.GET("/attributes", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), attributeController.FindAll)

	api.POST("/attribute-items", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), attributeItemController.Create)
	api.PUT("/attribute-items/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), attributeItemController.Update)
	api.DELETE("/attribute-items/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), attributeItemController.Delete)
	api.GET("/attribute-items/:id", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), attributeItemController.FindById)
	api.GET("/attribute-items", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), attributeItemController.FindAll)

	api.GET("/carts/active", authMiddleware.AuthMiddleware(authService, userService, customerService, "user"), cartController.CheckActiveCart)

	return router
}
