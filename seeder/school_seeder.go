package seeder

import (
	"log"
	"time"

	"backend/config"
	"backend/models"
)

// SeedSchools inserts 3 example schools
func SeedSchools() {
	schools := []models.School{
		{
			Code:            "SCH001",
			Name:            "Kuuhaku International School",
			MassarID:        "M001",
			Email:           "contact@kuuhaku.school",
			Phone:           "+212600000001",
			Phone1:          "+212600000002",
			Phone2:          "+212600000003",
			Fax:             "+212600000004",
			InterfaceLang:   "en",
			County:          "Rabat County",
			State:           "Rabat-Salé-Kénitra",
			City:            "Rabat",
			Address:         "123 Main Street, Rabat",
			Zip:             "10000",
			SiteWeb:         "https://kuuhaku.school",
			StartDate:       parseDate("2010-09-01"),
			Signatory:       "John Administrator",
			TitleSignatory:  "Principal",
			Logo:            "logo.png",
			LogoDark:        "logo_dark.png",
			LogoLight:       "logo_light.png",
			Facebook:        "kuuhaku.fb",
			Instagram:       "kuuhaku.ig",
			Snapchat:        "kuuhaku.sc",
			Discord:         "kuuhaku.dc",
			Whatsapp:        "+212600000001",
			X:               "@kuuhaku",
			RegistrationNum: "REG123",
			CNSS:            "CNSS001",
			RCE:             "RCE001",
			TVA:             "TVA001",
			Order:           1,
			Status:          true,
		},
		{
			Code:            "SCH002",
			Name:            "Zenith High School",
			MassarID:        "M002",
			Email:           "info@zenith.school",
			Phone:           "+212600000101",
			Phone1:          "+212600000102",
			Phone2:          "+212600000103",
			Fax:             "+212600000104",
			InterfaceLang:   "fr",
			County:          "Casablanca County",
			State:           "Casablanca-Settat",
			City:            "Casablanca",
			Address:         "45 Zenith Avenue, Casablanca",
			Zip:             "20000",
			SiteWeb:         "https://zenith.school",
			StartDate:       parseDate("2015-01-15"),
			Signatory:       "Alice Principal",
			TitleSignatory:  "Headmaster",
			Logo:            "logo_zenith.png",
			LogoDark:        "logo_zenith_dark.png",
			LogoLight:       "logo_zenith_light.png",
			Facebook:        "zenith.fb",
			Instagram:       "zenith.ig",
			Snapchat:        "zenith.sc",
			Discord:         "zenith.dc",
			Whatsapp:        "+212600000101",
			X:               "@zenith",
			RegistrationNum: "REG456",
			CNSS:            "CNSS002",
			RCE:             "RCE002",
			TVA:             "TVA002",
			Order:           2,
			Status:          true,
		},
		{
			Code:            "SCH003",
			Name:            "Morocco Academy",
			MassarID:        "M003",
			Email:           "hello@moroccoacademy.school",
			Phone:           "+212600000201",
			Phone1:          "+212600000202",
			Phone2:          "+212600000203",
			Fax:             "+212600000204",
			InterfaceLang:   "ar",
			County:          "Marrakech County",
			State:           "Marrakech-Safi",
			City:            "Marrakech",
			Address:         "789 Academy Road, Marrakech",
			Zip:             "40000",
			SiteWeb:         "https://moroccoacademy.school",
			StartDate:       parseDate("2018-06-10"),
			Signatory:       "Mohamed Director",
			TitleSignatory:  "Director",
			Logo:            "logo_morocco.png",
			LogoDark:        "logo_morocco_dark.png",
			LogoLight:       "logo_morocco_light.png",
			Facebook:        "morocco.fb",
			Instagram:       "morocco.ig",
			Snapchat:        "morocco.sc",
			Discord:         "morocco.dc",
			Whatsapp:        "+212600000201",
			X:               "@moroccoacademy",
			RegistrationNum: "REG789",
			CNSS:            "CNSS003",
			RCE:             "RCE003",
			TVA:             "TVA003",
			Order:           3,
			Status:          true,
		},
	}

	for _, s := range schools {
		var existing models.School
		if err := config.DB.Where("code = ?", s.Code).First(&existing).Error; err == nil {
			continue // already exists
		}

		if err := config.DB.Create(&s).Error; err != nil {
			log.Fatal("Failed to seed school:", err)
		}
	}

	log.Println("School seeding completed!")
}

func parseDate(s string) *time.Time {
	t, _ := time.Parse("2006-01-02", s)
	return &t
}
