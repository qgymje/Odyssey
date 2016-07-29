package models

import (
	"fmt"
	"testing"

	"gopkg.in/pg.v4/orm"
)

type TestProfile struct {
	Id     int
	Lang   string
	Active bool
	UserId int
}

// User has many profiles.
type TestUser struct {
	Id       int
	Name     string
	Profiles []*TestProfile
}

func TestHasMany(t *testing.T) {
	qs := []string{
		"CREATE TEMP TABLE test_users (id int, name text)",
		"CREATE TEMP TABLE test_profiles (id int, lang text, active bool, user_id int)",
		"INSERT INTO test_users VALUES (1, 'user 1')",
		"INSERT INTO test_profiles VALUES (1, 'en', TRUE, 1), (2, 'ru', TRUE, 1), (3, 'md', FALSE, 1)",
	}
	for _, q := range qs {
		_, err := db.Exec(q)
		if err != nil {
			panic(err)

		}

	}

	// Select user and all his active profiles with following queries:
	//
	// SELECT "user".* FROM "users" AS "user" ORDER BY "user"."id" LIMIT 1
	//
	// SELECT "profile".* FROM "profiles" AS "profile"
	// WHERE (active IS TRUE) AND (("profile"."user_id") IN ((1)))

	var tuser TestUser
	err := db.Model(&tuser).
		Column("tuser.*", "Profiles").
		Relation("Profiles", func(q *orm.Query) *orm.Query {
			return q.Where("active IS TRUE")

		}).
		First()
	if err != nil {
		panic(err)

	}

	fmt.Println(tuser.Id, tuser.Name, tuser.Profiles[0], tuser.Profiles[1])

}
