// Code generated by SQLBoiler 4.19.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testBuildings(t *testing.T) {
	t.Parallel()

	query := Buildings()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBuildingsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Building{}
	if err = randomize.Struct(seed, o, buildingDBTypes, true, buildingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
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

	count, err := Buildings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBuildingsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Building{}
	if err = randomize.Struct(seed, o, buildingDBTypes, true, buildingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Buildings().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Buildings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBuildingsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Building{}
	if err = randomize.Struct(seed, o, buildingDBTypes, true, buildingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BuildingSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Buildings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBuildingsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Building{}
	if err = randomize.Struct(seed, o, buildingDBTypes, true, buildingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BuildingExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Building exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BuildingExists to return true, but got false.")
	}
}

func testBuildingsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Building{}
	if err = randomize.Struct(seed, o, buildingDBTypes, true, buildingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	buildingFound, err := FindBuilding(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if buildingFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBuildingsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Building{}
	if err = randomize.Struct(seed, o, buildingDBTypes, true, buildingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Buildings().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBuildingsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Building{}
	if err = randomize.Struct(seed, o, buildingDBTypes, true, buildingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Buildings().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBuildingsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	buildingOne := &Building{}
	buildingTwo := &Building{}
	if err = randomize.Struct(seed, buildingOne, buildingDBTypes, false, buildingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}
	if err = randomize.Struct(seed, buildingTwo, buildingDBTypes, false, buildingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = buildingOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = buildingTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Buildings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBuildingsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	buildingOne := &Building{}
	buildingTwo := &Building{}
	if err = randomize.Struct(seed, buildingOne, buildingDBTypes, false, buildingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}
	if err = randomize.Struct(seed, buildingTwo, buildingDBTypes, false, buildingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = buildingOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = buildingTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Buildings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func buildingBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Building) error {
	*o = Building{}
	return nil
}

func buildingAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Building) error {
	*o = Building{}
	return nil
}

func buildingAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Building) error {
	*o = Building{}
	return nil
}

func buildingBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Building) error {
	*o = Building{}
	return nil
}

func buildingAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Building) error {
	*o = Building{}
	return nil
}

func buildingBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Building) error {
	*o = Building{}
	return nil
}

func buildingAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Building) error {
	*o = Building{}
	return nil
}

func buildingBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Building) error {
	*o = Building{}
	return nil
}

func buildingAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Building) error {
	*o = Building{}
	return nil
}

func testBuildingsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Building{}
	o := &Building{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, buildingDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Building object: %s", err)
	}

	AddBuildingHook(boil.BeforeInsertHook, buildingBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	buildingBeforeInsertHooks = []BuildingHook{}

	AddBuildingHook(boil.AfterInsertHook, buildingAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	buildingAfterInsertHooks = []BuildingHook{}

	AddBuildingHook(boil.AfterSelectHook, buildingAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	buildingAfterSelectHooks = []BuildingHook{}

	AddBuildingHook(boil.BeforeUpdateHook, buildingBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	buildingBeforeUpdateHooks = []BuildingHook{}

	AddBuildingHook(boil.AfterUpdateHook, buildingAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	buildingAfterUpdateHooks = []BuildingHook{}

	AddBuildingHook(boil.BeforeDeleteHook, buildingBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	buildingBeforeDeleteHooks = []BuildingHook{}

	AddBuildingHook(boil.AfterDeleteHook, buildingAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	buildingAfterDeleteHooks = []BuildingHook{}

	AddBuildingHook(boil.BeforeUpsertHook, buildingBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	buildingBeforeUpsertHooks = []BuildingHook{}

	AddBuildingHook(boil.AfterUpsertHook, buildingAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	buildingAfterUpsertHooks = []BuildingHook{}
}

func testBuildingsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Building{}
	if err = randomize.Struct(seed, o, buildingDBTypes, true, buildingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Buildings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBuildingsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Building{}
	if err = randomize.Struct(seed, o, buildingDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(strmangle.SetMerge(buildingPrimaryKeyColumns, buildingColumnsWithoutDefault)...)); err != nil {
		t.Error(err)
	}

	count, err := Buildings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBuildingToManyApartments(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Building
	var b, c Apartment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildingDBTypes, true, buildingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, apartmentDBTypes, false, apartmentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, apartmentDBTypes, false, apartmentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.BuildingID = a.ID
	c.BuildingID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Apartments().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.BuildingID == b.BuildingID {
			bFound = true
		}
		if v.BuildingID == c.BuildingID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BuildingSlice{&a}
	if err = a.L.LoadApartments(ctx, tx, false, (*[]*Building)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Apartments); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Apartments = nil
	if err = a.L.LoadApartments(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Apartments); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBuildingToManyAddOpApartments(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Building
	var b, c, d, e Apartment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildingDBTypes, false, strmangle.SetComplement(buildingPrimaryKeyColumns, buildingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Apartment{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, apartmentDBTypes, false, strmangle.SetComplement(apartmentPrimaryKeyColumns, apartmentColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Apartment{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddApartments(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.BuildingID {
			t.Error("foreign key was wrong value", a.ID, first.BuildingID)
		}
		if a.ID != second.BuildingID {
			t.Error("foreign key was wrong value", a.ID, second.BuildingID)
		}

		if first.R.Building != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Building != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Apartments[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Apartments[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Apartments().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testBuildingsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Building{}
	if err = randomize.Struct(seed, o, buildingDBTypes, true, buildingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
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

func testBuildingsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Building{}
	if err = randomize.Struct(seed, o, buildingDBTypes, true, buildingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BuildingSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBuildingsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Building{}
	if err = randomize.Struct(seed, o, buildingDBTypes, true, buildingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Buildings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	buildingDBTypes = map[string]string{`ID`: `bigint`, `Name`: `text`, `Address`: `text`}
	_               = bytes.MinRead
)

func testBuildingsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(buildingPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(buildingAllColumns) == len(buildingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Building{}
	if err = randomize.Struct(seed, o, buildingDBTypes, true, buildingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Buildings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, buildingDBTypes, true, buildingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBuildingsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(buildingAllColumns) == len(buildingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Building{}
	if err = randomize.Struct(seed, o, buildingDBTypes, true, buildingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Buildings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, buildingDBTypes, true, buildingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(buildingAllColumns, buildingPrimaryKeyColumns) {
		fields = buildingAllColumns
	} else {
		fields = strmangle.SetComplement(
			buildingAllColumns,
			buildingPrimaryKeyColumns,
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

	slice := BuildingSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBuildingsUpsert(t *testing.T) {
	t.Parallel()

	if len(buildingAllColumns) == len(buildingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Building{}
	if err = randomize.Struct(seed, &o, buildingDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Building: %s", err)
	}

	count, err := Buildings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, buildingDBTypes, false, buildingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Building: %s", err)
	}

	count, err = Buildings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
