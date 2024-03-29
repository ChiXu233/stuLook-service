package model

type UserMysql struct {
	Id       int64     `json:"id" gorm:"column:id;type:bigint;primarykey;not null"`
	Name     string    `json:"name" gorm:"column:name;type:varchar(20);not null"`
	Account  string    `json:"account" gorm:"column:account;type:varchar(20);not null"`
	Password string    `json:"password" gorm:"column:password;type:varchar(60);not null"`
	Salt     string    `json:"salt" gorm:"column:salt;type:varchar(20);not null"`
	RoleId   int64     `json:"role_id" gorm:"column:role_id;type:varchar(20);not null;foreignKey:RoleId"` //属于那个角色
	Role     RoleMysql `gorm:"foreignKey:RoleId"`
}

type RoleMysql struct {
	Id   int64       `json:"id" gorm:"column:id;type:bigint;primarykey;not null"`
	Name string      `json:"name" gorm:"column:name;type:varchar(20);not null;"`
	User []UserMysql `gorm:"foreignKey:RoleId"`
	Api  []ApiMysql  `gorm:"many2many:role_apis;"`
}

type ApiMysql struct {
	Id        int64       `json:"id" gorm:"index;column:id;type:bigint(20);not null"`
	Name      string      `json:"name" gorm:"column:name;type:varchar(20);not null"`
	Url       string      `json:"url" gorm:"column:url;type:varchar(20);not null"`
	Method    string      `json:"method" gorm:"column:method;type:varchar(10);not null"`
	Descc     string      `json:"descc" gorm:"column:descc;type:varchar(144)"`
	RoleMysql []RoleMysql `gorm:"many2many:role_apis;"`
}
