package core

import "discount-service/models"

type DiscountRuleModel struct {
	DiscountRule *models.DiscountRule `json:"discount_rule"`
	Attrs []DiscountRuleAttrsModel `json:"attrs"`
}

type DiscountRuleAttrsModel struct{
	DiscountAttributes *models.DiscountAttribute `json:"discount_attributes"`
	DiscountMetaAttributes []*models.DiscountMetaAttribute `json:"discount_meta_attributes"`
}