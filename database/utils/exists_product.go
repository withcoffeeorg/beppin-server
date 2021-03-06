package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
)

// ExistsProduct - Checks if exists a product.
func ExistsProduct(db *sql.DB, product models.Product) (exists bool, err error) {
	identifier := product.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to check product: %w (product)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		SELECT
			EXISTS(
				SELECT
					id
				FROM
					products
				WHERE
					id = $1
			)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the exists (%v) product statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(product.ID).Scan(&exists)
	if err != nil {
		err = fmt.Errorf("failed to select the exists (%v) product statement: %v", identifier, err)
	}
	return
}
