package user

import (
	"asyncservice/rpc/user/pb"
)

type User struct {
	UID          int64  `json:"uid"`
	Nickname     string `json:"nickname"`
	Introduction string `json:"introduction"`
	Gender       int32  `json:"gender"`
}

func toUser(user *user_service.UserInfo) *User {
	return &User{
		UID:          user.Uid,
		Gender:       user.Gender,
		Nickname:     user.Nickname,
		Introduction: user.Introduction,
	}
}

func toUserinfoRequest(uid int64) *user_service.UserInfoRequest {
	return &user_service.UserInfoRequest{
		Uid: uid,
	}
}

func toBatchUserinfoRequest(uids []int64) *user_service.UserInfoBatchRequest {
	return &user_service.UserInfoBatchRequest{
		Uids: uids,
	}
}

func toCreateUserRequest(uid int64, nickname, introduction string, gender int32) *user_service.CreateUserRequest {
	return &user_service.CreateUserRequest{
		Userinfo: toUserItem(uid, nickname, introduction, gender),
	}
}

func toUserItem(uid int64, nickname, introduction string, gender int32) *user_service.UserInfo {
	return &user_service.UserInfo{
		Uid:          uid,
		Nickname:     nickname,
		Introduction: introduction,
		Gender:       gender,
	}
}
