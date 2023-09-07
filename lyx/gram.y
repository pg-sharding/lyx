
%{
package lyx

import (
	"crypto/rand"
	"encoding/hex"
)

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
	  return "", err
	}
	return hex.EncodeToString(bytes), nil
}


type LyxParser yyParser

func NewLyxParser() LyxParser {
	return yyNewParser()
}

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union {
	str                    string
    strlist                []string
	byte                   byte
	bytes                  []byte
	int                    int
	bool                   bool
	empty                  struct{}

	node				   Node
    from_list              []FromClauseNode
    from                   FromClauseNode
    tableelt               []TableElt
    tableref               FromClauseNode

    nodeList              []Node
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
//%type <val> expr number


// CMDS
%type <node> command
%type <node> routable_statement

// same for terminals
%token <str> SCONST IDENT

%type<str> any_id any_val table_name

%type<bool> opt_program

%type<node> where_clause

%type<from_list> from_clause from_list

%type<node> DeallocateStmt execute_param_clause ExecuteStmt PreparableStmt

%type<from> table_ref 
%type<tableref> relation_expr joined_table

%type<node> func_arg_expr
%type<bool> opt_ordinality copy_from
%type<nodeList> func_arg_list
%type<node> func_alias_clause

%type<node> a_expr c_expr b_expr

/* CIUD */
%token<str> CREATE ALTER

%token<str> ROLE DATABASE NATURAL

%token<str> TABLE INDEX
/* SQL */
%token<str> SELECT UPDATE INSERT DELETE FROM WHERE AND VALUES OR

/* basic words */
%token<str> CALLED ABORT_P COMMENTS EXTRACT THEN CONTENT_P RULE ELSE NFKC FUNCTION UNKNOWN XMLELEMENT DECLARE SUBSCRIPTION INCLUDING ADD_P CHAIN EXPRESSION INSENSITIVE SAVEPOINT CONCURRENTLY INPUT_P REINDEX CONSTRAINT DEFERRED CONNECTION REPLACE ASSIGNMENT FINALIZE RECHECK XMLFOREST CURRENT_TIMESTAMP INITIALLY COMPRESSION READ CAST SCHEMAS POSITION LEAKPROOF RESTRICT UNLISTEN USER CHARACTERISTICS IMPORT_P SCROLL SHARE CURRENT_DATE REASSIGN ROUTINES XML_P EXTERNAL ROLLUP IDENTITY_P INCREMENT LARGE_P MINVALUE NAME_P YES_P GENERATED GROUPS PLANS SYSTEM_P AUTHORIZATION LOCATION SYSID JSON XMLTABLE CURRENT_ROLE STORAGE OPERATOR OWNED OPTION PASSWORD SIMPLE SNAPSHOT TEXT_P INOUT WRITE FAMILY RESTART TYPE_P GRANT UNION DETACH ACTION FOLLOWING UNIQUE GRANTED SEQUENCE TIES GROUPING INDENT CONVERSION_P ENCRYPTED TABLES STABLE JSON_ARRAY LEAST ATOMIC NFD STRICT_P ALWAYS MATERIALIZED PREPARED TRANSFORM CURRENT_P LOGGED VALIDATOR CASCADE DISABLE_P REFERENCING SESSION MAPPING NOTIFY PARTITION STATEMENT JSON_ARRAYAGG CURRENT_SCHEMA CUBE FILTER SQL_P OVERLAPS NAMES ROWS NORMALIZE EVENT OTHERS BETWEEN RETURN SERVER TEMPLATE GROUP_P AFTER POLICY UNBOUNDED UNTIL EXCEPT ILIKE IN_P IMMEDIATE NEXT CONFIGURATION INCLUDE NEW SHOW EXTENSION LABEL JSON_OBJECTAGG DICTIONARY LISTEN NULLS_P AT DOMAIN_P OF REPLICA VALIDATE TABLESAMPLE END_P COLUMNS INHERIT SERIALIZABLE BREADTH OWNER SETS VIEWS MATCH ACCESS LOCK_P OPTIONS CURSOR UESCAPE SOME CATALOG_P COST LOCKED RECURSIVE CURRENT_CATALOG PARAMETER PRIOR PLACING AGGREGATE EXCLUSIVE CONTINUE_P DEFINER MODE REPEATABLE REFRESH SCALAR VALUE_P XMLSERIALIZE DEPTH NOTHING OVER START WHEN FORMAT NONE OFF STANDALONE_P HOLD SCHEMA WORK LOCALTIME SKIP HANDLER KEYS LEADING LEVEL CHECK INLINE_P LIKE OLD SESSION_USER CACHE ISOLATION WHITESPACE_P SUBSTRING DISCARD PROCEDURE LANGUAGE OVERRIDING EXCLUDE TRAILING ALSO ASENSITIVE INDEXES REF_P SECURITY VARIADIC COMMITTED FORWARD MERGE TRANSACTION FUNCTIONS INVOKER XMLATTRIBUTES SYSTEM_USER PARSER PUBLICATION MAXVALUE ROUTINE SUPPORT BOTH BACKWARD RELATIVE_P GREATEST XMLNAMESPACES CLASS NOWAIT REVOKE XMLEXISTS IMMUTABLE SEARCH TRUSTED DEFERRABLE ASSERTION CHECKPOINT EACH ATTRIBUTE METHOD NFKD VERSION_P WRAPPER ENABLE_P SEQUENCES XMLCONCAT LOCALTIMESTAMP RETURNS CONSTRAINTS MATCHED JSON_OBJECT ATTACH BEFORE BEGIN_P PRIVILEGES VOLATILE OBJECT_P TREAT HAVING COMMENT CYCLE LOAD COALESCE NFC RELEASE ANY CLOSE OUT_P PASSING STATISTICS CURRENT_TIME CONFLICT PROCEDURES RENAME UNENCRYPTED ABSOLUTE_P PRECEDING JSON_SERIALIZE XMLPARSE INHERITS VALID CURRENT_USER ABSENT CALL COLLATION CASE PARALLEL UNCOMMITTED IMPLICIT_P MOVE OVERLAY SIMILAR SYMMETRIC TRIM ANALYSE NULLIF XMLROOT VERBOSE ADMIN PARTIAL VIEW XMLPI CASCADED INSTEAD STORED TRIGGER TYPES_P DEFAULTS DEPENDS RANGE DO INTERSECT COLUMN ENUM_P EXCLUDING STRIP_P WITHIN JSON_SCALAR PROCEDURAL 

/* misc */
%token<str> EXPLAIN
%token<str> EXECUTE PREPARE

/* tx management */
%token<str> BEGIN COMMIT ROLLBACK

%token<str> VACUUM CLUSTER ANALYZE

/* insert into */
%token<str> INTO

/* JOINS */
%token<str> OUTER_P FULL RIGHT LEFT INNER_P

/* drop/trtuncate stmts */
%token<str> DROP TRUNCATE

%token<str> FETCH FOR ORDER GROUP BY OFFSET LIMIT WINDOW

/* CROSS */
%token<str> CROSS

/* variables */
%token<str> SET RESET LOCAL ALL

/* for tables */
%token<str> PRIMARY KEY FOREIGN REFERENCES

/* uning for delete */
%token<str> USING

/* ':' & ',' */
%token<str> TCOLON TCOMMA

/* '(' & ')' */
%token<str> TOPENBR TCLOSEBR


/* '[' & ']' */
%token<str> TSQOPENBR TSQCLOSEBR

/* ';' != */
%token<str> TSEMICOLON TNOT_EQUALS

/* '.' */
%token<str> TDOT

%token<str> NULL_P ISNULL ON JOIN

/* + - * / */
%token<str> TPLUS TMINUS TMUL TDIV

/* % ^ < <= >=  */
%token<str>  TGREATER TGREATER_EQUALS TLESS TLESS_EQUALS TMOD TPOW

/* copy */
%token<str> DELIMITERS PROGRAM STDIN

%token<str> FALSE_P TRUE_P 
%token<strlist> copy_generic_opt_list

/* copy opt item  */
%token<str> BINARY FREEZE DELIMITER CSV HEADER_P QUOTE ESCAPE FORCE ENCODING


/* '=' */
%token<str> TEQ

/*

Operator:
 							if (c == '~' || c == '!' || c == '@' ||
								c == '#' || c == '^' || c == '&' ||
								c == '|' || c == '`' || c == '?' ||
								c == '%')
								break; */
%token<str> OP

/* for update */
%token<str> ONLY

/*  */
%token<str> RETURNING

/* COPY  */
%token<str> COPY 


/* ARRAY  */
%token<str> ARRAY 

/* ROW  */
%token<str> ROW 

/* TYPES */ 
%token<str> SETOF INT_P  INTEGER  SMALLINT BIGINT REAL  FLOAT_P DOUBLE_P DECIMAL_P  DEC NUMERIC BOOLEAN_P BIT

/* TYPES */ 
%token<str> YEAR_P MONTH_P DAY_P HOUR_P MINUTE_P SECOND_P

%token<str> CHARACTER CHAR_P VARCHAR NATIONAL NCHAR

%token<str> PRECISION  VARYING TIMESTAMP TIME INTERVAL WITHOUT ZONE

%token<str> IF_P DEALLOCATE EXISTS

%token<str> DELETE_P OIDS PRESERVE TABLESPACE

%token<int> PARAM

/* IS NOT NULL */
%token<str> IS NOT NULL DISTINCT DEFAULT

/* NORMALIZED */
%token<str> NORMALIZED

/* NOTNULL */
%token<str> NOTNULL

/* DOCUMENT_P */
%token<str> DOCUMENT_P

/* ASYMMETRIC */
%token<str> ASYMMETRIC

/* order by  */
%token<str> ASC DESC NULLS_LA FIRST_P LAST_P

/* */
%token<str> TO STDOUT

%token<str> TEMPORARY GLOBAL TEMP UNLOGGED

/* LATERAL */
%token<str> LATERAL_P

/* WITH_LA ORDINALITY */
%token<str> ORDINALITY WITH_LA WITH

/* COLLATE */
%token<str> COLLATE

%token<str> AS

%token<str> DATA_P NO


%type<str> reserved_keyword col_name_keyword type_func_name_keyword bare_label_keyword unreserved_keyword


/**/

%type<node> ExecuteStmt
%type<node> SelectStmt
%type<node> InsertStmt
%type<node> UpdateStmt
%type<node> DeleteStmt
%type<node> PrepareStmt
%type<node> varset_stmt
%type<node> begin_stmt 
%type<node> commit_stmt
%type<node> rollback_stmt
%type<node> create_stmt alter_stmt
%type<node> vacuum_stmt cluster_stmt analyze_stmt
%type<node> truncate_stmt drop_stmt
%type<str> semicolon_opt

%type<node> PreparableStmt

%type<str> opt_collate_clause

%type<node> func_application
%type<node> func_table
%type<node> func_expr_windowless
%type<node> AexprConst

%type<node> select_clause select_no_parens select_with_parens SelectStmt values_clause locked_rels_list 
%type<node> for_locking_strength for_locking_item for_locking_items opt_for_locking_clause having_clause
%type<node> grouping_sets_clause
%type<node> cube_clause rollup_clause empty_grouping_set group_by_item group_by_list group_clause
%type<node> first_or_next row_or_rows I_or_F_const select_fetch_first_value
%type<node> select_offset_value select_limit_value offset_clause limit_clause opt_select_limit select_limit
%type<node> opt_distinct_clause opt_all_clause distinct_clause simple_select window_clause window_definition_list

%type<node> copy_stmt

%type<node> target_el

%type<nodeList> opt_target_list
%type<nodeList> target_list

%type<node> ColRef qualColRef

%type<nodeList> expr_list

%type<str> opt_with_data

%type<str> ColId columnref

%type<str> NonReservedWord name database_name access_method index_name

%type<strlist> name_list 

%type<str> reloptions reloption_list opt_reloptions

%type<strlist> comma_separated_col_refs insert_col_refs

%type<nodeList> insert_comma_separated_tuples insert_tuples

%type<strlist> opt_insert_col_refs columnList opt_column_list copy_gengeneric_opt_list copy_generic_opt_arg_list copy_opt_list
%type<str> any_tok  columnElem copy_gengeneric_opt_elem copy_generic_opt_arg_list_item copy_opt_item

%type<str> OptTemp

%type<strlist> copy_options

%type<str> attr_name ColLabel file_name
%type<from	> qualified_name

%type<strlist> opt_using delete_comma_separated_using_refs
%type<strlist> set_clause_list

%type<tableelt> create_stmt_coldefs

%type<str> opt_only opt_alias_clause alias_clause

%type<str> anything set_clause

%type<str> operator

%type<str> TableFuncElement
%type<strlist> TableFuncElementList OptTableFuncElementList

/*  TYPES */
%type<str>  interval_second opt_interval opt_timezone ConstInterval ConstDatetime opt_varying character CharacterWithoutLength CharacterWithLength ConstCharacter Character
%type<str>  BitWithoutLength BitWithLength ConstBit Bit opt_float Numeric opt_type_modifiers  GenericType ConstTypename SimpleTypename opt_array_bounds Typename
%type<str> type_function_name

%type<str> type_list prep_type_clause

%type<str> opt_indirection indirection opt_slice_bound indirection_el 

/* Precedence: lowest to highest */
%nonassoc	SET				/* see relation_expr_opt_alias */
%left		UNION EXCEPT
%left		INTERSECT
%left		OR
%left		AND
%right		NOT
%nonassoc	IS ISNULL NOTNULL	/* IS sets precedence for IS NULL, etc */
%nonassoc	TLESS TGREATER TEQ TLESS_EQUALS TGREATER_EQUALS TNOT_EQUALS
%nonassoc	BETWEEN IN_P LIKE ILIKE SIMILAR NOT_LA
%nonassoc	ESCAPE			/* ESCAPE must be just above LIKE/ILIKE/SIMILAR */
/*
 * To support target_el without AS, it used to be necessary to assign IDENT an
 * explicit precedence just less than Op.  While that's not really necessary
 * since we removed postfix operators, it's still helpful to do so because
 * there are some other unreserved keywords that need precedence assignments.
 * If those keywords have the same precedence as IDENT then they clearly act
 * the same as non-keywords, reducing the risk of unwanted precedence effects.
 *
 * We need to do this for PARTITION, RANGE, ROWS, and GROUPS to support
 * opt_existing_window_name (see comment there).
 *
 * The frame_bound productions UNBOUNDED PRECEDING and UNBOUNDED FOLLOWING
 * are even messier: since UNBOUNDED is an unreserved keyword (per spec!),
 * there is no principled way to distinguish these from the productions
 * a_expr PRECEDING/FOLLOWING.  We hack this up by giving UNBOUNDED slightly
 * lower precedence than PRECEDING and FOLLOWING.  At present this doesn't
 * appear to cause UNBOUNDED to be treated differently from other unreserved
 * keywords anywhere else in the grammar, but it's definitely risky.  We can
 * blame any funny behavior of UNBOUNDED on the SQL standard, though.
 *
 * To support CUBE and ROLLUP in GROUP BY without reserving them, we give them
 * an explicit priority lower than '(', so that a rule with CUBE '(' will shift
 * rather than reducing a conflicting rule that takes CUBE as a function name.
 * Using the same precedence as IDENT seems right for the reasons given above.
 */
%nonassoc	UNBOUNDED		/* ideally would have same precedence as IDENT */
%nonassoc	IDENT PARTITION RANGE ROWS GROUPS PRECEDING FOLLOWING CUBE ROLLUP
%left		Op OPERATOR		/* multi-character ops and user-defined operators */
%left		TPLUS TMINUS
%left		TMUL TDIV TMOD
%left		TPOW
/* Unary Operators */
%left		AT				/* sets precedence for AT TIME ZONE */
%left		COLLATE
%right		UMINUS
%left		TSQOPENBR TSQCLOSEBR
%left		TOPENBR TCLOSEBR
%left		TYPECAST
%left		TDOT
/*
 * These might seem to be low-precedence, but actually they are not part
 * of the arithmetic hierarchy at all in their use as JOIN operators.
 * We make them high-precedence to support their use as function names.
 * They wouldn't be given a precedence at all, were it not that we need
 * left-associativity among the JOIN rules themselves.
 */
%left		JOIN CROSS LEFT FULL RIGHT INNER_P NATURAL



/* unroutable query parts */
%type<str> opt_group_by_clause opt_window_clause sort_clause opt_limit_clause opt_offset_clause opt_fetch_clause opt_for_clause 

%start any_command

%%


any_command:
    command semicolon_opt

semicolon_opt:
/*empty*/ {}
| TSEMICOLON {
    
}


/*
 * Keyword category lists.  Generally, every keyword present in
 * the Postgres grammar should appear in exactly one of these lists.
 *
 * Put a new keyword into the first list that it can go into without causing
 * shift or reduce conflicts.  The earlier lists define "less reserved"
 * categories of keywords.
 *
 * Make sure that each keyword's category in kwlist.h matches where
 * it is listed here.  (Someday we may be able to generate these lists and
 * kwlist.h's table from one source of truth.)
 */

/* "Unreserved" keywords --- available for use as any kind of name.
 */
unreserved_keyword:
			  ABORT_P
			| ABSENT
			| ABSOLUTE_P
			| ACCESS
			| ACTION
			| ADD_P
			| ADMIN
			| AFTER
			| AGGREGATE
			| ALSO
			| ALTER
			| ALWAYS
			| ASENSITIVE
			| ASSERTION
			| ASSIGNMENT
			| AT
			| ATOMIC
			| ATTACH
			| ATTRIBUTE
			| BACKWARD
			| BEFORE
			| BEGIN_P
			| BREADTH
			| BY
			| CACHE
			| CALL
			| CALLED
			| CASCADE
			| CASCADED
			| CATALOG_P
			| CHAIN
			| CHARACTERISTICS
			| CHECKPOINT
			| CLASS
			| CLOSE
			| CLUSTER
			| COLUMNS
			| COMMENT
			| COMMENTS
			| COMMIT
			| COMMITTED
			| COMPRESSION
			| CONFIGURATION
			| CONFLICT
			| CONNECTION
			| CONSTRAINTS
			| CONTENT_P
			| CONTINUE_P
			| CONVERSION_P
			| COPY
			| COST
			| CSV
			| CUBE
			| CURRENT_P
			| CURSOR
			| CYCLE
			| DATA_P
			| DATABASE
			| DAY_P
			| DEALLOCATE
			| DECLARE
			| DEFAULTS
			| DEFERRED
			| DEFINER
			| DELETE_P
			| DELIMITER
			| DELIMITERS
			| DEPENDS
			| DEPTH
			| DETACH
			| DICTIONARY
			| DISABLE_P
			| DISCARD
			| DOCUMENT_P
			| DOMAIN_P
			| DOUBLE_P
			| DROP
			| EACH
			| ENABLE_P
			| ENCODING
			| ENCRYPTED
			| ENUM_P
			| ESCAPE
			| EVENT
			| EXCLUDE
			| EXCLUDING
			| EXCLUSIVE
			| EXECUTE
			| EXPLAIN
			| EXPRESSION
			| EXTENSION
			| EXTERNAL
			| FAMILY
			| FILTER
			| FINALIZE
			| FIRST_P
			| FOLLOWING
			| FORCE
			| FORMAT
			| FORWARD
			| FUNCTION
			| FUNCTIONS
			| GENERATED
			| GLOBAL
			| GRANTED
			| GROUPS
			| HANDLER
			| HEADER_P
			| HOLD
			| HOUR_P
			| IDENTITY_P
			| IF_P
			| IMMEDIATE
			| IMMUTABLE
			| IMPLICIT_P
			| IMPORT_P
			| INCLUDE
			| INCLUDING
			| INCREMENT
			| INDENT
			| INDEX
			| INDEXES
			| INHERIT
			| INHERITS
			| INLINE_P
			| INPUT_P
			| INSENSITIVE
			| INSERT
			| INSTEAD
			| INVOKER
			| ISOLATION
			| KEY
			| KEYS
			| LABEL
			| LANGUAGE
			| LARGE_P
			| LAST_P
			| LEAKPROOF
			| LEVEL
			| LISTEN
			| LOAD
			| LOCAL
			| LOCATION
			| LOCK_P
			| LOCKED
			| LOGGED
			| MAPPING
			| MATCH
			| MATCHED
			| MATERIALIZED
			| MAXVALUE
			| MERGE
			| METHOD
			| MINUTE_P
			| MINVALUE
			| MODE
			| MONTH_P
			| MOVE
			| NAME_P
			| NAMES
			| NEW
			| NEXT
			| NFC
			| NFD
			| NFKC
			| NFKD
			| NO
			| NORMALIZED
			| NOTHING
			| NOTIFY
			| NOWAIT
			| NULLS_P
			| OBJECT_P
			| OF
			| OFF
			| OIDS
			| OLD
			| OPERATOR
			| OPTION
			| OPTIONS
			| ORDINALITY
			| OTHERS
			| OVER
			| OVERRIDING
			| OWNED
			| OWNER
			| PARALLEL
			| PARAMETER
			| PARSER
			| PARTIAL
			| PARTITION
			| PASSING
			| PASSWORD
			| PLANS
			| POLICY
			| PRECEDING
			| PREPARE
			| PREPARED
			| PRESERVE
			| PRIOR
			| PRIVILEGES
			| PROCEDURAL
			| PROCEDURE
			| PROCEDURES
			| PROGRAM
			| PUBLICATION
			| QUOTE
			| RANGE
			| READ
			| REASSIGN
			| RECHECK
			| RECURSIVE
			| REF_P
			| REFERENCING
			| REFRESH
			| REINDEX
			| RELATIVE_P
			| RELEASE
			| RENAME
			| REPEATABLE
			| REPLACE
			| REPLICA
			| RESET
			| RESTART
			| RESTRICT
			| RETURN
			| RETURNS
			| REVOKE
			| ROLE
			| ROLLBACK
			| ROLLUP
			| ROUTINE
			| ROUTINES
			| ROWS
			| RULE
			| SAVEPOINT
			| SCALAR
			| SCHEMA
			| SCHEMAS
			| SCROLL
			| SEARCH
			| SECOND_P
			| SECURITY
			| SEQUENCE
			| SEQUENCES
			| SERIALIZABLE
			| SERVER
			| SESSION
			| SET
			| SETS
			| SHARE
			| SHOW
			| SIMPLE
			| SKIP
			| SNAPSHOT
			| SQL_P
			| STABLE
			| STANDALONE_P
			| START
			| STATEMENT
			| STATISTICS
			| STDIN
			| STDOUT
			| STORAGE
			| STORED
			| STRICT_P
			| STRIP_P
			| SUBSCRIPTION
			| SUPPORT
			| SYSID
			| SYSTEM_P
			| TABLES
			| TABLESPACE
			| TEMP
			| TEMPLATE
			| TEMPORARY
			| TEXT_P
			| TIES
			| TRANSACTION
			| TRANSFORM
			| TRIGGER
			| TRUNCATE
			| TRUSTED
			| TYPE_P
			| TYPES_P
			| UESCAPE
			| UNBOUNDED
			| UNCOMMITTED
			| UNENCRYPTED
			| UNKNOWN
			| UNLISTEN
			| UNLOGGED
			| UNTIL
			| UPDATE
			| VACUUM
			| VALID
			| VALIDATE
			| VALIDATOR
			| VALUE_P
			| VARYING
			| VERSION_P
			| VIEW
			| VIEWS
			| VOLATILE
			| WHITESPACE_P
			| WITHIN
			| WITHOUT
			| WORK
			| WRAPPER
			| WRITE
			| XML_P
			| YEAR_P
			| YES_P
			| ZONE
		;

/* Column identifier --- keywords that can be column, table, etc names.
 *
 * Many of these keywords will in fact be recognized as type or function
 * names too; but they have special productions for the purpose, and so
 * can't be treated as "generic" type or function names.
 *
 * The type names appearing here are not usable as function names
 * because they can be followed by '(' in typename productions, which
 * looks too much like a function call for an LR(1) parser.
 */
col_name_keyword:
			  BETWEEN	{$$=$1}
			| BIGINT	{$$=$1}
			| BIT	{$$=$1}
			| BOOLEAN_P	{$$=$1}
			| CHAR_P	{$$=$1}
			| CHARACTER	{$$=$1}
			| COALESCE	{$$=$1}
			| DEC	{$$=$1}
			| DECIMAL_P	{$$=$1}
			| EXISTS	{$$=$1}
			| EXTRACT	{$$=$1}
			| FLOAT_P	{$$=$1}
			| GREATEST	{$$=$1}
			| GROUPING	{$$=$1}
			| INOUT	{$$=$1}
			| INT_P	{$$=$1}
			| INTEGER	{$$=$1}
			| INTERVAL	{$$=$1}
			| JSON	{$$=$1}
			| JSON_ARRAY	{$$=$1}
			| JSON_ARRAYAGG	{$$=$1}
			| JSON_OBJECT	{$$=$1}
			| JSON_OBJECTAGG	{$$=$1}
			| JSON_SCALAR	{$$=$1}
			| JSON_SERIALIZE	{$$=$1}
			| LEAST	{$$=$1}
			| NATIONAL	{$$=$1}
			| NCHAR	{$$=$1}
			| NONE	{$$=$1}
			| NORMALIZE	{$$=$1}
			| NULLIF	{$$=$1}
			| NUMERIC	{$$=$1}
			| OUT_P	{$$=$1}
			| OVERLAY	{$$=$1}
			| POSITION	{$$=$1}
			| PRECISION	{$$=$1}
			| REAL	{$$=$1}
			| ROW	{$$=$1}
			| SETOF	{$$=$1}
			| SMALLINT	{$$=$1}
			| SUBSTRING	{$$=$1}
			| TIME	{$$=$1}
			| TIMESTAMP	{$$=$1}
			| TREAT	{$$=$1}
			| TRIM	{$$=$1}
			| VALUES	{$$=$1}
			| VARCHAR	{$$=$1}
			| XMLATTRIBUTES	{$$=$1}
			| XMLCONCAT	{$$=$1}
			| XMLELEMENT	{$$=$1}
			| XMLEXISTS	{$$=$1}
			| XMLFOREST	{$$=$1}
			| XMLNAMESPACES	{$$=$1}
			| XMLPARSE	{$$=$1}
			| XMLPI	{$$=$1}
			| XMLROOT	{$$=$1}
			| XMLSERIALIZE	{$$=$1}
			| XMLTABLE	{$$=$1}
		;

/* Type/function identifier --- keywords that can be type or function names.
 *
 * Most of these are keywords that are used as operators in expressions;
 * in general such keywords can't be column names because they would be
 * ambiguous with variables, but they are unambiguous as function identifiers.
 *
 * Do not include POSITION, SUBSTRING, etc here since they have explicit
 * productions in a_expr to support the goofy SQL9x argument syntax.
 * - thomas 2000-11-28
 */
type_func_name_keyword:
			  AUTHORIZATION	{$$=$1}
			| BINARY	{$$=$1}
			| COLLATION	{$$=$1}
			| CONCURRENTLY	{$$=$1}
			| CROSS	{$$=$1}
			| CURRENT_SCHEMA	{$$=$1}
			| FREEZE	{$$=$1}
			| FULL	{$$=$1}
			| ILIKE	{$$=$1}
			| INNER_P	{$$=$1}
			| IS	{$$=$1}
			| ISNULL	{$$=$1}
			| JOIN	{$$=$1}
			| LEFT	{$$=$1}
			| LIKE	{$$=$1}
			| NATURAL	{$$=$1}
			| NOTNULL	{$$=$1}
			| OUTER_P	{$$=$1}
			| OVERLAPS	{$$=$1}
			| RIGHT	{$$=$1}
			| SIMILAR	{$$=$1}
			| TABLESAMPLE	{$$=$1}
			| VERBOSE	{$$=$1}
		;

/* Reserved keyword --- these keywords are usable only as a ColLabel.
 *
 * Keywords appear here if they could not be distinguished from variable,
 * type, or function names in some contexts.  Don't put things here unless
 * forced to.
 */
reserved_keyword:
			  ANALYSE	{$$=$1}
			| ALL	{$$=$1}
			| UPDATE {$$=$1}
			| INSERT {$$=$1}
			| DELETE {$$=$1}
			| FROM {$$=$1}
			| WHERE {$$=$1}
			| AND {$$=$1}
			| OR {$$=$1}
			| VALUES {$$=$1}
			| TABLE {$$=$1}
			| EXPLAIN {$$=$1}
			| EXECUTE {$$=$1}
			| PREPARE {$$=$1}
			| BEGIN {$$=$1}
			| COMMIT{ $$=$1}
			| ROLLBACK {$$=$1}
			| SET { $$=$1}
			| RESET { $$=$1}
			| LOCAL {$$=$1}
			| ALL {$$=$1}
			| CREATE {$$=$1}
			| ALTER {$$=$1}
			| VACUUM {$$=$1}
			| CLUSTER {$$=$1}
			| ANALYZE {$$=$1}
			| DROP {$$=$1}
			| TRUNCATE {$$=$1}
			| INDEX {$$=$1}
			| ROLE {$$=$1}
			| DATABASE {$$=$1}
			| ORDER {$$=$1}
			| GROUP {$$=$1}
			| BY {$$=$1}
			| INTO {$$=$1}
			| LIMIT {$$=$1}
			| OFFSET {$$=$1}
			| FOR  {$$=$1}
			| FETCH {$$=$1}
			| TOPENBR {$$=$1}
			| TCLOSEBR {$$=$1}
			| TCOMMA {$$=$1}
			| TCOLON {$$=$1}
			| ONLY {$$=$1}
			| AS {$$=$1}
			| RETURNING {$$=$1}
			| KEY {$$=$1}
			| FOREIGN {$$=$1}
			| PRIMARY {$$=$1}
			| REFERENCES {$$=$1}
			| TMUL {$$=$1}
			| TPLUS {$$=$1}
			| TMOD {$$=$1}
			| TPOW {$$=$1}
			| TDIV {$$=$1}
			| JOIN {$$=$1}
			| CROSS {$$=$1}
			| IS  {$$=$1}
			| TRUE_P {$$=$1}
			| FALSE_P {$$=$1}
			| ON {$$=$1}
			| OP {$$=$1}
			| INT_P {$$=$1}
			| FLOAT_P {$$=$1}
			| DOUBLE_P {$$=$1}
			| INTEGER {$$=$1}
			| ANALYZE	{$$=$1}
			| AND	{$$=$1}
			| ANY	{$$=$1}
			| ARRAY	{$$=$1}
			| AS	{$$=$1}
			| ASC	{$$=$1}
			| ASYMMETRIC	{$$=$1}
			| BOTH	{$$=$1}
			| CASE	{$$=$1}
			| CAST	{$$=$1}
			| CHECK	{$$=$1}
			| COLLATE	{$$=$1}
			| COLUMN	{$$=$1}
			| CONSTRAINT	{$$=$1}
			| CREATE	{$$=$1}
			| CURRENT_CATALOG	{$$=$1}
			| CURRENT_DATE	{$$=$1}
			| CURRENT_ROLE	{$$=$1}
			| CURRENT_TIME	{$$=$1}
			| CURRENT_TIMESTAMP	{$$=$1}
			| CURRENT_USER	{$$=$1}
			| DEFAULT	{$$=$1}
			| DEFERRABLE	{$$=$1}
			| DESC	{$$=$1}
			| DISTINCT	{$$=$1}
			| DO	{$$=$1}
			| ELSE	{$$=$1}
			| END_P	{$$=$1}
			| EXCEPT	{$$=$1}
			| FALSE_P	{$$=$1}
			| FETCH	{$$=$1}
			| FOR	{$$=$1}
			| FOREIGN	{$$=$1}
			| FROM	{$$=$1}
			| GRANT	{$$=$1}
			| GROUP_P	{$$=$1}
			| HAVING	{$$=$1}
			| IN_P	{$$=$1}
			| INITIALLY	{$$=$1}
			| INTERSECT	{$$=$1}
			| INTO	{$$=$1}
			| LATERAL_P	{$$=$1}
			| LEADING	{$$=$1}
			| LIMIT	{$$=$1}
			| LOCALTIME	{$$=$1}
			| LOCALTIMESTAMP	{$$=$1}
			| NOT	{$$=$1}
			| NULL_P	{$$=$1}
			| OFFSET	{$$=$1}
			| ON	{$$=$1}
			| ONLY	{$$=$1}
			| OR	{$$=$1}
			| ORDER	{$$=$1}
			| PLACING	{$$=$1}
			| PRIMARY	{$$=$1}
			| REFERENCES	{$$=$1}
			| RETURNING	{$$=$1}
			| SELECT	{$$=$1}
			| SESSION_USER	{$$=$1}
			| SOME	{$$=$1}
			| SYMMETRIC	{$$=$1}
			| SYSTEM_USER	{$$=$1}
			| TABLE	{$$=$1}
			| THEN	{$$=$1}
			| TO	{$$=$1}
			| TRAILING	{$$=$1}
			| TRUE_P	{$$=$1}
			| UNION	{$$=$1}
			| UNIQUE	{$$=$1}
			| USER	{$$=$1}
			| USING	{$$=$1}
			| VARIADIC	{$$=$1}
			| WHEN	{$$=$1}
			| WHERE	{$$=$1}
			| WINDOW	{$$=$1}
			| WITH	{$$=$1}
		;

/*
 * While all keywords can be used as column labels when preceded by AS,
 * not all of them can be used as a "bare" column label without AS.
 * Those that can be used as a bare label must be listed here,
 * in addition to appearing in one of the category lists above.
 *
 * Always add a new keyword to this list if possible.  Mark it BARE_LABEL
 * in kwlist.h if it is included here, or AS_LABEL if it is not.
 */
bare_label_keyword:
			  ABORT_P	{$$=$1}
			| ABSENT	{$$=$1}
			| ABSOLUTE_P	{$$=$1}
			| ACCESS	{$$=$1}
			| ACTION	{$$=$1}
			| ADD_P	{$$=$1}
			| ADMIN	{$$=$1}
			| AFTER	{$$=$1}
			| AGGREGATE	{$$=$1}
			| ALL	{$$=$1}
			| ALSO	{$$=$1}
			| ALTER	{$$=$1}
			| ALWAYS	{$$=$1}
			| ANALYSE	{$$=$1}
			| ANALYZE	{$$=$1}
			| AND	{$$=$1}
			| ANY	{$$=$1}
			| ASC	{$$=$1}
			| ASENSITIVE	{$$=$1}
			| ASSERTION	{$$=$1}
			| ASSIGNMENT	{$$=$1}
			| ASYMMETRIC	{$$=$1}
			| AT	{$$=$1}
			| ATOMIC	{$$=$1}
			| ATTACH	{$$=$1}
			| ATTRIBUTE	{$$=$1}
			| AUTHORIZATION	{$$=$1}
			| BACKWARD	{$$=$1}
			| BEFORE	{$$=$1}
			| BEGIN_P	{$$=$1}
			| BETWEEN	{$$=$1}
			| BIGINT	{$$=$1}
			| BINARY	{$$=$1}
			| BIT	{$$=$1}
			| BOOLEAN_P	{$$=$1}
			| BOTH	{$$=$1}
			| BREADTH	{$$=$1}
			| BY	{$$=$1}
			| CACHE	{$$=$1}
			| CALL	{$$=$1}
			| CALLED	{$$=$1}
			| CASCADE	{$$=$1}
			| CASCADED	{$$=$1}
			| CASE	{$$=$1}
			| CAST	{$$=$1}
			| CATALOG_P	{$$=$1}
			| CHAIN	{$$=$1}
			| CHARACTERISTICS	{$$=$1}
			| CHECK	{$$=$1}
			| CHECKPOINT	{$$=$1}
			| CLASS	{$$=$1}
			| CLOSE	{$$=$1}
			| CLUSTER	{$$=$1}
			| COALESCE	{$$=$1}
			| COLLATE	{$$=$1}
			| COLLATION	{$$=$1}
			| COLUMN	{$$=$1}
			| COLUMNS	{$$=$1}
			| COMMENT	{$$=$1}
			| COMMENTS	{$$=$1}
			| COMMIT	{$$=$1}
			| COMMITTED	{$$=$1}
			| COMPRESSION	{$$=$1}
			| CONCURRENTLY	{$$=$1}
			| CONFIGURATION	{$$=$1}
			| CONFLICT	{$$=$1}
			| CONNECTION	{$$=$1}
			| CONSTRAINT	{$$=$1}
			| CONSTRAINTS	{$$=$1}
			| CONTENT_P	{$$=$1}
			| CONTINUE_P	{$$=$1}
			| CONVERSION_P	{$$=$1}
			| COPY	{$$=$1}
			| COST	{$$=$1}
			| CROSS	{$$=$1}
			| CSV	{$$=$1}
			| CUBE	{$$=$1}
			| CURRENT_P	{$$=$1}
			| CURRENT_CATALOG	{$$=$1}
			| CURRENT_DATE	{$$=$1}
			| CURRENT_ROLE	{$$=$1}
			| CURRENT_SCHEMA	{$$=$1}
			| CURRENT_TIME	{$$=$1}
			| CURRENT_TIMESTAMP	{$$=$1}
			| CURRENT_USER	{$$=$1}
			| CURSOR	{$$=$1}
			| CYCLE	{$$=$1}
			| DATA_P	{$$=$1}
			| DATABASE	{$$=$1}
			| DEALLOCATE	{$$=$1}
			| DEC	{$$=$1}
			| DECIMAL_P	{$$=$1}
			| DECLARE	{$$=$1}
			| DEFAULT	{$$=$1}
			| DEFAULTS	{$$=$1}
			| DEFERRABLE	{$$=$1}
			| DEFERRED	{$$=$1}
			| DEFINER	{$$=$1}
			| DELETE_P	{$$=$1}
			| DELIMITER	{$$=$1}
			| DELIMITERS	{$$=$1}
			| DEPENDS	{$$=$1}
			| DEPTH	{$$=$1}
			| DESC	{$$=$1}
			| DETACH	{$$=$1}
			| DICTIONARY	{$$=$1}
			| DISABLE_P	{$$=$1}
			| DISCARD	{$$=$1}
			| DISTINCT	{$$=$1}
			| DO	{$$=$1}
			| DOCUMENT_P	{$$=$1}
			| DOMAIN_P	{$$=$1}
			| DOUBLE_P	{$$=$1}
			| DROP	{$$=$1}
			| EACH	{$$=$1}
			| ELSE	{$$=$1}
			| ENABLE_P	{$$=$1}
			| ENCODING	{$$=$1}
			| ENCRYPTED	{$$=$1}
			| END_P	{$$=$1}
			| ENUM_P	{$$=$1}
			| ESCAPE	{$$=$1}
			| EVENT	{$$=$1}
			| EXCLUDE	{$$=$1}
			| EXCLUDING	{$$=$1}
			| EXCLUSIVE	{$$=$1}
			| EXECUTE	{$$=$1}
			| EXISTS	{$$=$1}
			| EXPLAIN	{$$=$1}
			| EXPRESSION	{$$=$1}
			| EXTENSION	{$$=$1}
			| EXTERNAL	{$$=$1}
			| EXTRACT	{$$=$1}
			| FALSE_P	{$$=$1}
			| FAMILY	{$$=$1}
			| FINALIZE	{$$=$1}
			| FIRST_P	{$$=$1}
			| FLOAT_P	{$$=$1}
			| FOLLOWING	{$$=$1}
			| FORCE	{$$=$1}
			| FOREIGN	{$$=$1}
			| FORMAT	{$$=$1}
			| FORWARD	{$$=$1}
			| FREEZE	{$$=$1}
			| FULL	{$$=$1}
			| FUNCTION	{$$=$1}
			| FUNCTIONS	{$$=$1}
			| GENERATED	{$$=$1}
			| GLOBAL	{$$=$1}
			| GRANTED	{$$=$1}
			| GREATEST	{$$=$1}
			| GROUPING	{$$=$1}
			| GROUPS	{$$=$1}
			| HANDLER	{$$=$1}
			| HEADER_P	{$$=$1}
			| HOLD	{$$=$1}
			| IDENTITY_P	{$$=$1}
			| IF_P	{$$=$1}
			| ILIKE	{$$=$1}
			| IMMEDIATE	{$$=$1}
			| IMMUTABLE	{$$=$1}
			| IMPLICIT_P	{$$=$1}
			| IMPORT_P	{$$=$1}
			| IN_P	{$$=$1}
			| INCLUDE	{$$=$1}
			| INCLUDING	{$$=$1}
			| INCREMENT	{$$=$1}
			| INDENT	{$$=$1}
			| INDEX	{$$=$1}
			| INDEXES	{$$=$1}
			| INHERIT	{$$=$1}
			| INHERITS	{$$=$1}
			| INITIALLY	{$$=$1}
			| INLINE_P	{$$=$1}
			| INNER_P	{$$=$1}
			| INOUT	{$$=$1}
			| INPUT_P	{$$=$1}
			| INSENSITIVE	{$$=$1}
			| INSERT	{$$=$1}
			| INSTEAD	{$$=$1}
			| INT_P	{$$=$1}
			| INTEGER	{$$=$1}
			| INTERVAL	{$$=$1}
			| INVOKER	{$$=$1}
			| IS	{$$=$1}
			| ISOLATION	{$$=$1}
			| JOIN	{$$=$1}
			| JSON	{$$=$1}
			| JSON_ARRAY	{$$=$1}
			| JSON_ARRAYAGG	{$$=$1}
			| JSON_OBJECT	{$$=$1}
			| JSON_OBJECTAGG	{$$=$1}
			| JSON_SCALAR	{$$=$1}
			| JSON_SERIALIZE	{$$=$1}
			| KEY	{$$=$1}
			| KEYS	{$$=$1}
			| LABEL	{$$=$1}
			| LANGUAGE	{$$=$1}
			| LARGE_P	{$$=$1}
			| LAST_P	{$$=$1}
			| LATERAL_P	{$$=$1}
			| LEADING	{$$=$1}
			| LEAKPROOF	{$$=$1}
			| LEAST	{$$=$1}
			| LEFT	{$$=$1}
			| LEVEL	{$$=$1}
			| LIKE	{$$=$1}
			| LISTEN	{$$=$1}
			| LOAD	{$$=$1}
			| LOCAL	{$$=$1}
			| LOCALTIME	{$$=$1}
			| LOCALTIMESTAMP	{$$=$1}
			| LOCATION	{$$=$1}
			| LOCK_P	{$$=$1}
			| LOCKED	{$$=$1}
			| LOGGED	{$$=$1}
			| MAPPING	{$$=$1}
			| MATCH	{$$=$1}
			| MATCHED	{$$=$1}
			| MATERIALIZED	{$$=$1}
			| MAXVALUE	{$$=$1}
			| MERGE	{$$=$1}
			| METHOD	{$$=$1}
			| MINVALUE	{$$=$1}
			| MODE	{$$=$1}
			| MOVE	{$$=$1}
			| NAME_P	{$$=$1}
			| NAMES	{$$=$1}
			| NATIONAL	{$$=$1}
			| NATURAL	{$$=$1}
			| NCHAR	{$$=$1}
			| NEW	{$$=$1}
			| NEXT	{$$=$1}
			| NFC	{$$=$1}
			| NFD	{$$=$1}
			| NFKC	{$$=$1}
			| NFKD	{$$=$1}
			| NO	{$$=$1}
			| NONE	{$$=$1}
			| NORMALIZE	{$$=$1}
			| NORMALIZED	{$$=$1}
			| NOT	{$$=$1}
			| NOTHING	{$$=$1}
			| NOTIFY	{$$=$1}
			| NOWAIT	{$$=$1}
			| NULL_P	{$$=$1}
			| NULLIF	{$$=$1}
			| NULLS_P	{$$=$1}
			| NUMERIC	{$$=$1}
			| OBJECT_P	{$$=$1}
			| OF	{$$=$1}
			| OFF	{$$=$1}
			| OIDS	{$$=$1}
			| OLD	{$$=$1}
			| ONLY	{$$=$1}
			| OPERATOR	{$$=$1}
			| OPTION	{$$=$1}
			| OPTIONS	{$$=$1}
			| OR	{$$=$1}
			| ORDINALITY	{$$=$1}
			| OTHERS	{$$=$1}
			| OUT_P	{$$=$1}
			| OUTER_P	{$$=$1}
			| OVERLAY	{$$=$1}
			| OVERRIDING	{$$=$1}
			| OWNED	{$$=$1}
			| OWNER	{$$=$1}
			| PARALLEL	{$$=$1}
			| PARAMETER	{$$=$1}
			| PARSER	{$$=$1}
			| PARTIAL	{$$=$1}
			| PARTITION	{$$=$1}
			| PASSING	{$$=$1}
			| PASSWORD	{$$=$1}
			| PLACING	{$$=$1}
			| PLANS	{$$=$1}
			| POLICY	{$$=$1}
			| POSITION	{$$=$1}
			| PRECEDING	{$$=$1}
			| PREPARE	{$$=$1}
			| PREPARED	{$$=$1}
			| PRESERVE	{$$=$1}
			| PRIMARY	{$$=$1}
			| PRIOR	{$$=$1}
			| PRIVILEGES	{$$=$1}
			| PROCEDURAL	{$$=$1}
			| PROCEDURE	{$$=$1}
			| PROCEDURES	{$$=$1}
			| PROGRAM	{$$=$1}
			| PUBLICATION	{$$=$1}
			| QUOTE	{$$=$1}
			| RANGE	{$$=$1}
			| READ	{$$=$1}
			| REAL	{$$=$1}
			| REASSIGN	{$$=$1}
			| RECHECK	{$$=$1}
			| RECURSIVE	{$$=$1}
			| REF_P	{$$=$1}
			| REFERENCES	{$$=$1}
			| REFERENCING	{$$=$1}
			| REFRESH	{$$=$1}
			| REINDEX	{$$=$1}
			| RELATIVE_P	{$$=$1}
			| RELEASE	{$$=$1}
			| RENAME	{$$=$1}
			| REPEATABLE	{$$=$1}
			| REPLACE	{$$=$1}
			| REPLICA	{$$=$1}
			| RESET	{$$=$1}
			| RESTART	{$$=$1}
			| RESTRICT	{$$=$1}
			| RETURN	{$$=$1}
			| RETURNS	{$$=$1}
			| REVOKE	{$$=$1}
			| RIGHT	{$$=$1}
			| ROLE	{$$=$1}
			| ROLLBACK	{$$=$1}
			| ROLLUP	{$$=$1}
			| ROUTINE	{$$=$1}
			| ROUTINES	{$$=$1}
			| ROW	{$$=$1}
			| ROWS	{$$=$1}
			| RULE	{$$=$1}
			| SAVEPOINT	{$$=$1}
			| SCALAR	{$$=$1}
			| SCHEMA	{$$=$1}
			| SCHEMAS	{$$=$1}
			| SCROLL	{$$=$1}
			| SEARCH	{$$=$1}
			| SECURITY	{$$=$1}
			| SELECT	{$$=$1}
			| SEQUENCE	{$$=$1}
			| SEQUENCES	{$$=$1}
			| SERIALIZABLE	{$$=$1}
			| SERVER	{$$=$1}
			| SESSION	{$$=$1}
			| SESSION_USER	{$$=$1}
			| SET	{$$=$1}
			| SETOF	{$$=$1}
			| SETS	{$$=$1}
			| SHARE	{$$=$1}
			| SHOW	{$$=$1}
			| SIMILAR	{$$=$1}
			| SIMPLE	{$$=$1}
			| SKIP	{$$=$1}
			| SMALLINT	{$$=$1}
			| SNAPSHOT	{$$=$1}
			| SOME	{$$=$1}
			| SQL_P	{$$=$1}
			| STABLE	{$$=$1}
			| STANDALONE_P	{$$=$1}
			| START	{$$=$1}
			| STATEMENT	{$$=$1}
			| STATISTICS	{$$=$1}
			| STDIN	{$$=$1}
			| STDOUT	{$$=$1}
			| STORAGE	{$$=$1}
			| STORED	{$$=$1}
			| STRICT_P	{$$=$1}
			| STRIP_P	{$$=$1}
			| SUBSCRIPTION	{$$=$1}
			| SUBSTRING	{$$=$1}
			| SUPPORT	{$$=$1}
			| SYMMETRIC	{$$=$1}
			| SYSID	{$$=$1}
			| SYSTEM_P	{$$=$1}
			| SYSTEM_USER	{$$=$1}
			| TABLE	{$$=$1}
			| TABLES	{$$=$1}
			| TABLESAMPLE	{$$=$1}
			| TABLESPACE	{$$=$1}
			| TEMP	{$$=$1}
			| TEMPLATE	{$$=$1}
			| TEMPORARY	{$$=$1}
			| TEXT_P	{$$=$1}
			| THEN	{$$=$1}
			| TIES	{$$=$1}
			| TIME	{$$=$1}
			| TIMESTAMP	{$$=$1}
			| TRAILING	{$$=$1}
			| TRANSACTION	{$$=$1}
			| TRANSFORM	{$$=$1}
			| TREAT	{$$=$1}
			| TRIGGER	{$$=$1}
			| TRIM	{$$=$1}
			| TRUE_P	{$$=$1}
			| TRUNCATE	{$$=$1}
			| TRUSTED	{$$=$1}
			| TYPE_P	{$$=$1}
			| TYPES_P	{$$=$1}
			| UESCAPE	{$$=$1}
			| UNBOUNDED	{$$=$1}
			| UNCOMMITTED	{$$=$1}
			| UNENCRYPTED	{$$=$1}
			| UNIQUE	{$$=$1}
			| UNKNOWN	{$$=$1}
			| UNLISTEN	{$$=$1}
			| UNLOGGED	{$$=$1}
			| UNTIL	{$$=$1}
			| UPDATE	{$$=$1}
			| USER	{$$=$1}
			| USING	{$$=$1}
			| VACUUM	{$$=$1}
			| VALID	{$$=$1}
			| VALIDATE	{$$=$1}
			| VALIDATOR	{$$=$1}
			| VALUE_P	{$$=$1}
			| VALUES	{$$=$1}
			| VARCHAR	{$$=$1}
			| VARIADIC	{$$=$1}
			| VERBOSE	{$$=$1}
			| VERSION_P	{$$=$1}
			| VIEW	{$$=$1}
			| VIEWS	{$$=$1}
			| VOLATILE	{$$=$1}
			| WHEN	{$$=$1}
			| WHITESPACE_P	{$$=$1}
			| WORK	{$$=$1}
			| WRAPPER	{$$=$1}
			| WRITE	{$$=$1}
			| XML_P	{$$=$1}
			| XMLATTRIBUTES	{$$=$1}
			| XMLCONCAT	{$$=$1}
			| XMLELEMENT	{$$=$1}
			| XMLEXISTS	{$$=$1}
			| XMLFOREST	{$$=$1}
			| XMLNAMESPACES	{$$=$1}
			| XMLPARSE	{$$=$1}
			| XMLPI	{$$=$1}
			| XMLROOT	{$$=$1}
			| XMLSERIALIZE	{$$=$1}
			| XMLTABLE	{$$=$1}
			| YES_P	{$$=$1}
			| ZONE	{$$=$1}
		;

any_tok:
    reserved_keyword {$$=$1} | SCONST {$$=$1} |  IDENT {$$=$1} 


anything:
    /*nothing*/ {}
    | any_tok anything {}




command:
    /* noting */ { $$ = nil } | 
	routable_statement
	{
		setParseTree(yylex, $1)
	} | begin_stmt {
		setParseTree(yylex, $1)
    } | commit_stmt {
		setParseTree(yylex, $1)
    } | rollback_stmt {
		setParseTree(yylex, $1)
    } | ExecuteStmt {
		setParseTree(yylex, $1)
    } | varset_stmt {
		setParseTree(yylex, $1)
    } | PrepareStmt {
		setParseTree(yylex, $1)
	} | DeallocateStmt {
		setParseTree(yylex, $1)
    } | vacuum_stmt {
		setParseTree(yylex, $1)
    } | cluster_stmt {
		setParseTree(yylex, $1)
    } | analyze_stmt {
		setParseTree(yylex, $1)
    } | truncate_stmt {
		setParseTree(yylex, $1)
    } | drop_stmt {
		setParseTree(yylex, $1)
    } | create_stmt {
		setParseTree(yylex, $1)
    } | alter_stmt {
		setParseTree(yylex, $1)
    }

any_id:
	IDENT
	{
		$$ = string($1)
	}

any_val:
    reserved_keyword {
        $$ = $1
    }
    | SCONST
	{
		$$ = string($1)
	}

table_name:
	IDENT
	{
		$$ = string($1)
	}

operator:
    IDENT {
        $$ = $1
    } | AND {
        $$ = "AND"
    } | OR {
        $$ = "OR"
    } | TNOT_EQUALS {
        $$ = "!="
    } | TEQ {
        $$ = "="
    } | TLESS {
        $$ = "<"
    } | TGREATER {
       $$ = ">"
    }| TGREATER_EQUALS {
       $$ = ">="
    }| TLESS_EQUALS {
       $$ = "<="
    } | OP {
        $$ = $1
    }


/*****************************************************************************
 *
 *	Type syntax
 *		SQL introduces a large amount of type-specific syntax.
 *		Define individual clauses to handle these cases, and use
 *		 the generic case to handle regular type-extensible Postgres syntax.
 *		- thomas 1997-10-10
 *
 *****************************************************************************/

Typename:	SimpleTypename opt_array_bounds
				{
					$$ = $1
				}
			| SETOF SimpleTypename opt_array_bounds
				{
					$$ = $2
				}
			/* SQL standard syntax, currently only one-dimensional */
			| SimpleTypename ARRAY TSQOPENBR SCONST  TSQCLOSEBR
				{
					$$ = $1
				}
			| SETOF SimpleTypename ARRAY TSQOPENBR SCONST  TSQCLOSEBR
				{
					$$ = $2
				}
			| SimpleTypename ARRAY
				{

					$$ = $1
				}
			| SETOF SimpleTypename ARRAY
				{
					$$ = $2
				}
		;


opt_array_bounds:
			opt_array_bounds TSQOPENBR TSQCLOSEBR
					{ }
			| opt_array_bounds TSQOPENBR SCONST TSQCLOSEBR
					{  }
			| /*EMPTY*/
					{  }
		;

SimpleTypename:
			GenericType								{ 
					$$ = $1 }
			| Numeric								{ 
					$$ = $1 }
			| Bit									{
					$$ = $1 }
			| Character								{ 
					$$ = $1}
			| ConstDatetime							{ 
					$$ = $1}
			| ConstInterval opt_interval
				{

				}
			| ConstInterval TOPENBR SCONST  TCLOSEBR
				{

				}
		;
/* We have a separate ConstTypename to allow defaulting fixed-length
 * types such as CHAR() and BIT() to an unspecified length.
 * SQL9x requires that these default to a length of one, but this
 * makes no sense for constructs like CHAR 'hi' and BIT '0101',
 * where there is an obvious better choice to make.
 * Note that ConstInterval is not included here since it must
 * be pushed up higher in the rules to accommodate the postfix
 * options (e.g. INTERVAL '1' YEAR). Likewise, we have to handle
 * the generic-type-name case in AExprConst to avoid premature
 * reduce/reduce conflicts against function names.
 */
ConstTypename:
			Numeric									{ $$ = $1; }
			| ConstBit								{ $$ = $1; }
			| ConstCharacter						{ $$ = $1; }
			| ConstDatetime							{ $$ = $1; }
		;


type_function_name:	IDENT							{ 
					$$ = $1}


/*
 * GenericType covers all type names that don't have special syntax mandated
 * by the standard, including qualified names.  We also allow type modifiers.
 * To avoid parsing conflicts against function invocations, the modifiers
 * have to be shown as expr_list here, but parse analysis will only accept
 * constants for them.
 */
GenericType:
			type_function_name opt_type_modifiers
				{

					$$ = $1
				}
			// | type_function_name attrs opt_type_modifiers
			// 	{
			// 	}
		;

opt_type_modifiers: TOPENBR expr_list TCLOSEBR				{ }
					| /* EMPTY */					{ }
		;

/*
 * SQL numeric data types
 */
Numeric:	INT_P
				{
				
					$$ = $1
				}
			| INTEGER
				{
					$$ = $1
				}
			| SMALLINT
				{
					$$ = $1
				}
			| BIGINT
				{
					$$ = $1
				}
			| REAL
				{
					$$ = $1
				}
			| FLOAT_P opt_float
				{
					$$ = $1
				}
			| DOUBLE_P PRECISION
				{
					$$ = $1
				}
			| DECIMAL_P opt_type_modifiers
				{
					$$ = $1
				}
			| DEC opt_type_modifiers
				{
					$$ = $1
				}
			| NUMERIC opt_type_modifiers
				{
					$$ = $1
				}
			| BOOLEAN_P
				{
					$$ = $1
				}
		;

opt_float:	TOPENBR SCONST TCLOSEBR
				{
					/*
					 * Check FLOAT() precision limits assuming IEEE floating
					 * types - thomas 1997-09-18
					 */

				}
			| /*EMPTY*/
				{
				}
		;

/*
 * SQL bit-field data types
 * The following implements BIT() and BIT VARYING().
 */
Bit:		BitWithLength
				{
				}
			| BitWithoutLength
				{
				}
		;

/* ConstBit is like Bit except "BIT" defaults to unspecified length */
/* See notes for ConstCharacter, which addresses same issue for "CHAR" */
ConstBit:	BitWithLength
				{
				}
			| BitWithoutLength
				{
				}
		;

BitWithLength:
			BIT opt_varying TOPENBR expr_list TCLOSEBR
				{

				}
		;

BitWithoutLength:
			BIT opt_varying
				{
					/* bit defaults to bit(1), varbit to no limit */

				}
		;



reloptions:
			TOPENBR reloption_list TCLOSEBR			{ $$ = $2; }
		;

opt_reloptions:		WITH reloptions					{ $$ = $2; }
			 |		/* EMPTY */						{  }
		;

reloption_list:
			reloption_elem							{}
			| reloption_list TCOMMA reloption_elem		{  }
		;

/* This should match def_elem and also allow qualified names */
reloption_elem:
			ColLabel TEQ def_arg
				{
					
				}
			| ColLabel
				{
				}
			| ColLabel TDOT ColLabel TEQ def_arg
				{
				}
			| ColLabel TDOT ColLabel
				{
				}
		;


/* Note: any simple identifier will be returned as a type name! */
def_arg:
	// func_type						{ $$ = (Node *)$1; }
			// | reserved_keyword				{ $$ = (Node *)makeString(pstrdup($1)); }
			// | qual_all_Op					{ $$ = (Node *)$1; }
			// | NumericOnly					{ $$ = (Node *)$1; }
			SCONST						{  }
			// | NONE							{ $$ = (Node *)makeString(pstrdup($1)); }
		;



/*
 * SQL character data types
 * The following implements CHAR() and VARCHAR().
 */
Character:  CharacterWithLength
				{
				}
			| CharacterWithoutLength
				{
				}
		;

ConstCharacter:  CharacterWithLength
				{
				}
			| CharacterWithoutLength
				{
					/* Length was not specified so allow to be unrestricted.
					 * This handles problems with fixed-length (bpchar) strings
					 * which in column definitions must default to a length
					 * of one, but should not be constrained if the length
					 * was not specified.
					 */

				}
		;


CharacterWithLength:  character TOPENBR SCONST TCLOSEBR
				{

				}
		;

CharacterWithoutLength:	 character
				{
					/* char defaults to char(1), varchar to no limit */

				}
		;


character:	CHARACTER opt_varying
										{  }
			| CHAR_P opt_varying
										{ }
			| VARCHAR
										{ }
			| NATIONAL CHARACTER opt_varying
										{  }
			| NATIONAL CHAR_P opt_varying
										{  }
			| NCHAR opt_varying
										{ }
		;

opt_varying:
			VARYING									{ }
			| /*EMPTY*/								{ }
		;

/*
 * SQL date/time types
 */

ConstDatetime:
			TIMESTAMP TOPENBR SCONST /* Iconst */ TCLOSEBR opt_timezone
				{

				}
			| TIMESTAMP opt_timezone
				{

				}
			| TIME TOPENBR SCONST /* Iconst */ TCLOSEBR opt_timezone
				{

				}
			| TIME opt_timezone
				{

				}
		;


ConstInterval:
			INTERVAL
				{

				}
		;

opt_timezone:
			WITH_LA TIME ZONE						{  }
			| WITHOUT TIME ZONE						{  }
			| /*EMPTY*/								{  }
		;

opt_interval:
			YEAR_P
				{  }
			| MONTH_P
				{ }
			| DAY_P
				{ }
			| HOUR_P
				{  }
			| MINUTE_P
				{  }
			| interval_second
				{  }
			| YEAR_P TO MONTH_P
				{
			
				}
			| DAY_P TO HOUR_P
				{
		
				}
			| DAY_P TO MINUTE_P
				{
	
				}
			| DAY_P TO interval_second
				{
		
				}
			| HOUR_P TO MINUTE_P
				{
		
				}
			| HOUR_P TO interval_second
				{

				}
			| MINUTE_P TO interval_second
				{

				}
			| /*EMPTY*/
				{ }
		;

interval_second:
			SECOND_P
				{
				}
			| SECOND_P TOPENBR SCONST /* Iconst */ TCLOSEBR
				{
				}
		;



/*
 * Redundancy here is needed to avoid shift/reduce conflicts,
 * since TEMP is not a reserved word.  See also OptTempTableName.
 *
 * NOTE: we accept both GLOBAL and LOCAL options.  They currently do nothing,
 * but future versions might consider GLOBAL to request SQL-spec-compliant
 * temp table behavior, so warn about that.  Since we have no modules the
 * LOCAL keyword is really meaningless; furthermore, some other products
 * implement LOCAL as meaning the same as our default temp table behavior,
 * so we'll probably continue to treat LOCAL as a noise word.
 */
OptTemp:	TEMPORARY					{  }
			| TEMP						{  }
			| LOCAL TEMPORARY			{ }
			| LOCAL TEMP				{ }
			| GLOBAL TEMPORARY
				{
				}
			| GLOBAL TEMP
				{
				}
			| UNLOGGED					{ }
			| /*EMPTY*/					{ }
		;



create_as_target:
			qualified_name opt_column_list OptWith OnCommitOption OptTableSpace
				{

				}
		;

opt_with_data:
			WITH DATA_P								{ }
			| WITH NO DATA_P						{ }
			| /*EMPTY*/								{ }
		;

/*
 * General expressions
 * This is the heart of the expression syntax.
 *
 * We have two expression types: a_expr is the unrestricted kind, and
 * b_expr is a subset that must be used in some places to avoid shift/reduce
 * conflicts.  For example, we can't do BETWEEN as "BETWEEN a_expr AND a_expr"
 * because that use of AND conflicts with AND as a boolean operator.  So,
 * b_expr is used in BETWEEN and we remove boolean keywords from b_expr.
 *
 * Note that '(' a_expr ')' is a b_expr, so an unrestricted expression can
 * always be used by surrounding it with parens.
 *
 * c_expr is all the productions that are common to a_expr and b_expr;
 * it's factored out just to eliminate redundant coding.
 *
 * Be careful of productions involving more than one terminal token.
 * By default, bison will assign such productions the precedence of their
 * last terminal, but in nearly all cases you want it to be the precedence
 * of the first terminal instead; otherwise you will not get the behavior
 * you expect!  So we use %prec annotations freely to set precedences.
 */

opt_collate_clause:
			COLLATE any_val
				{

				}
			| /* EMPTY */				{ }
		;


func_name: IDENT


OptTableFuncElementList:
			TableFuncElementList				{ $$ = $1; }
			| /*EMPTY*/							{ $$ = nil; }
		;

TableFuncElementList:
			TableFuncElement
				{
					$$ = []string{$1};
				}
			| TableFuncElementList TCOMMA TableFuncElement
				{
					$$ = append($1, $3);
				}
		;

TableFuncElement:	ColId any_val opt_collate_clause
				{
					$$ = $1
				}
		|  ColId IDENT opt_collate_clause
				{
					$$ = $1
				}
		;



func_application: 
		func_name TOPENBR TCLOSEBR {
			$$ = nil
		} 
		| func_name TOPENBR func_arg_expr TCLOSEBR {
			$$ = nil
		}
		/* select count(*) from x */
		| func_name TOPENBR TMUL TCLOSEBR {
			$$ = nil
		}

func_arg_expr: a_expr {
	$$ = nil
}

func_arg_list:  func_arg_expr
				{
					$$ = []Node{$1}
				}
			| func_arg_list TCOMMA func_arg_expr
				{
					$$ = append($1, $3)
				}
		;

/*
 * Constants
 */
AexprConst: SCONST {
	$$ = &AExprConst{
		Value: $1,
	}
} | NULL_P {
	$$ = &AExprConst{
		Value: $1,
	}
}| TRUE_P {
	$$ = &AExprConst{
		Value: $1,
	}
}| FALSE_P {
	$$ = &AExprConst{
		Value: $1,
	}
}

func_expr: func_application

expr_list:	a_expr
				{
					$$ = []Node{$1};
				}
			| expr_list TCOMMA a_expr
				{
					$$ = append($1, $3);
				}
		;

opt_asymmetric: ASYMMETRIC
			| /*EMPTY*/
		;

type_list:	Typename								{ }
			| type_list TCOMMA Typename				{ }
		;

array_expr: TSQOPENBR expr_list TSQCLOSEBR
				{
					
				}
			// | TSQOPENBR array_expr_list TSQCLOSEBR
			// 	{
					
			// 	}
			| TSQOPENBR TSQCLOSEBR
				{
				
				}
		;

explicit_row:	ROW TOPENBR expr_list TCLOSEBR				{  }
			| ROW TOPENBR TCLOSEBR							{  }
		;


implicit_row:	TOPENBR expr_list TCOMMA a_expr TCLOSEBR		{  }
		;


indirection_el:
			TDOT attr_name
				{
				}
			| TDOT TMUL
				{
				}
			| TSQOPENBR a_expr TSQCLOSEBR
				{

				}
			| TSQOPENBR opt_slice_bound TCOLON opt_slice_bound TSQCLOSEBR
				{

				}
		;

opt_slice_bound:
			a_expr									{  }
			| /*EMPTY*/								{  }
		;

indirection:
			indirection_el							{ }
			| indirection indirection_el			{  }
		;

opt_indirection:
			/*EMPTY*/								{  }
			| opt_indirection indirection_el		{ }
		;



/*
 * Productions that can be used in both a_expr and b_expr.
 *
 * Note: productions that refer recursively to a_expr or b_expr mostly
 * cannot appear here.	However, it's OK to refer to a_exprs that occur
 * inside parentheses, such as function arguments; that cannot introduce
 * ambiguity to the b_expr syntax.
 */
c_expr:	

	// columnref								
	// 			{ 
	// 				$$ = &AExprConst{
	// 					Value: $1,
	// 				}
	// 			}
			AexprConst
				{ $$ = $1; }
			| ColRef {
					$$ = $1
			}
			| PARAM opt_indirection
				{
					$$ = &ParamRef {
						Number: $1,
					}
				}
            | TOPENBR a_expr TCLOSEBR {
                $$ = $2
            }
            | func_expr {}
			| ARRAY array_expr
				{
				}
			| implicit_row {} 
			| explicit_row {}

a_expr:		
            c_expr									{ $$ = $1; }
			| a_expr TYPECAST Typename
					{  }
			// | a_expr COLLATE any_name
			// 	{
			// 		CollateClause *n = makeNode(CollateClause);

			// 		n->arg = $1;
			// 		n->collname = $3;
			// 		n->location = @2;
			// 		$$ = (Node *) n;
			// 	}
			// | a_expr AT TIME ZONE a_expr			%prec AT
			// 	{
			// 		$$ = (Node *) makeFuncCall(SystemFuncName("timezone"),
			// 								   list_make2($5, $1),
			// 								   COERCE_SQL_SYNTAX,
			// 								   @2);
			// 	}
		/*
		 * These operators must be called out explicitly in order to make use
		 * of bison's automatic operator-precedence handling.  All other
		 * operator names are handled by the generic productions using "Op",
		 * below; and all those operators will have the same precedence.
		 *
		 * If you add more explicitly-known operators, be sure to add them
		 * also to b_expr and to the MathOp list below.
		 */
            |
			 TPLUS a_expr					%prec TMINUS
				{ $$ = $2 }
			| TMINUS a_expr					%prec TMINUS
					{ $$ = $2 }
			| a_expr TPLUS a_expr
				{
                     $$ = &AExprOp{
                        Left: $1,
                        Right: $3,
                        Op: $2,
                    } 
                }
			| a_expr TMINUS a_expr
				{
                     $$ = &AExprOp{
                        Left: $1,
                        Right: $3,
                        Op: $2,
                    } 
                }
			| a_expr TMUL a_expr
				{
                     $$ = &AExprOp{
                        Left: $1,
                        Right: $3,
                        Op: $2,
                    } 
                }
			| a_expr TDIV a_expr
				{
                     $$ = &AExprOp{
                        Left: $1,
                        Right: $3,
                        Op: $2,
                    } 
                }
			| a_expr TMOD a_expr
				{
                     $$ = &AExprOp{
                        Left: $1,
                        Right: $3,
                        Op: $2,
                    } 
                }
			| a_expr TPOW a_expr
				{
                     $$ = &AExprOp{
                        Left: $1,
                        Right: $3,
                        Op: $2,
                    } 
                }
			| a_expr TLESS a_expr
					{
                     $$ = &AExprOp{
                        Left: $1,
                        Right: $3,
                        Op: $2,
                    } 
                }
			| a_expr TGREATER a_expr
				{
                     $$ = &AExprOp{
                        Left: $1,
                        Right: $3,
                        Op: $2,
                    } 
                }
			| a_expr TEQ a_expr
						{
                     $$ = &AExprOp{
                        Left: $1,
                        Right: $3,
                        Op: $2,
                    } 
                }
			| a_expr TLESS_EQUALS a_expr
				{
                     $$ = &AExprOp{
                        Left: $1,
                        Right: $3,
                        Op: $2,
                    } 
                }
			| a_expr TGREATER_EQUALS a_expr
				{
                     $$ = &AExprOp{
                        Left: $1,
                        Right: $3,
                        Op: $2,
                    } 
                }
			| a_expr TNOT_EQUALS a_expr
				{
                     $$ = &AExprOp{
                        Left: $1,
                        Right: $3,
                        Op: $2,
                    } 
                }

			| a_expr OP a_expr				%prec OP
				{
                     $$ = &AExprOp{
                        Left: $1,
                        Right: $3,
                        Op: $2,
                    } 
                }
			// | a_expr qual_Op a_expr				%prec Op
			// 	{ $$ = (Node *) makeA_Expr(AEXPR_OP, $2, $1, $3, @2); }
			// | qual_Op a_expr					%prec Op
			// 	{ $$ = (Node *) makeA_Expr(AEXPR_OP, $1, NULL, $2, @1); }

			| a_expr AND a_expr
				{
                     $$ = &AExprOp{
                        Left: $1,
                        Right: $3,
                        Op: $2,
                    } 
                }
			| a_expr OR a_expr
				{
                     $$ = &AExprOp{
                        Left: $1,
                        Right: $3,
                        Op: $2,
                    } 
                }
			| NOT a_expr
				{
                     $$ = $2
                }
			// | NOT_LA a_expr						%prec NOT
			// 	{ $$ = makeNotExpr($2, @1); }

			// | a_expr LIKE a_expr
			// 	{
			// 		$$ = (Node *) makeSimpleA_Expr(AEXPR_LIKE, "~~",
			// 									   $1, $3, @2);
			// 	}
			// | a_expr LIKE a_expr ESCAPE a_expr					%prec LIKE
			// 	{
			// 		FuncCall   *n = makeFuncCall(SystemFuncName("like_escape"),
			// 									 list_make2($3, $5),
			// 									 COERCE_EXPLICIT_CALL,
			// 									 @2);
			// 		$$ = (Node *) makeSimpleA_Expr(AEXPR_LIKE, "~~",
			// 									   $1, (Node *) n, @2);
			// 	}
			// | a_expr NOT_LA LIKE a_expr							%prec NOT_LA
			// 	{
			// 		$$ = (Node *) makeSimpleA_Expr(AEXPR_LIKE, "!~~",
			// 									   $1, $4, @2);
			// 	}
			// | a_expr NOT_LA LIKE a_expr ESCAPE a_expr			%prec NOT_LA
			// 	{
			// 		FuncCall   *n = makeFuncCall(SystemFuncName("like_escape"),
			// 									 list_make2($4, $6),
			// 									 COERCE_EXPLICIT_CALL,
			// 									 @2);
			// 		$$ = (Node *) makeSimpleA_Expr(AEXPR_LIKE, "!~~",
			// 									   $1, (Node *) n, @2);
			// 	}
			// | a_expr ILIKE a_expr
			// 	{
			// 		$$ = (Node *) makeSimpleA_Expr(AEXPR_ILIKE, "~~*",
			// 									   $1, $3, @2);
			// 	}
			// | a_expr ILIKE a_expr ESCAPE a_expr					%prec ILIKE
			// 	{
			// 		FuncCall   *n = makeFuncCall(SystemFuncName("like_escape"),
			// 									 list_make2($3, $5),
			// 									 COERCE_EXPLICIT_CALL,
			// 									 @2);
			// 		$$ = (Node *) makeSimpleA_Expr(AEXPR_ILIKE, "~~*",
			// 									   $1, (Node *) n, @2);
			// 	}
			// | a_expr NOT_LA ILIKE a_expr						%prec NOT_LA
			// 	{
			// 		$$ = (Node *) makeSimpleA_Expr(AEXPR_ILIKE, "!~~*",
			// 									   $1, $4, @2);
			// 	}
			// | a_expr NOT_LA ILIKE a_expr ESCAPE a_expr			%prec NOT_LA
			// 	{
			// 		FuncCall   *n = makeFuncCall(SystemFuncName("like_escape"),
			// 									 list_make2($4, $6),
			// 									 COERCE_EXPLICIT_CALL,
			// 									 @2);
			// 		$$ = (Node *) makeSimpleA_Expr(AEXPR_ILIKE, "!~~*",
			// 									   $1, (Node *) n, @2);
			// 	}

			// | a_expr SIMILAR TO a_expr							%prec SIMILAR
			// 	{
			// 		FuncCall   *n = makeFuncCall(SystemFuncName("similar_to_escape"),
			// 									 list_make1($4),
			// 									 COERCE_EXPLICIT_CALL,
			// 									 @2);
			// 		$$ = (Node *) makeSimpleA_Expr(AEXPR_SIMILAR, "~",
			// 									   $1, (Node *) n, @2);
			// 	}
			// | a_expr SIMILAR TO a_expr ESCAPE a_expr			%prec SIMILAR
			// 	{
			// 		FuncCall   *n = makeFuncCall(SystemFuncName("similar_to_escape"),
			// 									 list_make2($4, $6),
			// 									 COERCE_EXPLICIT_CALL,
			// 									 @2);
			// 		$$ = (Node *) makeSimpleA_Expr(AEXPR_SIMILAR, "~",
			// 									   $1, (Node *) n, @2);
			// 	}
			// | a_expr NOT_LA SIMILAR TO a_expr					%prec NOT_LA
			// 	{
			// 		FuncCall   *n = makeFuncCall(SystemFuncName("similar_to_escape"),
			// 									 list_make1($5),
			// 									 COERCE_EXPLICIT_CALL,
			// 									 @2);
			// 		$$ = (Node *) makeSimpleA_Expr(AEXPR_SIMILAR, "!~",
			// 									   $1, (Node *) n, @2);
			// 	}
			// | a_expr NOT_LA SIMILAR TO a_expr ESCAPE a_expr		%prec NOT_LA
			// 	{
			// 		FuncCall   *n = makeFuncCall(SystemFuncName("similar_to_escape"),
			// 									 list_make2($5, $7),
			// 									 COERCE_EXPLICIT_CALL,
			// 									 @2);
			// 		$$ = (Node *) makeSimpleA_Expr(AEXPR_SIMILAR, "!~",
			// 									   $1, (Node *) n, @2);
			// 	}

			/* NullTest clause
			 * Define SQL-style Null test clause.
			 * Allow two forms described in the standard:
			 *	a IS NULL
			 *	a IS NOT NULL
			 * Allow two SQL extensions
			 *	a ISNULL
			 *	a NOTNULL
			 */
			| a_expr IS NULL_P							%prec IS
                { $$ = $1 }
			| a_expr ISNULL
                 { $$ = $1 }
			| a_expr IS NOT NULL_P						%prec IS
                    { $$ = $1 }
			| a_expr NOTNULL
				 { $$ = $1 }
			// | row OVERLAPS row
			// 	{
			// 		if (list_length($1) != 2)
			// 			ereport(ERROR,
			// 					(errcode(ERRCODE_SYNTAX_ERROR),
			// 					 errmsg("wrong number of parameters on left side of OVERLAPS expression"),
			// 					 parser_errposition(@1)));
			// 		if (list_length($3) != 2)
			// 			ereport(ERROR,
			// 					(errcode(ERRCODE_SYNTAX_ERROR),
			// 					 errmsg("wrong number of parameters on right side of OVERLAPS expression"),
			// 					 parser_errposition(@3)));
			// 		$$ = (Node *) makeFuncCall(SystemFuncName("overlaps"),
			// 								   list_concat($1, $3),
			// 								   COERCE_SQL_SYNTAX,
			// 								   @2);
			// 	}
			| a_expr IS TRUE_P							%prec IS
				 { $$ = $1 }
			| a_expr IS NOT TRUE_P						%prec IS
				 { $$ = $1 }
			| a_expr IS FALSE_P							%prec IS
				 { $$ = $1 }
			| a_expr IS NOT FALSE_P						%prec IS
				 { $$ = $1 }
			// | a_expr IS UNKNOWN							%prec IS
			// 	{
			// 		BooleanTest *b = makeNode(BooleanTest);

			// 		b->arg = (Expr *) $1;
			// 		b->booltesttype = IS_UNKNOWN;
			// 		b->location = @2;
			// 		$$ = (Node *) b;
			// 	}
			// | a_expr IS NOT UNKNOWN						%prec IS
			// 	{
			// 		BooleanTest *b = makeNode(BooleanTest);

			// 		b->arg = (Expr *) $1;
			// 		b->booltesttype = IS_NOT_UNKNOWN;
			// 		b->location = @2;
			// 		$$ = (Node *) b;
			// 	}
			| a_expr IS DISTINCT FROM a_expr			%prec IS
				{
                     $$ = &AExprOp{
                        Left: $1,
                        Right: $5,
                        Op: $3,
                    } 
                }
			| a_expr IS NOT DISTINCT FROM a_expr		%prec IS
				{
                     $$ = &AExprOp{
                        Left: $1,
                        Right: $6,
                        Op: $3,
                    } 
                }
			| a_expr BETWEEN opt_asymmetric b_expr AND a_expr		%prec BETWEEN
				{
                     $$ = &AExprOp{
                        Left: $1,
                        Right: &AExprList{
							List: []Node{
								$4,
								$6,
							},
						},
                        Op: "BETWEEN",
                    }
				}
			// | a_expr NOT_LA BETWEEN opt_asymmetric b_expr AND a_expr %prec NOT_LA
			// 	{
			// 		$$ = (Node *) makeSimpleA_Expr(AEXPR_NOT_BETWEEN,
			// 									   "NOT BETWEEN",
			// 									   $1,
			// 									   (Node *) list_make2($5, $7),
			// 									   @2);
			// 	}
			// | a_expr BETWEEN SYMMETRIC b_expr AND a_expr			%prec BETWEEN
			// 	{
			// 		$$ = (Node *) makeSimpleA_Expr(AEXPR_BETWEEN_SYM,
			// 									   "BETWEEN SYMMETRIC",
			// 									   $1,
			// 									   (Node *) list_make2($4, $6),
			// 									   @2);
			// 	}
			// | a_expr NOT_LA BETWEEN SYMMETRIC b_expr AND a_expr		%prec NOT_LA
			// 	{
			// 		$$ = (Node *) makeSimpleA_Expr(AEXPR_NOT_BETWEEN_SYM,
			// 									   "NOT BETWEEN SYMMETRIC",
			// 									   $1,
			// 									   (Node *) list_make2($5, $7),
			// 									   @2);
			// 	}
			// | a_expr IN_P in_expr
			// 	{
			// 		/* in_expr returns a SubLink or a list of a_exprs */
			// 		if (IsA($3, SubLink))
			// 		{
			// 			/* generate foo = ANY (subquery) */
			// 			SubLink	   *n = (SubLink *) $3;

			// 			n->subLinkType = ANY_SUBLINK;
			// 			n->subLinkId = 0;
			// 			n->testexpr = $1;
			// 			n->operName = NIL;		/* show it's IN not = ANY */
			// 			n->location = @2;
			// 			$$ = (Node *) n;
			// 		}
			// 		else
			// 		{
			// 			/* generate scalar IN expression */
			// 			$$ = (Node *) makeSimpleA_Expr(AEXPR_IN, "=", $1, $3, @2);
			// 		}
			// 	}
			// | a_expr NOT_LA IN_P in_expr						%prec NOT_LA
			// 	{
			// 		/* in_expr returns a SubLink or a list of a_exprs */
			// 		if (IsA($4, SubLink))
			// 		{
			// 			/* generate NOT (foo = ANY (subquery)) */
			// 			/* Make an = ANY node */
			// 			SubLink	   *n = (SubLink *) $4;

			// 			n->subLinkType = ANY_SUBLINK;
			// 			n->subLinkId = 0;
			// 			n->testexpr = $1;
			// 			n->operName = NIL;		/* show it's IN not = ANY */
			// 			n->location = @2;
			// 			/* Stick a NOT on top; must have same parse location */
			// 			$$ = makeNotExpr((Node *) n, @2);
			// 		}
			// 		else
			// 		{
			// 			/* generate scalar NOT IN expression */
			// 			$$ = (Node *) makeSimpleA_Expr(AEXPR_IN, "<>", $1, $4, @2);
			// 		}
			// 	}
			// | a_expr subquery_Op sub_type select_with_parens	%prec Op
			// 	{
			// 		SubLink	   *n = makeNode(SubLink);

			// 		n->subLinkType = $3;
			// 		n->subLinkId = 0;
			// 		n->testexpr = $1;
			// 		n->operName = $2;
			// 		n->subselect = $4;
			// 		n->location = @2;
			// 		$$ = (Node *) n;
			// 	}
			// | a_expr subquery_Op sub_type '(' a_expr ')'		%prec Op
			// 	{
			// 		if ($3 == ANY_SUBLINK)
			// 			$$ = (Node *) makeA_Expr(AEXPR_OP_ANY, $2, $1, $5, @2);
			// 		else
			// 			$$ = (Node *) makeA_Expr(AEXPR_OP_ALL, $2, $1, $5, @2);
			// 	}
			// | UNIQUE opt_unique_null_treatment select_with_parens
			// 	{
			// 		/* Not sure how to get rid of the parentheses
			// 		 * but there are lots of shift/reduce errors without them.
			// 		 *
			// 		 * Should be able to implement this by plopping the entire
			// 		 * select into a node, then transforming the target expressions
			// 		 * from whatever they are into count(*), and testing the
			// 		 * entire result equal to one.
			// 		 * But, will probably implement a separate node in the executor.
			// 		 */
			// 		ereport(ERROR,
			// 				(errcode(ERRCODE_FEATURE_NOT_SUPPORTED),
			// 				 errmsg("UNIQUE predicate is not yet implemented"),
			// 				 parser_errposition(@1)));
			// 	}
			| a_expr IS DOCUMENT_P					%prec IS
				    { $$ = $1 }
			| a_expr IS NOT DOCUMENT_P				%prec IS
				    {  $$ = $1 }
			| a_expr IS NORMALIZED								%prec IS
				    {  $$ = $1}
			// | a_expr IS unicode_normal_form NORMALIZED			%prec IS
			// 	    { /* result not matter */ }
			| a_expr IS NOT NORMALIZED							%prec IS
				    {  $$ = $1}
			// | a_expr IS NOT unicode_normal_form NORMALIZED		%prec IS
			// 	{
			// 		$$ = makeNotExpr((Node *) makeFuncCall(SystemFuncName("is_normalized"),
			// 											   list_make2($1, makeStringConst($4, @4)),
			// 											   COERCE_SQL_SYNTAX,
			// 											   @2),
			// 						 @2);
			// 	}
			| DEFAULT
				{
					/*
					 * The SQL spec only allows DEFAULT in "contextually typed
					 * expressions", but for us, it's easier to allow it in
					 * any a_expr and then throw error during parse analysis
					 * if it's in an inappropriate context.  This way also
					 * lets us say something smarter than "syntax error".
					 */

				}
		;


qualColRef:
    IDENT TDOT ColId {
        $$ = &ColumnRef{
            ColName: $3,
            TableAlias: $1, 
        }
    }

ColRef:
    ColId {
        $$ = &ColumnRef{
            ColName: $1,
        }
    } | qualColRef {
        $$ = $1
    }

columnref:	ColId
				{
					$$ = $1
				}
			// TODO: support
			// | ColId indirection
			// 	{
			// 		$$ = $1
			// 	}
		;


b_expr:
			c_expr
				{ $$ = $1; }
			| TPLUS b_expr					//%prec UMINUS
				{ $$ = $2; }
			| TMINUS b_expr					//%prec UMINUS
					{ $$ = $2; }
			| b_expr TPLUS b_expr
					{ /* result not matter */ }
			| b_expr TMINUS b_expr
					{ /* result not matter */ }
			| b_expr TMUL b_expr
					{ /* result not matter */ }
			| b_expr TDIV b_expr
					{ /* result not matter */ }
			| b_expr TMOD b_expr
				    { /* result not matter */ }
			| b_expr TPOW b_expr
				    { /* result not matter */ }
			| b_expr TLESS b_expr
				    { /* result not matter */ }
			| b_expr TGREATER b_expr
				    { /* result not matter */ }
			| b_expr TEQ b_expr
				    { /* result not matter */ }
			| b_expr TLESS_EQUALS b_expr
				    { /* result not matter */ }
			| b_expr TGREATER_EQUALS b_expr
				    { /* result not matter */ }
			| b_expr TNOT_EQUALS b_expr
				    { /* result not matter */ }
            | func_expr {/* result not matter */}


where_clause:
    /* empty */
    {
        $$ = &AExprEmpty{}
    }
    | WHERE a_expr
    {
        $$ = $2
    }


opt_target_list: target_list						{  $$ = $1 }
			| /* EMPTY */							{  $$ = nil }
		;

target_list:
			target_el								{ $$ = []Node{$1} }
			| target_list TCOMMA target_el				{  $$ = append($1, $3) }
		;

target_el:	a_expr AS IDENT
				{
					$$ = $1
				}
			| a_expr IDENT
				{
					$$ = $1
				}
			| a_expr
				{
					$$ = $1
				}
			| TMUL /* '*' */
				{
					$$ = &AExprEmpty{}
				}
		;


varset_stmt:
    SET any_id TEQ any_val {
        $$ = &VarSet {
            Name: $2,
            Value: $4,
        }
    }
    | SET LOCAL any_id TEQ any_val {
        $$ = &VarSet {
            IsLocal: true,
            Name: $3,
            Value: $5,
        }
    }
    | RESET any_id {}
    | RESET ALL {}


begin_stmt:
    BEGIN
    {
        $$ = &Begin{

        }
    }


commit_stmt:
    COMMIT
    {
        $$ = &Commit{
            
        }
    }


rollback_stmt:
    ROLLBACK
    {
        $$ = &Rollback{
            
        }
    }

routable_statement:
    SelectStmt  semicolon_opt {
        $$ = $1
    } | InsertStmt semicolon_opt {
        $$ = $1
    } | UpdateStmt semicolon_opt {
        $$ = $1
    } | DeleteStmt semicolon_opt {
        $$ = $1
    } | copy_stmt semicolon_opt {
        $$ = $1
    } | EXPLAIN routable_statement {
        $$ = &Explain{
            Stmt: $2,
        }
    } 

opt_pk:
    /* nothing */ {} |
    PRIMARY KEY {}

opt_ref:
    /* nothing */ {} |
    REFERENCES any_id TOPENBR any_id TCLOSEBR {}

create_stmt_coldefs:
    IDENT Typename opt_pk {
        $$ = []TableElt {{
                ColName: $1,
                ColType: $2,
            },
        }
    }
    | FOREIGN KEY TOPENBR any_id TCLOSEBR opt_ref {
        $$ = nil
    }
    | create_stmt_coldefs TCOMMA create_stmt_coldefs {
        $$ = append($1, $3...)
    }

create_stmt: 
    CREATE TABLE table_name TOPENBR create_stmt_coldefs TCLOSEBR anything {
        $$ = &CreateTable {
            TableName: $3,
            TableElts: $5,
        }
    } | CREATE INDEX anything {
        $$ = &Index{

        }
    } | CREATE ROLE anything {
        $$ = &CreateRole {

        }
    } | CREATE DATABASE anything {
        $$ = &CreateDatabase {

        }
    }

alter_stmt: 
    ALTER anything {
        $$ = &Alter{
            
        }
    }

vacuum_stmt:
    VACUUM anything {
        $$ = &Vacuum {

        }
    }

cluster_stmt:
    CLUSTER anything {
        $$ = &Cluster {
            
        }
    }


analyze_stmt:
    ANALYZE anything {
        $$ = &Analyze {
            
        }
    }

drop_stmt:
    DROP anything {
        $$ = &Drop {
            
        }
    }


truncate_stmt:
    TRUNCATE anything {
        $$ = &Truncate {
            
        }
    }



opt_asc_desc: ASC							{  }
			| DESC							{  }
			| /*EMPTY*/						{ }
		;


opt_nulls_order: NULLS_LA FIRST_P			{}
			| NULLS_LA LAST_P				{ }
			| /*EMPTY*/						{  }
		;



opt_sort_clause:
			sort_clause								{  }
			| /*EMPTY*/								{  }
		;

sort_clause:
			ORDER BY sortby_list					{  }
		;

sortby_list:
			sortby									{ }
			| sortby_list TCOMMA sortby				{ }
		;

sortby:
// sortby:		a_expr USING qual_all_Op opt_nulls_order
// 				{
// 				}
//          |
			 a_expr opt_asc_desc opt_nulls_order
				{
				/* no operator */
				}
		;



opt_group_by_clause: /*empty*/ {} | GROUP BY anything {}
opt_window_clause:  /*empty*/ {} | window_clause {}
opt_limit_clause: /*empty*/ {} | LIMIT anything {}
opt_offset_clause: /*empty*/ {} | OFFSET anything {}
opt_fetch_clause: /*empty*/ {} | FETCH anything {}
opt_for_clause:  /*empty*/ {} |FOR anything {}

/*
 * Name classification hierarchy.
 *
 * IDENT is the lexeme returned by the lexer for identifiers that match
 * no known keyword.  In most cases, we can accept certain keywords as
 * names, not only IDENTs.	We prefer to accept as many such keywords
 * as possible to minimize the impact of "reserved words" on programmers.
 * So, we divide names into several possible classes.  The classification
 * is chosen in part to make keywords acceptable as names wherever possible.
 */

/* Column identifier --- names that can be column, table, etc names.
 */
ColId:		IDENT									{ $$ = $1; }
			| unreserved_keyword					{ $$ = $1; }
			| col_name_keyword						{ $$ = $1; }
		;

/* Type/function identifier --- names that can be type or function names.
 */
type_function_name:	IDENT							{ $$ = $1; }
			| unreserved_keyword					{ $$ = $1; }
			| type_func_name_keyword				{ $$ = $1; }
		;

/* Any not-fully-reserved word --- these names can be, eg, role names.
 */
NonReservedWord:	IDENT							{ $$ = $1; }
			| unreserved_keyword					{ $$ = $1; }
			| col_name_keyword						{ $$ = $1; }
			| type_func_name_keyword				{ $$ = $1; }
		;

/* Column label --- allowed labels in "AS" clauses.
 * This presently includes *all* Postgres keywords.
 */
ColLabel:	IDENT									{ $$ = $1; }
			| unreserved_keyword					{ $$ = $1; }
			| col_name_keyword						{ $$ = $1; }
			| type_func_name_keyword				{ $$ = $1; }
			| reserved_keyword						{ $$ = $1; }
		;

/*
 * The production for a qualified relation name has to exactly match the
 * production for a qualified func_name, because in a FROM clause we cannot
 * tell which we are parsing until we see what comes after it ('(' for a
 * func_name, something else for a relation). Therefore we allow 'indirection'
 * which may contain subscripts, and reject that case in the C code.
 */
qualified_name:
			ColId {
			$$ = &RangeVar {
                    SchemaName: "",
                    RelationName: $1,
                    Alias: "",
                }
            }
			| ColId TDOT attr_name
				{
					$$ = &RangeVar {
						SchemaName: $1,
						RelationName: $3,
						Alias: "",
					}
				}
			// | ColId indirection
			// 	{
			// 		$$ = makeRangeVarFromQualifiedName($1, $2, @1, yyscanner);
			// 	}
		;


from_clause:
			FROM from_list							{ $$ = $2; }
			| /*EMPTY*/								{ }
		;

from_list:
			table_ref								{ $$ = []FromClauseNode{$1}; }
			| from_list TCOMMA table_ref				{ $$ = append($1, $3); }
		;

relation_expr:
			qualified_name
				{
					/* inheritance query, implicitly */
					$$ = $1;
				}
			// | extended_relation_expr
			// 	{
			// 		$$ = $1;
			// 	}
		;

alias_clause:
    AS any_id {
        $$ = $2
    }
    | 
     any_id {
        $$ = $1
    }


opt_alias_clause:
    /* nothing */ {
        $$ = ""
    } | alias_clause

/*
 * As func_expr but does not accept WINDOW functions directly
 * (but they can still be contained in arguments for functions etc).
 * Use this when window expressions are not allowed, where needed to
 * disambiguate the grammar (e.g. in CREATE INDEX).
 */
func_expr_windowless:
			func_application						{ $$ = $1; }
			// | func_expr_common_subexpr				{ $$ = $1; }
			// | json_aggregate_func					{ $$ = $1; }
		;

opt_ordinality: WITH_LA ORDINALITY					{ $$ = true; }
			| /*EMPTY*/								{ $$ = false; }
		;



/*
 * func_alias_clause can include both an Alias and a coldeflist, so we make it
 * return a 2-element list that gets disassembled by calling production.
 */
func_alias_clause:
			alias_clause
				{
					$$ = nil
				}
			| AS TOPENBR TableFuncElementList TCLOSEBR
				{
					$$ = nil
				}
			| AS ColId TOPENBR TableFuncElementList TCLOSEBR
				{

				}
			| ColId TOPENBR TableFuncElementList TCLOSEBR
				{

				}
			| /*EMPTY*/
				{
					$$ = nil
				}
		;

/*
 * func_table represents a function invocation in a FROM list. It can be
 * a plain function call, like "foo(...)", or a ROWS FROM expression with
 * one or more function calls, "ROWS FROM (foo(...), bar(...))",
 * optionally with WITH ORDINALITY attached.
 * In the ROWS FROM syntax, a column definition list can be given for each
 * function, for example:
 *     ROWS FROM (foo() AS (foo_res_a text, foo_res_b text),
 *                bar() AS (bar_res_a text, bar_res_b text))
 * It's also possible to attach a column definition list to the RangeFunction
 * as a whole, but that's handled by the table_ref production.
 */
func_table: func_expr_windowless opt_ordinality
				{
					$$ = $1
				}

/*
 * table_ref is where an alias clause can be attached.
 */
table_ref:	relation_expr opt_alias_clause
				{
					$1.SetAlias($2)
					$$ = $1;
				}
			| func_table func_alias_clause
				{

				}
			| LATERAL_P func_table func_alias_clause
				{

				}
			| joined_table
				{
					$$ = $1;
				}
			| TOPENBR joined_table TCLOSEBR alias_clause
				{
					$2.SetAlias($4)
					$$ = $2;
				}
		;


/* OUTER is just noise... */
opt_outer: OUTER_P {}
			| /*EMPTY*/ {}
		;


join_type:	FULL opt_outer							{  }
			| LEFT opt_outer						{ }
			| RIGHT opt_outer						{  }
			| INNER_P								{ }
		;


join_qual: 
            // USING TOPENBR name_list TCLOSEBR opt_alias_clause_for_join_using
			// 	{
					
			// 	}
			// |
             ON a_expr
				{
                    
				}
		;


joined_table:
			TOPENBR joined_table TCLOSEBR
				{
					$$ = $2;
				}
			| table_ref CROSS JOIN table_ref
				{
					/* CROSS JOIN is same as unqualified inner join */
					$$ = &JoinExpr{
                        Larg: $1,
                        Rarg: $4,
                    };
				}
			| table_ref join_type JOIN table_ref join_qual
				{
                    $$ = &JoinExpr{
                        Larg: $1,
                        Rarg: $4,
                    };
				}
			| table_ref JOIN table_ref join_qual
				{
					/* letting join_type reduce to empty doesn't work */
                    $$ = &JoinExpr{
                        Larg: $1,
                        Rarg: $3,
                    };
				}
			| table_ref NATURAL join_type JOIN table_ref
				{
                    $$ = &JoinExpr{
                        Larg: $1,
                        Rarg: $5,
                    };
				}
			| table_ref NATURAL JOIN table_ref
				{
					/* letting join_type reduce to empty doesn't work */
                    $$ = &JoinExpr{
                        Larg: $1,
                        Rarg: $4,
                    };
				}
		;


/*
 * Window Definitions
 */
window_clause:
			WINDOW window_definition_list			{ $$ = $2; }
			| /*EMPTY*/								{ $$ = NIL; }
		;

window_definition_list:
		anything {}
		// 	window_definition						{ $$ = list_make1($1); }
		// 	| window_definition_list ',' window_definition
		// 											{ $$ = lappend($1, $3); }
		// ;

/*
 * This rule parses SELECT statements that can appear within set operations,
 * including UNION, INTERSECT and EXCEPT.  '(' and ')' can be used to specify
 * the ordering of the set operations.	Without '(' and ')' we want the
 * operations to be ordered per the precedence specs at the head of this file.
 *
 * As with select_no_parens, simple_select cannot have outer parentheses,
 * but can have parenthesized subclauses.
 *
 * It might appear that we could fold the first two alternatives into one
 * by using opt_distinct_clause.  However, that causes a shift/reduce conflict
 * against INSERT ... SELECT ... ON CONFLICT.  We avoid the ambiguity by
 * requiring SELECT DISTINCT [ON] to be followed by a non-empty target_list.
 *
 * Note that sort clauses cannot be included at this level --- SQL requires
 *		SELECT foo UNION SELECT bar ORDER BY baz
 * to be parsed as
 *		(SELECT foo UNION SELECT bar) ORDER BY baz
 * not
 *		SELECT foo UNION (SELECT bar ORDER BY baz)
 * Likewise for WITH, FOR UPDATE and LIMIT.  Therefore, those clauses are
 * described as part of the select_no_parens production, not simple_select.
 * This does not limit functionality, because you can reintroduce these
 * clauses inside parentheses.
 *
 * NOTE: only the leftmost component SelectStmt should have INTO.
 * However, this is not checked by the grammar; parse analysis must check it.
 */
simple_select:
			SELECT opt_all_clause opt_target_list
			into_clause from_clause where_clause
			group_clause having_clause window_clause
				{
					$$ = &Select{
						TargetList: $3,
						FromClause: $4,
						Where: $5,
					}
				}
			| SELECT distinct_clause target_list
			into_clause from_clause where_clause
			group_clause having_clause window_clause
				{
					$$ = &Select{
						TargetList: $3,
						FromClause: $4,
						Where: $5,
					}
				}
			| values_clause							{ $$ = $1; }
		// 	| TABLE relation_expr
		// 		{
		// 			/* same as SELECT * FROM relation_expr */
		// 			ColumnRef  *cr = makeNode(ColumnRef);
		// 			ResTarget  *rt = makeNode(ResTarget);
		// 			SelectStmt *n = makeNode(SelectStmt);

		// 			cr->fields = list_make1(makeNode(A_Star));
		// 			cr->location = -1;

		// 			rt->name = NULL;
		// 			rt->indirection = NIL;
		// 			rt->val = (Node *) cr;
		// 			rt->location = -1;

		// 			n->targetList = list_make1(rt);
		// 			n->fromClause = list_make1($2);
		// 			$$ = (Node *) n;
		// 		}
		// 	| select_clause UNION set_quantifier select_clause
		// 		{
		// 			$$ = makeSetOp(SETOP_UNION, $3 == SET_QUANTIFIER_ALL, $1, $4);
		// 		}
		// 	| select_clause INTERSECT set_quantifier select_clause
		// 		{
		// 			$$ = makeSetOp(SETOP_INTERSECT, $3 == SET_QUANTIFIER_ALL, $1, $4);
		// 		}
		// 	| select_clause EXCEPT set_quantifier select_clause
		// 		{
		// 			$$ = makeSetOp(SETOP_EXCEPT, $3 == SET_QUANTIFIER_ALL, $1, $4);
		// 		}
		// ;


/* We use (NIL) as a placeholder to indicate that all target expressions
 * should be placed in the DISTINCT list during parsetree analysis.
 */
distinct_clause:
			DISTINCT								{  }
			| DISTINCT ON TOPENBR expr_list TCLOSEBR		{ $$ = $4; }
		;

opt_all_clause:
			ALL
			| /*EMPTY*/
		;

opt_distinct_clause:
			distinct_clause							{ $$ = $1; }
			| opt_all_clause						{ $$ = NIL; }
		;


select_limit:
			limit_clause offset_clause
				{
				}
			| offset_clause limit_clause
				{
				}
			| limit_clause
				{
				}
			| offset_clause
				{
				}
		;


opt_select_limit:
			select_limit						{ $$ = $1; }
			| /* EMPTY */						{  }
		;

limit_clause:
			LIMIT select_limit_value
				{

				}
			| LIMIT select_limit_value ',' select_offset_value
				{
					// XXXX: todo forbid

				}
			/* SQL:2008 syntax */
			/* to avoid shift/reduce conflicts, handle the optional value with
			 * a separate production rather than an opt_ expression.  The fact
			 * that ONLY is fully reserved means that this way, we defer any
			 * decision about what rule reduces ROW or ROWS to the point where
			 * we can see the ONLY token in the lookahead slot.
			 */
			| FETCH first_or_next select_fetch_first_value row_or_rows ONLY
				{
				}
			| FETCH first_or_next select_fetch_first_value row_or_rows WITH TIES
				{
				}
			| FETCH first_or_next row_or_rows ONLY
				{
				}
			| FETCH first_or_next row_or_rows WITH TIES
				{
				}
		;

offset_clause:
			OFFSET select_offset_value
				{ $$ = $2; }
			/* SQL:2008 syntax */
			| OFFSET select_fetch_first_value row_or_rows
				{ $$ = $2; }
		;

select_limit_value:
			a_expr									{ $$ = $1; }
			| ALL
				{
				}
		;

select_offset_value:
			a_expr									{ $$ = $1; }
		;

/*
 * Allowing full expressions without parentheses causes various parsing
 * problems with the trailing ROW/ROWS key words.  SQL spec only calls for
 * <simple value specification>, which is either a literal or a parameter (but
 * an <SQL parameter reference> could be an identifier, bringing up conflicts
 * with ROW/ROWS). We solve this by leveraging the presence of ONLY (see above)
 * to determine whether the expression is missing rather than trying to make it
 * optional in this rule.
 *
 * c_expr covers almost all the spec-required cases (and more), but it doesn't
 * cover signed numeric literals, which are allowed by the spec. So we include
 * those here explicitly. We need FCONST as well as ICONST because values that
 * don't fit in the platform's "long", but do fit in bigint, should still be
 * accepted here. (This is possible in 64-bit Windows as well as all 32-bit
 * builds.)
 */
select_fetch_first_value:
			c_expr									{ $$ = $1; }
			| TPLUS I_or_F_const
				{  }
			| TMINUS I_or_F_const
				{}
		;

I_or_F_const:
		SCONST {};
		// 	Iconst									{ }
		// 	| FCONST								{  }
		// ;

/* noise words */
row_or_rows: ROW									{  }
			| ROWS									{ }
		;

first_or_next: FIRST_P								{  }
			| NEXT									{  }
		;


/*
 * This syntax for group_clause tries to follow the spec quite closely.
 * However, the spec allows only column references, not expressions,
 * which introduces an ambiguity between implicit row constructors
 * (a,b) and lists of column references.
 *
 * We handle this by using the a_expr production for what the spec calls
 * <ordinary grouping set>, which in the spec represents either one column
 * reference or a parenthesized list of column references. Then, we check the
 * top node of the a_expr to see if it's an implicit RowExpr, and if so, just
 * grab and use the list, discarding the node. (this is done in parse analysis,
 * not here)
 *
 * (we abuse the row_format field of RowExpr to distinguish implicit and
 * explicit row constructors; it's debatable if anyone sanely wants to use them
 * in a group clause, but if they have a reason to, we make it possible.)
 *
 * Each item in the group_clause list is either an expression tree or a
 * GroupingSet node of some type.
 */
group_clause:
			GROUP_P BY set_quantifier group_by_list
				{

				}
			| /*EMPTY*/
				{
				}
		;

group_by_list:
			group_by_item							{ }
			| group_by_list TCOMMA group_by_item		{ }
		;

group_by_item:
			a_expr									{ $$ = $1; }
			| empty_grouping_set					{ $$ = $1; }
			| cube_clause							{ $$ = $1; }
			| rollup_clause							{ $$ = $1; }
			| grouping_sets_clause					{ $$ = $1; }
		;

empty_grouping_set:
			TOPENBR TCLOSEBR
				{
				}
		;

/*
 * These hacks rely on setting precedence of CUBE and ROLLUP below that of '(',
 * so that they shift in these rules rather than reducing the conflicting
 * unreserved_keyword rule.
 */

rollup_clause:
			ROLLUP TOPENBR expr_list TCLOSEBR
				{
				}
		;

cube_clause:
			CUBE TOPENBR expr_list TCLOSEBR
				{
				}
		;

grouping_sets_clause:
			GROUPING SETS TOPENBR group_by_list TCLOSEBR
				{
				}
		;

having_clause:
			HAVING a_expr							{  }
			| /*EMPTY*/								{ }
		;

for_locking_clause:
			for_locking_items						{ }
			| FOR READ ONLY							{ }
		;

opt_for_locking_clause:
			for_locking_clause						{ }
			| /* EMPTY */							{ }
		;

for_locking_items:
			for_locking_item						{  }
			| for_locking_items for_locking_item	{ }
		;

for_locking_item:
			for_locking_strength locked_rels_list opt_nowait_or_skip
				{
				}
		;

for_locking_strength:
			FOR UPDATE							{ }
			| FOR NO KEY UPDATE					{ }
			| FOR SHARE							{  }
			| FOR KEY SHARE						{ }
		;


locked_rels_list:
			OF qualified_name_list					{ $$ = $2; }
			| /* EMPTY */							{ }
		;

/*
 * We should allow ROW '(' expr_list ')' too, but that seems to require
 * making VALUES a fully reserved word, which will probably break more apps
 * than allowing the noise-word is worth.
 */
values_clause:
			VALUES TOPENBR expr_list TCLOSEBR
				{
					$$ = $3
				}
			| values_clause TCOMMA TOPENBR expr_list TCLOSEBR
				{
					$$ = append($1, $3...)
				}
		;



/*****************************************************************************
 *
 *		QUERY:
 *				SELECT STATEMENTS
 *
 *****************************************************************************/

/* A complete SELECT statement looks like this.
 *
 * The rule returns either a single SelectStmt node or a tree of them,
 * representing a set-operation tree.
 *
 * There is an ambiguity when a sub-SELECT is within an a_expr and there
 * are excess parentheses: do the parentheses belong to the sub-SELECT or
 * to the surrounding a_expr?  We don't really care, but bison wants to know.
 * To resolve the ambiguity, we are careful to define the grammar so that
 * the decision is staved off as long as possible: as long as we can keep
 * absorbing parentheses into the sub-SELECT, we will do so, and only when
 * it's no longer possible to do that will we decide that parens belong to
 * the expression.	For example, in "SELECT (((SELECT 2)) + 3)" the extra
 * parentheses are treated as part of the sub-select.  The necessity of doing
 * it that way is shown by "SELECT (((SELECT 2)) UNION SELECT 2)".	Had we
 * parsed "((SELECT 2))" as an a_expr, it'd be too late to go back to the
 * SELECT viewpoint when we see the UNION.
 *
 * This approach is implemented by defining a nonterminal select_with_parens,
 * which represents a SELECT with at least one outer layer of parentheses,
 * and being careful to use select_with_parens, never '(' SelectStmt ')',
 * in the expression grammar.  We will then have shift-reduce conflicts
 * which we can resolve in favor of always treating '(' <select> ')' as
 * a select_with_parens.  To resolve the conflicts, the productions that
 * conflict with the select_with_parens productions are manually given
 * precedences lower than the precedence of ')', thereby ensuring that we
 * shift ')' (and then reduce to select_with_parens) rather than trying to
 * reduce the inner <select> nonterminal to something else.  We use UMINUS
 * precedence for this, which is a fairly arbitrary choice.
 *
 * To be able to define select_with_parens itself without ambiguity, we need
 * a nonterminal select_no_parens that represents a SELECT structure with no
 * outermost parentheses.  This is a little bit tedious, but it works.
 *
 * In non-expression contexts, we use SelectStmt which can represent a SELECT
 * with or without outer parentheses.
 */

SelectStmt: select_no_parens			
			| select_with_parens	
		;

select_with_parens:
			TOPENBR select_no_parens TCLOSEBR				{ $$ = $2; }
			| TOPENBR select_with_parens TCLOSEBR			{ $$ = $2; }
		;

/*
 * This rule parses the equivalent of the standard's <query expression>.
 * The duplicative productions are annoying, but hard to get rid of without
 * creating shift/reduce conflicts.
 *
 *	The locking clause (FOR UPDATE etc) may be before or after LIMIT/OFFSET.
 *	In <=7.2.X, LIMIT/OFFSET had to be after FOR UPDATE
 *	We now support both orderings, but prefer LIMIT/OFFSET before the locking
 * clause.
 *	2002-08-28 bjm
 */
select_no_parens:
			simple_select						{ $$ = $1; }
			| select_clause sort_clause
				{
					$$ = $1;
				}
			| select_clause opt_sort_clause for_locking_clause opt_select_limit
				{
					$$ = $1;
				}
			| select_clause opt_sort_clause select_limit opt_for_locking_clause
				{
					$$ = $1;
				}
			| with_clause select_clause
				{
					$$ = $2;
				}
			| with_clause select_clause sort_clause
				{
					$$ = $2;
				}
			| with_clause select_clause opt_sort_clause for_locking_clause opt_select_limit
				{
					$$ = $2;
				}
			| with_clause select_clause opt_sort_clause select_limit opt_for_locking_clause
				{
					$$ = $2;
				}
		;

select_clause:
			simple_select							{ $$ = $1; }
			| select_with_parens					{ $$ = $1; }
		;


comma_separated_col_refs:
    ColId { $$ = []string{$1} }
    | comma_separated_col_refs TCOMMA ColId {
        $$ = append($1, $3)
    }

opt_comma_separated_col_refs:
    /*nothing*/{} | comma_separated_col_refs {}

insert_col_refs:
    TOPENBR comma_separated_col_refs TCLOSEBR {
        $$ = $2
    }


opt_column_list:
			TOPENBR columnList TCLOSEBR				{ $$ = $2; }
			| /*EMPTY*/								{ $$ = nil; }
		;


columnElem: ColId
				{
					$$ = $1
				}
		;

columnList:
			columnElem								{ $$ = []string{$1}; }
			| columnList TCOMMA columnElem				{ $$ = append($1, $3); }
		;


opt_insert_col_refs:
    /* nothing */ {
		$$ = nil
	} |
    insert_col_refs {
		$$ = $1
	}    




insert_comma_separated_tuples: 
    a_expr {
        $$ = []Node{$1}
    } | insert_comma_separated_tuples TCOMMA a_expr {
        $$ = append($1, $3)
    }

insert_tuples:
    TOPENBR insert_comma_separated_tuples TCLOSEBR {
        $$ = $2
    }


opt_returning:
    /* nothing */ {}
    | RETURNING anything {}


/*****************************************************************************
 *
 *		QUERY:
 *				PREPARE <plan_name> [(args, ...)] AS <query>
 *
 *****************************************************************************/

PrepareStmt: PREPARE name prep_type_clause AS PreparableStmt
				{
				
					$$ = &PrepareStmt{
						Name: $2,
						Statement: $5,
					};
				}
		;

prep_type_clause: TOPENBR type_list TCLOSEBR	{  }
				| /* EMPTY */				{  }
		;

PreparableStmt:
			SelectStmt
			| InsertStmt
			| UpdateStmt
			| DeleteStmt					/* by default all are $$=$1 */
		;

/*****************************************************************************
 *
 * EXECUTE <plan_name> [(params, ...)]
 * CREATE TABLE <name> AS EXECUTE <plan_name> [(params, ...)]
 *
 *****************************************************************************/

ExecuteStmt: EXECUTE name execute_param_clause
				{
					$$ = &ExecuteStmt{
						Name: $2,
					};
				}
			| CREATE OptTemp TABLE create_as_target AS
				EXECUTE name execute_param_clause opt_with_data
				{
				}
			| CREATE OptTemp TABLE IF_P NOT EXISTS create_as_target AS
				EXECUTE name execute_param_clause opt_with_data
				{

				}
		;

execute_param_clause: TOPENBR expr_list TCLOSEBR		{  }
					| /* EMPTY */					{ }
					;

/*****************************************************************************
 *
 *		QUERY:
 *				DEALLOCATE [PREPARE] <plan_name>
 *
 *****************************************************************************/

DeallocateStmt: DEALLOCATE name
					{
						$$ = &DeallocateStmt{
							$2,
						};
					}
				| DEALLOCATE PREPARE name
					{						
						$$ = &DeallocateStmt{
							$3,
						};
					}
				| DEALLOCATE ALL
					{
						$$ = &DeallocateStmt{
						};
					}
				| DEALLOCATE PREPARE ALL
					{
						$$ = &DeallocateStmt{
						};
					}
		;



/* https://www.postgresql.org/docs/current/sql-insert.html */
InsertStmt: 
    /* consider only first tuple from values */
    INSERT INTO relation_expr insert_col_refs VALUES insert_tuples anything {
        $$ = &Insert{
            TableRef: $3,
            Columns: $4,
            Values: $6,
        }
    } | INSERT INTO relation_expr opt_insert_col_refs SelectStmt {
        $$ = &Insert{
            TableRef: $3,
            Columns: $4,
            SubSelect: $5,
        }
    }

opt_only:
    /* nothing */ {}
    | ONLY {}

/* https://www.postgresql.org/docs/current/sql-update.html */


set_clause_list:
			set_clause							{ $$ = []string{$1}; }
			| set_clause_list TCOMMA set_clause	{ $$ = append($1, $3); }
		

set_clause:
			any_id TEQ a_expr
				{
					$$ = $1;
				}
			// | TOPENBR set_target_list TCLOSEBR TEQ a_expr
			// 	{
			// 		$$ = $2;
			// 	}
		

UpdateStmt:
    UPDATE opt_only relation_expr SET set_clause_list from_clause where_clause opt_returning {
        $$ = &Update {
            TableRef: $3,
            Where: $7,
        }
    }


/* https://www.postgresql.org/docs/current/sql-delete.html */

delete_comma_separated_using_refs:
    any_id { $$ = []string{$1} }
    | any_id TCOMMA delete_comma_separated_using_refs {
        $$ = append($3, $1)
    }

opt_using:
    /* nothing */ {}
    | USING delete_comma_separated_using_refs {}


DeleteStmt:
    DELETE FROM opt_only relation_expr opt_using where_clause opt_returning {
        $$ = &Delete{
            TableRef: $4,
            Where: $6,
        }
    }


opt_boolean_or_string:
			TRUE_P									{  }
			| FALSE_P								{}
			| ON									{ }
			/*
			 * OFF is also accepted as a boolean value, but is handled by
			 * the NonReservedWord rule.  The action for booleans and strings
			 * is the same, so we don't need to distinguish them here.
			 */
			| SCONST				{  }
		;

opt_binary:
			BINARY
				{
				}
			| /*EMPTY*/								{  }
		

/* new COPY option syntax */
copy_gengeneric_opt_list:
			copy_gengeneric_opt_elem
				{
					$$ = []string{$1};
				}
			| copy_gengeneric_opt_list TCOMMA copy_gengeneric_opt_elem
				{
					$$ = append($1, $3);
				}
		;

copy_gengeneric_opt_elem:
			ColLabel copy_generic_opt_arg
				{

				}
		;

copy_generic_opt_arg:
			opt_boolean_or_string			{  }
			// | NumericOnly					{  }
			| TMUL							{  }
			| TOPENBR copy_generic_opt_arg_list TCLOSEBR		{ }
			| /* EMPTY */					{  }
		;

copy_generic_opt_arg_list:
			  copy_generic_opt_arg_list_item
				{
					$$ = []string{$1};
				}
			| copy_generic_opt_arg_list TCOMMA copy_generic_opt_arg_list_item
				{
					$$ = append($1, $3);
				}
		;

/* beware of emitting non-string list elements here; see commands/define.c */
copy_generic_opt_arg_list_item:
			opt_boolean_or_string	{  }
		;


copy_options: copy_opt_list							{ $$ = $1; }
			| TOPENBR copy_generic_opt_list TCLOSEBR			{ $$ = $2; }
		;

/* old COPY option syntax */
copy_opt_list:
			copy_opt_list copy_opt_item				{ $$ = append($1, $2); }
			| /* EMPTY */							{ $$ = nil; }
		;

opt_as:		AS
			| /* EMPTY */
		;

copy_opt_item:
			BINARY
				{
				}
			| FREEZE
				{
				}
			| DELIMITER opt_as SCONST
				{
				}
			| NULL_P opt_as SCONST
				{
				}
			| CSV
				{
				}
			| HEADER_P
				{
				}
			| QUOTE opt_as SCONST
				{
				}
			| ESCAPE opt_as SCONST
				{
				}
			| FORCE QUOTE columnList
				{
				}
			| FORCE QUOTE TMUL
				{
				}
			| FORCE NOT NULL_P columnList
				{
				}
			| FORCE NULL_P columnList
				{
				}
			| ENCODING SCONST
				{
				}
		;


copy_from:
			FROM									{ $$ = true; }
			| TO									{ $$ = false; }
		;

opt_program:
			PROGRAM									{ $$ = true; }
			| /* EMPTY */							{ $$ = false; }
		;


opt_with:	WITH									{}
			| WITH_LA								{}
			| /*EMPTY*/								{}
		;


/* WITH (options) is preferred, WITH OIDS and WITHOUT OIDS are legacy forms */
OptWith:
			WITH reloptions				{ }
			| WITH OIDS					{ }
			| WITHOUT OIDS				{ }
			| /*EMPTY*/					{  }
		;

OnCommitOption:  ON COMMIT DROP				{ }
			| ON COMMIT DELETE_P ROWS		{ }
			| ON COMMIT PRESERVE ROWS		{ }
			| /*EMPTY*/						{ }
		;

OptTableSpace:   TABLESPACE name					{ }
			| /*EMPTY*/								{ }
		;

OptConsTableSpace:   USING INDEX TABLESPACE name	{ }
			| /*EMPTY*/								{ }
		;

ExistingIndex:   USING INDEX index_name				{ }
		;


copy_file_name:
			SCONST									{ }
			| STDIN									{  }
			| STDOUT								{  }
		;


copy_delimiter:
			opt_using DELIMITERS SCONST
				{
				}
			| /*EMPTY*/								{ }
		;


/* https://www.postgresql.org/docs/current/sql-copy.html */
copy_stmt: 
	COPY opt_binary qualified_name opt_column_list
			copy_from opt_program copy_file_name copy_delimiter opt_with
			copy_options where_clause {
				$$ = &Copy {
					TableRef: $3,
					Where: $11,
					IsFrom: $5,
				}   
		}
		| COPY TOPENBR PreparableStmt TCLOSEBR TO opt_program copy_file_name opt_with copy_options
				{
					$$ = &Copy {
						IsFrom: false,
						SubStmt: $3,
					}   
				}
		;



name_list:	name
					{  }
			| name_list TCOMMA name
					{  }
		;


name:		ColId									{ $$ = $1; };

database_name:
			ColId									{ $$ = $1; };

access_method:
			ColId									{ $$ = $1; };

attr_name:	ColLabel								{ $$ = $1; };

index_name: ColId									{ $$ = $1; };

file_name:	SCONST									{ $$ = $1; };

%%


