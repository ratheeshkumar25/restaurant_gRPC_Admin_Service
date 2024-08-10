package services

import (
	menupb "github.com/ratheeshkumar/restaurant_admin_serviceV1/internal/menu/handler"
	pb "github.com/ratheeshkumar/restaurant_admin_serviceV1/internal/pb"
)

// CreateMenu implements interfaces.AdminServiceInter.
func (a *AdminService) CreateMenu(p *pb.AMenuItem) (*pb.AdminResponse, error) {
	result, err := menupb.CreateMenuHandler(a.client, p)
	if err != nil {
		return nil, err
	}

	return &pb.AdminResponse{
		Status:  "Success",
		Error:   result.Error,
		Message: result.Message,
	}, nil
}

// FetchAllMenu implements interfaces.AdminServiceInter.
func (a *AdminService) FetchAllMenu(p *pb.AdminNoParam) (*pb.AMenuList, error) {
	result, err := menupb.FetchAllMenuHandler(a.client, p)
	if err != nil {
		return nil, err
	}
	var menu []*pb.AMenuItem
	for _, i := range result.Item {
		menu = append(menu, &pb.AMenuItem{
			Id:        i.Id,
			Category:  i.Category,
			Name:      i.Name,
			Foodimage: i.Foodimage,
			Duration:  i.Duration,
		})
	}
	return &pb.AMenuList{
		Menus: menu,
	}, nil
}

// FetchMenuByID implements interfaces.AdminServiceInter.
func (a *AdminService) FetchMenuByID(p *pb.AMenuBYId) (*pb.AMenuItem, error) {
	result, err := menupb.FetchMenuByIDHandler(a.client, p)
	if err != nil {
		return nil, err
	}
	return &pb.AMenuItem{
		Id:        result.Id,
		Category:  result.Category,
		Name:      result.Name,
		Foodimage: result.Foodimage,
		Price:     result.Price,
		Duration:  result.Duration,
	}, nil
}

// FetchMenuByName implements interfaces.AdminServiceInter.
func (a *AdminService) FetchMenuByName(p *pb.AMenuBYName) (*pb.AMenuItem, error) {
	result, err := menupb.FetchMenuByNameHandler(a.client, p)
	if err != nil {
		return nil, err
	}
	return &pb.AMenuItem{
		Id:        result.Id,
		Category:  result.Category,
		Name:      result.Name,
		Foodimage: result.Foodimage,
		Price:     result.Price,
		Duration:  result.Duration,
	}, nil
}
