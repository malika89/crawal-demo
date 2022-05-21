package services


import (
	"crawl/handler/models"
)


type service struct {
	models.BaseModel
}
func GetService() *service  {
	var serviceApp service
	serviceApp.Table = "users"
	serviceApp.Model = models.User{}
	return &serviceApp
}


//Delete 支持批量删除
func(s *service) Delete() error {
	dbSession := s.GetSession()
	if _, err := dbSession.Where(s.Model).Delete(&models.User{}); err != nil {
		return err
	}
	return nil
}
