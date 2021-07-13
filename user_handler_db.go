package go_lambda_api

import (
	"context"
	"errors"

	"github.com/guregu/dynamo"
)

func scanUsers(ctx context.Context) ([]User, error) {
	var res []User
	table := gdb.Table(usersTable)
	if err := table.Scan().AllWithContext(ctx, &res); err != nil {
		if errors.Is(err, dynamo.ErrNotFound) {
			return nil, nil
		}
		return res, err
	}
	return res, nil
}
