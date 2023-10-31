package querystring

import (
	"github.com/MichealJl/querystring/condition"
	"reflect"
	"testing"
)

func TestNumberCondition(t *testing.T) {
	cond, err := Parse(`total:(123)`)
	if err != nil {
		t.Error(err)
	}
	target := condition.NewNumberCondition("123")
	target.SetField("total")

	if !reflect.DeepEqual(cond, target) {
		t.Fatal("Parse failed")
	}

	t.Log("Parse success")
}

func TestStringCondition(t *testing.T) {
	cond, err := Parse(`title:(小米)`)
	if err != nil {
		t.Error(err)
	}
	target := condition.NewStringCondition("小米")
	target.SetField("title")

	if !reflect.DeepEqual(cond, target) {
		t.Fatal("Parse failed")
	}

	t.Log("Parse success")
}

func TestPhraseCondition(t *testing.T) {
	cond, err := Parse(`title:("小米")`)
	if err != nil {
		t.Error(err)
	}
	target := condition.NewPhraseCondition("小米")
	target.SetField("title")

	if !reflect.DeepEqual(cond, target) {
		t.Fatal("Parse failed")
	}

	t.Log("Parse success")
}

// 空格将被转换成OR查询
func TestSpaceCondition(t *testing.T) {
	cond, err := Parse(`title:("小米" "华为")`)
	if err != nil {
		t.Error(err)
	}
	cond2, err := Parse(`title:(小米 华为)`)
	if err != nil {
		t.Error(err)
	}
	target := condition.NewOrCondition(
		condition.NewPhraseCondition("小米").SetField("title"),
		condition.NewPhraseCondition("华为").SetField("title"),
	)
	target2 := condition.NewOrCondition(
		condition.NewStringCondition("小米").SetField("title"),
		condition.NewStringCondition("华为").SetField("title"),
	)

	if !reflect.DeepEqual(cond, target) {
		t.Fatal("Parse failed")
	}
	if !reflect.DeepEqual(cond2, target2) {
		t.Fatal("Parse failed")
	}

	t.Log("Parse success")
}

func TestRangeCondition(t *testing.T) {
	cond, err := Parse(`total:[0 TO 0.111] OR total:[3] OR total:[a TO z]`)
	if err != nil {
		t.Error(err)
	}
	first := condition.NewRangeCondition(pointer("0"), pointer("0.111")).SetField("total")
	second := condition.NewRangeCondition(pointer("3"), nil).SetField("total")
	third := condition.NewRangeCondition(pointer("a"), pointer("z")).SetField("total")
	target := condition.NewOrCondition(condition.NewOrCondition(first, second), third)

	if !reflect.DeepEqual(cond, target) {
		t.Fatal("Parse failed")
	}

	t.Log("Parse success")
}

func TestAndCondition(t *testing.T) {
	cond, err := Parse(`title:("小米" AND "华为")`)
	if err != nil {
		t.Error(err)
	}
	left := condition.NewPhraseCondition("小米").SetField("title")
	right := condition.NewPhraseCondition("华为").SetField("title")
	target := condition.NewAndCondition(left, right)

	if !reflect.DeepEqual(cond, target) {
		t.Fatal("Parse failed")
	}

	t.Log("Parse success")
}

func TestOrCondition(t *testing.T) {
	cond, err := Parse(`title:("小米" OR "华为")`)
	if err != nil {
		t.Error(err)
	}
	left := condition.NewPhraseCondition("小米").SetField("title")
	right := condition.NewPhraseCondition("华为").SetField("title")
	target := condition.NewOrCondition(left, right)

	if !reflect.DeepEqual(cond, target) {
		t.Fatal("Parse failed")
	}

	t.Log("Parse success")
}

func TestNotCondition(t *testing.T) {
	cond, err := Parse(`title:("小米" NOT "华为")`)
	if err != nil {
		t.Error(err)
	}
	cond2, err1 := Parse(`title:("小米" OR "华为") NOT type:("电脑")`)
	if err1 != nil {
		t.Error(err1)
	}
	cond3, err2 := Parse(`NOT title:(IPhone)`)
	if err2 != nil {
		t.Error(err2)
	}

	target := condition.NewNotCondition(
		condition.NewPhraseCondition("小米").SetField("title"),
		condition.NewPhraseCondition("华为").SetField("title"),
	)
	target2 := condition.NewNotCondition(
		condition.NewOrCondition(
			condition.NewPhraseCondition("小米").SetField("title"),
			condition.NewPhraseCondition("华为").SetField("title"),
		),
		condition.NewPhraseCondition("电脑").SetField("type"),
	)
	target3 := condition.NewNotCondition(
		nil,
		condition.NewStringCondition("IPhone").SetField("title"),
	)

	if !reflect.DeepEqual(cond, target) {
		t.Fatal("Parse failed")
	}
	if !reflect.DeepEqual(cond2, target2) {
		t.Fatal("Parse failed")
	}
	if !reflect.DeepEqual(cond3, target3) {
		t.Fatal("Parse failed")
	}

	t.Log("Parse success")
}

// 优先级条件测试
func TestPriorityCondition(t *testing.T) {
	// 最终会被解析为 title:((小米 OR (华为 AND VIVO)) NOT IPhone)
	cond, err := Parse(`title:(小米 OR 华为 AND VIVO NOT IPhone)`)
	if err != nil {
		t.Error(err)
	}
	target := condition.NewNotCondition(
		condition.NewOrCondition(
			condition.NewStringCondition("小米").SetField("title"),
			condition.NewAndCondition(
				condition.NewStringCondition("华为").SetField("title"),
				condition.NewStringCondition("VIVO").SetField("title"),
			),
		),
		condition.NewStringCondition("IPhone").SetField("title"),
	)

	if !reflect.DeepEqual(cond, target) {
		t.Fatal("Parse failed")
	}
	t.Log("Parse success")
}

func TestMultiCondition(t *testing.T) {
	expr := `mobile:((小米 OR 华为) NOT (小米*Pro OR 华为*Pro)) AND type:("手机" OR "电脑") NOT RAM:(8)`
	cond, err := Parse(expr)
	if err != nil {
		t.Error(err)
	}

	left := condition.NewAndCondition(
		condition.NewNotCondition(
			condition.NewOrCondition(
				condition.NewStringCondition("小米").SetField("mobile"),
				condition.NewStringCondition("华为").SetField("mobile"),
			),
			condition.NewOrCondition(
				condition.NewStringCondition("小米*Pro").SetField("mobile"),
				condition.NewStringCondition("华为*Pro").SetField("mobile"),
			),
		),
		condition.NewOrCondition(
			condition.NewPhraseCondition("手机").SetField("type"),
			condition.NewPhraseCondition("电脑").SetField("type"),
		),
	)
	right := condition.NewNumberCondition("8").SetField("RAM")
	target := condition.NewNotCondition(left, right)
	if !reflect.DeepEqual(cond, target) {
		t.Fatal("Parse failed")
	}

	t.Log("Parse success")

}

func pointer(v string) *string {
	return &v
}
