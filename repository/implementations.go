package repository

import (
	"context"
	"errors"
	"fmt"
)

func (r *Repository) RegisterUser(ctx context.Context, input RegisterUserInput) (output int, err error) {
	err = r.Db.QueryRowContext(ctx, "INSERT INTO users (phone_number, full_name, \"password\", created_at) VALUES($1, $2, $3, NOW()) RETURNING id", input.PhoneNumber, input.FullName, input.Password).Scan(&output)
	if err != nil {
		return 0, errors.New("no data found")
	}

	return output, nil
}

func (r *Repository) GetUserDataByPhoneNumber(ctx context.Context, input string) (output UserData, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT id, phone_number, full_name, \"password\" FROM users WHERE phone_number = $1", input).Scan(&output.Id, &output.PhoneNumber, &output.FullName, &output.Password)
	if err != nil {
		return output, errors.New("no data found")
	}

	return output, nil
}

func (r *Repository) GetUserDataByID(ctx context.Context, input int) (output UserData, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT id, phone_number, full_name, \"password\" FROM users WHERE id = $1", input).Scan(&output.Id, &output.PhoneNumber, &output.FullName, &output.Password)
	if err != nil {
		return output, errors.New("no data found")
	}

	return output, nil
}

func (r *Repository) UpdateUserDataByID(ctx context.Context, input UserData) (err error) {
	rows, err := r.Db.ExecContext(ctx, "UPDATE users SET phone_number = $1, full_name = $2, updated_at = NOW() WHERE id = $3", input.PhoneNumber, input.FullName, input.Id)
	if err != nil {
		return errors.New("update failed")
	}

	count, err := rows.RowsAffected()
	if err != nil || count == 0 {
		return errors.New("update failed")
	}

	return nil
}

func (r *Repository) UpdateLoginActivity(ctx context.Context, input int) (err error) {
	rows, err := r.Db.ExecContext(ctx, "INSERT INTO activity (user_id, last_login, login_attempt, created_at) VALUES($1, NOW(), 1, NOW()) ON CONFLICT (user_id) DO UPDATE SET last_login = NOW(), login_attempt = activity.login_attempt+1, updated_at = NOW()", input)
	if err != nil {
		fmt.Println("err query", err)
		return errors.New("no data found")
	}

	count, err := rows.RowsAffected()
	if err != nil || count == 0 {
		fmt.Println("err rows", count, err)
		return errors.New("update failed")
	}

	return nil
}
