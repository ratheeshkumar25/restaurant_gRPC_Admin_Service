package interfaces

import (
	pb "github.com/ratheeshkumar/restaurant_admin_serviceV1/internal/pb"
)

type AdminServiceInter interface {
	AdminLoginService(p *pb.AdminRequest) (*pb.AdminResponse, error)
	CreateMenu(p *pb.AMenuItem) (*pb.AdminResponse, error)
	FetchMenuByID(p *pb.AMenuBYId) (*pb.AMenuItem, error)
	FetchMenuByName(p *pb.AMenuBYName) (*pb.AMenuItem, error)
	FetchAllMenu(p *pb.AdminNoParam) (*pb.AMenuList, error)
}
