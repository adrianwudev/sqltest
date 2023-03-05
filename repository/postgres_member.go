package member

import (
	"database/sql"
	"fmt"

	"github.com/adrianwudev/sqltest/domain"
	"github.com/sirupsen/logrus"
)

type postgresMemberRepository struct {
	Conn *sql.DB
}

func NewPostgresMemberRepository(conn *sql.DB) domain.MemberRepository {
	return &postgresMemberRepository{conn}
}

// FetchAll implements domain.MemberRepository
func (p *postgresMemberRepository) FetchAll() (result []domain.Member, err error) {
	rows, err := p.Conn.Query("SELECT * FROM mem")
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()

	result = make([]domain.Member, 0)
	for rows.Next() {
		m := domain.Member{}
		err = rows.Scan(&m.Name, &m.Age, &m.Id)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, m)
	}

	return result, nil
}

func (p *postgresMemberRepository) Add(m domain.Member) (err error) {
	isExisted, err := p.CheckIfExists(m)
	if err != nil {
		logrus.Error(err)
		return err
	}

	if isExisted {
		fmt.Printf("The member has already been inserted, name: %s, age: %d\n", m.Name, m.Age)
		return nil
	}
	statement, err := p.Conn.Prepare("INSERT INTO mem (name, age) VALUES ($1, $2)")
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(m.Name, m.Age)
	if err != nil {
		logrus.Error(err)
		return err
	}

	fmt.Println("successfully inserted")

	return nil
}

func (p *postgresMemberRepository) CheckIfExists(m domain.Member) (bool, error) {
	var count int
	statement, err := p.Conn.Prepare("SELECT COUNT(*) FROM mem WHERE name=$1 AND age=$2")
	if err != nil {
		logrus.Error(err)
		return false, err
	}
	defer statement.Close()

	err = statement.QueryRow(m.Name, m.Age).Scan(&count)
	if err != nil {
		logrus.Error(err)
		return false, err
	}
	fmt.Println()
	fmt.Println("Check IF Exists, Count: ", count)

	return count > 0, nil
}
