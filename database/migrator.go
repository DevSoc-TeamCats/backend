package database

import (
	"github.com/BearTS/go-gin-monolith/database/tables"
	"gorm.io/gorm"
)

type Migrate struct {
	TableName string
	Run       func(*gorm.DB) error
}

func AutoMigrate(db *gorm.DB) []Migrate {
	var users tables.Users
	var otpVerifications tables.OtpVerifications
	var devices tables.Devices
	var transactions tables.Transactions
	var reports tables.Reports
	var subscriptionPlans tables.SubscriptionPlans

	usersM := Migrate{TableName: "users",
		Run: func(d *gorm.DB) error { return db.AutoMigrate(&users) }}
	otpVerificationM := Migrate{TableName: "otp_verifications",
		Run: func(d *gorm.DB) error { return db.AutoMigrate(&otpVerifications) }}
	devicesM := Migrate{TableName: "devices",
		Run: func(d *gorm.DB) error { return db.AutoMigrate(&devices) }}
	transactionsM := Migrate{TableName: "transactions",
		Run: func(d *gorm.DB) error { return db.AutoMigrate(&transactions) }}
	subscriptionPlansM := Migrate{TableName: "subscription_plans",
		Run: func(d *gorm.DB) error { return db.AutoMigrate(&subscriptionPlans) }}
	reportsM := Migrate{TableName: "reports",
		Run: func(d *gorm.DB) error { return db.AutoMigrate(&reports) }}

	return []Migrate{
		usersM,
		otpVerificationM,
		devicesM,
		transactionsM,
		subscriptionPlansM,
		reportsM,
	}
}
