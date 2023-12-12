package db

/*
 This file was created by mkdb-client.
 The intention is not to modify thils file, but you may extend the struct DBPercentAlert
 in a seperate file (so that you can regenerate this one from time to time)
*/

/*
 PRIMARY KEY: ID
*/

/*
 postgres:
 create sequence percentalert_seq;

Main Table:

 CREATE TABLE percentalert (id integer primary key default nextval('percentalert_seq'),totalmetric text not null  ,countmetric text not null  ,effects integer not null  );

Alter statements:
ALTER TABLE percentalert ADD COLUMN totalmetric text not null default '';
ALTER TABLE percentalert ADD COLUMN countmetric text not null default '';
ALTER TABLE percentalert ADD COLUMN effects integer not null default 0;


Archive Table: (structs can be moved from main to archive using Archive() function)

 CREATE TABLE percentalert_archive (id integer unique not null,totalmetric text not null,countmetric text not null,effects integer not null);
*/

import (
	"context"
	gosql "database/sql"
	"fmt"
	savepb "golang.conradwood.net/apis/promconfig"
	"golang.conradwood.net/go-easyops/sql"
	"os"
)

var (
	default_def_DBPercentAlert *DBPercentAlert
)

type DBPercentAlert struct {
	DB                  *sql.DB
	SQLTablename        string
	SQLArchivetablename string
}

func DefaultDBPercentAlert() *DBPercentAlert {
	if default_def_DBPercentAlert != nil {
		return default_def_DBPercentAlert
	}
	psql, err := sql.Open()
	if err != nil {
		fmt.Printf("Failed to open database: %s\n", err)
		os.Exit(10)
	}
	res := NewDBPercentAlert(psql)
	ctx := context.Background()
	err = res.CreateTable(ctx)
	if err != nil {
		fmt.Printf("Failed to create table: %s\n", err)
		os.Exit(10)
	}
	default_def_DBPercentAlert = res
	return res
}
func NewDBPercentAlert(db *sql.DB) *DBPercentAlert {
	foo := DBPercentAlert{DB: db}
	foo.SQLTablename = "percentalert"
	foo.SQLArchivetablename = "percentalert_archive"
	return &foo
}

// archive. It is NOT transactionally save.
func (a *DBPercentAlert) Archive(ctx context.Context, id uint64) error {

	// load it
	p, err := a.ByID(ctx, id)
	if err != nil {
		return err
	}

	// now save it to archive:
	_, e := a.DB.ExecContext(ctx, "archive_DBPercentAlert", "insert into "+a.SQLArchivetablename+" (id,totalmetric, countmetric, effects) values ($1,$2, $3, $4) ", p.ID, p.TotalMetric, p.CountMetric, p.Effects)
	if e != nil {
		return e
	}

	// now delete it.
	a.DeleteByID(ctx, id)
	return nil
}

// Save (and use database default ID generation)
func (a *DBPercentAlert) Save(ctx context.Context, p *savepb.PercentAlert) (uint64, error) {
	qn := "DBPercentAlert_Save"
	rows, e := a.DB.QueryContext(ctx, qn, "insert into "+a.SQLTablename+" (totalmetric, countmetric, effects) values ($1, $2, $3) returning id", p.TotalMetric, p.CountMetric, p.Effects)
	if e != nil {
		return 0, a.Error(ctx, qn, e)
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, a.Error(ctx, qn, fmt.Errorf("No rows after insert"))
	}
	var id uint64
	e = rows.Scan(&id)
	if e != nil {
		return 0, a.Error(ctx, qn, fmt.Errorf("failed to scan id after insert: %s", e))
	}
	p.ID = id
	return id, nil
}

// Save using the ID specified
func (a *DBPercentAlert) SaveWithID(ctx context.Context, p *savepb.PercentAlert) error {
	qn := "insert_DBPercentAlert"
	_, e := a.DB.ExecContext(ctx, qn, "insert into "+a.SQLTablename+" (id,totalmetric, countmetric, effects) values ($1,$2, $3, $4) ", p.ID, p.TotalMetric, p.CountMetric, p.Effects)
	return a.Error(ctx, qn, e)
}

func (a *DBPercentAlert) Update(ctx context.Context, p *savepb.PercentAlert) error {
	qn := "DBPercentAlert_Update"
	_, e := a.DB.ExecContext(ctx, qn, "update "+a.SQLTablename+" set totalmetric=$1, countmetric=$2, effects=$3 where id = $4", p.TotalMetric, p.CountMetric, p.Effects, p.ID)

	return a.Error(ctx, qn, e)
}

// delete by id field
func (a *DBPercentAlert) DeleteByID(ctx context.Context, p uint64) error {
	qn := "deleteDBPercentAlert_ByID"
	_, e := a.DB.ExecContext(ctx, qn, "delete from "+a.SQLTablename+" where id = $1", p)
	return a.Error(ctx, qn, e)
}

// get it by primary id
func (a *DBPercentAlert) ByID(ctx context.Context, p uint64) (*savepb.PercentAlert, error) {
	qn := "DBPercentAlert_ByID"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,totalmetric, countmetric, effects from "+a.SQLTablename+" where id = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByID: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByID: error scanning (%s)", e))
	}
	if len(l) == 0 {
		return nil, a.Error(ctx, qn, fmt.Errorf("No PercentAlert with id %v", p))
	}
	if len(l) != 1 {
		return nil, a.Error(ctx, qn, fmt.Errorf("Multiple (%d) PercentAlert with id %v", len(l), p))
	}
	return l[0], nil
}

