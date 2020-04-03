// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
)

type Slice_NgingFileEmbedded []*NgingFileEmbedded

func (s Slice_NgingFileEmbedded) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingFileEmbedded) RangeRaw(fn func(m *NgingFileEmbedded) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingFileEmbedded) GroupBy(keyField string) map[string][]*NgingFileEmbedded {
	r := map[string][]*NgingFileEmbedded{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingFileEmbedded{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingFileEmbedded) KeyBy(keyField string) map[string]*NgingFileEmbedded {
	r := map[string]*NgingFileEmbedded{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingFileEmbedded) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingFileEmbedded) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingFileEmbedded) FromList(data interface{}) Slice_NgingFileEmbedded {
	values, ok := data.([]*NgingFileEmbedded)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingFileEmbedded{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

// NgingFileEmbedded 嵌入文件
type NgingFileEmbedded struct {
	base    factory.Base
	objects []*NgingFileEmbedded

	Id        uint64 `db:"id,omitempty,pk" bson:"id,omitempty" comment:"主键" json:"id" xml:"id"`
	Project   string `db:"project" bson:"project" comment:"项目名" json:"project" xml:"project"`
	TableId   string `db:"table_id" bson:"table_id" comment:"表主键" json:"table_id" xml:"table_id"`
	TableName string `db:"table_name" bson:"table_name" comment:"表名称" json:"table_name" xml:"table_name"`
	FieldName string `db:"field_name" bson:"field_name" comment:"字段名" json:"field_name" xml:"field_name"`
	FileIds   string `db:"file_ids" bson:"file_ids" comment:"文件id列表" json:"file_ids" xml:"file_ids"`
	Embedded  string `db:"embedded" bson:"embedded" comment:"是否(Y/N)为内嵌文件" json:"embedded" xml:"embedded"`
}

// - base function

func (a *NgingFileEmbedded) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingFileEmbedded) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingFileEmbedded) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingFileEmbedded) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingFileEmbedded) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingFileEmbedded) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingFileEmbedded) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingFileEmbedded) SetNamer(namer func(string) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingFileEmbedded) Namer() func(string) string {
	return a.base.Namer()
}

func (a *NgingFileEmbedded) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingFileEmbedded) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingFileEmbedded) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingFileEmbedded) Objects() []*NgingFileEmbedded {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingFileEmbedded) XObjects() Slice_NgingFileEmbedded {
	return Slice_NgingFileEmbedded(a.Objects())
}

func (a *NgingFileEmbedded) NewObjects() factory.Ranger {
	return &Slice_NgingFileEmbedded{}
}

func (a *NgingFileEmbedded) InitObjects() *[]*NgingFileEmbedded {
	a.objects = []*NgingFileEmbedded{}
	return &a.objects
}

func (a *NgingFileEmbedded) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingFileEmbedded) Short_() string {
	return "nging_file_embedded"
}

func (a *NgingFileEmbedded) Struct_() string {
	return "NgingFileEmbedded"
}

func (a *NgingFileEmbedded) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingFileEmbedded) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingFileEmbedded) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	base := a.base
	if !a.base.Eventable() {
		err = a.Param(mw, args...).SetRecv(a).One()
		a.base = base
		return
	}
	queryParam := a.Param(mw, args...).SetRecv(a)
	if err = DBI.FireReading(a, queryParam); err != nil {
		return
	}
	err = queryParam.One()
	a.base = base
	if err == nil {
		err = DBI.FireReaded(a, queryParam)
	}
	return
}

