package pkg

import (
	"database/sql"
	"errors"
	"fmt"
)

var ErrVersionConflict = errors.New("version conflict")

func UpdateBalance(db *sql.DB, user_id int, amount float64) error {
	var currenctBalance float64
	var currentVersion int

	err := db.QueryRow("select balance,version from users where user_id=$1", user_id).Scan(&currenctBalance, &currentVersion)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("User not found %v", err)
		}
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	resp, err := tx.Exec("update users set balance=$1, version=version+1 where user_id=$2 and vesion=$3", amount, user_id, currentVersion)

	if err != nil {
		return err
	}
	rowsAffected, _ := resp.RowsAffected()
	if rowsAffected == 0 {
		return ErrVersionConflict
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
