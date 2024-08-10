package handlers

import (
	"context"
	"log"

	pb "github.com/ratheeshkumar/restaurant_admin_serviceV1/internal/pb"
)

func (a *AdminHandler) CreateMenu(ctx context.Context, p *pb.AMenuItem) (*pb.AdminResponse, error) {
	result, err := a.AdminService.CreateMenu(p)
	if err != nil {
		log.Println("error while creating menu")
		return nil, err
	}
	return result, err
}

func (a *AdminHandler) FetchByMenuID(ctx context.Context, p *pb.AMenuBYId) (*pb.AMenuItem, error) {
	result, err := a.AdminService.FetchMenuByID(p)

	if err != nil {
		log.Println("error while fetching menuby id")
		return nil, err
	}
	return result, err
}

func (a *AdminHandler) FetchByName(ctx context.Context, p *pb.AMenuBYName) (*pb.AMenuItem, error) {
	result, err := a.AdminService.FetchMenuByName(p)

	if err != nil {
		log.Println("error while fetching foodby name")
		return nil, err
	}
	return result, err
}

func (a *AdminHandler) FetchMenus(ctx context.Context, p *pb.AdminNoParam) (*pb.AMenuList, error) {
	result, err := a.AdminService.FetchAllMenu(p)

	if err != nil {
		log.Println("error while fetching whole menulist")
		return nil, err
	}
	return result, err
}
