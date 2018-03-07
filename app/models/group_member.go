package models

import (
	"github.com/ken-aio/go-echo-sqlboiler/app/infra/db"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// GroupWithMemberList group and member info
type GroupWithMemberList struct {
	GroupID   uint64             `json:"group_id"`
	GroupName string             `json:"group_name"`
	Members   []*groupMemberList `json:"members"`
}

type groupMemberList struct {
	IsAdmin  bool   `json:"is_admin"`
	UserID   uint64 `json:"user_id"`
	UserName string `json:"user_name"`
}

// GetGroupMemberList get member list
func GetGroupMemberList(groupID uint64) (*GroupWithMemberList, error) {
	group, err := db.GroupsG(
		qm.Where("id = ?", groupID),
	).One()
	if err != nil {
		return nil, err
	}

	members, err := group.GroupMembersG(
		qm.Load("User"),
		qm.OrderBy("id asc"),
	).All()
	if err != nil {
		return nil, err
	}

	resMembers := make([]*groupMemberList, len(members))
	for i := 0; i < len(members); i++ {
		member := members[i]
		resMember := &groupMemberList{}
		resMember.IsAdmin = member.IsAdmin == 1
		resMember.UserID = member.R.User.ID
		resMember.UserName = member.R.User.Name
		resMembers[i] = resMember
	}

	res := &GroupWithMemberList{}
	res.GroupID = group.ID
	res.GroupName = group.Name
	res.Members = resMembers

	return res, nil
}
