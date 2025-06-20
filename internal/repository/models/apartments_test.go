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

func testApartments(t *testing.T) {
	t.Parallel()

	query := Apartments()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testApartmentsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Apartment{}
	if err = randomize.Struct(seed, o, apartmentDBTypes, true, apartmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
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

	count, err := Apartments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testApartmentsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Apartment{}
	if err = randomize.Struct(seed, o, apartmentDBTypes, true, apartmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Apartments().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Apartments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testApartmentsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Apartment{}
	if err = randomize.Struct(seed, o, apartmentDBTypes, true, apartmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ApartmentSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Apartments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testApartmentsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Apartment{}
	if err = randomize.Struct(seed, o, apartmentDBTypes, true, apartmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ApartmentExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Apartment exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ApartmentExists to return true, but got false.")
	}
}

func testApartmentsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Apartment{}
	if err = randomize.Struct(seed, o, apartmentDBTypes, true, apartmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	apartmentFound, err := FindApartment(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if apartmentFound == nil {
		t.Error("want a record, got nil")
	}
}

func testApartmentsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Apartment{}
	if err = randomize.Struct(seed, o, apartmentDBTypes, true, apartmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Apartments().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testApartmentsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Apartment{}
	if err = randomize.Struct(seed, o, apartmentDBTypes, true, apartmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Apartments().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testApartmentsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	apartmentOne := &Apartment{}
	apartmentTwo := &Apartment{}
	if err = randomize.Struct(seed, apartmentOne, apartmentDBTypes, false, apartmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
	}
	if err = randomize.Struct(seed, apartmentTwo, apartmentDBTypes, false, apartmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = apartmentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = apartmentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Apartments().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testApartmentsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	apartmentOne := &Apartment{}
	apartmentTwo := &Apartment{}
	if err = randomize.Struct(seed, apartmentOne, apartmentDBTypes, false, apartmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
	}
	if err = randomize.Struct(seed, apartmentTwo, apartmentDBTypes, false, apartmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = apartmentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = apartmentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Apartments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func apartmentBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Apartment) error {
	*o = Apartment{}
	return nil
}

func apartmentAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Apartment) error {
	*o = Apartment{}
	return nil
}

func apartmentAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Apartment) error {
	*o = Apartment{}
	return nil
}

func apartmentBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Apartment) error {
	*o = Apartment{}
	return nil
}

func apartmentAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Apartment) error {
	*o = Apartment{}
	return nil
}

func apartmentBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Apartment) error {
	*o = Apartment{}
	return nil
}

func apartmentAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Apartment) error {
	*o = Apartment{}
	return nil
}

func apartmentBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Apartment) error {
	*o = Apartment{}
	return nil
}

func apartmentAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Apartment) error {
	*o = Apartment{}
	return nil
}

func testApartmentsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Apartment{}
	o := &Apartment{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, apartmentDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Apartment object: %s", err)
	}

	AddApartmentHook(boil.BeforeInsertHook, apartmentBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	apartmentBeforeInsertHooks = []ApartmentHook{}

	AddApartmentHook(boil.AfterInsertHook, apartmentAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	apartmentAfterInsertHooks = []ApartmentHook{}

	AddApartmentHook(boil.AfterSelectHook, apartmentAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	apartmentAfterSelectHooks = []ApartmentHook{}

	AddApartmentHook(boil.BeforeUpdateHook, apartmentBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	apartmentBeforeUpdateHooks = []ApartmentHook{}

	AddApartmentHook(boil.AfterUpdateHook, apartmentAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	apartmentAfterUpdateHooks = []ApartmentHook{}

	AddApartmentHook(boil.BeforeDeleteHook, apartmentBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	apartmentBeforeDeleteHooks = []ApartmentHook{}

	AddApartmentHook(boil.AfterDeleteHook, apartmentAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	apartmentAfterDeleteHooks = []ApartmentHook{}

	AddApartmentHook(boil.BeforeUpsertHook, apartmentBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	apartmentBeforeUpsertHooks = []ApartmentHook{}

	AddApartmentHook(boil.AfterUpsertHook, apartmentAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	apartmentAfterUpsertHooks = []ApartmentHook{}
}

func testApartmentsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Apartment{}
	if err = randomize.Struct(seed, o, apartmentDBTypes, true, apartmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Apartments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testApartmentsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Apartment{}
	if err = randomize.Struct(seed, o, apartmentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(strmangle.SetMerge(apartmentPrimaryKeyColumns, apartmentColumnsWithoutDefault)...)); err != nil {
		t.Error(err)
	}

	count, err := Apartments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testApartmentToOneBuildingUsingBuilding(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Apartment
	var foreign Building

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, apartmentDBTypes, false, apartmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, buildingDBTypes, false, buildingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.BuildingID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Building().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddBuildingHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Building) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := ApartmentSlice{&local}
	if err = local.L.LoadBuilding(ctx, tx, false, (*[]*Apartment)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Building == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Building = nil
	if err = local.L.LoadBuilding(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Building == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testApartmentToOneSetOpBuildingUsingBuilding(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Apartment
	var b, c Building

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, apartmentDBTypes, false, strmangle.SetComplement(apartmentPrimaryKeyColumns, apartmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, buildingDBTypes, false, strmangle.SetComplement(buildingPrimaryKeyColumns, buildingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, buildingDBTypes, false, strmangle.SetComplement(buildingPrimaryKeyColumns, buildingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Building{&b, &c} {
		err = a.SetBuilding(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Building != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Apartments[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.BuildingID != x.ID {
			t.Error("foreign key was wrong value", a.BuildingID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.BuildingID))
		reflect.Indirect(reflect.ValueOf(&a.BuildingID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.BuildingID != x.ID {
			t.Error("foreign key was wrong value", a.BuildingID, x.ID)
		}
	}
}

func testApartmentsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Apartment{}
	if err = randomize.Struct(seed, o, apartmentDBTypes, true, apartmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
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

func testApartmentsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Apartment{}
	if err = randomize.Struct(seed, o, apartmentDBTypes, true, apartmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ApartmentSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testApartmentsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Apartment{}
	if err = randomize.Struct(seed, o, apartmentDBTypes, true, apartmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Apartments().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	apartmentDBTypes = map[string]string{`ID`: `bigint`, `BuildingID`: `bigint`, `Number`: `text`, `Floor`: `integer`, `SQMeters`: `real`}
	_                = bytes.MinRead
)

func testApartmentsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(apartmentPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(apartmentAllColumns) == len(apartmentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Apartment{}
	if err = randomize.Struct(seed, o, apartmentDBTypes, true, apartmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Apartments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, apartmentDBTypes, true, apartmentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testApartmentsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(apartmentAllColumns) == len(apartmentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Apartment{}
	if err = randomize.Struct(seed, o, apartmentDBTypes, true, apartmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Apartments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, apartmentDBTypes, true, apartmentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(apartmentAllColumns, apartmentPrimaryKeyColumns) {
		fields = apartmentAllColumns
	} else {
		fields = strmangle.SetComplement(
			apartmentAllColumns,
			apartmentPrimaryKeyColumns,
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

	slice := ApartmentSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testApartmentsUpsert(t *testing.T) {
	t.Parallel()

	if len(apartmentAllColumns) == len(apartmentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Apartment{}
	if err = randomize.Struct(seed, &o, apartmentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Apartment: %s", err)
	}

	count, err := Apartments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, apartmentDBTypes, false, apartmentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Apartment struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Apartment: %s", err)
	}

	count, err = Apartments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
