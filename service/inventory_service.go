package service

import (
	"clockwork-server/model"
	"clockwork-server/repository"
	"clockwork-server/request"
	"errors"
)

type InventoryService interface {
	Create(request request.InventoryCreateInput) (model.Inventory, error)
	Update(inputID request.InventoryFindById, request request.InventoryUpdateInput) (model.Inventory, error)
	FindById(inventoryId int) (model.Inventory, error)
	FindByProductId(productId int) (model.Inventory, error)
	FindAll(page int, page_size int, q string) ([]model.Inventory, error)
	Delete(inventoryId int) (model.Inventory, error)
	syncInventoryAfterUpdateCartItem(productId int, cartItem model.CartItem, cartItemReq request.CartItemUpdateRequest) (model.Inventory, error)
	syncInventoryAfterDeleteItemFromCart(cartItem model.CartItem) (model.Inventory, error)
}

type inventoryService struct {
	repository   repository.InventoryRepository
	cartItemRepo repository.CartItemRepository
}

func NewInventoryService(repository repository.InventoryRepository, cartItemRepo repository.CartItemRepository) InventoryService {
	return &inventoryService{
		repository,
		cartItemRepo,
	}
}

const INVENTORY_ALREADY_EXIST_NOTIF = "create inventory failed, inventory already exist"

func (s *inventoryService) Create(request request.InventoryCreateInput) (model.Inventory, error) {
	inventory := model.Inventory{}
	inventory.StockQty = request.StockQty
	inventory.SalableQty = request.SalableQty
	inventory.ReservedQty = request.ReservedQty
	inventory.IsInStock = request.IsInStock

	// checkInventory, err := s.repository.FindByProductId(request.ProductID)
	// if err != nil {
	// 	return inventory, err
	// }

	// if checkInventory.ID != 0 {
	// 	return inventory, errors.New(INVENTORY_ALREADY_EXIST_NOTIF)
	// }

	newInventory, err := s.repository.Create(inventory)
	if err != nil {
		return newInventory, err
	}

	return newInventory, nil
}

func (s *inventoryService) Update(inputID request.InventoryFindById, request request.InventoryUpdateInput) (model.Inventory, error) {
	inventory, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return inventory, err
	}

	inventory.StockQty = request.StockQty
	inventory.IsInStock = request.IsInStock

	updatedInventory, err := s.repository.Update(inventory)
	if err != nil {
		return updatedInventory, err
	}

	return updatedInventory, nil

}

func (s *inventoryService) FindById(inventoryId int) (model.Inventory, error) {
	inventory, err := s.repository.FindById(inventoryId)
	if err != nil {
		return inventory, err
	}

	if inventory.ID == 0 {
		return inventory, errors.New("Inventory not found")
	}

	return inventory, nil
}

func (s *inventoryService) FindByProductId(productId int) (model.Inventory, error) {
	inventory, err := s.repository.FindByProductId(productId)
	if err != nil {
		return inventory, err
	}

	if inventory.ID == 0 {
		return inventory, errors.New("Inventory not found")
	}

	return inventory, nil
}

func (s *inventoryService) FindAll(page int, pageSize int, q string) ([]model.Inventory, error) {
	inventorys, err := s.repository.FindAll(page, pageSize, q)
	if err != nil {
		return inventorys, err
	}
	return inventorys, nil
}

func (s *inventoryService) Delete(inventoryId int) (model.Inventory, error) {
	inventory, err := s.repository.Delete(inventoryId)
	if err != nil {
		return inventory, err
	}

	return inventory, nil
}

func (s *inventoryService) syncInventoryAfterUpdateCartItem(productId int, cartItem model.CartItem, cartItemReq request.CartItemUpdateRequest) (model.Inventory, error) {
	inventory, err := s.FindByProductId(productId)
	if err != nil {
		return inventory, err
	}

	if cartItem.Qty > cartItemReq.Qty {
		inventory.ReservedQty -= cartItem.Qty - cartItemReq.Qty
		inventory.SalableQty += cartItem.Qty - cartItemReq.Qty
	} else if cartItem.Qty < cartItemReq.Qty {
		inventory.ReservedQty += cartItemReq.Qty - cartItem.Qty
		inventory.SalableQty -= cartItemReq.Qty - cartItem.Qty
	}

	updatedInventory, err := s.repository.Update(inventory)
	if err != nil {
		return inventory, err
	}

	return updatedInventory, err
}

func (s *inventoryService) syncInventoryAfterDeleteItemFromCart(cartItem model.CartItem) (model.Inventory, error) {
	inventory, err := s.FindByProductId(int(cartItem.ProductID))
	if err != nil {
		return inventory, err
	}

	inventory.ReservedQty -= cartItem.Qty
	inventory.SalableQty += cartItem.Qty

	updateInventory, err := s.repository.Update(inventory)
	if err != nil {
		return updateInventory, err
	}

	return updateInventory, nil
}
