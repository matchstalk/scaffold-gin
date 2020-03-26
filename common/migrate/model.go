package migrate

import "time"

type Migrate struct {
	Version 	string	`gorm:"primary_key"`
	ApplyTime	time.Time
}