// get all rows
func (a *DBPercentAlert) All(ctx context.Context) ([]*savepb.PercentAlert, error) {
	qn := "DBPercentAlert_all"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,totalmetric, countmetric, effects from "+a.SQLTablename+" order by id")
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("All: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, fmt.Errorf("All: error scanning (%s)", e)
	}
	return l, nil
}

/**********************************************************************
* GetBy[FIELD] functions
**********************************************************************/

// get all "DBPercentAlert" rows with matching TotalMetric
func (a *DBPercentAlert) ByTotalMetric(ctx context.Context, p string) ([]*savepb.PercentAlert, error) {
	qn := "DBPercentAlert_ByTotalMetric"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,totalmetric, countmetric, effects from "+a.SQLTablename+" where totalmetric = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByTotalMetric: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByTotalMetric: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBPercentAlert) ByLikeTotalMetric(ctx context.Context, p string) ([]*savepb.PercentAlert, error) {
	qn := "DBPercentAlert_ByLikeTotalMetric"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,totalmetric, countmetric, effects from "+a.SQLTablename+" where totalmetric ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByTotalMetric: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByTotalMetric: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBPercentAlert" rows with matching CountMetric
func (a *DBPercentAlert) ByCountMetric(ctx context.Context, p string) ([]*savepb.PercentAlert, error) {
	qn := "DBPercentAlert_ByCountMetric"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,totalmetric, countmetric, effects from "+a.SQLTablename+" where countmetric = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByCountMetric: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByCountMetric: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBPercentAlert) ByLikeCountMetric(ctx context.Context, p string) ([]*savepb.PercentAlert, error) {
	qn := "DBPercentAlert_ByLikeCountMetric"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,totalmetric, countmetric, effects from "+a.SQLTablename+" where countmetric ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByCountMetric: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByCountMetric: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBPercentAlert" rows with matching Effects
func (a *DBPercentAlert) ByEffects(ctx context.Context, p uint32) ([]*savepb.PercentAlert, error) {
	qn := "DBPercentAlert_ByEffects"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,totalmetric, countmetric, effects from "+a.SQLTablename+" where effects = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByEffects: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByEffects: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBPercentAlert) ByLikeEffects(ctx context.Context, p uint32) ([]*savepb.PercentAlert, error) {
	qn := "DBPercentAlert_ByLikeEffects"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,totalmetric, countmetric, effects from "+a.SQLTablename+" where effects ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByEffects: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByEffects: error scanning (%s)", e))
	}
	return l, nil
}

/**********************************************************************
* Helper to convert from an SQL Query
**********************************************************************/

// from a query snippet (the part after WHERE)
func (a *DBPercentAlert) FromQuery(ctx context.Context, query_where string, args ...interface{}) ([]*savepb.PercentAlert, error) {
	rows, err := a.DB.QueryContext(ctx, "custom_query_"+a.Tablename(), "select "+a.SelectCols()+" from "+a.Tablename()+" where "+query_where, args...)
	if err != nil {
		return nil, err
	}
	return a.FromRows(ctx, rows)
}

/**********************************************************************
* Helper to convert from an SQL Row to struct
**********************************************************************/
func (a *DBPercentAlert) Tablename() string {
	return a.SQLTablename
}

func (a *DBPercentAlert) SelectCols() string {
	return "id,totalmetric, countmetric, effects"
}
func (a *DBPercentAlert) SelectColsQualified() string {
	return "" + a.SQLTablename + ".id," + a.SQLTablename + ".totalmetric, " + a.SQLTablename + ".countmetric, " + a.SQLTablename + ".effects"
}

func (a *DBPercentAlert) FromRows(ctx context.Context, rows *gosql.Rows) ([]*savepb.PercentAlert, error) {
	var res []*savepb.PercentAlert
	for rows.Next() {
		foo := savepb.PercentAlert{}
		err := rows.Scan(&foo.ID, &foo.TotalMetric, &foo.CountMetric, &foo.Effects)
		if err != nil {
			return nil, a.Error(ctx, "fromrow-scan", err)
		}
		res = append(res, &foo)
	}
	return res, nil
}

/**********************************************************************
* Helper to create table and columns
**********************************************************************/
func (a *DBPercentAlert) CreateTable(ctx context.Context) error {
	csql := []string{
		`create sequence if not exists ` + a.SQLTablename + `_seq;`,
		`CREATE TABLE if not exists ` + a.SQLTablename + ` (id integer primary key default nextval('` + a.SQLTablename + `_seq'),totalmetric text not null  ,countmetric text not null  ,effects integer not null  );`,
		`CREATE TABLE if not exists ` + a.SQLTablename + `_archive (id integer primary key default nextval('` + a.SQLTablename + `_seq'),totalmetric text not null  ,countmetric text not null  ,effects integer not null  );`,
	}
	for i, c := range csql {
		_, e := a.DB.ExecContext(ctx, fmt.Sprintf("create_"+a.SQLTablename+"_%d", i), c)
		if e != nil {
			return e
		}
	}
	return nil
}

/**********************************************************************
* Helper to meaningful errors
**********************************************************************/
func (a *DBPercentAlert) Error(ctx context.Context, q string, e error) error {
	if e == nil {
		return nil
	}
	return fmt.Errorf("[table="+a.SQLTablename+", query=%s] Error: %s", q, e)
}




