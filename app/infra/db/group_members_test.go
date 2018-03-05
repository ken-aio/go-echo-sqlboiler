// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testGroupMembers(t *testing.T) {
	t.Parallel()

	query := GroupMembers(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testGroupMembersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	groupMember := &GroupMember{}
	if err = randomize.Struct(seed, groupMember, groupMemberDBTypes, true, groupMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = groupMember.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = groupMember.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := GroupMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGroupMembersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	groupMember := &GroupMember{}
	if err = randomize.Struct(seed, groupMember, groupMemberDBTypes, true, groupMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = groupMember.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = GroupMembers(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := GroupMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGroupMembersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	groupMember := &GroupMember{}
	if err = randomize.Struct(seed, groupMember, groupMemberDBTypes, true, groupMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = groupMember.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := GroupMemberSlice{groupMember}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := GroupMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testGroupMembersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	groupMember := &GroupMember{}
	if err = randomize.Struct(seed, groupMember, groupMemberDBTypes, true, groupMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = groupMember.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := GroupMemberExists(tx, groupMember.ID)
	if err != nil {
		t.Errorf("Unable to check if GroupMember exists: %s", err)
	}
	if !e {
		t.Errorf("Expected GroupMemberExistsG to return true, but got false.")
	}
}
func testGroupMembersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	groupMember := &GroupMember{}
	if err = randomize.Struct(seed, groupMember, groupMemberDBTypes, true, groupMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = groupMember.Insert(tx); err != nil {
		t.Error(err)
	}

	groupMemberFound, err := FindGroupMember(tx, groupMember.ID)
	if err != nil {
		t.Error(err)
	}

	if groupMemberFound == nil {
		t.Error("want a record, got nil")
	}
}
func testGroupMembersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	groupMember := &GroupMember{}
	if err = randomize.Struct(seed, groupMember, groupMemberDBTypes, true, groupMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = groupMember.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = GroupMembers(tx).Bind(groupMember); err != nil {
		t.Error(err)
	}
}

func testGroupMembersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	groupMember := &GroupMember{}
	if err = randomize.Struct(seed, groupMember, groupMemberDBTypes, true, groupMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = groupMember.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := GroupMembers(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testGroupMembersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	groupMemberOne := &GroupMember{}
	groupMemberTwo := &GroupMember{}
	if err = randomize.Struct(seed, groupMemberOne, groupMemberDBTypes, false, groupMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}
	if err = randomize.Struct(seed, groupMemberTwo, groupMemberDBTypes, false, groupMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = groupMemberOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = groupMemberTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := GroupMembers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testGroupMembersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	groupMemberOne := &GroupMember{}
	groupMemberTwo := &GroupMember{}
	if err = randomize.Struct(seed, groupMemberOne, groupMemberDBTypes, false, groupMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}
	if err = randomize.Struct(seed, groupMemberTwo, groupMemberDBTypes, false, groupMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = groupMemberOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = groupMemberTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := GroupMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func groupMemberBeforeInsertHook(e boil.Executor, o *GroupMember) error {
	*o = GroupMember{}
	return nil
}

func groupMemberAfterInsertHook(e boil.Executor, o *GroupMember) error {
	*o = GroupMember{}
	return nil
}

func groupMemberAfterSelectHook(e boil.Executor, o *GroupMember) error {
	*o = GroupMember{}
	return nil
}

func groupMemberBeforeUpdateHook(e boil.Executor, o *GroupMember) error {
	*o = GroupMember{}
	return nil
}

func groupMemberAfterUpdateHook(e boil.Executor, o *GroupMember) error {
	*o = GroupMember{}
	return nil
}

func groupMemberBeforeDeleteHook(e boil.Executor, o *GroupMember) error {
	*o = GroupMember{}
	return nil
}

func groupMemberAfterDeleteHook(e boil.Executor, o *GroupMember) error {
	*o = GroupMember{}
	return nil
}

func groupMemberBeforeUpsertHook(e boil.Executor, o *GroupMember) error {
	*o = GroupMember{}
	return nil
}

func groupMemberAfterUpsertHook(e boil.Executor, o *GroupMember) error {
	*o = GroupMember{}
	return nil
}

func testGroupMembersHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &GroupMember{}
	o := &GroupMember{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, groupMemberDBTypes, false); err != nil {
		t.Errorf("Unable to randomize GroupMember object: %s", err)
	}

	AddGroupMemberHook(boil.BeforeInsertHook, groupMemberBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	groupMemberBeforeInsertHooks = []GroupMemberHook{}

	AddGroupMemberHook(boil.AfterInsertHook, groupMemberAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	groupMemberAfterInsertHooks = []GroupMemberHook{}

	AddGroupMemberHook(boil.AfterSelectHook, groupMemberAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	groupMemberAfterSelectHooks = []GroupMemberHook{}

	AddGroupMemberHook(boil.BeforeUpdateHook, groupMemberBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	groupMemberBeforeUpdateHooks = []GroupMemberHook{}

	AddGroupMemberHook(boil.AfterUpdateHook, groupMemberAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	groupMemberAfterUpdateHooks = []GroupMemberHook{}

	AddGroupMemberHook(boil.BeforeDeleteHook, groupMemberBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	groupMemberBeforeDeleteHooks = []GroupMemberHook{}

	AddGroupMemberHook(boil.AfterDeleteHook, groupMemberAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	groupMemberAfterDeleteHooks = []GroupMemberHook{}

	AddGroupMemberHook(boil.BeforeUpsertHook, groupMemberBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	groupMemberBeforeUpsertHooks = []GroupMemberHook{}

	AddGroupMemberHook(boil.AfterUpsertHook, groupMemberAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	groupMemberAfterUpsertHooks = []GroupMemberHook{}
}
func testGroupMembersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	groupMember := &GroupMember{}
	if err = randomize.Struct(seed, groupMember, groupMemberDBTypes, true, groupMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = groupMember.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := GroupMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGroupMembersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	groupMember := &GroupMember{}
	if err = randomize.Struct(seed, groupMember, groupMemberDBTypes, true); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = groupMember.Insert(tx, groupMemberColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := GroupMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGroupMemberToOneGroupUsingGroup(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local GroupMember
	var foreign Group

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, groupMemberDBTypes, false, groupMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, groupDBTypes, false, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.GroupID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Group(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := GroupMemberSlice{&local}
	if err = local.L.LoadGroup(tx, false, (*[]*GroupMember)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.Group == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Group = nil
	if err = local.L.LoadGroup(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Group == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testGroupMemberToOneUserUsingUser(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local GroupMember
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, groupMemberDBTypes, false, groupMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.User(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := GroupMemberSlice{&local}
	if err = local.L.LoadUser(tx, false, (*[]*GroupMember)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testGroupMemberToOneSetOpGroupUsingGroup(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a GroupMember
	var b, c Group

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, groupMemberDBTypes, false, strmangle.SetComplement(groupMemberPrimaryKeyColumns, groupMemberColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, groupDBTypes, false, strmangle.SetComplement(groupPrimaryKeyColumns, groupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, groupDBTypes, false, strmangle.SetComplement(groupPrimaryKeyColumns, groupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Group{&b, &c} {
		err = a.SetGroup(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Group != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.GroupMembers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.GroupID != x.ID {
			t.Error("foreign key was wrong value", a.GroupID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.GroupID))
		reflect.Indirect(reflect.ValueOf(&a.GroupID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.GroupID != x.ID {
			t.Error("foreign key was wrong value", a.GroupID, x.ID)
		}
	}
}
func testGroupMemberToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a GroupMember
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, groupMemberDBTypes, false, strmangle.SetComplement(groupMemberPrimaryKeyColumns, groupMemberColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.GroupMembers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}
func testGroupMembersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	groupMember := &GroupMember{}
	if err = randomize.Struct(seed, groupMember, groupMemberDBTypes, true, groupMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = groupMember.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = groupMember.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testGroupMembersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	groupMember := &GroupMember{}
	if err = randomize.Struct(seed, groupMember, groupMemberDBTypes, true, groupMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = groupMember.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := GroupMemberSlice{groupMember}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testGroupMembersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	groupMember := &GroupMember{}
	if err = randomize.Struct(seed, groupMember, groupMemberDBTypes, true, groupMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = groupMember.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := GroupMembers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	groupMemberDBTypes = map[string]string{`CreatedAt`: `datetime`, `CreatedBy`: `varchar`, `GroupID`: `bigint`, `ID`: `bigint`, `IsAdmin`: `tinyint`, `UpdatedAt`: `datetime`, `UpdatedBy`: `varchar`, `UserID`: `bigint`}
	_                  = bytes.MinRead
)

func testGroupMembersUpdate(t *testing.T) {
	t.Parallel()

	if len(groupMemberColumns) == len(groupMemberPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	groupMember := &GroupMember{}
	if err = randomize.Struct(seed, groupMember, groupMemberDBTypes, true, groupMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = groupMember.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := GroupMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, groupMember, groupMemberDBTypes, true, groupMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}

	if err = groupMember.Update(tx); err != nil {
		t.Error(err)
	}
}

func testGroupMembersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(groupMemberColumns) == len(groupMemberPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	groupMember := &GroupMember{}
	if err = randomize.Struct(seed, groupMember, groupMemberDBTypes, true, groupMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = groupMember.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := GroupMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, groupMember, groupMemberDBTypes, true, groupMemberPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(groupMemberColumns, groupMemberPrimaryKeyColumns) {
		fields = groupMemberColumns
	} else {
		fields = strmangle.SetComplement(
			groupMemberColumns,
			groupMemberPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(groupMember))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := GroupMemberSlice{groupMember}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testGroupMembersUpsert(t *testing.T) {
	t.Parallel()

	if len(groupMemberColumns) == len(groupMemberPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	groupMember := GroupMember{}
	if err = randomize.Struct(seed, &groupMember, groupMemberDBTypes, true); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = groupMember.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert GroupMember: %s", err)
	}

	count, err := GroupMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &groupMember, groupMemberDBTypes, false, groupMemberPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GroupMember struct: %s", err)
	}

	if err = groupMember.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert GroupMember: %s", err)
	}

	count, err = GroupMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
