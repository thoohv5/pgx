package pgtype_test

import (
	"context"
	"testing"

	"github.com/thoohv5/pgx/pgtype"
	"github.com/thoohv5/pgx/pgxtest"
)

func TestBoxCodec(t *testing.T) {
	skipCockroachDB(t, "Server does not support box type")

	pgxtest.RunValueRoundTripTests(context.Background(), t, defaultConnTestRunner, nil, "box", []pgxtest.ValueRoundTripTest{
		{
			pgtype.Box{
				P:     [2]pgtype.Vec2{{7.1, 5.2345678}, {3.14, 1.678}},
				Valid: true,
			},
			new(pgtype.Box),
			isExpectedEq(pgtype.Box{
				P:     [2]pgtype.Vec2{{7.1, 5.2345678}, {3.14, 1.678}},
				Valid: true,
			}),
		},
		{
			pgtype.Box{
				P:     [2]pgtype.Vec2{{7.1, 5.2345678}, {-13.14, -5.234}},
				Valid: true,
			},
			new(pgtype.Box),
			isExpectedEq(pgtype.Box{
				P:     [2]pgtype.Vec2{{7.1, 5.2345678}, {-13.14, -5.234}},
				Valid: true,
			}),
		},
		{pgtype.Box{}, new(pgtype.Box), isExpectedEq(pgtype.Box{})},
		{nil, new(pgtype.Box), isExpectedEq(pgtype.Box{})},
	})
}
