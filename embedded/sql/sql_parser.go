// Code generated by goyacc -l -o sql_parser.go sql_grammar.y. DO NOT EDIT.
package sql

import __yyfmt__ "fmt"

func setResult(l yyLexer, stmts []SQLStmt) {
	l.(*lexer).result = stmts
}

type yySymType struct {
	yys      int
	stmts    []SQLStmt
	stmt     SQLStmt
	colsSpec []*ColSpec
	colSpec  *ColSpec
	cols     []*ColSelector
	rows     []*RowSpec
	row      *RowSpec
	values   []ValueExp
	value    ValueExp
	id       string
	number   uint64
	str      string
	boolean  bool
	blob     []byte
	sqlType  SQLValueType
	aggFn    AggregateFn
	ids      []string
	col      *ColSelector
	sel      Selector
	sels     []Selector
	distinct bool
	ds       DataSource
	tableRef *TableRef
	joins    []*JoinSpec
	join     *JoinSpec
	joinType JoinType
	boolExp  ValueExp
	err      error
	ordcols  []*OrdCol
	opt_ord  Comparison
	logicOp  LogicOperator
	cmpOp    CmpOperator
}

const CREATE = 57346
const USE = 57347
const DATABASE = 57348
const SNAPSHOT = 57349
const SINCE = 57350
const UP = 57351
const TO = 57352
const TABLE = 57353
const INDEX = 57354
const ON = 57355
const ALTER = 57356
const ADD = 57357
const COLUMN = 57358
const PRIMARY = 57359
const KEY = 57360
const BEGIN = 57361
const TRANSACTION = 57362
const COMMIT = 57363
const UPSERT = 57364
const INTO = 57365
const VALUES = 57366
const SELECT = 57367
const DISTINCT = 57368
const FROM = 57369
const BEFORE = 57370
const TX = 57371
const JOIN = 57372
const HAVING = 57373
const WHERE = 57374
const GROUP = 57375
const BY = 57376
const LIMIT = 57377
const ORDER = 57378
const ASC = 57379
const DESC = 57380
const AS = 57381
const NOT = 57382
const LIKE = 57383
const EXISTS = 57384
const NULL = 57385
const JOINTYPE = 57386
const LOP = 57387
const CMPOP = 57388
const IDENTIFIER = 57389
const TYPE = 57390
const NUMBER = 57391
const VARCHAR = 57392
const BOOLEAN = 57393
const BLOB = 57394
const AGGREGATE_FUNC = 57395
const ERROR = 57396
const STMT_SEPARATOR = 57397

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"CREATE",
	"USE",
	"DATABASE",
	"SNAPSHOT",
	"SINCE",
	"UP",
	"TO",
	"TABLE",
	"INDEX",
	"ON",
	"ALTER",
	"ADD",
	"COLUMN",
	"PRIMARY",
	"KEY",
	"BEGIN",
	"TRANSACTION",
	"COMMIT",
	"UPSERT",
	"INTO",
	"VALUES",
	"SELECT",
	"DISTINCT",
	"FROM",
	"BEFORE",
	"TX",
	"JOIN",
	"HAVING",
	"WHERE",
	"GROUP",
	"BY",
	"LIMIT",
	"ORDER",
	"ASC",
	"DESC",
	"AS",
	"NOT",
	"LIKE",
	"EXISTS",
	"NULL",
	"JOINTYPE",
	"LOP",
	"CMPOP",
	"IDENTIFIER",
	"TYPE",
	"NUMBER",
	"VARCHAR",
	"BOOLEAN",
	"BLOB",
	"AGGREGATE_FUNC",
	"ERROR",
	"','",
	"'.'",
	"STMT_SEPARATOR",
	"'('",
	"')'",
	"'@'",
	"'*'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 208

