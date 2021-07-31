package main

import (
	"github.com/nullism/bqb"
)

func basic() {
	q := bqb.New(bqb.PGSQL).
		Select("id, name, email").
		From("users").
		Where("email LIKE '%@yahoo.com'")
	q.Print()
}

func complexQuery() {
	q := bqb.New(bqb.PGSQL).
		Select("t.name", "t.id", bqb.Valf("(SELECT * FROM my_t WHERE id=?) as name", 123)).
		From("my_table t").
		Join("my_other_table ot ON t.id = ot.id").
		Join(bqb.Valf("users u ON t.id = ?", 7)).
		Where(
			bqb.Valf("ST_Distance(t.geom, ot.geom) < ?", 101),
			bqb.Valf("t.name LIKE ?", "william%"),
			bqb.Valf("ST_Distance(t.geom, GeomFromEWKT(?)) < ?", "SRID=4326;POINT(44 -111)", 102),
		).
		Limit(10).
		Offset(2).
		OrderBy("t.name ASC, ot.name DESC").
		GroupBy("t.name", "t.id").
		Having(
			bqb.Valf("COUNT(t.name) > ?", 2),
			bqb.Valf("COUNT(ot.name) > ?", 5),
		)

	q.Print()
}

func join() {
	q := bqb.New(bqb.PGSQL).
		Select("uuidv3_generate() as uuid", "u.id", "UPPER(u.name) as screamname", "u.age", "e.email").
		From("users u").
		Join("emails e ON e.user_id = u.id").
		Where(
			bqb.Valf("u.id IN (?, ?, ?) AND e.email LIKE ?", 1, 3, 5, "%@gmail.com"),
			bqb.Valf("u.id IN (?, ?, ?) AND e.email LIKE ?", 2, 4, 6, "%@hotmail.com"),
			bqb.Valf("u.id IN (?)", []int{7, 8, 9, 10, 11, 12}),
		).
		OrderBy("u.age DESC").
		Limit(10)
	q.Print()
}

func andOr() {
	q := bqb.New(bqb.PGSQL).Select("*").From("patrons").
		Where(
			"drivers_license IS NOT NULL AND (age > 20 AND age < 60)",
			"drivers_license IS NULL AND age >= 60",
			"is_known = true",
		)
	q.Print()
}

func valf() {
	email := "foo@bar.com"
	password := "p4ssw0rd"
	q := bqb.New(bqb.PGSQL).
		Select("*").
		From("users").
		Where(
			bqb.Valf("email = ?", email),
			bqb.Valf("password = ?", password),
		)
	q.Print()
}

func main() {

	// join()
	// complexQuery()
	andOr()
	// basic()
	// valf()

	q := bqb.New(bqb.PGSQL).
		Select("*").From("users").Where(
		bqb.Valf(
			"(name LIKE ? AND age < ?) OR (name LIKE ? AND (age > ? AND email LIKE ?)",
			"smith%", 20, "barb%", 22, "%@gmail.com",
		),
	)
	q.Print()

}
