package Domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type GroupRepository interface {
	FindAll() ([]Group, error)
	FindById(id string) (*Group, *AppError)
}

// We wrap the sql db stuff in  a struct GroupRepoDB
type GroupRepoDB struct {
	db *sql.DB
}

// We implement the interface GroupRepository
// By using the struct GroupRepoDB (the real implementation)
func (ch GroupRepoDB) FindAll() ([]Group, error) {
	findall_sql := "SELECT * FROM Group"
	rows, err := ch.db.Query(findall_sql)

	if err != nil {
		log.Println("error executing sql")
	}

	Groups := make([]Group, 0)
	for rows.Next() {
		var Group Group
		err = rows.Scan(&Group.ID, &Group.Name)
		if err != nil {
			log.Println("Error scanning rows" + err.Error())
		}
		Groups = append(Groups, Group)
	}
	return Groups, nil
}

// We implement the interface GroupRepository
// By using the struct GroupRepoDB (the real implementation)
func (ch GroupRepoDB) FindById(ID string) (*Group, *AppError) {
	var Group Group
	err := ch.db.QueryRow("SELECT * FROM Group WHERE ID=?", ID).Scan(&Group.ID, &Group.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NotFoundError("Group not found")
		} else {
			log.Println("Error scanning rows ById" + err.Error())
			return nil, UnexpectedError("Unexpected db error")
		}
	}
	return &Group, nil
}

func NewGroupRepositoryDB() GroupRepoDB {

	//conn, err := sql.Open("mysql", "root:change-me@tcp(localhost:3306)/Groups")
	db, err := sql.Open("mysql", "root:change-me@tcp(localhost:3306)/Groups")

	if err != nil {
		panic(err.Error())
	}

	//seed method
	//defer db.Close()

	//_, err = db.Exec("INSERT INTO Group (ID, name) VALUES (13, 'Frank')")

	//if err != nil {
	//	panic(err)
	//}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 3)

	return GroupRepoDB{db}
}
