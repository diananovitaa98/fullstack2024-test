package repository

import (
	"context"
	"database/sql"
	"fullstacktest/entity"
)

type ClientRepoItf interface {
	InsertClient(ctx context.Context, data *entity.MyClient) (*int, error)
	SelectAllClient(ctx context.Context) ([]entity.MyClient, error)
	UpdateClient(ctx context.Context, data *entity.MyClient) error
	DeleteClient(ctx context.Context, slug string) error
}

type ClientRepoImpl struct {
	db *sql.DB
}

func NewClientRepo(db *sql.DB) ClientRepoImpl {
	return ClientRepoImpl{
		db: db,
	}
}

func (cr ClientRepoImpl) InsertClient(ctx context.Context, data *entity.MyClient) (*int, error) {
	query := `
		INSERT INTO my_client (
			name, slug, is_project, self_capture, client_prefix,
			client_logo, address, phone_number, city,
			created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW())
		RETURNING id
	`

	var id int
	err := cr.db.QueryRowContext(
		ctx,
		query,
		data.Name,
		data.Slug,
		data.IsProject,
		data.SelfCapture,
		data.ClientPrefix,
		data.ClientLogo,
		data.Address,
		data.PhoneNumber,
		data.City,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (cr ClientRepoImpl) SelectAllClient(ctx context.Context) ([]entity.MyClient, error) {
	var clients []entity.MyClient

	q := `
		SELECT id, name, slug, is_project, self_capture, client_prefix,
		       client_logo, address, phone_number, city,
		       created_at, updated_at, deleted_at
		FROM my_client
		WHERE deleted_at IS NULL
	`

	rows, err := cr.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var client entity.MyClient
		err := rows.Scan(
			&client.ID,
			&client.Name,
			&client.Slug,
			&client.IsProject,
			&client.SelfCapture,
			&client.ClientPrefix,
			&client.ClientLogo,
			&client.Address,
			&client.PhoneNumber,
			&client.City,
			&client.CreatedAt,
			&client.UpdatedAt,
			&client.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return clients, nil
}

func (cr ClientRepoImpl) UpdateClient(ctx context.Context, data *entity.MyClient) error {
	query := `
		UPDATE my_client
		SET name = $1,
		    is_project = $2,
		    self_capture = $3,
		    client_prefix = $4,
		    client_logo = $5,
		    address = $6,
		    phone_number = $7,
		    city = $8,
		    updated_at = NOW()
		WHERE slug = $9 AND deleted_at IS NULL
	`

	_, err := cr.db.ExecContext(
		ctx,
		query,
		data.Name,
		data.IsProject,
		data.SelfCapture,
		data.ClientPrefix,
		data.ClientLogo,
		data.Address,
		data.PhoneNumber,
		data.City,
		data.Slug,
	)

	return err
}

func (cr ClientRepoImpl) DeleteClient(ctx context.Context, slug string) error {
	query := `
		UPDATE my_client
		SET deleted_at = NOW()
		WHERE slug = $1 AND deleted_at IS NULL
	`

	_, err := cr.db.ExecContext(ctx, query, slug)
	return err
}
