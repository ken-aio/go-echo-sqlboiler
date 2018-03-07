package v1

import (
	"net/http"
	"strconv"

	"github.com/ken-aio/go-echo-sqlboiler/app/services"
	"github.com/labstack/echo"
)

// GroupMemberList GroupMember list API
// @Summary GroupMemberList GroupMember list API
// @Description list group mmebers
// @Accept  json
// @Produce  json
// @Param   group_id     path    uint64     true        "group id"
// @Success 200 {object} models.GroupWithMemberList	""
// @Router /api/v1/groups/{group_id}/members [get]
func GroupMemberList(c echo.Context) error {
	groupID, err := strconv.ParseUint(c.Param("group_id"), 10, 64)
	if err != nil {
		return err
	}

	res, err := services.MemberList(groupID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
