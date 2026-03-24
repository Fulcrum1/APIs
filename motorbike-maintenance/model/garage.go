// model/garage.go
package model

type Garage struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	UserID uint   `json:"user_id" gorm:"not null;index"`
	Notes  string `json:"notes" gorm:"type:text"`

	// Relationships
	User  User   `json:"user" gorm:"foreignKey:UserID"`
	Bikes []Bike `json:"bikes" gorm:"foreignKey:GarageID"`
}