var yyAct = [...]int{

	180, 34, 51, 115, 132, 6, 117, 65, 74, 57,
	85, 118, 66, 120, 127, 167, 7, 47, 125, 166,
	121, 122, 123, 124, 35, 36, 127, 137, 138, 119,
	151, 126, 121, 122, 123, 124, 100, 3, 172, 70,
	99, 160, 44, 126, 164, 162, 146, 54, 48, 143,
	45, 110, 106, 92, 91, 71, 143, 133, 142, 76,
	61, 67, 55, 53, 18, 93, 62, 54, 49, 179,
	171, 90, 148, 89, 94, 159, 175, 36, 116, 108,
	88, 48, 83, 35, 97, 2, 78, 95, 98, 16,
	36, 17, 19, 9, 137, 138, 50, 103, 105, 147,
	33, 144, 112, 109, 30, 107, 31, 75, 129, 75,
	96, 82, 128, 81, 72, 45, 69, 56, 45, 43,
	40, 38, 140, 141, 37, 138, 87, 139, 52, 68,
	181, 182, 153, 169, 64, 170, 136, 114, 156, 154,
	150, 157, 158, 102, 135, 4, 104, 77, 161, 163,
	59, 58, 22, 165, 9, 111, 29, 63, 20, 12,
	13, 130, 79, 60, 145, 12, 13, 39, 28, 14,
	42, 174, 177, 178, 173, 14, 46, 15, 26, 27,
	8, 183, 23, 15, 184, 152, 9, 24, 25, 176,
	168, 113, 134, 101, 86, 84, 41, 21, 32, 149,
	131, 155, 80, 73, 11, 10, 5, 1,
}
var yyPact = [...]int{

	-20, -1000, 161, -20, -1000, 7, -20, -1000, 138, 126,
	-1000, -1000, 176, 172, 157, 133, -1000, -1000, -20, -1000,
	-20, 30, -1000, 77, 74, 154, 73, 162, 72, 71,
	161, 155, 41, 89, -1000, 5, 11, -1000, 4, 70,
	-1000, 123, 121, 148, 2, 10, -1000, 136, -20, 3,
	30, -1000, 69, -22, 67, 60, 1, -1000, 118, 37,
	146, 66, 64, -1000, 155, 82, -1000, 68, 89, -1000,
	-5, -6, 9, 19, -1000, 39, 63, 35, -1000, 60,
	-19, -1000, -1000, -1000, 111, -1000, 82, 116, 123, -7,
	-1000, -1000, -1000, 58, 62, -1000, -8, -1000, -1000, 131,
	55, 104, -29, -1000, 3, 89, -1000, -1000, 143, -1000,
	-1000, -1, -1000, 113, 102, 49, 86, -1000, -29, -29,
	0, -1000, -1000, -1000, -1000, -9, 54, -1000, 151, -13,
	52, 17, -1000, -17, 96, -29, 43, -29, -29, 25,
	79, -18, 129, -14, -1000, -29, -1000, -15, -1, -40,
	-1000, -2, 98, 101, 49, 15, -1000, 79, -1000, -1000,
	-1000, -21, -1000, 49, -1000, -1000, -1000, -17, 89, 27,
	43, 43, -1000, -1000, -1000, -1000, 14, 93, -1000, 43,
	-1000, -1000, -1000, 93, -1000,
}
var yyPgo = [...]int{

	0, 207, 145, 17, 206, 16, 205, 204, 5, 203,
	8, 202, 201, 200, 4, 199, 6, 78, 198, 1,
	197, 7, 12, 196, 9, 195, 10, 194, 3, 193,
	192, 191, 190, 2, 189, 185, 0, 85,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 37, 37, 4, 4, 5,
	5, 3, 3, 6, 6, 6, 6, 6, 6, 23,
	23, 7, 13, 13, 14, 11, 11, 12, 12, 15,
	15, 16, 16, 16, 16, 16, 16, 16, 9, 9,
	10, 8, 20, 20, 18, 18, 17, 17, 17, 19,
	19, 19, 21, 21, 21, 22, 22, 24, 24, 25,
	25, 26, 26, 27, 29, 29, 31, 31, 30, 30,
	32, 32, 35, 35, 34, 34, 36, 36, 36, 33,
	33, 28, 28, 28, 28, 28, 28, 28, 28,
}
var yyR2 = [...]int{

	0, 2, 2, 2, 4, 0, 2, 1, 5, 1,
	1, 2, 3, 3, 3, 4, 10, 7, 6, 0,
	3, 8, 1, 3, 3, 1, 3, 1, 3, 1,
	3, 1, 1, 1, 1, 3, 2, 1, 1, 3,
	2, 12, 0, 1, 2, 4, 1, 4, 4, 1,
	3, 5, 1, 5, 3, 1, 3, 0, 3, 0,
	1, 1, 2, 5, 0, 2, 0, 3, 0, 2,
	0, 2, 0, 3, 2, 4, 0, 1, 1, 0,
	2, 1, 1, 2, 3, 3, 3, 3, 4,
}
var yyChk = [...]int{

	-1000, -1, -37, 57, -2, -4, -8, -5, 19, 25,
	-6, -7, 4, 5, 14, 22, -37, -37, 57, -37,
	20, -20, 26, 6, 11, 12, 6, 7, 11, 23,
	-37, -37, -18, -17, -19, 53, 47, 47, 47, 13,
	47, -23, 8, 47, -22, 47, -2, -3, -5, 27,
	55, -33, 39, 58, 56, 58, 47, -24, 28, 29,
	15, 58, 56, 21, -37, -21, -22, 58, -17, 47,
	61, -19, 47, -9, -10, 47, 58, 29, 49, 16,
	-11, 47, 47, -3, -25, -26, -27, 44, -22, -8,
	-33, 59, 59, 56, 55, 48, 47, 49, -10, 59,
	55, -29, 32, -26, 30, -24, 59, 47, 17, -10,
	59, 24, 47, -31, 33, -28, -17, -16, 40, 58,
	42, 49, 50, 51, 52, 47, 60, 43, -21, -33,
	18, -13, -14, 58, -30, 31, 34, 45, 46, 41,
	-28, -28, 58, 58, 47, 13, 59, 47, 55, -15,
	-16, 47, -35, 36, -28, -12, -19, -28, -28, 50,
	59, -8, 59, -28, 59, -14, 59, 55, -32, 35,
	34, 55, 59, -16, -33, 49, -34, -19, -19, 55,
	-36, 37, 38, -19, -36,
}
var yyDef = [...]int{

	5, -2, 0, 5, 1, 5, 5, 7, 0, 42,
	9, 10, 0, 0, 0, 0, 6, 2, 5, 3,
	5, 0, 43, 0, 0, 0, 0, 19, 0, 0,
	6, 0, 0, 79, 46, 0, 49, 13, 0, 0,
	14, 57, 0, 0, 0, 55, 4, 0, 5, 0,
	0, 44, 0, 0, 0, 0, 0, 15, 0, 0,
	0, 0, 0, 8, 11, 59, 52, 0, 79, 80,
	0, 0, 50, 0, 38, 0, 0, 0, 20, 0,
	0, 25, 56, 12, 64, 60, 61, 0, 57, 0,
	45, 47, 48, 0, 0, 40, 0, 58, 18, 0,
	0, 66, 0, 62, 0, 79, 54, 51, 0, 39,
	17, 0, 26, 68, 0, 65, 81, 82, 0, 0,
	0, 31, 32, 33, 34, 49, 0, 37, 0, 0,
	0, 21, 22, 0, 72, 0, 0, 0, 0, 0,
	83, 0, 0, 0, 36, 0, 53, 0, 0, 0,
	29, 0, 70, 0, 69, 67, 27, 86, 87, 85,
	84, 0, 35, 63, 16, 23, 24, 0, 79, 0,
	0, 0, 88, 30, 41, 71, 73, 76, 28, 0,
	74, 77, 78, 76, 75,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	58, 59, 61, 3, 55, 3, 56, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 60,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 57,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmts = yyDollar[2].stmts
			setResult(yylex, yyDollar[2].stmts)
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmts = []SQLStmt{yyDollar[1].stmt}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmts = []SQLStmt{yyDollar[1].stmt}
		}
	case 4:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.stmts = append([]SQLStmt{yyDollar[1].stmt}, yyDollar[4].stmts...)
		}
	case 5:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
		}
	case 8:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.stmt = &TxStmt{stmts: yyDollar[4].stmts}
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmts = []SQLStmt{yyDollar[1].stmt}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmts = append([]SQLStmt{yyDollar[1].stmt}, yyDollar[3].stmts...)
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmt = &CreateDatabaseStmt{DB: yyDollar[3].id}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmt = &UseDatabaseStmt{DB: yyDollar[3].id}
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.stmt = &UseSnapshotStmt{sinceTx: yyDollar[3].number, asBefore: yyDollar[4].number}
		}
	case 16:
		yyDollar = yyS[yypt-10 : yypt+1]
		{
			yyVAL.stmt = &CreateTableStmt{table: yyDollar[3].id, colsSpec: yyDollar[5].colsSpec, pk: yyDollar[9].id}
		}
	case 17:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.stmt = &CreateIndexStmt{table: yyDollar[4].id, col: yyDollar[6].id}
		}
	case 18:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.stmt = &AddColumnStmt{table: yyDollar[3].id, colSpec: yyDollar[6].colSpec}
		}
	case 19:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.number = 0
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.number = yyDollar[3].number
		}
	case 21:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt = &UpsertIntoStmt{tableRef: yyDollar[3].tableRef, cols: yyDollar[5].ids, rows: yyDollar[8].rows}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.rows = []*RowSpec{yyDollar[1].row}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.rows = append(yyDollar[1].rows, yyDollar[3].row)
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.row = &RowSpec{Values: yyDollar[2].values}
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ids = []string{yyDollar[1].id}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.ids = append(yyDollar[1].ids, yyDollar[3].id)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.cols = []*ColSelector{yyDollar[1].col}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.cols = append(yyDollar[1].cols, yyDollar[3].col)
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.values = []ValueExp{yyDollar[1].value}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].value)
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Number{val: yyDollar[1].number}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Varchar{val: yyDollar[1].str}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Bool{val: yyDollar[1].boolean}
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Blob{val: yyDollar[1].blob}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.value = &SysFn{fn: yyDollar[1].id}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.value = &Param{id: yyDollar[2].id}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &NullValue{}
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.colsSpec = []*ColSpec{yyDollar[1].colSpec}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.colsSpec = append(yyDollar[1].colsSpec, yyDollar[3].colSpec)
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.colSpec = &ColSpec{colName: yyDollar[1].id, colType: yyDollar[2].sqlType}
		}
	case 41:
		yyDollar = yyS[yypt-12 : yypt+1]
		{
			yyVAL.stmt = &SelectStmt{
				distinct:  yyDollar[2].distinct,
				selectors: yyDollar[3].sels,
				ds:        yyDollar[5].ds,
				joins:     yyDollar[6].joins,
				where:     yyDollar[7].boolExp,
				groupBy:   yyDollar[8].cols,
				having:    yyDollar[9].boolExp,
				orderBy:   yyDollar[10].ordcols,
				limit:     yyDollar[11].number,
				as:        yyDollar[12].id,
			}
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.distinct = false
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.distinct = true
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyDollar[1].sel.setAlias(yyDollar[2].id)
			yyVAL.sels = []Selector{yyDollar[1].sel}
		}
	case 45:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyDollar[3].sel.setAlias(yyDollar[4].id)
			yyVAL.sels = append(yyDollar[1].sels, yyDollar[3].sel)
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.sel = yyDollar[1].col
		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.sel = &AggColSelector{aggFn: yyDollar[1].aggFn, col: "*"}
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.sel = &AggColSelector{aggFn: yyDollar[1].aggFn, db: yyDollar[3].col.db, table: yyDollar[3].col.table, col: yyDollar[3].col.col}
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.col = &ColSelector{col: yyDollar[1].id}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.col = &ColSelector{table: yyDollar[1].id, col: yyDollar[3].id}
		}
	case 51:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.col = &ColSelector{db: yyDollar[1].id, table: yyDollar[3].id, col: yyDollar[5].id}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ds = yyDollar[1].tableRef
		}
	case 53:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyDollar[2].tableRef.asBefore = yyDollar[3].number
			yyDollar[2].tableRef.as = yyDollar[4].id
			yyVAL.ds = yyDollar[2].tableRef
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.ds = yyDollar[2].stmt.(*SelectStmt)
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.tableRef = &TableRef{table: yyDollar[1].id}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.tableRef = &TableRef{db: yyDollar[1].id, table: yyDollar[3].id}
		}
	case 57:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.number = 0
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.number = yyDollar[3].number
		}
	case 59:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.joins = nil
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.joins = yyDollar[1].joins
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.joins = []*JoinSpec{yyDollar[1].join}
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.joins = append([]*JoinSpec{yyDollar[1].join}, yyDollar[2].joins...)
		}
	case 63:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.join = &JoinSpec{joinType: yyDollar[1].joinType, ds: yyDollar[3].ds, cond: yyDollar[5].boolExp}
		}
	case 64:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolExp = nil
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.boolExp = yyDollar[2].boolExp
		}
	case 66:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.cols = nil
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.cols = yyDollar[3].cols
		}
	case 68:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolExp = nil
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.boolExp = yyDollar[2].boolExp
		}
	case 70:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.number = 0
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.number = yyDollar[2].number
		}
	case 72:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.ordcols = nil
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.ordcols = yyDollar[3].ordcols
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.ordcols = []*OrdCol{{sel: yyDollar[1].col, cmp: yyDollar[2].opt_ord}}
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.ordcols = append(yyDollar[1].ordcols, &OrdCol{sel: yyDollar[3].col, cmp: yyDollar[4].opt_ord})
		}
	case 76:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.opt_ord = GreaterOrEqualTo
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.opt_ord = GreaterOrEqualTo
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.opt_ord = LowerOrEqualTo
		}
	case 79:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.id = ""
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.id = yyDollar[2].id
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.boolExp = yyDollar[1].sel
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.boolExp = yyDollar[1].value
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.boolExp = &NotBoolExp{exp: yyDollar[2].boolExp}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.boolExp = yyDollar[2].boolExp
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.boolExp = &LikeBoolExp{sel: yyDollar[1].sel, pattern: yyDollar[3].str}
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.boolExp = &BinBoolExp{op: yyDollar[2].logicOp, left: yyDollar[1].boolExp, right: yyDollar[3].boolExp}
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.boolExp = &CmpBoolExp{op: yyDollar[2].cmpOp, left: yyDollar[1].boolExp, right: yyDollar[3].boolExp}
		}
	case 88:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.boolExp = &ExistsBoolExp{q: (yyDollar[3].stmt).(*SelectStmt)}
		}
	}
	goto yystack /* stack new state and value */
}