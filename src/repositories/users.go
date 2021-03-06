package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// Users it is a users's repository
type Users struct {
	db *sql.DB
}

// NewUserRepositories create new users's repository
func NewUserRepositories(db *sql.DB) *Users {
	return &Users{db}
}

// Create create a new user at the database
func (repository Users) Create(user models.User) (uint64, error) {
	statement, error := repository.db.Prepare("insert into users (name, nick, email, pass) values(?, ?, ?, ?)")
	if error != nil {
		return 0, nil
	}
	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Pass)
	if error != nil {
		return 0, error
	}

	lastID, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(lastID), nil
}

// Update update the user informations
func (repository Users) Update(user models.User) (int64, error) {
	statement, error := repository.db.Prepare("update users set name=?, nick=?, email=? where id=?")
	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.ID)
	if error != nil {
		return 0, error
	}

	rowsAffected, error := result.RowsAffected()
	if error != nil {
		return 0, nil
	}

	return rowsAffected, nil
}

// Find find users by name or nick
func (repository Users) Find(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, error := repository.db.Query("select id, name, nick, email, created_at from users where name like ? or nick like ?", nameOrNick, nameOrNick)
	if error != nil {
		return nil, error
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if error = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); error != nil {
			return nil, error
		}

		users = append(users, user)
	}

	return users, nil
}

// FindByID find users by user id
func (repository Users) FindByID(userID uint64) (models.User, error) {
	rows, error := repository.db.Query("select id, name, nick, email, created_at from users where id = ?", userID)
	if error != nil {
		return models.User{}, error
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		if error = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); error != nil {
			return models.User{}, error
		}
	}

	return user, nil
}

// Delete remove a user by id from the database
func (repository Users) Delete(userID uint64) error {
	statement, error := repository.db.Prepare("delete from users where id=?")
	if error != nil {
		return error
	}
	defer statement.Close()

	result, error := statement.Exec(userID)
	if error != nil {
		return error
	}

	_, error = result.RowsAffected()
	if error != nil {
		return error
	}

	return nil
}

// FindByEmail search user by email return ID and password
func (repository Users) FindByEmail(email string) (models.User, error) {
	rows, error := repository.db.Query("select id, pass from users where email=?", email)
	if error != nil {
		return models.User{}, error
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		if error = rows.Scan(
			&user.ID,
			&user.Pass,
		); error != nil {
			return models.User{}, error
		}

	}

	return user, nil
}

// Follow persiste the follower
func (repository Users) Follow(userID, followedUserID uint64) error {
	statement, error := repository.db.Prepare("insert ignore into followers (following_id, follower_id) values (?,?)")
	if error != nil {
		return error
	}

	_, error = statement.Exec(followedUserID, userID)
	if error != nil {
		return error
	}

	return nil
}

// Unfollow remove follow user
func (repository Users) Unfollow(followerUserID, followedUserID uint64) error {
	statement, error := repository.db.Prepare("delete from followers where follower_id=? and following_id=?")
	if error != nil {
		return error
	}

	fmt.Println(followerUserID)
	fmt.Println(followedUserID)

	_, error = statement.Exec(followerUserID, followedUserID)
	if error != nil {
		return error
	}

	return nil
}

//Followers search all user's followers
func (repository Users) Followers(userID uint64) ([]models.User, error) {
	statement, err := repository.db.Prepare("select u.id, u.name, u.nick, u.email from users u join followers f on f.follower_id = u.id where f.following_id=?")
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query(userID)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

//Followings search all user's followings
func (repository Users) Followings(userID uint64) ([]models.User, error) {
	statement, err := repository.db.Prepare("select u.id, u.name, u.nick, u.email from users u join followers f on f.following_id = u.id where f.follower_id=?")
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query(userID)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repository Users) GetUserPassword(userID uint64) (string, error) {
	statement, err := repository.db.Prepare("select pass from users where id=?")
	if err != nil {
		return "", err
	}
	defer statement.Close()

	rows, err := statement.Query(userID)
	if err != nil {
		return "", err
	}

	var password string
	if rows.Next() {
		if err = rows.Scan(&password); err != nil {
			return "", err
		}
	}

	return password, nil
}

func (repository Users) UpdatePassword(userID uint64, newPassword string) (bool, error) {
	statement, err := repository.db.Prepare("update users set pass=? where id=?")
	if err != nil {
		return false, err
	}
	defer statement.Close()

	result, err := statement.Exec(newPassword, userID)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected == 1, nil
}
