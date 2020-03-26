package search

import (
	"github.com/matchstalk/scaffold-gin/tools"
	"reflect"
	"strings"
)

const FromQueryTag = "search"

type resolveSearchTag struct {
	Type   string
	Column string
}

/**
 * 	exact / iexact 等于
 * 	contains / icontains 包含
 *	gt / gte 大于 / 大于等于
 *	lt / lte 小于 / 小于等于
 *	startswith / istartswith 以…起始
 *	endswith / iendswith 以…结束
 *	in
 *	isnull
 *  order 排序		e.g. order[key]=desc     order[key]=asc
 */
func ResolveSearchQuery(driver string, q interface{}, condition map[string][]interface{}, order []string) (map[string][]interface{}, []string) {
	if condition == nil {
		condition = make(map[string][]interface{})
	}
	if order == nil {
		order = make([]string, 0)
	}
	qType := reflect.TypeOf(q)
	qValue := reflect.ValueOf(q)
	var tag string
	var ok bool
	var t resolveSearchTag
	for i := 0; i < qType.NumField(); i++ {
		tag, ok = "", false
		tag, ok = qType.Field(i).Tag.Lookup(FromQueryTag)
		if !ok {
			//递归调用
			condition, order = ResolveSearchQuery(driver, qValue.Field(i).Interface(), condition, order)
			continue
		}
		t = resolveTagValue(tag)
		//解析
		switch t.Type {
		case "exact", "iexact":
			condition[t.Column+" = ?"] = []interface{}{qValue.Field(i).Interface()}
		case "contains", "icontains":
			//注意:mysql不支持ilike
			if driver == tools.Postgres && t.Type == "icontains" {
				condition[t.Column+" ilike ?"] = []interface{}{"%" + qValue.Field(i).String() + "%"}
			} else {
				condition[t.Column+" like ?"] = []interface{}{"%" + qValue.Field(i).String() + "%"}
			}
		case "gt":
			condition[t.Column+" > ?"] = []interface{}{qValue.Field(i).Interface()}
		case "gte":
			condition[t.Column+" >= ?"] = []interface{}{qValue.Field(i).Interface()}
		case "lt":
			condition[t.Column+" < ?"] = []interface{}{qValue.Field(i).Interface()}
		case "lte":
			condition[t.Column+" <= ?"] = []interface{}{qValue.Field(i).Interface()}
		case "startswith", "istartswith":
			if driver == tools.Postgres && t.Type == "istartswith" {
				condition[t.Column+" ilike ?"] = []interface{}{qValue.Field(i).String() + "%"}
			} else {
				condition[t.Column+" like ?"] = []interface{}{qValue.Field(i).String() + "%"}
			}
		case "endswith", "iendswith":
			if driver == tools.Postgres && t.Type == "iendswith" {
				condition[t.Column+" ilike ?"] = []interface{}{"%" + qValue.Field(i).String()}
			} else {
				condition[t.Column+" like ?"] = []interface{}{"%" + qValue.Field(i).String()}
			}
		case "in":
			condition[t.Column+" in (?)"] = []interface{}{qValue.Field(i).Interface()}
		case "isnull":
			if !(qValue.Field(i).IsZero() && qValue.Field(i).IsNil()) {
				condition[t.Column+" isnull"] = make([]interface{}, 0)
			}
		}
	}
	return condition, order
}

func resolveTagValue(tag string) resolveSearchTag {
	var r resolveSearchTag
	tags := strings.Split(tag, ";")
	var ts []string
	for _, t := range tags {
		ts = strings.Split(t, ":")
		if len(ts) == 0 {
			continue
		}
		switch ts[0] {
		case "type":
			if len(ts) > 1 {
				r.Type = ts[1]
			}
		case "column":
			if len(ts) > 1 {
				r.Column = ts[1]
			}
		}
	}
	return r
}
