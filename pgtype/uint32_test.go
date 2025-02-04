package pgtype_test

import (
	"context"
	"testing"

	"github.com/thoohv5/pgx/pgtype"
	"github.com/thoohv5/pgx/pgxtest"
)

func TestUint32Codec(t *testing.T) {
	pgxtest.RunValueRoundTripTests(context.Background(), t, defaultConnTestRunner, pgxtest.KnownOIDQueryExecModes, "oid", []pgxtest.ValueRoundTripTest{
		{
			pgtype.Uint32{Uint32: pgtype.TextOID, Valid: true},
			new(pgtype.Uint32),
			isExpectedEq(pgtype.Uint32{Uint32: pgtype.TextOID, Valid: true}),
		},
		{pgtype.Uint32{}, new(pgtype.Uint32), isExpectedEq(pgtype.Uint32{})},
		{nil, new(pgtype.Uint32), isExpectedEq(pgtype.Uint32{})},
	})
}
