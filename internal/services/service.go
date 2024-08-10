package services

import (
	"errors"

	menupb "github.com/ratheeshkumar/restaurant_admin_serviceV1/internal/menu/pb"
	pb "github.com/ratheeshkumar/restaurant_admin_serviceV1/internal/pb"
	inter "github.com/ratheeshkumar/restaurant_admin_serviceV1/internal/repositories/interfaces"
	"github.com/ratheeshkumar/restaurant_admin_serviceV1/internal/services/interfaces"
	"github.com/ratheeshkumar/restaurant_admin_serviceV1/internal/utils"
)

type AdminService struct {
	AdminRepo inter.AdminRepoInter
	client    menupb.MenuServiceClient
}

// AdminLogin implements interfaces.AdminServiceInter.
func (a *AdminService) AdminLoginService(p *pb.AdminRequest) (*pb.AdminResponse, error) {
	admin, err := a.AdminRepo.FindAdmin(p.Username)
	if err != nil {
		return nil, err
	}

	if admin.Password != p.Password {
		return nil, errors.New("incorrect password")
	}

	token, err := utils.GenerateToken(p.Username)
	if err != nil {
		return nil, err
	}
	response := pb.AdminResponse{
		Status:  "Success",
		Error:   "",
		Message: "Logged successfully,Token:" + token,
	}

	return &response, nil
}

func NewAdminService(repo inter.AdminRepoInter, client menupb.MenuServiceClient) interfaces.AdminServiceInter {
	return &AdminService{
		AdminRepo: repo,
		client:    client,
	}
}
