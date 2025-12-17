package models

import (
	"time"

	"gorm.io/gorm"
)

type School struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Code            string         `gorm:"size:255;unique;not null" json:"code"`
	Name            string         `gorm:"size:255;not null" json:"name"`
	MassarID        string         `gorm:"size:255" json:"massar_id,omitempty"`
	Email           string         `gorm:"size:255" json:"email,omitempty"`
	Phone           string         `gorm:"size:50" json:"phone,omitempty"`
	Phone1          string         `gorm:"size:50" json:"phone_1,omitempty"`
	Phone2          string         `gorm:"size:50" json:"phone_2,omitempty"`
	Fax             string         `gorm:"size:50" json:"fax,omitempty"`
	InterfaceLang   string         `gorm:"size:50" json:"interface_language,omitempty"`
	County          string         `gorm:"size:100" json:"county,omitempty"`
	State           string         `gorm:"size:100" json:"state,omitempty"`
	City            string         `gorm:"size:100" json:"city,omitempty"`
	Address         string         `gorm:"type:text" json:"address,omitempty"`
	Zip             string         `gorm:"size:20" json:"zip,omitempty"`
	SiteWeb         string         `gorm:"size:255" json:"site_web,omitempty"`
	StartDate       *time.Time     `json:"start_date,omitempty"`
	Signatory       string         `gorm:"size:255" json:"signatory,omitempty"`
	TitleSignatory  string         `gorm:"size:255" json:"title_signatory,omitempty"`
	Logo            string         `gorm:"size:255" json:"logo,omitempty"`
	LogoDark        string         `gorm:"size:255" json:"logo_dark,omitempty"`
	LogoLight       string         `gorm:"size:255" json:"logo_light,omitempty"`
	Facebook        string         `gorm:"size:255" json:"facebook,omitempty"`
	Instagram       string         `gorm:"size:255" json:"instagram,omitempty"`
	Snapchat        string         `gorm:"size:255" json:"snapchat,omitempty"`
	Discord         string         `gorm:"size:255" json:"discord,omitempty"`
	Whatsapp        string         `gorm:"size:255" json:"whatsapp,omitempty"`
	X               string         `gorm:"size:255" json:"x,omitempty"`
	RegistrationNum string         `gorm:"size:255" json:"registration_number,omitempty"`
	CNSS            string         `gorm:"size:255" json:"CNSS,omitempty"`
	RCE             string         `gorm:"size:255" json:"RCE,omitempty"`
	TVA             string         `gorm:"size:255" json:"TVA,omitempty"`
	Order           int            `gorm:"default:0" json:"order"`
	Status          bool           `gorm:"default:true" json:"status"`
	ParentID        *uint          `json:"school_id,omitempty"`
	CreatedByID     *uint          `json:"created_by,omitempty"`
	UpdatedByID     *uint          `json:"updated_by,omitempty"`
	DeletedByID     *uint          `json:"deleted_by,omitempty"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
