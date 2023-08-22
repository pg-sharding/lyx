
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

    selectStmt              *Select

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

/* LATERAL */
%token<str> LATERAL_P

/* WITH_LA ORDINALITY */
%token<str> ORDINALITY WITH_LA

/* COLLATE */
%token<str> COLLATE

%token<str> AS


%type<str> reserved_keyword


/**/

%type<node> execute_stmt
%type<node> select_stmt
%type<node> insert_stmt
%type<node> update_stmt
%type<node> delete_stmt
%type<node> varset_stmt
%type<node> prepare_stmt
%type<node> begin_stmt 
%type<node> commit_stmt
%type<node> rollback_stmt
%type<node> prepare_stmt
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

%type<node> copy_stmt

%type<node> target_el

%type<nodeList> opt_target_list
%type<nodeList> target_list

%type<node> ColRef qualColRef

%type<nodeList> expr_list

%type<str> ColId

%type<strlist> comma_separated_col_refs insert_col_refs

%type<nodeList> insert_comma_separated_tuples insert_tuples

%type<strlist> opt_insert_col_refs columnList opt_column_list copy_gengeneric_opt_list copy_generic_opt_arg_list copy_opt_list
%type<str> any_tok  columnElem copy_gengeneric_opt_elem copy_generic_opt_arg_list_item copy_opt_item

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


