package librepository

import (
	"database/sql"
	libentity "github.com/Nikitastarikov/practice-on-golang/exercise4/lib/entity"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) (*Repository, error) {
	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) CreatePhone(number string) (*libentity.Phone, error) {
	q := `
		INSERT INTO phones (number)
		VALUES ($1)
		RETURNING *;
	`

	phone := new(libentity.Phone)

	err := r.db.QueryRow(q, number).Scan(
		&phone.Id,
		&phone.Number,
	)

	if err != nil {
		return nil, err
	}

	return phone, nil
}

func (r *Repository) GetPhoneById(id int) (*libentity.Phone, error) {
	q := `
		SELECT *
		FROM phones
		WHERE id = $1;
	`

	phone := new(libentity.Phone)

	err := r.db.QueryRow(q, id).Scan(
		&phone.Id,
		&phone.Number,
	)

	if err != nil {
		return nil, err
	}

	return phone, nil
}

func (r *Repository) UpdatePhoneById(id int, number string) (*libentity.Phone, error) {
	q := `
		UPDATE phones
		SET number = $2
		WHERE id = $1
		RETURNING *;
	`

	phone := new(libentity.Phone)

	row := r.db.QueryRow(q, id, number)

	err := row.Scan(
		&phone.Id,
		&phone.Number,
	)

	if err != nil {
		return nil, err
	}

	return phone, nil
}

func (r *Repository) DeletePhoneById(id int) error {
	q := `
		DELETE FROM
		phones
		WHERE id = $1;
	`

	_, err := r.db.Exec(q, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetPhoneList(limit, page *int) ([]*libentity.Phone, error) {
	_limit := 100
	_page := 1

	if limit != nil {
		_limit = *limit
	}

	if page != nil {
		_page = *page
	}

	offset := _limit * (_page - 1)

	q := `
		SELECT *
		FROM phones
		OFFSET $1
		LIMIT $2
	`

	rows, err := r.db.Query(q, offset, _limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	phoneList := make([]*libentity.Phone, 0)

	for rows.Next() {
		phone := new(libentity.Phone)

		err = rows.Scan(
			&phone.Id,
			&phone.Number,
		)

		if err != nil {
			return nil, err
		}

		phoneList = append(phoneList, phone)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return phoneList, nil
}
