package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"

	"gitea.tursom.cn/tursom/kvs/kv"
)

type (
	sqliteKVS struct {
		lang.BaseObject
		db    *sql.DB
		table string
	}
)

func New(db *sql.DB, table string) (kv.Store[[]byte, []byte], exceptions.Exception) {
	kvs := &sqliteKVS{
		db:    db,
		table: table,
	}
	if err := kvs.createTable(); err != nil {
		return nil, err
	}

	return kvs, nil
}

func (s *sqliteKVS) createTable() exceptions.Exception {
	if _, err := s.db.Exec("create table if not exists " + s.table + " (" +
		"k blob primary key not null," +
		"v blob" +
		")"); err != nil {
		return exceptions.Package(err)
	}

	return nil
}

func (s *sqliteKVS) Get(key []byte) ([]byte, exceptions.Exception) {
	rows, err := s.db.Query("select v from "+s.table+" where k = ?", key)
	if err != nil {
		return nil, exceptions.Package(err)
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var data []byte
	err = rows.Scan(&data)
	if err != nil {
		return nil, exceptions.Package(err)
	}

	return data, nil
}

func (s *sqliteKVS) Put(key []byte, value []byte) exceptions.Exception {
	if value == nil {
		value = []byte{}
	}

	exec, err := s.db.Exec("update "+s.table+" set v=? where k=?", value, key)
	if err != nil {
		return exceptions.Package(err)
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		return exceptions.Package(err)
	}

	if affected != 0 {
		return nil
	}

	if _, err = s.db.Exec("insert into "+s.table+" (k,v) values (?,?)", key, value); err != nil {
		return exceptions.Package(err)
	}

	return nil
}

func (s *sqliteKVS) Delete(key []byte) exceptions.Exception {
	if _, err := s.db.Exec("delete from "+s.table+" where k=?", key); err != nil {
		return exceptions.Package(err)
	}

	return nil
}
