package controllers

type PlanService interface {
}

type PlanController struct {
	planService PlanService
}

func NewPlanController(planService PlanService) *PlanController {
	return &PlanController{
		planService: planService,
	}
}

func (c *PlanController) GetPlans() {
	
}
