package model

import "miaosha-jjl/common/global"

type Signup struct {
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
	Email    string  `gorm:"column:email:type:varchar(30);comment:邮箱注册;" json:"email"`
}

type Email struct {
	Emaile   string `gorm:"column:emaile;type:varchar(50);comment:邮箱;default:NULL;" json:"emaile"`     // 邮箱
	Password string `gorm:"column:password;type:varchar(11);comment:密码;default:NULL;" json:"password"` // 密码
	Uid      int32  `gorm:"column:uid;type:int;comment:用户名称;primaryKey;not null;" json:"uid"`          // 用户名称
}

func (s *Signup) TableName() string {
	return "user"
}

// 账号登录//jj

func (s *Signup) LoginUser(account string) error {
	return global.DB.Table("user").Where("account = ?", account).Limit(1).Find(&s).Error

}
func (s *Signup) Create() error {
	return global.DB.Create(&s).Error
}

// 修改密码//jj

func (s *Signup) Update(account string, pwd string) error {
	return global.DB.Model(&User{}).Where("account = ?", account).Update("pwd", pwd).Error
}

func (s *Signup) Signup(email string) error {
	return global.DB.Table("user").Where("email=?", email).Limit(1).Find(&s).Error
}

func (s *Signup) AddSignupEnter() error {
	return global.DB.Debug().Table("signup_signup").Create(&s).Error
}

func (e *Email) RegisterUser() error {
	return global.DB.Debug().Table("email").Create(&e).Error
}

//func (e *Email) Create() error {
//	return global.DB.Debug().Create(&e).Error
//
//}
