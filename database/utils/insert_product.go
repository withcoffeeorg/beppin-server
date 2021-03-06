package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
)

// InsertProduct - Insert a product.
func InsertProduct(db *sql.DB, product models.Product) (err error) {
	exists, err := ExistsUser(db, models.User{ID: product.UserID})
	if err != nil {
		return
	}

	if !exists {
		err = fmt.Errorf("failed to check (%d) user: %w", product.UserID, errs.ErrNotExistentObject)
		return
	}

	query := `
		INSERT INTO
			products(user_id, name, description)
		VALUES
			($1, $2, $3)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the insert product statement: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		product.UserID,
		product.Name,
		product.Description,
	)
	if err != nil {
		err = fmt.Errorf("failed to execute insert product statement: %v", err)
	}
	return
}
