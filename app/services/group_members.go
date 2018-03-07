package services

import "github.com/ken-aio/go-echo-sqlboiler/app/models"

// MemberList member list
func MemberList(groupID uint64) (*models.GroupWithMemberList, error) {
	return models.GetGroupMemberList(groupID)
}
