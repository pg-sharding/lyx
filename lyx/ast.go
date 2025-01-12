package lyx

// Node represents any query.
type Node interface {
	iNode()
}

type ColumnRef struct {
	TableAlias string
	ColName    string
}

func (*ColumnRef) iNode() {}

// /*
//  * RangeVar - range variable, used in FROM clauses
//  *
//  * Also used to represent table names in utility statements; there, the alias
//  * field is not used, and inh tells whether to apply the operation
//  * recursively to child tables.  In some contexts it is also useful to carry
//  * a TEMP table indication here.
//  */
//  typedef struct RangeVar
//  {
// 	 NodeTag		type;
// 	 char	   *catalogname;	/* the catalog (database) name, or NULL */
// 	 char	   *schemaname;		/* the schema name, or NULL */
// 	 char	   *relname;		/* the relation/sequence name */
// 	 bool		inh;			/* expand rel by inheritance? recursively act
// 								  * on children? */
// 	 char		relpersistence; /* see RELPERSISTENCE_* in pg_class.h */
// 	 Alias	   *alias;			/* table alias & optional column aliases */
// 	 int			location;		/* token location, or -1 if unknown */
//  } RangeVar;

type FromClauseNode interface {
	SetAlias(string)
}

type RangeVar struct {
	SchemaName   string
	RelationName string
	Alias        string
}

func (r *RangeVar) SetAlias(s string) {
	r.Alias = s
}

type SubLink struct {
	SubSelect Node
}

func (r *RangeVar) iNode() {}

// /*----------
//  * JoinExpr - for SQL JOIN expressions
//  *
//  * isNatural, usingClause, and quals are interdependent.  The user can write
//  * only one of NATURAL, USING(), or ON() (this is enforced by the grammar).
//  * If he writes NATURAL then parse analysis generates the equivalent USING()
//  * list, and from that fills in "quals" with the right equality comparisons.
//  * If he writes USING() then "quals" is filled with equality comparisons.
//  * If he writes ON() then only "quals" is set.  Note that NATURAL/USING
//  * are not equivalent to ON() since they also affect the output column list.
//  *
//  * alias is an Alias node representing the AS alias-clause attached to the
//  * join expression, or NULL if no clause.  NB: presence or absence of the
//  * alias has a critical impact on semantics, because a join with an alias
//  * restricts visibility of the tables/columns inside it.
//  *
//  * join_using_alias is an Alias node representing the join correlation
//  * name that SQL:2016 and later allow to be attached to JOIN/USING.
//  * Its column alias list includes only the common column names from USING,
//  * and it does not restrict visibility of the join's input tables.
//  *
//  * During parse analysis, an RTE is created for the Join, and its index
//  * is filled into rtindex.  This RTE is present mainly so that Vars can
//  * be created that refer to the outputs of the join.  The planner sometimes
//  * generates JoinExprs internally; these can have rtindex = 0 if there are
//  * no join alias variables referencing such joins.
//  *----------
//  */
//  typedef struct JoinExpr
//  {
// 	 NodeTag		type;
// 	 JoinType	jointype;		/* type of join */
// 	 bool		isNatural;		/* Natural join? Will need to shape table */
// 	 Node	   *larg;			/* left subtree */
// 	 Node	   *rarg;			/* right subtree */
// 	 List	   *usingClause;	/* USING clause, if any (list of String) */
// 	 Alias	   *join_using_alias;	/* alias attached to USING clause, if any */
// 	 Node	   *quals;			/* qualifiers on join, if any */
// 	 Alias	   *alias;			/* user-written alias clause, if any */
// 	 int			rtindex;		/* RT index assigned for join, or 0 */
//  } JoinExpr;

type JoinExpr struct {
	Larg FromClauseNode
	Rarg FromClauseNode

	Alias string
}

type AExprEmpty struct {
}

func (*AExprEmpty) iNode() {
}

type AExprList struct {
	List []Node
}

func (*AExprList) iNode() {
}

type AExprSConst struct {
	Value string
}

type AExprIConst struct {
	Value int
}

type AExprBConst struct {
	Value bool
}

// NULL
type AExprNConst struct {
	Value bool
}

func (*AExprIConst) iNode() {
}
func (*AExprSConst) iNode() {
}

func (*AExprBConst) iNode() {
}

func (*AExprNConst) iNode() {
}

type AExprOp struct {
	Left  Node
	Right Node

	Op string
}

func (*AExprOp) iNode() {
}

func (r *JoinExpr) SetAlias(s string) {
	r.Alias = s
}

type CommonTableExpr struct {
	Name     string
	SubQuery Node
}

type Select struct {
	FromClause []FromClauseNode
	WithClause []*CommonTableExpr
	Where      Node
	TargetList []Node

	// Used in set operations
	Op   SetOperation
	LArg Node
	RArg Node
}

