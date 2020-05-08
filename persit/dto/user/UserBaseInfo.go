package user

/**
 * @Author: Tomonori
 * @Date: 2020/4/15 15:47
 * @Title: 平台用户基础信息
 * --- --- ---
 * @Desc:
 */
type UserBaseInfo struct {
	Uid int64
	Phone string
	Email string
	RealName string
	NickName string
	AvatarUri string
	Description string
	Gender int //0-位置 1-男 2-女

	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64

	Status int //0-未激活 1-正常 2-ban
}
