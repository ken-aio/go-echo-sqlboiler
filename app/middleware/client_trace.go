package middleware

import (
	"strings"

	"github.com/ken-aio/go-echo-sqlboiler/app/infra/db"
	"github.com/labstack/echo"
	"github.com/volatiletech/sqlboiler/boil"
)

func ClientTrace(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ua := strings.Join(c.Request().Header["User-Agent"], ",")
		if len(ua) > 20 {
			ua = ua[0:20]
		}
		trace := c.Request().Method + " " + c.Path() + " " + ua

		// users
		db.AddUserHook(boil.BeforeInsertHook, func(_ boil.Executor, m *db.User) error {
			m.CreatedBy = trace
			m.UpdatedBy = trace
			return nil
		})
		db.AddUserHook(boil.BeforeUpdateHook, func(_ boil.Executor, m *db.User) error {
			m.UpdatedBy = trace
			return nil
		})

		// group_members
		db.AddGroupMemberHook(boil.BeforeInsertHook, func(_ boil.Executor, m *db.GroupMember) error {
			m.CreatedBy = trace
			m.UpdatedBy = trace
			return nil
		})
		db.AddGroupMemberHook(boil.BeforeUpdateHook, func(_ boil.Executor, m *db.GroupMember) error {
			m.UpdatedBy = trace
			return nil
		})

		// groups
		db.AddGroupHook(boil.BeforeInsertHook, func(_ boil.Executor, m *db.Group) error {
			m.CreatedBy = trace
			m.UpdatedBy = trace
			return nil
		})
		db.AddGroupHook(boil.BeforeUpdateHook, func(_ boil.Executor, m *db.Group) error {
			m.UpdatedBy = trace
			return nil
		})
		return next(c)
	}
}
