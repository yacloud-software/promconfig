package db

/*
 This file was created by mkdb-client.
 The intention is not to modify this file, but you may extend the struct DBPercentAlert
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
ALTER TABLE percentalert ADD COLUMN IF NOT EXISTS totalmetric text not null default '';
ALTER TABLE percentalert ADD COLUMN IF NOT EXISTS countmetric text not null default '';
ALTER TABLE percentalert ADD COLUMN IF NOT EXISTS effects integer not null default 0;


Archive Table: (structs can be moved from main to archive using Archive() function)

 CREATE TABLE percentalert_archive (id integer unique not null,totalmetric text not null,countmetric text not null,effects integer not null);
*/

import (
	"context"
	gosql "database/sql"
	"fmt"
	savepb "golang.conradwood.net/apis/promconfig"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/sql"
	"os"
	"sync"
)

var (
	default_def_DBPercentAlert *DBPercentAlert
)

type DBPercentAlert struct {
	DB                   *sql.DB
	SQLTablename         string
	SQLArchivetablename  string
	customColumnHandlers []CustomColumnHandler
	lock                 sync.Mutex
}

func init() {
	RegisterDBHandlerFactory(func() Handler {
		return DefaultDBPercentAlert()
	})
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

func (a *DBPercentAlert) GetCustomColumnHandlers() []CustomColumnHandler {
	return a.customColumnHandlers
}
func (a *DBPercentAlert) AddCustomColumnHandler(w CustomColumnHandler) {
	a.lock.Lock()
	a.customColumnHandlers = append(a.customColumnHandlers, w)
	a.lock.Unlock()
}

func (a *DBPercentAlert) NewQuery() *Query {
	return newQuery(a)
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

// return a map with columnname -> value_from_proto
func (a *DBPercentAlert) buildSaveMap(ctx context.Context, p *savepb.PercentAlert) (map[string]interface{}, error) {
	extra, err := extraFieldsToStore(ctx, a, p)
	if err != nil {
		return nil, err
	}
	res := make(map[string]interface{})
	res["id"] = a.get_col_from_proto(p, "id")
	res["totalmetric"] = a.get_col_from_proto(p, "totalmetric")
	res["countmetric"] = a.get_col_from_proto(p, "countmetric")
	res["effects"] = a.get_col_from_proto(p, "effects")
	if extra != nil {
		for k, v := range extra {
			res[k] = v
		}
	}
	return res, nil
}

func (a *DBPercentAlert) Save(ctx context.Context, p *savepb.PercentAlert) (uint64, error) {
	qn := "save_DBPercentAlert"
	smap, err := a.buildSaveMap(ctx, p)
	if err != nil {
		return 0, err
	}
	delete(smap, "id") // save without id
	return a.saveMap(ctx, qn, smap, p)
}

// Save using the ID specified
func (a *DBPercentAlert) SaveWithID(ctx context.Context, p *savepb.PercentAlert) error {
	qn := "insert_DBPercentAlert"
	smap, err := a.buildSaveMap(ctx, p)
	if err != nil {
		return err
	}
	_, err = a.saveMap(ctx, qn, smap, p)
	return err
}

// use a hashmap of columnname->values to store to database (see buildSaveMap())
func (a *DBPercentAlert) saveMap(ctx context.Context, queryname string, smap map[string]interface{}, p *savepb.PercentAlert) (uint64, error) {
	// Save (and use database default ID generation)

	var rows *gosql.Rows
	var e error

	q_cols := ""
	q_valnames := ""
	q_vals := make([]interface{}, 0)
	deli := ""
	i := 0
	// build the 2 parts of the query (column names and value names) as well as the values themselves
	for colname, val := range smap {
		q_cols = q_cols + deli + colname
		i++
		q_valnames = q_valnames + deli + fmt.Sprintf("$%d", i)
		q_vals = append(q_vals, val)
		deli = ","
	}
	rows, e = a.DB.QueryContext(ctx, queryname, "insert into "+a.SQLTablename+" ("+q_cols+") values ("+q_valnames+") returning id", q_vals...)
	if e != nil {
		return 0, a.Error(ctx, queryname, e)
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, a.Error(ctx, queryname, errors.Errorf("No rows after insert"))
	}
	var id uint64
	e = rows.Scan(&id)
	if e != nil {
		return 0, a.Error(ctx, queryname, errors.Errorf("failed to scan id after insert: %s", e))
	}
	p.ID = id
	return id, nil
}

// if ID==0 save, otherwise update
func (a *DBPercentAlert) SaveOrUpdate(ctx context.Context, p *savepb.PercentAlert) error {
	if p.ID == 0 {
		_, err := a.Save(ctx, p)
		return err
	}
	return a.Update(ctx, p)
}
func (a *DBPercentAlert) Update(ctx context.Context, p *savepb.PercentAlert) error {
	qn := "DBPercentAlert_Update"
	_, e := a.DB.ExecContext(ctx, qn, "update "+a.SQLTablename+" set totalmetric=$1, countmetric=$2, effects=$3 where id = $4", a.get_TotalMetric(p), a.get_CountMetric(p), a.get_Effects(p), p.ID)

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
	l, e := a.fromQuery(ctx, qn, "id = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("ByID: error scanning (%s)", e))
	}
	if len(l) == 0 {
		return nil, a.Error(ctx, qn, errors.Errorf("No PercentAlert with id %v", p))
	}
	if len(l) != 1 {
		return nil, a.Error(ctx, qn, errors.Errorf("Multiple (%d) PercentAlert with id %v", len(l), p))
	}
	return l[0], nil
}

// get it by primary id (nil if no such ID row, but no error either)
func (a *DBPercentAlert) TryByID(ctx context.Context, p uint64) (*savepb.PercentAlert, error) {
	qn := "DBPercentAlert_TryByID"
	l, e := a.fromQuery(ctx, qn, "id = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("TryByID: error scanning (%s)", e))
	}
	if len(l) == 0 {
		return nil, nil
	}
	if len(l) != 1 {
		return nil, a.Error(ctx, qn, errors.Errorf("Multiple (%d) PercentAlert with id %v", len(l), p))
	}
	return l[0], nil
}

// get it by multiple primary ids
func (a *DBPercentAlert) ByIDs(ctx context.Context, p []uint64) ([]*savepb.PercentAlert, error) {
	qn := "DBPercentAlert_ByIDs"
	l, e := a.fromQuery(ctx, qn, "id in $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("TryByID: error scanning (%s)", e))
	}
	return l, nil
}

// get all rows
func (a *DBPercentAlert) All(ctx context.Context) ([]*savepb.PercentAlert, error) {
	qn := "DBPercentAlert_all"
	l, e := a.fromQuery(ctx, qn, "true")
	if e != nil {
		return nil, errors.Errorf("All: error scanning (%s)", e)
	}
	return l, nil
}

/**********************************************************************
* GetBy[FIELD] functions
**********************************************************************/

// get all "DBPercentAlert" rows with matching TotalMetric
func (a *DBPercentAlert) ByTotalMetric(ctx context.Context, p string) ([]*savepb.PercentAlert, error) {
	qn := "DBPercentAlert_ByTotalMetric"
	l, e := a.fromQuery(ctx, qn, "totalmetric = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("ByTotalMetric: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBPercentAlert" rows with multiple matching TotalMetric
func (a *DBPercentAlert) ByMultiTotalMetric(ctx context.Context, p []string) ([]*savepb.PercentAlert, error) {
	qn := "DBPercentAlert_ByTotalMetric"
	l, e := a.fromQuery(ctx, qn, "totalmetric in $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("ByTotalMetric: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBPercentAlert) ByLikeTotalMetric(ctx context.Context, p string) ([]*savepb.PercentAlert, error) {
	qn := "DBPercentAlert_ByLikeTotalMetric"
	l, e := a.fromQuery(ctx, qn, "totalmetric ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("ByTotalMetric: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBPercentAlert" rows with matching CountMetric
func (a *DBPercentAlert) ByCountMetric(ctx context.Context, p string) ([]*savepb.PercentAlert, error) {
	qn := "DBPercentAlert_ByCountMetric"
	l, e := a.fromQuery(ctx, qn, "countmetric = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("ByCountMetric: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBPercentAlert" rows with multiple matching CountMetric
func (a *DBPercentAlert) ByMultiCountMetric(ctx context.Context, p []string) ([]*savepb.PercentAlert, error) {
	qn := "DBPercentAlert_ByCountMetric"
	l, e := a.fromQuery(ctx, qn, "countmetric in $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("ByCountMetric: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBPercentAlert) ByLikeCountMetric(ctx context.Context, p string) ([]*savepb.PercentAlert, error) {
	qn := "DBPercentAlert_ByLikeCountMetric"
	l, e := a.fromQuery(ctx, qn, "countmetric ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("ByCountMetric: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBPercentAlert" rows with matching Effects
func (a *DBPercentAlert) ByEffects(ctx context.Context, p uint32) ([]*savepb.PercentAlert, error) {
	qn := "DBPercentAlert_ByEffects"
	l, e := a.fromQuery(ctx, qn, "effects = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("ByEffects: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBPercentAlert" rows with multiple matching Effects
func (a *DBPercentAlert) ByMultiEffects(ctx context.Context, p []uint32) ([]*savepb.PercentAlert, error) {
	qn := "DBPercentAlert_ByEffects"
	l, e := a.fromQuery(ctx, qn, "effects in $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("ByEffects: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBPercentAlert) ByLikeEffects(ctx context.Context, p uint32) ([]*savepb.PercentAlert, error) {
	qn := "DBPercentAlert_ByLikeEffects"
	l, e := a.fromQuery(ctx, qn, "effects ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("ByEffects: error scanning (%s)", e))
	}
	return l, nil
}

/**********************************************************************
* The field getters
**********************************************************************/

// getter for field "ID" (ID) [uint64]
func (a *DBPercentAlert) get_ID(p *savepb.PercentAlert) uint64 {
	return uint64(p.ID)
}

// getter for field "TotalMetric" (TotalMetric) [string]
func (a *DBPercentAlert) get_TotalMetric(p *savepb.PercentAlert) string {
	return string(p.TotalMetric)
}

// getter for field "CountMetric" (CountMetric) [string]
func (a *DBPercentAlert) get_CountMetric(p *savepb.PercentAlert) string {
	return string(p.CountMetric)
}

// getter for field "Effects" (Effects) [uint32]
func (a *DBPercentAlert) get_Effects(p *savepb.PercentAlert) uint32 {
	return uint32(p.Effects)
}

/**********************************************************************
* Helper to convert from an SQL Query
**********************************************************************/

// from a query snippet (the part after WHERE)
func (a *DBPercentAlert) ByDBQuery(ctx context.Context, query *Query) ([]*savepb.PercentAlert, error) {
	extra_fields, err := extraFieldsToQuery(ctx, a)
	if err != nil {
		return nil, err
	}
	i := 0
	for col_name, value := range extra_fields {
		i++
		/*
		   efname:=fmt.Sprintf("EXTRA_FIELD_%d",i)
		   query.Add(col_name+" = "+efname,QP{efname:value})
		*/
		query.AddEqual(col_name, value)
	}

	gw, paras := query.ToPostgres()
	queryname := "custom_dbquery"
	rows, err := a.DB.QueryContext(ctx, queryname, "select "+a.SelectCols()+" from "+a.Tablename()+" where "+gw, paras...)
	if err != nil {
		return nil, err
	}
	res, err := a.FromRows(ctx, rows)
	rows.Close()
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (a *DBPercentAlert) FromQuery(ctx context.Context, query_where string, args ...interface{}) ([]*savepb.PercentAlert, error) {
	return a.fromQuery(ctx, "custom_query_"+a.Tablename(), query_where, args...)
}

// from a query snippet (the part after WHERE)
func (a *DBPercentAlert) fromQuery(ctx context.Context, queryname string, query_where string, args ...interface{}) ([]*savepb.PercentAlert, error) {
	extra_fields, err := extraFieldsToQuery(ctx, a)
	if err != nil {
		return nil, err
	}
	eq := ""
	if extra_fields != nil && len(extra_fields) > 0 {
		eq = " AND ("
		// build the extraquery "eq"
		i := len(args)
		deli := ""
		for col_name, value := range extra_fields {
			i++
			eq = eq + deli + col_name + fmt.Sprintf(" = $%d", i)
			deli = " AND "
			args = append(args, value)
		}
		eq = eq + ")"
	}
	rows, err := a.DB.QueryContext(ctx, queryname, "select "+a.SelectCols()+" from "+a.Tablename()+" where ( "+query_where+") "+eq, args...)
	if err != nil {
		return nil, err
	}
	res, err := a.FromRows(ctx, rows)
	rows.Close()
	if err != nil {
		return nil, err
	}
	return res, nil
}

/**********************************************************************
* Helper to convert from an SQL Row to struct
**********************************************************************/
func (a *DBPercentAlert) get_col_from_proto(p *savepb.PercentAlert, colname string) interface{} {
	if colname == "id" {
		return a.get_ID(p)
	} else if colname == "totalmetric" {
		return a.get_TotalMetric(p)
	} else if colname == "countmetric" {
		return a.get_CountMetric(p)
	} else if colname == "effects" {
		return a.get_Effects(p)
	}
	panic(fmt.Sprintf("in table \"%s\", column \"%s\" cannot be resolved to proto field name", a.Tablename(), colname))
}

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
		// SCANNER:
		foo := &savepb.PercentAlert{}
		// create the non-nullable pointers
		// create variables for scan results
		scanTarget_0 := &foo.ID
		scanTarget_1 := &foo.TotalMetric
		scanTarget_2 := &foo.CountMetric
		scanTarget_3 := &foo.Effects
		err := rows.Scan(scanTarget_0, scanTarget_1, scanTarget_2, scanTarget_3)
		// END SCANNER

		if err != nil {
			return nil, a.Error(ctx, "fromrow-scan", err)
		}
		res = append(res, foo)
	}
	return res, nil
}

/**********************************************************************
* Helper to create table and columns
**********************************************************************/
func (a *DBPercentAlert) CreateTable(ctx context.Context) error {
	csql := []string{
		`create sequence if not exists ` + a.SQLTablename + `_seq;`,
		`CREATE TABLE if not exists ` + a.SQLTablename + ` (id integer primary key default nextval('` + a.SQLTablename + `_seq'),totalmetric text not null ,countmetric text not null ,effects integer not null );`,
		`CREATE TABLE if not exists ` + a.SQLTablename + `_archive (id integer primary key default nextval('` + a.SQLTablename + `_seq'),totalmetric text not null ,countmetric text not null ,effects integer not null );`,
		`ALTER TABLE ` + a.SQLTablename + ` ADD COLUMN IF NOT EXISTS totalmetric text not null default '';`,
		`ALTER TABLE ` + a.SQLTablename + ` ADD COLUMN IF NOT EXISTS countmetric text not null default '';`,
		`ALTER TABLE ` + a.SQLTablename + ` ADD COLUMN IF NOT EXISTS effects integer not null default 0;`,

		`ALTER TABLE ` + a.SQLTablename + `_archive  ADD COLUMN IF NOT EXISTS totalmetric text not null  default '';`,
		`ALTER TABLE ` + a.SQLTablename + `_archive  ADD COLUMN IF NOT EXISTS countmetric text not null  default '';`,
		`ALTER TABLE ` + a.SQLTablename + `_archive  ADD COLUMN IF NOT EXISTS effects integer not null  default 0;`,
	}

	for i, c := range csql {
		_, e := a.DB.ExecContext(ctx, fmt.Sprintf("create_"+a.SQLTablename+"_%d", i), c)
		if e != nil {
			return e
		}
	}

	// these are optional, expected to fail
	csql = []string{
		// Indices:

		// Foreign keys:

	}
	for i, c := range csql {
		a.DB.ExecContextQuiet(ctx, fmt.Sprintf("create_"+a.SQLTablename+"_%d", i), c)
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
	return errors.Errorf("[table="+a.SQLTablename+", query=%s] Error: %s", q, e)
}

