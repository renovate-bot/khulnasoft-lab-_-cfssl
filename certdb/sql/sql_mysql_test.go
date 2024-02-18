//go:build mysql
// +build mysql

package sql

import (
	"testing"

	"github.com/khulnasoft-lab/cfssl/certdb/testdb"
)

func TestMySQL(t *testing.T) {
	db := testdb.MySQLDB()
	ta := TestAccessor{
		Accessor: NewAccessor(db),
		DB:       db,
	}
	testEverything(ta, t)
}
