package core

import (
	"context"
	"strconv"
	"strings"
	"time"
)

func (dc *DiscountCoreContext) CalculateDiscount(ctx context.Context, accountId string, amount int) () {
	discountRule, err := dc.GetDiscountRule(ctx, accountId)
	if err != nil {
		return
	}

	var points []int
	var clubIds []string

	if discountRule.DiscountRule.IsClubLists {
		//TODO fill clubIds with merchant clubs

	} else {
		for _, attr := range discountRule.Attrs {
			if attr.DiscountAttributes.Name == "ClubList" {
				for _, value := range attr.DiscountMetaAttributes {
					clubIds = append(clubIds, value.Value)
				}
			}
		}
	}
	for _, attr := range discountRule.Attrs {
		switch attr.DiscountAttributes.Name {
		case "TotalNumberOfTransactionsInMerchant":
			for _, value := range attr.DiscountMetaAttributes {
				if value.Name == "min-max" {
					xy := strings.Split(value.Value, "-")
					x, err := strconv.Atoi(xy[0])
					if err != nil {
						continue
					}
					y, err := strconv.Atoi(xy[1])
					if err != nil {
						continue
					}
					cal
					if cal >= x && cal < y {
						points = append(points, value.Point)
						continue
					}
				}
			}
		case "TotalAmountOfTransactionInMerchant":
			for _, value := range attr.DiscountMetaAttributes {
				if value.Name == "min-max" {
					xy := strings.Split(value.Value, "-")
					x, err := strconv.Atoi(xy[0])
					if err != nil {
						continue
					}
					y, err := strconv.Atoi(xy[1])
					if err != nil {
						continue
					}
					cal
					if cal >= x && cal < y {
						points = append(points, value.Point)
						continue
					}
				}
			}
		case "LastTransactionDateInMerchant":
			for _, value := range attr.DiscountMetaAttributes {
				if value.Name == "min-max" {
					xy := strings.Split(value.Value, "-")
					layout := "2006-01-02T15:04:05.000Z"
					x, err := time.Parse(layout, xy[0])
					if err != nil {
						continue
					}
					y, err := time.Parse(layout, xy[1])
					if err != nil {
						continue
					}
					cal
					if cal.After(x) && cal.Before(y) {
						points = append(points, value.Point)
						continue
					}
				}
			}
		case "PaymentTypeInMerchant":
			for _, value := range attr.DiscountMetaAttributes {
				cal
				if value.Name == "online" && cal == "online" {
					points = append(points, value.Point)
					continue
				}
				if value.Name == "offline" && cal == "offline" {
					points = append(points, value.Point)
					continue
				}
			}
		case "TotalNumberOfTransactionsInClub":
			for _, value := range attr.DiscountMetaAttributes {
				if value.Name == "min-max" {
					xy := strings.Split(value.Value, "-")
					x, err := strconv.Atoi(xy[0])
					if err != nil {
						continue
					}
					y, err := strconv.Atoi(xy[1])
					if err != nil {
						continue
					}
					cal
					if cal >= x && cal < y {
						points = append(points, value.Point)
						continue
					}
				}
			}
		case "TotalAmountOfTransactionInClub":
			for _, value := range attr.DiscountMetaAttributes {
				if value.Name == "min-max" {
					xy := strings.Split(value.Value, "-")
					x, err := strconv.Atoi(xy[0])
					if err != nil {
						continue
					}
					y, err := strconv.Atoi(xy[1])
					if err != nil {
						continue
					}
					cal
					if cal >= x && cal < y {
						points = append(points, value.Point)
						continue
					}
				}
			}
		case "LastTransactionDateInClub":
			for _, value := range attr.DiscountMetaAttributes {
				if value.Name == "min-max" {
					xy := strings.Split(value.Value, "-")
					layout := "2006-01-02T15:04:05.000Z"
					x, err := time.Parse(layout, xy[0])
					if err != nil {
						continue
					}
					y, err := time.Parse(layout, xy[1])
					if err != nil {
						continue
					}
					cal
					if cal.After(x) && cal.Before(y) {
						points = append(points, value.Point)
						continue
					}
				}
			}
		case "PaymentTypeInClub":
			for _, value := range attr.DiscountMetaAttributes {
				cal
				if value.Name == "online" && cal == "online" {
					points = append(points, value.Point)
					continue
				}
				if value.Name == "offline" && cal == "offline" {
					points = append(points, value.Point)
					continue
				}
			}
		case "GroupList":
		case "Points":

		default:

		}
	}
}
