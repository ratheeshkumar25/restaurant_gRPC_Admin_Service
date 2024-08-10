package interfaces

import "github.com/ratheeshkumar/restaurant_admin_serviceV1/internal/model"

type AdminRepoInter interface {
	FindAdmin(username string) (*model.AdminModel, error)
}
