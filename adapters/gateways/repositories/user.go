package repositories

import (
	"database/sql"
	"strconv"

	"github.com/lucasasoaresmar/clean-go/adapters/gateways/repositories/queries"
	"github.com/lucasasoaresmar/clean-go/entities"
)

// User repository ...
type User struct{}

// GetByEmailAndPassword method
func (u *User) GetByEmailAndPassword(email string, password string) (user *entities.User, err error) {
	exec(func(db *sql.DB) {

		var id int
		user = new(entities.User)

		err := db.QueryRow(queries.GetUserByEmailAndPassword(), email, password).
			Scan(&id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return
		}

		user.ID = strconv.Itoa(id)
	})

	return
}
