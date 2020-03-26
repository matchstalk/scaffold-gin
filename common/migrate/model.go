package migrate

import "time"

type Migrate struct {
	Version   string `global:"primary_key"`
	ApplyTime time.Time
}
