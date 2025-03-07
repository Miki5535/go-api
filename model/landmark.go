package model

type Landmark struct {
	Idx     int    `gorm:"column:idx;primary_key;AUTO_INCREMENT"`
	Name    string `gorm:"column:name;NOT NULL"`
	Country int    `gorm:"column:country;NOT NULL"`
	Detail  string `gorm:"column:detail;NOT NULL"`
	Url     string `gorm:"column:url;NOT NULL"`
}

func (m *Landmark) TableName() string {
	return "landmark"
}
