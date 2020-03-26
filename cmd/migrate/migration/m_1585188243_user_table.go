package migration

import "fmt"

func (e Migration) M1585188243UserTable() {
	version := "m_1585188243_user_table"
	var err error
	e.db = e.getDb().Begin()
	tx := e.db
	defer func() {
		if err != nil {
			tx.Rollback()
			fmt.Printf("%s run failed!", version)
		} else {
			tx.Commit()
			fmt.Printf("%s run finished!", version)
		}
	}()
	err = e.AddMigrateVersion(version)
	if err != nil {
		return
	}
	// todo migration block
	//
	// todo migration block
}
