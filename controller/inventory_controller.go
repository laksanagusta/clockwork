package controller

import (
	"clockwork-server/helper"
	"clockwork-server/request"
	"clockwork-server/response"
	"clockwork-server/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const SAVE_INVENTORY_FAILED = "Failed to save inventory"
const GET_INVENTORY_FAILED = "Failed to get inventory"

type InventoryControllerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindByProductId(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type inventoryController struct {
	service service.InventoryService
}

func NewInventoryController(service service.InventoryService) InventoryControllerInterface {
	return &inventoryController{service}
}

func (_inventoryController *inventoryController) Create(c *gin.Context) {
	var input request.InventoryCreateInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse(SAVE_INVENTORY_FAILED, http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inventory, err := _inventoryController.service.Create(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse(SAVE_INVENTORY_FAILED, http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success save inventory !", http.StatusOK, helper.SUCCESS, response.FormatInventory(inventory))
	c.JSON(http.StatusOK, response)
}

func (_inventoryController *inventoryController) Update(c *gin.Context) {
	var inputID request.InventoryFindById
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse(SAVE_INVENTORY_FAILED, http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData request.InventoryUpdateInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse(SAVE_INVENTORY_FAILED, http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedInventory, err := _inventoryController.service.Update(inputID, inputData)
	if err != nil {
		response := helper.APIResponse(SAVE_INVENTORY_FAILED, http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success create inventory", http.StatusOK, helper.SUCCESS, response.FormatInventory(updatedInventory))
	c.JSON(http.StatusOK, response)
}

func (_inventoryController *inventoryController) FindById(c *gin.Context) {
	var input request.InventoryFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse(GET_INVENTORY_FAILED, http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inventory, err := _inventoryController.service.FindById(input.ID)
	if err != nil {
		response := helper.APIResponse(GET_INVENTORY_FAILED, http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get inventory !", http.StatusOK, helper.SUCCESS, response.FormatInventory(inventory))
	c.JSON(http.StatusOK, response)
}

func (_inventoryController *inventoryController) FindByProductId(c *gin.Context) {
	var input request.InventoryFindByProductId
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse(GET_INVENTORY_FAILED, http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inventorySingle, err := _inventoryController.service.FindByProductId(input.ProductID)
	if err != nil {
		response := helper.APIResponse(GET_INVENTORY_FAILED, http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get inventory !", http.StatusOK, helper.SUCCESS, response.FormatInventory(inventorySingle))
	c.JSON(http.StatusOK, response)
}

func (_inventoryController *inventoryController) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("page_size"))
	s := q.Get("q")
	inventorys, err := _inventoryController.service.FindAll(page, pageSize, s)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse(GET_INVENTORY_FAILED, http.StatusOK, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of inventorys", http.StatusOK, helper.SUCCESS, response.FormatInventories(inventorys))
	c.JSON(http.StatusOK, response)
}

func (_inventoryController *inventoryController) Delete(c *gin.Context) {
	var input request.InventoryFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete inventory", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inventory, err := _inventoryController.service.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete inventory", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete inventory !", http.StatusOK, helper.SUCCESS, response.FormatInventory(inventory))
	c.JSON(http.StatusOK, response)
}
