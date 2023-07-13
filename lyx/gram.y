
%{
package routerparser

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

	statement              Statement
    from_list              []FromClauseNode
    from                   FromClauseNode
    tableelt               []TableElt
    tableref               FromClauseNode

    selectStmt              *Select

    colref                 ColumnRef


    aexpr                  AExpr

    aexprList               []AExpr
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
//%type <val> expr number


// CMDS
%type <statement> command
%type <statement> routable_statement

// same for terminals
%token <str> STRING

%type<str> any_id any_val table_name

%type<aexpr> where_clause

%type<from_list> from_clause from_list


%type<from> table_ref 
%type<tableref> relation_expr joined_table

%type<str> target_list

%type<str> func_arg_list func_arg_expr

%type<aexpr> a_expr c_expr b_expr

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

/* ';' != */
%token<str> TSEMICOLON TNOT_EQUALS

/* '.' */
%token<str> TDOT

%token<str> NULL_P ISNULL ON JOIN

/* + - * / */
%token<str> TPLUS TMINUS TMUL TDIV

/* % ^ < <= >=  */
%token<str>  TGREATER TGREATER_EQUALS TLESS TLESS_EQUALS TMOD TPOW


%token<str> FALSE_P TRUE_P 

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

/* IS NOT NULL */
%token<str> IS NOT NULL DISTINCT DEFAULT

/* NORMALIZED */
%token<str> NORMALIZED

/* NOTNULL */
%token<str> NOTNULL

/* DOCUMENT_P */
%token<str> DOCUMENT_P

/* */
%token<str> TO STDOUT

%token<str> AS


%type<str> reserved_keyword


/**/

%type<statement> execute_stmt
%type<selectStmt> select_stmt
%type<statement> insert_stmt
%type<statement> update_stmt
%type<statement> delete_stmt
%type<statement> varset_stmt
%type<statement> prepare_stmt
%type<statement> begin_stmt 
%type<statement> commit_stmt
%type<statement> rollback_stmt
%type<statement> prepare_stmt
%type<statement> create_stmt alter_stmt
%type<statement> vacuum_stmt cluster_stmt analyze_stmt
%type<statement> truncate_stmt drop_stmt
%type<str> semicolon_opt

%type<statement> copy_stmt

%type<colref> ColRef qualColRef

%type<strlist> comma_separated_col_refs insert_col_refs

%type<aexprList> insert_comma_separated_tuples insert_tuples

%type<str> opt_insert_col_refs
%type<str> any_tok 

%type<strlist> opt_using delete_comma_separated_using_refs
%type<strlist> set_clause_list

%type<tableelt> create_stmt_coldefs

%type<str> opt_only opt_alias_clause alias_clause

%type<str> anything set_clause

%type<str> operator


/* unroutable query parts */
%type<str> opt_group_by_clause opt_window_clause opt_order_by_clause opt_limit_clause opt_offset_clause opt_fetch_clause opt_for_clause 

%start any_command

%%


any_command:
    command semicolon_opt

semicolon_opt:
/*empty*/ {}
| TSEMICOLON {
    
}


reserved_keyword:
    SELECT {$$=$1}
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
    | TSEMICOLON {$$=$1}
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
    | TLESS {$$=$1}
    | TLESS_EQUALS {$$=$1}
    | TGREATER {$$=$1}
    | TGREATER_EQUALS {$$=$1}
    | TNOT_EQUALS {$$=$1}
    | OP {$$=$1}

any_tok:
    reserved_keyword {$$=$1} | STRING {$$=$1}


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
	STRING
	{
		$$ = string($1)
	}

any_val:
    reserved_keyword {
        $$ = $1
    }
    | STRING
	{
		$$ = string($1)
	}

table_name:
	STRING
	{
		$$ = string($1)
	}

operator:
    STRING {
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

func_application: any_id TOPENBR TCLOSEBR {} |  any_id TOPENBR func_arg_expr TCLOSEBR {}

func_arg_expr: a_expr {}

func_arg_list:  func_arg_expr
				{
					$$ = ""
				}
			| func_arg_list TCOMMA func_arg_expr
				{
					$$ = ""
				}
		;


func_expr: func_application

c_expr:		any_tok	{
                $$ = &AExprLeaf{
                    Value: $1,
                }
            } 
            | TOPENBR a_expr TCLOSEBR {
                $$ = $2
            }
            | func_expr {}

a_expr:		
            c_expr									{ $$ = $1; }
			// | a_expr TYPECAST Typename
			// 		{ $$ = makeTypeCast($1, $3, @2); }
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
			 TPLUS a_expr					//%prec UMINUS
				{ $$ = $2 }
			| TMINUS a_expr					//%prec UMINUS
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

			| a_expr OP a_expr				//%prec Op
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
                { $$ = 1 }
			| a_expr ISNULL
                 { $$ = 1 }
			| a_expr IS NOT NULL_P						%prec IS
                    { $$ = 1 }
			| a_expr NOTNULL
				 { $$ = 1 }
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
				 { $$ = 1 }
			| a_expr IS NOT TRUE_P						%prec IS
				 { $$ = 1 }
			| a_expr IS FALSE_P							%prec IS
				 { $$ = 1 }
			| a_expr IS NOT FALSE_P						%prec IS
				 { $$ = 1 }
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
			// | a_expr BETWEEN opt_asymmetric b_expr AND a_expr		%prec BETWEEN
			// 	{
			// 		$$ = (Node *) makeSimpleA_Expr(AEXPR_BETWEEN,
			// 									   "BETWEEN",
			// 									   $1,
			// 									   (Node *) list_make2($4, $6),
			// 									   @2);
			// 	}
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
    any_val TDOT any_val {
        $$ = ColumnRef{
            ColName: $3,
            TableAlias: $1, 
        }
    }

ColRef:
    any_val {
        $$ = ColumnRef{
            ColName: $1,
        }
    } |  qualColRef {
        $$ = $1
    }



b_expr:
            TPLUS a_expr					//%prec UMINUS
				{ /* result not matter */ }
			| TMINUS a_expr					//%prec UMINUS
					{ /* result not matter */ }
			| a_expr TPLUS a_expr
					{ /* result not matter */ }
			| a_expr TMINUS a_expr
					{ /* result not matter */ }
			| a_expr TMUL a_expr
					{ /* result not matter */ }
			| a_expr TDIV a_expr
					{ /* result not matter */ }
			| a_expr TMOD a_expr
				    { /* result not matter */ }
			| a_expr TPOW a_expr
				    { /* result not matter */ }
			| a_expr TLESS a_expr
				    { /* result not matter */ }
			| a_expr TGREATER a_expr
				    { /* result not matter */ }
			| a_expr TEQ a_expr
				    { /* result not matter */ }
			| a_expr TLESS_EQUALS a_expr
				    { /* result not matter */ }
			| a_expr TGREATER_EQUALS a_expr
				    { /* result not matter */ }
			| a_expr TNOT_EQUALS a_expr
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


opt_target_list: target_list						{  }
			| /* EMPTY */							{  }
		;

target_list:
			target_el								{ }
			| target_list TCOMMA target_el				{  }
		;

target_el:	a_expr AS any_id
				{

				}
			| a_expr any_id
				{

				}
			| a_expr
				{

				}
			| TMUL /* '*' */
				{

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
    any_id any_val opt_pk {
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


opt_group_by_clause: /*empty*/ {} | GROUP BY anything {}
opt_window_clause:  /*empty*/ {} |WINDOW anything {}
opt_order_by_clause:  /*empty*/ {} |ORDER BY anything {}
opt_limit_clause: /*empty*/ {} | LIMIT anything {}
opt_offset_clause: /*empty*/ {} | OFFSET anything {}
opt_fetch_clause: /*empty*/ {} | FETCH anything {}
opt_for_clause:  /*empty*/ {} |FOR anything {}




from_clause:
			FROM from_list							{ $$ = $2; }
			| /*EMPTY*/								{ }
		;

from_list:
			table_ref								{ $$ = []FromClauseNode{$1}; }
			| from_list TCOMMA table_ref				{ $$ = append($1, $3); }
		;

relation_expr:
			any_id {
                $$ = &RangeVar {
                    SchemaName: "",
                    RelationName: $1,
                    Alias: "",
                }
            }
            | any_id TDOT any_id {
                $$ = &RangeVar {
                    SchemaName: $1,
                    RelationName: $3,
                    Alias: "",
                }
            }
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
 * table_ref is where an alias clause can be attached.
 */
table_ref:	relation_expr opt_alias_clause
				{
					$1.SetAlias($2)
					$$ = $1;
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
	SELECT target_list from_clause where_clause opt_group_by_clause opt_window_clause opt_order_by_clause opt_limit_clause opt_offset_clause opt_fetch_clause opt_for_clause 
	{
		$$ = &Select{
            FromClause: $3,
            Where: $4,
        }
    } | /*simple select */ SELECT target_list {
		$$ = &Select{
        }
    }


comma_separated_col_refs:
    any_id { $$ = []string{$1} }
    | any_id TCOMMA comma_separated_col_refs {
        $$ = append($3, $1)
    }

opt_comma_separated_col_refs:
    /*nothing*/{} | comma_separated_col_refs {}

insert_col_refs:
    TOPENBR comma_separated_col_refs TCLOSEBR {
        $$ = $2
    }


opt_insert_col_refs:
    /* nothing */ {} |
    insert_col_refs {}    



insert_comma_separated_tuples: 
    a_expr {
        $$ = []AExpr{$1}
    } | a_expr TCOMMA insert_comma_separated_tuples {
        $$ = append($3, $1)
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
            SubSelect: $5,
        }
    }

opt_only:
    /* nothing */ {}
    | ONLY {}

opt_upd_from:
    /* nothing */ {}
    | FROM table_name {}

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
    UPDATE opt_only relation_expr SET set_clause_list opt_upd_from where_clause opt_returning {
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

/* https://www.postgresql.org/docs/current/sql-copy.html */
copy_stmt: 
    COPY relation_expr opt_comma_separated_col_refs TO STDOUT where_clause {
        $$ = &Copy {
            TableRef: $2,
            Where: $6,
            IsFrom: false,
        }   
    }

%%

