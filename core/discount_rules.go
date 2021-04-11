package core

import (
	"context"
	"discount-service/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)


func (dc *DiscountCoreContext) ValidateDiscountCreationRules(ctx context.Context, input DiscountRuleModel) (bool, error) {
	return true, nil
}

func (dc *DiscountCoreContext) CreateDiscountRule(ctx context.Context, discountRule *DiscountRuleModel) error {
	err := discountRule.DiscountRule.Insert(ctx, dc.MysqlDb, boil.Infer())
	if err != nil {
		return err
	}

	for _, value := range discountRule.Attrs {
		err = discountRule.DiscountRule.AddDiscountAttributes(ctx, dc.MysqlDb, true, value.DiscountAttributes)
		if err != nil {
			return err
		}
		err = value.DiscountAttributes.AddDiscountMetaAttributes(ctx, dc.MysqlDb, true, value.DiscountMetaAttributes...)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dc *DiscountCoreContext) GetDiscountRule(ctx context.Context, accountId string) (DiscountRuleModel, error) {
	discountRule, err := models.DiscountRules(
		models.DiscountRuleWhere.UserID.EQ(accountId),
		qm.Load(qm.Rels(models.DiscountRuleRels.DiscountAttributes, models.DiscountAttributeRels.DiscountMetaAttributes)),
	).One(ctx, dc.MysqlDb)
	if err != nil {
		return DiscountRuleModel{}, nil
	}
	var discountRuleModel DiscountRuleModel
	discountRuleModel.DiscountRule = discountRule
	for _, discountAttribute := range discountRule.R.DiscountAttributes {
		var attr DiscountRuleAttrsModel
		attr.DiscountAttributes = discountAttribute
		for _, discountMetaAttribute := range discountAttribute.R.DiscountMetaAttributes {
			attr.DiscountMetaAttributes = append(attr.DiscountMetaAttributes, discountMetaAttribute)
		}
		discountRuleModel.Attrs = append(discountRuleModel.Attrs, attr)
	}
	return discountRuleModel, nil
}