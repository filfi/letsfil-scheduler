package models

import "gorm.io/gorm"

func SetupDatabase(db *gorm.DB) (*gorm.DB, error) {

	if err := migrateSchemas(db); err != nil {
		return nil, err
	}
	return db, nil
}

func migrateSchemas(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&RewardSync{},
		&MinerInfo{},
		&ServiceProvider{},
		&RaisingPlan{},
		&DepositDetails{},
		&FundDetails{},
	); err != nil {
		return err
	}
	return nil
}
