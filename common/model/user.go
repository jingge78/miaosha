package model

import "miaosha-jjl/common/global"

type User struct {
	Uid      uint32  `gorm:"column:uid;type:int UNSIGNED;comment:用户id;primaryKey;not null;" json:"uid"`                // 用户id
	Account  string  `gorm:"column:account;type:varchar(32);comment:用户账号;not null;" json:"account"`                    // 用户账号
	Pwd      string  `gorm:"column:pwd;type:varchar(32);comment:用户密码;not null;" json:"pwd"`                            // 用户密码
	RealName string  `gorm:"column:real_name;type:varchar(25);comment:真实姓名;" json:"real_name"`                         // 真实姓名
	Birthday int32   `gorm:"column:birthday;type:int;comment:生日;default:0;" json:"birthday"`                           // 生日
	CardId   string  `gorm:"column:card_id;type:varchar(20);comment:身份证号码;" json:"card_id"`                            // 身份证号码
	Nickname string  `gorm:"column:nickname;type:varchar(60);comment:用户昵称;default:NULL;" json:"nickname"`              // 用户昵称
	Avatar   string  `gorm:"column:avatar;type:varchar(256);comment:用户头像;default:NULL;" json:"avatar"`                 // 用户头像
	Phone    string  `gorm:"column:phone;type:char(15);comment:手机号码;default:NULL;" json:"phone"`                       // 手机号码
	AddTime  uint32  `gorm:"column:add_time;type:int UNSIGNED;comment:添加时间;default:0;" json:"add_time"`                // 添加时间
	NowMoney float64 `gorm:"column:now_money;type:decimal(8, 2) UNSIGNED;comment:用户余额;default:0.00;" json:"now_money"` // 用户余额
	Status   int8    `gorm:"column:status;type:tinyint(1);comment:1为正常，0为禁止;default:1;" json:"status"`                 // 1为正常，0为禁止
	UserType string  `gorm:"column:user_type;type:varchar(32);comment:用户类型;default:NULL;" json:"user_type"`            // 用户类型
	Addres   string  `gorm:"column:addres;type:varchar(255);comment:详细地址;" json:"addres"`                              // 详细地址
}

func (u *User) TableName() string {
	return "user"
}

// 账号登录//jj

func (u *User) LoginUser(account string) error {
	return global.DB.Debug().Table("user").Where("account = ?", account).Limit(1).Find(&u).Error
}
func (u *User) Create() error {
	return global.DB.Create(&u).Error
}

// 修改密码//jj

func (u *User) Update(account string, pwd string) error {
	return global.DB.Model(&User{}).Where("account = ?", account).Update("pwd", pwd).Error
}

func (u *User) GetUserUid(uid uint64) error {
	return global.DB.Debug().Table("user").Where("uid = ?", uid).Limit(1).Find(&u).Error
}

func (u *User) UserImproveInformation(uid uint64) error {
	return global.DB.Model(&User{}).Table("user").Where("uid = ?", uid).Updates(&u).Error
}
