package controllers

import (
	"discount-service/core"
	"discount-service/utils"
	"github.com/labstack/echo"
	"net/http"
)


func (cc *ControllerContext) CreateDiscountRule(ctx echo.Context) error {
	req := new(core.DiscountRuleModel)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.StandardHttpErrorResponse{Message: err.Error()})
	}
	var discountCoreContext = core.DiscountCoreContext{ApplicationContext: cc.ApplicationContext}

	ok, err := discountCoreContext.ValidateDiscountCreationRules(ctx.Request().Context(), *req)
	if err != nil || !ok {
		return ctx.JSON(http.StatusBadRequest, utils.StandardHttpErrorResponse{Message: "invalid discount rules data"})
	}

	err = discountCoreContext.CreateDiscountRule(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.StandardHttpErrorResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusNoContent, "")
}

func (cc *ControllerContext) GetDiscountRule(ctx echo.Context) error {
	accountId := ctx.Param("account_id")

	var discountCoreContext = core.DiscountCoreContext{ApplicationContext: cc.ApplicationContext}

	discountRule, err := discountCoreContext.GetDiscountRule(ctx.Request().Context(), accountId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.StandardHttpErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, utils.StandardHttpResponse{Result: discountRule})
}
