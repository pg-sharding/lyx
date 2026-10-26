package lyx_test

import (
	"testing"

	"github.com/pg-sharding/lyx/lyx"
	"github.com/stretchr/testify/require"
)

func TestAlterPostgresTargets(t *testing.T) {
	for _, query := range []string{
		"ALTER AGGREGATE sum(integer) RENAME TO total;",
		"ALTER COLLATION coll RENAME TO new_coll;",
		"ALTER CONVERSION conv RENAME TO new_conv;",
		"ALTER DATABASE db RENAME TO new_db;",
		"ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO reader;",
		"ALTER DOMAIN domain_name SET NOT NULL;",
		"ALTER EVENT TRIGGER event_trigger ENABLE;",
		"ALTER EXTENSION ext UPDATE;",
		"ALTER FOREIGN DATA WRAPPER fdw OPTIONS (ADD host 'localhost');",
		"ALTER FOREIGN TABLE entries ADD COLUMN id integer;",
		"ALTER FUNCTION func(integer) IMMUTABLE;",
		"ALTER GROUP readers ADD USER reader;",
		"ALTER INDEX entries_idx RENAME TO new_idx;",
		"ALTER LANGUAGE plpgsql RENAME TO new_lang;",
		"ALTER MATERIALIZED VIEW entries RENAME TO new_entries;",
		"ALTER OPERATOR + (integer, integer) OWNER TO owner;",
		"ALTER OPERATOR CLASS ops USING btree RENAME TO new_ops;",
		"ALTER OPERATOR FAMILY ops USING btree RENAME TO new_ops;",
		"ALTER POLICY policy ON entries RENAME TO new_policy;",
		"ALTER PROCEDURAL LANGUAGE plpgsql OWNER TO owner;",
		"ALTER PROCEDURE proc(integer) SECURITY DEFINER;",
		"ALTER PUBLICATION pub ADD TABLE entries;",
		"ALTER ROLE reader SET search_path TO public;",
		"ALTER ROUTINE routine(integer) RENAME TO new_routine;",
		"ALTER RULE rule_name ON entries RENAME TO new_rule;",
		"ALTER SCHEMA old_schema RENAME TO new_schema;",
		"ALTER SEQUENCE seq RENAME TO new_seq;",
		"ALTER SERVER server_name OPTIONS (SET host 'localhost');",
		"ALTER STATISTICS stats RENAME TO new_stats;",
		"ALTER SUBSCRIPTION sub DISABLE;",
		"ALTER SYSTEM SET work_mem TO '64MB';",
		"ALTER TABLE entries ADD COLUMN id integer;",
		"ALTER TABLESPACE space RENAME TO new_space;",
		"ALTER TEXT SEARCH CONFIGURATION cfg RENAME TO new_cfg;",
		"ALTER TEXT SEARCH DICTIONARY dict RENAME TO new_dict;",
		"ALTER TEXT SEARCH PARSER parser_name RENAME TO new_parser;",
		"ALTER TEXT SEARCH TEMPLATE template_name RENAME TO new_template;",
		"ALTER TRIGGER trigger_name ON entries RENAME TO new_trigger;",
		"ALTER TYPE type_name ADD VALUE 'new_value';",
		"ALTER USER reader SET search_path TO public;",
		"ALTER USER MAPPING FOR reader SERVER server_name OPTIONS (SET user 'reader');",
		"ALTER VIEW entries RENAME TO new_entries;",
		"/* comment */ alter /* target */ table entries add column id integer",
	} {
		t.Run(query, func(t *testing.T) {
			nodes, _, err := lyx.Parse(query)
			require.NoError(t, err)
			require.Len(t, nodes, 1)
			require.IsType(t, &lyx.Alter{}, nodes[0])
		})
	}
}

func TestAlterRejectsNonPostgresTargets(t *testing.T) {
	for _, query := range []string{
		"ALTER;",
		"ALTER SHARD sh1 OPTIONS (UNKNOWN USER 'user1');",
		"ALTER DISTRIBUTION ds1 ALTER RELATION a DISTRIBUTION KEY another SCHEMA schema_name;",
		"ALTER KEY RANGE kr1 ROUTE TO sh1;",
		"ALTER unknown RENAME TO other;",
		"ALTER ACCESS METHOD am RENAME TO new_am;",
		"ALTER FOREIGN SHARD sh1;",
		`ALTER FOREIGN "TABLE" entries;`,
		"ALTER EVENT event_name;",
		"ALTER MATERIALIZED TABLE entries;",
		"ALTER TEXT SEARCH SHARD sh1;",
		`ALTER TEXT SEARCH "CONFIGURATION" cfg RENAME TO new_cfg;`,
		"ALTER DEFAULT SHARD sh1;",
		"ALTER LARGE TABLE entries;",
		"ALTER PROCEDURAL ROUTINE func;",
		`ALTER "TABLE" entries ADD COLUMN id integer;`,
		"ALTER 'TABLE' entries ADD COLUMN id integer;",
		"/* comment */ alter /* target */ shard sh1;",
		"SELECT 1; ALTER SHARD sh1;",
	} {
		t.Run(query, func(t *testing.T) {
			_, _, err := lyx.Parse(query)
			require.Error(t, err)
		})
	}
}

func TestAlterWithReusedParser(t *testing.T) {
	parser := lyx.NewLyxParser()
	for _, query := range []string{
		"SELECT 1; ALTER TABLE entries ADD COLUMN id integer;",
		"/* different prefix length */ SELECT 2; ALTER VIEW entries RENAME TO new_entries;",
	} {
		nodes, _, err := lyx.ParseWithLexerParser(parser, lyx.NewStringTokenizer(""), query)
		require.NoError(t, err)
		require.Len(t, nodes, 2)
		require.IsType(t, &lyx.Select{}, nodes[0])
		require.IsType(t, &lyx.Alter{}, nodes[1])
	}
	_, _, err := lyx.ParseWithLexerParser(parser, lyx.NewStringTokenizer(""), "SELECT 1; ALTER SHARD sh1;")
	require.Error(t, err)
}
