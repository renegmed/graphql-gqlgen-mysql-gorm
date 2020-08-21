package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/renegmed/graphql-gqlgen-mysql-gorm/graph/generated"
	"github.com/renegmed/graphql-gqlgen-mysql-gorm/graph/model"
)

func (r *mutationResolver) UpdateItem(ctx context.Context, itemID int, item model.ItemInput) (*model.Item, error) {
	//log.Println("++++ Updated Item:\n\t", itemID)
	updatedItem := model.Item{
		ID:          itemID,
		ProductCode: item.ProductCode,
		ProductName: item.ProductName,
		Quantity:    item.Quantity,
		OrderID:     item.OrderID,
	}

	log.Println("++++ Updated item:\n\t", updatedItem)

	err := r.DB.Save(&updatedItem).Error
	if err != nil {
		return nil, err
	}
	return &updatedItem, nil
}

func (r *mutationResolver) CreateOrder(ctx context.Context, input model.OrderInput) (*model.Order, error) {
	order := model.Order{
		CustomerName: input.CustomerName,
		OrderAmount:  input.OrderAmount,
		Items:        mapItemsFromInput(input.Items),
	}

	log.Println("++++ Create order:\n\t", order)

	err := r.DB.Create(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *mutationResolver) UpdateOrder(ctx context.Context, orderID int, input model.OrderInput) (*model.Order, error) {
	updatedOrder := model.Order{
		ID:           orderID,
		CustomerName: input.CustomerName,
		OrderAmount:  input.OrderAmount,
		Items:        mapItemsFromInput(input.Items),
	}

	log.Println("++++ Updated order:\n\t", updatedOrder)

	err := r.DB.Save(&updatedOrder).Error
	if err != nil {
		return nil, err
	}
	return &updatedOrder, nil
}

func (r *mutationResolver) DeleteOrder(ctx context.Context, orderID int) (bool, error) {
	r.DB.Where("order_id = ?", orderID).Delete(&model.Item{})
	r.DB.Where("id = ?", orderID).Delete(&model.Order{})
	return true, nil
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	var orders []*model.Order
	// err := r.DB.Preload("Items").Find(&orders).Error
	err := r.DB.Set("gorm:auto_preload", true).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func mapItemsFromInput(itemsInput []*model.ItemInput) []*model.Item {
	var items []*model.Item
	for _, itemInput := range itemsInput {
		items = append(items, &model.Item{
			ProductCode: itemInput.ProductCode,
			ProductName: itemInput.ProductName,
			Quantity:    itemInput.Quantity,
			OrderID:     itemInput.OrderID,
		})
	}
	return items
}
