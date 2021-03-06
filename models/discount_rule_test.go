// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testDiscountRules(t *testing.T) {
	t.Parallel()

	query := DiscountRules()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDiscountRulesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountRule{}
	if err = randomize.Struct(seed, o, discountRuleDBTypes, true, discountRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DiscountRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscountRulesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountRule{}
	if err = randomize.Struct(seed, o, discountRuleDBTypes, true, discountRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DiscountRules().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DiscountRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscountRulesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountRule{}
	if err = randomize.Struct(seed, o, discountRuleDBTypes, true, discountRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DiscountRuleSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DiscountRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscountRulesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountRule{}
	if err = randomize.Struct(seed, o, discountRuleDBTypes, true, discountRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DiscountRuleExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DiscountRule exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DiscountRuleExists to return true, but got false.")
	}
}

func testDiscountRulesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountRule{}
	if err = randomize.Struct(seed, o, discountRuleDBTypes, true, discountRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	discountRuleFound, err := FindDiscountRule(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if discountRuleFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDiscountRulesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountRule{}
	if err = randomize.Struct(seed, o, discountRuleDBTypes, true, discountRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DiscountRules().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDiscountRulesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountRule{}
	if err = randomize.Struct(seed, o, discountRuleDBTypes, true, discountRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DiscountRules().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDiscountRulesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discountRuleOne := &DiscountRule{}
	discountRuleTwo := &DiscountRule{}
	if err = randomize.Struct(seed, discountRuleOne, discountRuleDBTypes, false, discountRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}
	if err = randomize.Struct(seed, discountRuleTwo, discountRuleDBTypes, false, discountRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = discountRuleOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = discountRuleTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DiscountRules().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDiscountRulesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	discountRuleOne := &DiscountRule{}
	discountRuleTwo := &DiscountRule{}
	if err = randomize.Struct(seed, discountRuleOne, discountRuleDBTypes, false, discountRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}
	if err = randomize.Struct(seed, discountRuleTwo, discountRuleDBTypes, false, discountRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = discountRuleOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = discountRuleTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DiscountRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func discountRuleBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DiscountRule) error {
	*o = DiscountRule{}
	return nil
}

func discountRuleAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DiscountRule) error {
	*o = DiscountRule{}
	return nil
}

func discountRuleAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DiscountRule) error {
	*o = DiscountRule{}
	return nil
}

func discountRuleBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DiscountRule) error {
	*o = DiscountRule{}
	return nil
}

func discountRuleAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DiscountRule) error {
	*o = DiscountRule{}
	return nil
}

func discountRuleBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DiscountRule) error {
	*o = DiscountRule{}
	return nil
}

func discountRuleAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DiscountRule) error {
	*o = DiscountRule{}
	return nil
}

func discountRuleBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DiscountRule) error {
	*o = DiscountRule{}
	return nil
}

func discountRuleAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DiscountRule) error {
	*o = DiscountRule{}
	return nil
}

func testDiscountRulesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DiscountRule{}
	o := &DiscountRule{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, discountRuleDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DiscountRule object: %s", err)
	}

	AddDiscountRuleHook(boil.BeforeInsertHook, discountRuleBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	discountRuleBeforeInsertHooks = []DiscountRuleHook{}

	AddDiscountRuleHook(boil.AfterInsertHook, discountRuleAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	discountRuleAfterInsertHooks = []DiscountRuleHook{}

	AddDiscountRuleHook(boil.AfterSelectHook, discountRuleAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	discountRuleAfterSelectHooks = []DiscountRuleHook{}

	AddDiscountRuleHook(boil.BeforeUpdateHook, discountRuleBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	discountRuleBeforeUpdateHooks = []DiscountRuleHook{}

	AddDiscountRuleHook(boil.AfterUpdateHook, discountRuleAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	discountRuleAfterUpdateHooks = []DiscountRuleHook{}

	AddDiscountRuleHook(boil.BeforeDeleteHook, discountRuleBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	discountRuleBeforeDeleteHooks = []DiscountRuleHook{}

	AddDiscountRuleHook(boil.AfterDeleteHook, discountRuleAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	discountRuleAfterDeleteHooks = []DiscountRuleHook{}

	AddDiscountRuleHook(boil.BeforeUpsertHook, discountRuleBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	discountRuleBeforeUpsertHooks = []DiscountRuleHook{}

	AddDiscountRuleHook(boil.AfterUpsertHook, discountRuleAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	discountRuleAfterUpsertHooks = []DiscountRuleHook{}
}

func testDiscountRulesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountRule{}
	if err = randomize.Struct(seed, o, discountRuleDBTypes, true, discountRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DiscountRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscountRulesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountRule{}
	if err = randomize.Struct(seed, o, discountRuleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(discountRuleColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DiscountRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscountRuleToManyDiscountAttributes(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DiscountRule
	var b, c DiscountAttribute

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discountRuleDBTypes, true, discountRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, discountAttributeDBTypes, false, discountAttributeColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, discountAttributeDBTypes, false, discountAttributeColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.DiscountRuleID, a.ID)
	queries.Assign(&c.DiscountRuleID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.DiscountAttributes().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.DiscountRuleID, b.DiscountRuleID) {
			bFound = true
		}
		if queries.Equal(v.DiscountRuleID, c.DiscountRuleID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DiscountRuleSlice{&a}
	if err = a.L.LoadDiscountAttributes(ctx, tx, false, (*[]*DiscountRule)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.DiscountAttributes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.DiscountAttributes = nil
	if err = a.L.LoadDiscountAttributes(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.DiscountAttributes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testDiscountRuleToManyAddOpDiscountAttributes(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DiscountRule
	var b, c, d, e DiscountAttribute

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discountRuleDBTypes, false, strmangle.SetComplement(discountRulePrimaryKeyColumns, discountRuleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DiscountAttribute{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, discountAttributeDBTypes, false, strmangle.SetComplement(discountAttributePrimaryKeyColumns, discountAttributeColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*DiscountAttribute{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddDiscountAttributes(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.DiscountRuleID) {
			t.Error("foreign key was wrong value", a.ID, first.DiscountRuleID)
		}
		if !queries.Equal(a.ID, second.DiscountRuleID) {
			t.Error("foreign key was wrong value", a.ID, second.DiscountRuleID)
		}

		if first.R.DiscountRule != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.DiscountRule != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.DiscountAttributes[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.DiscountAttributes[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.DiscountAttributes().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDiscountRuleToManySetOpDiscountAttributes(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DiscountRule
	var b, c, d, e DiscountAttribute

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discountRuleDBTypes, false, strmangle.SetComplement(discountRulePrimaryKeyColumns, discountRuleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DiscountAttribute{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, discountAttributeDBTypes, false, strmangle.SetComplement(discountAttributePrimaryKeyColumns, discountAttributeColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetDiscountAttributes(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.DiscountAttributes().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetDiscountAttributes(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.DiscountAttributes().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.DiscountRuleID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.DiscountRuleID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.DiscountRuleID) {
		t.Error("foreign key was wrong value", a.ID, d.DiscountRuleID)
	}
	if !queries.Equal(a.ID, e.DiscountRuleID) {
		t.Error("foreign key was wrong value", a.ID, e.DiscountRuleID)
	}

	if b.R.DiscountRule != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.DiscountRule != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.DiscountRule != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.DiscountRule != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.DiscountAttributes[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.DiscountAttributes[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testDiscountRuleToManyRemoveOpDiscountAttributes(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DiscountRule
	var b, c, d, e DiscountAttribute

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discountRuleDBTypes, false, strmangle.SetComplement(discountRulePrimaryKeyColumns, discountRuleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DiscountAttribute{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, discountAttributeDBTypes, false, strmangle.SetComplement(discountAttributePrimaryKeyColumns, discountAttributeColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddDiscountAttributes(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.DiscountAttributes().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveDiscountAttributes(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.DiscountAttributes().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.DiscountRuleID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.DiscountRuleID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.DiscountRule != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.DiscountRule != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.DiscountRule != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.DiscountRule != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.DiscountAttributes) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.DiscountAttributes[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.DiscountAttributes[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testDiscountRulesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountRule{}
	if err = randomize.Struct(seed, o, discountRuleDBTypes, true, discountRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDiscountRulesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountRule{}
	if err = randomize.Struct(seed, o, discountRuleDBTypes, true, discountRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DiscountRuleSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDiscountRulesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountRule{}
	if err = randomize.Struct(seed, o, discountRuleDBTypes, true, discountRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DiscountRules().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	discountRuleDBTypes = map[string]string{`ID`: `int`, `RuleType`: `enum('max','group','sorted','or','and')`, `UserID`: `varchar`, `IsPercentage`: `tinyint`, `IsClubLists`: `tinyint`}
	_                   = bytes.MinRead
)

func testDiscountRulesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(discountRulePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(discountRuleAllColumns) == len(discountRulePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DiscountRule{}
	if err = randomize.Struct(seed, o, discountRuleDBTypes, true, discountRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DiscountRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, discountRuleDBTypes, true, discountRulePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDiscountRulesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(discountRuleAllColumns) == len(discountRulePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DiscountRule{}
	if err = randomize.Struct(seed, o, discountRuleDBTypes, true, discountRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DiscountRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, discountRuleDBTypes, true, discountRulePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(discountRuleAllColumns, discountRulePrimaryKeyColumns) {
		fields = discountRuleAllColumns
	} else {
		fields = strmangle.SetComplement(
			discountRuleAllColumns,
			discountRulePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := DiscountRuleSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDiscountRulesUpsert(t *testing.T) {
	t.Parallel()

	if len(discountRuleAllColumns) == len(discountRulePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLDiscountRuleUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DiscountRule{}
	if err = randomize.Struct(seed, &o, discountRuleDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DiscountRule: %s", err)
	}

	count, err := DiscountRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, discountRuleDBTypes, false, discountRulePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscountRule struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DiscountRule: %s", err)
	}

	count, err = DiscountRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
