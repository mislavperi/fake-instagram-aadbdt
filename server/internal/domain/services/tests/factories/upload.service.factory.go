package factories

import (
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/services"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/services/tests/mocks"
)

func NewDailyUploadService(repo *mocks.DailyUploadRepository, pls *mocks.PlanLogServiceUpload, ps *mocks.PlanServiceUpload, us *mocks.UserServiceUpload, ls *mocks.LogServiceUpload) *services.DailyUploadService {
	return services.NewDailyUploadService(
		repo,
		pls,
		ps,
		us,
		ls,
	)
}
