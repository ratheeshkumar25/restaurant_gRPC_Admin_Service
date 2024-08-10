package handler

import (
	"context"
	"log"

	menupb "github.com/ratheeshkumar/restaurant_admin_serviceV1/internal/menu/pb"
	adminpb "github.com/ratheeshkumar/restaurant_admin_serviceV1/internal/pb"
)

func CreateMenuHandler(client menupb.MenuServiceClient, p *adminpb.AMenuItem) (*menupb.MenuResponse, error) {
	ctx := context.Background()

	response, err := client.CreateMenu(ctx, &menupb.MenuItem{
		Category:  p.Category,
		Name:      p.Name,
		Price:     p.Price,
		Foodimage: p.Foodimage,
		Duration:  p.Duration,
	})

	if err != nil {
		log.Printf("error while creating menuitems")
		return nil, err
	}
	return response, nil
}

func FetchMenuByIDHandler(client menupb.MenuServiceClient, p *adminpb.AMenuBYId) (*menupb.MenuItem, error) {
	ctx := context.Background()

	response, err := client.FetchByMenuID(ctx, &menupb.MenuID{Id: p.Id})
	if err != nil {
		log.Printf("error while loding menuby id ")
		return nil, err
	}
	return response, nil
}

func FetchMenuByNameHandler(client menupb.MenuServiceClient, p *adminpb.AMenuBYName) (*menupb.MenuItem, error) {
	ctx := context.Background()

	response, err := client.FetchByName(ctx, &menupb.FoodByName{Name: p.Name})
	if err != nil {
		log.Printf("error while loading menuby name")
		return nil, err
	}
	return response, nil
}

func FetchAllMenuHandler(client menupb.MenuServiceClient, p *adminpb.AdminNoParam) (*menupb.MenuList, error) {
	ctx := context.Background()

	response, err := client.FetchMenus(ctx, &menupb.NoParam{})
	if err != nil {
		log.Printf("error while loading menulist")
		return nil, err
	}
	return response, nil

}