type ValueClause struct {
	Values [][]Node
}

type Insert struct {
	TableRef FromClauseNode
	Columns  []string

	SubSelect Node
}

type Delete struct {
	TableRef FromClauseNode
	Where    Node
}

type Update struct {
	TableRef FromClauseNode
	Where    Node
}

type Explain struct {
	Stmt Node
}

type Execute struct {
	Id string
}

type Prepare struct {
	Id string
}

type TableElt struct {
	ColName string
	ColType string
}

type CreateTable struct {
	TableRv     FromClauseNode
	TableElts   []Node
	PartitionOf FromClauseNode
}

type CreateSchema struct {
}

type Alter struct {
}

type Analyze struct {
}

type Cluster struct {
}

type Vacuum struct {
}

type Truncate struct {
}

type Drop struct {
}

type Index struct {
}

type CreateRole struct {
}

type CreateDatabase struct {
}

type VarValue struct {
	Value string
}

type VarType string

const (
	VarTypeSet      = VarType("SET")
	VarTypeReset    = VarType("RESET")
	VarTypeResetAll = VarType("RESET ALL")
)

type VariableSetStmt struct {
	Kind    VarType
	Session bool
	IsLocal bool
	Default bool
	TxMode  []TransactionModeItem
	Name    string
	Value   []string
}

type VariableShowStmt struct {
	Name string
}

type SetOperation string

const (
	SetOpUnion     = SetOperation("UNION")
	SetOpIntersect = SetOperation("INTERSECT")
	SetOpExcept    = SetOperation("EXCEPT")
)

func (*Explain) iNode()          {}
func (*Select) iNode()           {}
func (*ValueClause) iNode()      {}
func (*Execute) iNode()          {}
func (*Prepare) iNode()          {}
func (*CreateTable) iNode()      {}
func (*Alter) iNode()            {}
func (*Analyze) iNode()          {}
func (*Cluster) iNode()          {}
func (*Vacuum) iNode()           {}
func (*Drop) iNode()             {}
func (*Truncate) iNode()         {}
func (*Index) iNode()            {}
func (*CreateRole) iNode()       {}
func (*CreateDatabase) iNode()   {}
func (*Insert) iNode()           {}
func (*Delete) iNode()           {}
func (*Update) iNode()           {}
func (*VarValue) iNode()         {}
func (*VariableSetStmt) iNode()  {}
func (*VariableShowStmt) iNode() {}
func (*SubLink) iNode()          {}
func (*CommonTableExpr) iNode()  {}
func (*TableElt) iNode()         {}
func (*CreateSchema) iNode()     {}

type TransactionStmtType int

const (
	TRANS_STMT_START             = TransactionStmtType(iota)
	TRANS_STMT_BEGIN             = TransactionStmtType(iota)
	TRANS_STMT_COMMIT            = TransactionStmtType(iota)
	TRANS_STMT_ROLLBACK          = TransactionStmtType(iota)
	TRANS_STMT_RELEASE           = TransactionStmtType(iota)
	TRANS_STMT_ROLLBACK_TO       = TransactionStmtType(iota)
	TRANS_STMT_SAVEPOINT         = TransactionStmtType(iota)
	TRANS_STMT_PREPARE           = TransactionStmtType(iota)
	TRANS_STMT_COMMIT_PREPARED   = TransactionStmtType(iota)
	TRANS_STMT_ROLLBACK_PREPARED = TransactionStmtType(iota)
)

type TransactionStmt struct {
	Kind          TransactionStmtType
	Name          string
	SavepointName string
	Gid           string
	Options       []TransactionModeItem
}

type TransactionModeItem int

const (
	TransactionIsolation     = TransactionModeItem(iota)
	TransactionReadOnly      = TransactionModeItem(iota)
	TransactionReadWrite     = TransactionModeItem(iota)
	TransactionDeferrable    = TransactionModeItem(iota)
	TransactionNotDeferrable = TransactionModeItem(iota)
)

func (*TransactionStmt) iNode() {}

type EmptyQuery struct{}

func (*EmptyQuery) iNode() {}

type Option struct {
	Name string
	Arg  Node
}

func (*Option) iNode() {
}

type Copy struct {
	TableRef FromClauseNode
	Where    Node
	IsFrom   bool
	SubStmt  Node
	Columns  []string
	Options  []Node
}

func (*Copy) iNode() {}

type ParamRef struct {
	Number int
}

func (*ParamRef) iNode() {}

type PrepareStmt struct {
	Name      string
	Statement Node
}

func (*PrepareStmt) iNode() {}

type ExecuteStmt struct {
	Name string
}

func (*ExecuteStmt) iNode() {}

type DeallocateStmt struct {
	Name string
}

func (*DeallocateStmt) iNode() {}

type FuncApplication struct {
	Name string
	Args []Node
}

func (*FuncApplication) iNode() {}