reserved_keyword:
	SCONST {$$=$1}
	| OFFSET {$$=$1}
	| ON {$$=$1}
	| PRECISION {$$=$1}
	| TCOMMA {$$=$1}
	| NUMERIC {$$=$1}
	| WITHOUT {$$=$1}
	| ZONE {$$=$1}
	| ALTER {$$=$1}
	| PREPARE {$$=$1}
	| INTO {$$=$1}
	| MONTH_P {$$=$1}
	| SECOND_P {$$=$1}
	| CSV {$$=$1}
	| DAY_P {$$=$1}
	| ROLLBACK {$$=$1}
	| DECIMAL_P {$$=$1}
	| EXECUTE {$$=$1}
	| VACUUM {$$=$1}
	| DELIMITERS {$$=$1}
	| SELECT {$$=$1}
	| ASYMMETRIC {$$=$1}
	| STDOUT {$$=$1}
	| DEFAULT {$$=$1}
	| LAST_P {$$=$1}
	| ROLE {$$=$1}
	| DATABASE {$$=$1}
	| FALSE_P {$$=$1}
	| FORCE {$$=$1}
	| MINUTE_P {$$=$1}
	| BIT {$$=$1}
	| GROUP {$$=$1}
	| TCOLON {$$=$1}
	| TRUE_P {$$=$1}
	| VALUES {$$=$1}
	| BY {$$=$1}
	| NORMALIZED {$$=$1}
	| ORDINALITY {$$=$1}
	| FETCH {$$=$1}
	| USING {$$=$1}
	| TIMESTAMP {$$=$1}
	| ANALYZE {$$=$1}
	| FIRST_P {$$=$1}
	| INSERT {$$=$1}
	| ORDER {$$=$1}
	| LOCAL {$$=$1}
	| COPY {$$=$1}
	| NULLS_LA {$$=$1}
	| ENCODING {$$=$1}
	| HOUR_P {$$=$1}
	| CHAR_P {$$=$1}
	| VARCHAR {$$=$1}
	| STDIN {$$=$1}
	| DELIMITER {$$=$1}
	| OP {$$=$1}
	| WITH_LA {$$=$1}
	| CLUSTER {$$=$1}
	| WINDOW {$$=$1}
	| KEY {$$=$1}
	| NULL_P {$$=$1}
	| SMALLINT {$$=$1}
	| QUOTE {$$=$1}
	| TIME {$$=$1}
	| INTERVAL {$$=$1}
	| INT_P {$$=$1}
	| CREATE {$$=$1}
	| UPDATE {$$=$1}
	| OUTER_P {$$=$1}
	| REFERENCES {$$=$1}
	| RETURNING {$$=$1}
	| BEGIN {$$=$1}
	| PROGRAM {$$=$1}
	| FREEZE {$$=$1}
	| DEC {$$=$1}
	| TRUNCATE {$$=$1}
	| RESET {$$=$1}
	| ROW {$$=$1}
	| FOR {$$=$1}
	| FOREIGN {$$=$1}
	| NATIONAL {$$=$1}
	| NULL {$$=$1}
	| FLOAT_P {$$=$1}
	| INDEX {$$=$1}
	| EXPLAIN {$$=$1}
	| PRIMARY {$$=$1}
	| LATERAL_P {$$=$1}
	| LIMIT {$$=$1}
	| ALL {$$=$1}
	| SETOF {$$=$1}
	| BOOLEAN_P {$$=$1}
	| DOCUMENT_P {$$=$1}
	| INTEGER {$$=$1}
	| CHARACTER {$$=$1}
	| TO {$$=$1}
	| ARRAY {$$=$1}
	| WHERE {$$=$1}
	| HEADER_P {$$=$1}
	| YEAR_P {$$=$1}
	| VARYING {$$=$1}
	| DESC {$$=$1}
	| FROM {$$=$1}
	| ONLY {$$=$1}
	| DOUBLE_P {$$=$1}
	| ASC {$$=$1}
	| AS {$$=$1}
	| TABLE {$$=$1}
	| TSEMICOLON {$$=$1}
	| BIGINT {$$=$1}
	| REAL {$$=$1}
	| NCHAR {$$=$1}
	| DELETE {$$=$1}
	| COMMIT {$$=$1}
	| DROP {$$=$1}
	| BINARY {$$=$1}
	| DISTINCT {$$=$1}

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
    } | execute_stmt {
		setParseTree(yylex, $1)
    } | varset_stmt {
		setParseTree(yylex, $1)
    } | prepare_stmt {
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


c_expr:		AexprConst {
				$$ = $1
			}
			| ColRef {
				$$ = $1
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

execute_stmt:
    EXECUTE any_id /* todo params */ anything
    {
        $$ = &Execute {
            Id: $2,
        }
    }

prepare_stmt:
    PREPARE any_id /* todo query */  anything
    {
        $$ = &Execute {
            Id: $2,
        }
    }

routable_statement:
    select_stmt  semicolon_opt {
        $$ = $1
    } | insert_stmt semicolon_opt {
        $$ = $1
    } | update_stmt semicolon_opt {
        $$ = $1
    } | delete_stmt semicolon_opt {
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
opt_window_clause:  /*empty*/ {} |WINDOW anything {}
opt_limit_clause: /*empty*/ {} | LIMIT anything {}
opt_offset_clause: /*empty*/ {} | OFFSET anything {}
opt_fetch_clause: /*empty*/ {} | FETCH anything {}
opt_for_clause:  /*empty*/ {} |FOR anything {}


/* Column identifier --- names that can be column, table, etc names.
 */
ColId:		IDENT									{ $$ = $1; }
			// | reserved_keyword						{ $$ = $1; }
			//| col_name_keyword						{ $$ = pstrdup($1); }
		;


/* Column label --- allowed labels in "AS" clauses.
 * This presently includes *all* Postgres keywords.
 */
ColLabel:	IDENT									{ $$ = $1; }
			// | unreserved_keyword					{ $$ = pstrdup($1); }
			// | col_name_keyword						{ $$ = pstrdup($1); }
			// | type_func_name_keyword				{ $$ = pstrdup($1); }
			// | reserved_keyword						{ $$ = pstrdup($1); }
		;

attr_name:	ColLabel								{ $$ = $1; };

file_name:	SCONST									{ $$ = $1; };

// indirection_el:
// 			TDOT attr_name
// 				{
// 					$$ = (Node *) makeString($2);
// 				}
			// | TDOT TMUL
			// 	{
			// 		$$ = (Node *) makeNode(A_Star);
			// 	}
			// | '[' a_expr ']'
			// 	{
			// 		A_Indices *ai = makeNode(A_Indices);

			// 		ai->is_slice = false;
			// 		ai->lidx = NULL;
			// 		ai->uidx = $2;
			// 		$$ = (Node *) ai;
			// 	}
			// | '[' opt_slice_bound ':' opt_slice_bound ']'
			// 	{
			// 		A_Indices *ai = makeNode(A_Indices);

			// 		ai->is_slice = true;
			// 		ai->lidx = $2;
			// 		ai->uidx = $4;
			// 		$$ = (Node *) ai;
			// 	}
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




/* https://www.postgresql.org/docs/current/sql-select.html */
select_stmt:
	SELECT target_list from_clause where_clause opt_group_by_clause opt_window_clause opt_sort_clause opt_limit_clause opt_sort_clause opt_offset_clause opt_fetch_clause opt_for_clause 
	{
		$$ = &Select{
            FromClause: $3,
            Where: $4,
			TargetList: $2,
        }
    } | /*simple select */ SELECT target_list {
		$$ = &Select{
			TargetList: $2,
        }
    }


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


/* https://www.postgresql.org/docs/current/sql-insert.html */
insert_stmt: 
    /* consider only first tuple from values */
    INSERT INTO relation_expr insert_col_refs VALUES insert_tuples anything {
        $$ = &Insert{
            TableRef: $3,
            Columns: $4,
            Values: $6,
        }
    } | INSERT INTO relation_expr opt_insert_col_refs select_stmt {
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
		

update_stmt:
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


delete_stmt:
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

opt_with:	WITH_LA
			| /*EMPTY*/
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


PreparableStmt:
			select_stmt
			| insert_stmt
			| update_stmt
			| delete_stmt
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


%%