func (a *NgingFileEmbedded) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingFileEmbedded:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFileEmbedded(*v))
		case []*NgingFileEmbedded:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFileEmbedded(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingFileEmbedded) GroupBy(keyField string, inputRows ...[]*NgingFileEmbedded) map[string][]*NgingFileEmbedded {
	var rows Slice_NgingFileEmbedded
	if len(inputRows) > 0 {
		rows = Slice_NgingFileEmbedded(inputRows[0])
	} else {
		rows = Slice_NgingFileEmbedded(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingFileEmbedded) KeyBy(keyField string, inputRows ...[]*NgingFileEmbedded) map[string]*NgingFileEmbedded {
	var rows Slice_NgingFileEmbedded
	if len(inputRows) > 0 {
		rows = Slice_NgingFileEmbedded(inputRows[0])
	} else {
		rows = Slice_NgingFileEmbedded(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingFileEmbedded) AsKV(keyField string, valueField string, inputRows ...[]*NgingFileEmbedded) param.Store {
	var rows Slice_NgingFileEmbedded
	if len(inputRows) > 0 {
		rows = Slice_NgingFileEmbedded(inputRows[0])
	} else {
		rows = Slice_NgingFileEmbedded(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingFileEmbedded) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingFileEmbedded:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFileEmbedded(*v))
		case []*NgingFileEmbedded:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFileEmbedded(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingFileEmbedded) Add() (pk interface{}, err error) {
	a.Id = 0
	if len(a.TableId) == 0 {
		a.TableId = "0"
	}
	if len(a.Embedded) == 0 {
		a.Embedded = "Y"
	}
	if a.base.Eventable() {
		err = DBI.Fire("creating", a, nil)
		if err != nil {
			return
		}
	}
	pk, err = a.Param(nil).SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint64(v)
		}
	}
	if err == nil && a.base.Eventable() {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *NgingFileEmbedded) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if len(a.TableId) == 0 {
		a.TableId = "0"
	}
	if len(a.Embedded) == 0 {
		a.Embedded = "Y"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Update()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *NgingFileEmbedded) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingFileEmbedded) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["table_id"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["table_id"] = "0"
		}
	}
	if val, ok := kvset["embedded"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["embedded"] = "Y"
		}
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(kvset).Update()
	}
	m := *a
	m.FromRow(kvset)
	var editColumns []string
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, editColumns, mw, args...)
}

func (a *NgingFileEmbedded) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		if len(a.TableId) == 0 {
			a.TableId = "0"
		}
		if len(a.Embedded) == 0 {
			a.Embedded = "Y"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Id = 0
		if len(a.TableId) == 0 {
			a.TableId = "0"
		}
		if len(a.Embedded) == 0 {
			a.Embedded = "Y"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint64(v)
		}
	}
	if err == nil && a.base.Eventable() {
		if pk == nil {
			err = DBI.Fire("updated", a, mw, args...)
		} else {
			err = DBI.Fire("created", a, nil)
		}
	}
	return
}

func (a *NgingFileEmbedded) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).Delete()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *NgingFileEmbedded) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingFileEmbedded) Reset() *NgingFileEmbedded {
	a.Id = 0
	a.Project = ``
	a.TableId = ``
	a.TableName = ``
	a.FieldName = ``
	a.FileIds = ``
	a.Embedded = ``
	return a
}

func (a *NgingFileEmbedded) AsMap() param.Store {
	r := param.Store{}
	r["Id"] = a.Id
	r["Project"] = a.Project
	r["TableId"] = a.TableId
	r["TableName"] = a.TableName
	r["FieldName"] = a.FieldName
	r["FileIds"] = a.FileIds
	r["Embedded"] = a.Embedded
	return r
}

func (a *NgingFileEmbedded) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint64(value)
		case "project":
			a.Project = param.AsString(value)
		case "table_id":
			a.TableId = param.AsString(value)
		case "table_name":
			a.TableName = param.AsString(value)
		case "field_name":
			a.FieldName = param.AsString(value)
		case "file_ids":
			a.FileIds = param.AsString(value)
		case "embedded":
			a.Embedded = param.AsString(value)
		}
	}
}

func (a *NgingFileEmbedded) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
	case map[string]interface{}:
		for kk, vv := range k {
			a.Set(kk, vv)
		}
	default:
		var (
			kk string
			vv interface{}
		)
		if k, y := key.(string); y {
			kk = k
		} else {
			kk = fmt.Sprint(key)
		}
		if len(value) > 0 {
			vv = value[0]
		}
		switch kk {
		case "Id":
			a.Id = param.AsUint64(vv)
		case "Project":
			a.Project = param.AsString(vv)
		case "TableId":
			a.TableId = param.AsString(vv)
		case "TableName":
			a.TableName = param.AsString(vv)
		case "FieldName":
			a.FieldName = param.AsString(vv)
		case "FileIds":
			a.FileIds = param.AsString(vv)
		case "Embedded":
			a.Embedded = param.AsString(vv)
		}
	}
}

func (a *NgingFileEmbedded) AsRow() param.Store {
	r := param.Store{}
	r["id"] = a.Id
	r["project"] = a.Project
	r["table_id"] = a.TableId
	r["table_name"] = a.TableName
	r["field_name"] = a.FieldName
	r["file_ids"] = a.FileIds
	r["embedded"] = a.Embedded
	return r
}

func (a *NgingFileEmbedded) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *NgingFileEmbedded) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
