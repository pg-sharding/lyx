package lyx_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pg-sharding/lyx/lyx"
)

func TestParseHandlesEscapedSingleQuotes(t *testing.T) {
	query := "SELECT 'value'';still''inside';"

	stmts, err := lyx.Parse(query)
	require.NoError(t, err)
	require.Len(t, stmts, 1)

	selectStmt, ok := stmts[0].(*lyx.Select)
	require.True(t, ok)

	literal, ok := selectStmt.TargetList[0].(*lyx.AExprSConst)
	require.True(t, ok)
	require.Equal(t, "value';still'inside", literal.Value)
}
