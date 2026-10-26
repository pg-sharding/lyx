package lyx

import "strings"

// isPostgresAlterTarget checks the ALTER prefixes from PostgreSQL's gram.y.
// The rest of the command is still passed through without validation.
func isPostgresAlterTarget(sql []byte) bool {
	lexer := NewLexer(sql)
	var token yySymType
	nextKeyword := func() string {
		if lexer.Lex(&token) == 0 {
			return ""
		}
		// Use the raw spelling so quoted identifiers and strings cannot
		// stand in for keywords, even when the lexer returns IDENT.
		return strings.ToUpper(string(sql[lexer.ts:lexer.te]))
	}
	if nextKeyword() != "ALTER" {
		return false
	}
	switch nextKeyword() {
	case "AGGREGATE", "COLLATION", "CONVERSION", "DATABASE", "DOMAIN",
		"EXTENSION", "FUNCTION", "GROUP", "INDEX", "LANGUAGE", "OPERATOR",
		"POLICY", "PROCEDURE", "PUBLICATION", "ROLE", "ROUTINE", "RULE",
		"SCHEMA", "SEQUENCE", "SERVER", "STATISTICS", "SUBSCRIPTION",
		"SYSTEM", "TABLE", "TABLESPACE", "TRIGGER", "TYPE", "USER", "VIEW":
		return true
	case "DEFAULT":
		return nextKeyword() == "PRIVILEGES"
	case "EVENT":
		return nextKeyword() == "TRIGGER"
	case "FOREIGN":
		switch nextKeyword() {
		case "TABLE":
			return true
		case "DATA":
			return nextKeyword() == "WRAPPER"
		}
	case "LARGE":
		return nextKeyword() == "OBJECT"
	case "MATERIALIZED":
		return nextKeyword() == "VIEW"
	case "PROCEDURAL":
		return nextKeyword() == "LANGUAGE"
	case "TEXT":
		if nextKeyword() == "SEARCH" {
			switch nextKeyword() {
			case "CONFIGURATION", "DICTIONARY", "PARSER", "TEMPLATE":
				return true
			}
		}
	}
	return false
}
