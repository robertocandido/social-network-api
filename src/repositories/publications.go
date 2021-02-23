package repositories

import (
	"api/src/models"
	"database/sql"
)

const (
	InsertPublication = `insert into publications(title, content, user_id)
								  values (?, ?, ?)`
	SelectPublicationByID = `select p.id, p.title, p.content, p.likes, p.user_id, u.nick, p.created_at
							   from publications p
									join users u on p.user_id = u.id
							  where p.id = ?`
	SelectUserAndFollowedUserPublicationsByUserID = `select p.id, p.title, p.content, p.likes, p.user_id, u.nick, p.created_at
													   from users u
													        join publications p on u.id = p.user_id
													  where u.id = ?
														 or u.id in (select f.following_id from followers f where f.follower_id = ?)
												   order by p.id`
)

type Publications struct {
	db *sql.DB
}

func NewPublicationRepositories(db *sql.DB) *Publications {
	return &Publications{db}
}

//Create persist a new publication
func (repository Publications) Create(pub models.Publication) (uint64, error) {
	statement, err := repository.db.Prepare(InsertPublication)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(pub.Title, pub.Content, pub.UserID)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertId), nil
}

//Find find a specific publication by id
func (repository Publications) Find(ID uint64) (publication models.Publication, err error) {
	rows, err := repository.db.Query(SelectPublicationByID, ID)

	if err != nil {
		return
	}
	defer rows.Close()

	if rows.Next() {
		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.Likes,
			&publication.UserID,
			&publication.UserNick,
			&publication.CreatedAt,
		); err != nil {
			return
		}
	}

	return
}

func (repository Publications) FindByUserAndFollowUsers(userID uint64) (publications []models.Publication, err error) {
	rows, err := repository.db.Query(SelectUserAndFollowedUserPublicationsByUserID, userID, userID)
	if err != nil {
		return
	}
	defer rows.Close()

	var publication models.Publication
	for rows.Next() {
		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.Likes,
			&publication.UserID,
			&publication.UserNick,
			&publication.CreatedAt,
		); err != nil {
			return
		}
		publications = append(publications, publication)
	}

	return
}