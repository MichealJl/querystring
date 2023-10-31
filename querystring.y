%{
package querystring

import (
	"github.com/MichealJl/querystring/condition"
)
%}

%union {
	s 	string
	n 	int
	q 	condition.Condition
}


%left tNOT
%left tOR
%left tAND
%token tSTRING tPHRASE tNUMBER
%token tTO tCOLON
%token tLEFTBRACKET tRIGHTBRACKET tLEFTRANGE tRIGHTRANGE

%type <s>                tSTRING
%type <s>                tPHRASE
%type <s>                tNUMBER
%type <q>                baseCondition searchRange searchBase searchLogicPart searchParts
%%

// 这里输入 开始
input:
searchParts {
	yylex.(*lexerWrapper).query = $1
};

// 这里要递归拼接条件
searchParts:
searchLogicPart {
	$$ = $1
};

// 递归这一层
searchLogicPart:
searchBase {
	$$ = $1
}
|
tLEFTBRACKET searchLogicPart tRIGHTBRACKET {
	$$ = $2
}
|
searchLogicPart tAND searchLogicPart {
	$$ = condition.NewAndCondition($1, $3)
}
|
searchLogicPart tOR searchLogicPart {
	$$ = condition.NewOrCondition($1, $3)
}
|
tNOT searchBase {
	$$ = condition.NewNotCondition(nil, $2)
}
|
searchLogicPart tNOT searchLogicPart {
	$$ = condition.NewNotCondition($1, $3)
};

searchBase:
baseCondition {
	$$ = $1
}
|
tSTRING tCOLON tLEFTBRACKET searchLogicPart tRIGHTBRACKET {
	q := $4
	q.SetField($1)
	$$ = q
}
|
tSTRING tCOLON searchRange {
	q := $3
	q.SetField($1)
	$$ = q
};

// 范围类型
searchRange:
tLEFTRANGE tSTRING tRIGHTRANGE {
	min := $2
	$$ = condition.NewRangeCondition(&min, nil)
}
|
tLEFTRANGE tTO tSTRING tRIGHTRANGE {
	max := $3
	$$ = condition.NewRangeCondition(nil, &max)
}
|
tLEFTRANGE tSTRING tTO tSTRING tRIGHTRANGE {
	min := $2
	max := $4
	$$ = condition.NewRangeCondition(&min, &max)
}
|tLEFTRANGE tNUMBER tRIGHTRANGE {
 	min := $2
 	$$ = condition.NewRangeCondition(&min, nil)
}
|
tLEFTRANGE tTO tNUMBER tRIGHTRANGE {
 	max := $3
 	$$ = condition.NewRangeCondition(nil, &max)
}
|
tLEFTRANGE tNUMBER tTO tNUMBER tRIGHTRANGE {
 	min := $2
 	max := $4
 	$$ = condition.NewRangeCondition(&min, &max)
};

// 基础类型
baseCondition:
tSTRING {
	$$ = condition.NewStringCondition($1)
}
|
tSTRING tSTRING {
	$$ = condition.NewOrCondition(condition.NewStringCondition($1), condition.NewStringCondition($2))
}
|
tPHRASE {
	$$ = condition.NewPhraseCondition($1)
}
|
tPHRASE tPHRASE {
	$$ = condition.NewOrCondition(condition.NewPhraseCondition($1), condition.NewPhraseCondition($2))
}
|
tNUMBER {
	$$ = condition.NewNumberCondition($1)
};