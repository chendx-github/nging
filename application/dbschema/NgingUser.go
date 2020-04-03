// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"time"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
)

type Slice_NgingUser []*NgingUser

func (s Slice_NgingUser) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingUser) RangeRaw(fn func(m *NgingUser) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingUser) GroupBy(keyField string) map[string][]*NgingUser {
	r := map[string][]*NgingUser{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingUser{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingUser) KeyBy(keyField string) map[string]*NgingUser {
	r := map[string]*NgingUser{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingUser) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingUser) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingUser) FromList(data interface{}) Slice_NgingUser {
	values, ok := data.([]*NgingUser)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingUser{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

// NgingUser 用户
type NgingUser struct {
	base    factory.Base
	objects []*NgingUser

	Id        uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	Username  string `db:"username" bson:"username" comment:"用户名" json:"username" xml:"username"`
	Email     string `db:"email" bson:"email" comment:"邮箱" json:"email" xml:"email"`
	Mobile    string `db:"mobile" bson:"mobile" comment:"手机号" json:"mobile" xml:"mobile"`
	Password  string `db:"password" bson:"password" comment:"密码" json:"password" xml:"password"`
	Salt      string `db:"salt" bson:"salt" comment:"盐值" json:"salt" xml:"salt"`
	SafePwd   string `db:"safe_pwd" bson:"safe_pwd" comment:"安全密码" json:"safe_pwd" xml:"safe_pwd"`
	Avatar    string `db:"avatar" bson:"avatar" comment:"头像" json:"avatar" xml:"avatar"`
	Gender    string `db:"gender" bson:"gender" comment:"性别(male-男;female-女;secret-保密)" json:"gender" xml:"gender"`
	LastLogin uint   `db:"last_login" bson:"last_login" comment:"最后登录时间" json:"last_login" xml:"last_login"`
	LastIp    string `db:"last_ip" bson:"last_ip" comment:"最后登录IP" json:"last_ip" xml:"last_ip"`
	Disabled  string `db:"disabled" bson:"disabled" comment:"状态" json:"disabled" xml:"disabled"`
	Online    string `db:"online" bson:"online" comment:"是否在线" json:"online" xml:"online"`
	RoleIds   string `db:"role_ids" bson:"role_ids" comment:"角色ID(多个用“,”分隔开)" json:"role_ids" xml:"role_ids"`
	Created   uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated   uint   `db:"updated" bson:"updated" comment:"更新时间" json:"updated" xml:"updated"`
	FileSize  uint64 `db:"file_size" bson:"file_size" comment:"上传文件总大小" json:"file_size" xml:"file_size"`
	FileNum   uint64 `db:"file_num" bson:"file_num" comment:"上传文件数量" json:"file_num" xml:"file_num"`
}

// - base function

func (a *NgingUser) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingUser) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingUser) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingUser) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingUser) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingUser) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingUser) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingUser) SetNamer(namer func(string) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingUser) Namer() func(string) string {
	return a.base.Namer()
}

func (a *NgingUser) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingUser) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingUser) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingUser) Objects() []*NgingUser {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingUser) XObjects() Slice_NgingUser {
	return Slice_NgingUser(a.Objects())
}

func (a *NgingUser) NewObjects() factory.Ranger {
	return &Slice_NgingUser{}
}

func (a *NgingUser) InitObjects() *[]*NgingUser {
	a.objects = []*NgingUser{}
	return &a.objects
}

func (a *NgingUser) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingUser) Short_() string {
	return "nging_user"
}

func (a *NgingUser) Struct_() string {
	return "NgingUser"
}

func (a *NgingUser) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingUser) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingUser) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
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

func (a *NgingUser) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingUser:
			err = DBI.FireReaded(a, queryParam, Slice_NgingUser(*v))
		case []*NgingUser:
			err = DBI.FireReaded(a, queryParam, Slice_NgingUser(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingUser) GroupBy(keyField string, inputRows ...[]*NgingUser) map[string][]*NgingUser {
	var rows Slice_NgingUser
	if len(inputRows) > 0 {
		rows = Slice_NgingUser(inputRows[0])
	} else {
		rows = Slice_NgingUser(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingUser) KeyBy(keyField string, inputRows ...[]*NgingUser) map[string]*NgingUser {
	var rows Slice_NgingUser
	if len(inputRows) > 0 {
		rows = Slice_NgingUser(inputRows[0])
	} else {
		rows = Slice_NgingUser(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingUser) AsKV(keyField string, valueField string, inputRows ...[]*NgingUser) param.Store {
	var rows Slice_NgingUser
	if len(inputRows) > 0 {
		rows = Slice_NgingUser(inputRows[0])
	} else {
		rows = Slice_NgingUser(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingUser) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingUser:
			err = DBI.FireReaded(a, queryParam, Slice_NgingUser(*v))
		case []*NgingUser:
			err = DBI.FireReaded(a, queryParam, Slice_NgingUser(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingUser) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.Gender) == 0 {
		a.Gender = "secret"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.Online) == 0 {
		a.Online = "N"
	}
	if a.base.Eventable() {
		err = DBI.Fire("creating", a, nil)
		if err != nil {
			return
		}
	}
	pk, err = a.Param(nil).SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil && a.base.Eventable() {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *NgingUser) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Gender) == 0 {
		a.Gender = "secret"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.Online) == 0 {
		a.Online = "N"
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

func (a *NgingUser) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingUser) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["gender"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["gender"] = "secret"
		}
	}
	if val, ok := kvset["disabled"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["disabled"] = "N"
		}
	}
	if val, ok := kvset["online"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["online"] = "N"
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

func (a *NgingUser) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		a.Updated = uint(time.Now().Unix())
		if len(a.Gender) == 0 {
			a.Gender = "secret"
		}
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if len(a.Online) == 0 {
			a.Online = "N"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.Gender) == 0 {
			a.Gender = "secret"
		}
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if len(a.Online) == 0 {
			a.Online = "N"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
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

func (a *NgingUser) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

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

func (a *NgingUser) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingUser) Reset() *NgingUser {
	a.Id = 0
	a.Username = ``
	a.Email = ``
	a.Mobile = ``
	a.Password = ``
	a.Salt = ``
	a.SafePwd = ``
	a.Avatar = ``
	a.Gender = ``
	a.LastLogin = 0
	a.LastIp = ``
	a.Disabled = ``
	a.Online = ``
	a.RoleIds = ``
	a.Created = 0
	a.Updated = 0
	a.FileSize = 0
	a.FileNum = 0
	return a
}

func (a *NgingUser) AsMap() param.Store {
	r := param.Store{}
	r["Id"] = a.Id
	r["Username"] = a.Username
	r["Email"] = a.Email
	r["Mobile"] = a.Mobile
	r["Password"] = a.Password
	r["Salt"] = a.Salt
	r["SafePwd"] = a.SafePwd
	r["Avatar"] = a.Avatar
	r["Gender"] = a.Gender
	r["LastLogin"] = a.LastLogin
	r["LastIp"] = a.LastIp
	r["Disabled"] = a.Disabled
	r["Online"] = a.Online
	r["RoleIds"] = a.RoleIds
	r["Created"] = a.Created
	r["Updated"] = a.Updated
	r["FileSize"] = a.FileSize
	r["FileNum"] = a.FileNum
	return r
}

func (a *NgingUser) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "username":
			a.Username = param.AsString(value)
		case "email":
			a.Email = param.AsString(value)
		case "mobile":
			a.Mobile = param.AsString(value)
		case "password":
			a.Password = param.AsString(value)
		case "salt":
			a.Salt = param.AsString(value)
		case "safe_pwd":
			a.SafePwd = param.AsString(value)
		case "avatar":
			a.Avatar = param.AsString(value)
		case "gender":
			a.Gender = param.AsString(value)
		case "last_login":
			a.LastLogin = param.AsUint(value)
		case "last_ip":
			a.LastIp = param.AsString(value)
		case "disabled":
			a.Disabled = param.AsString(value)
		case "online":
			a.Online = param.AsString(value)
		case "role_ids":
			a.RoleIds = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsUint(value)
		case "file_size":
			a.FileSize = param.AsUint64(value)
		case "file_num":
			a.FileNum = param.AsUint64(value)
		}
	}
}

func (a *NgingUser) Set(key interface{}, value ...interface{}) {
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
			a.Id = param.AsUint(vv)
		case "Username":
			a.Username = param.AsString(vv)
		case "Email":
			a.Email = param.AsString(vv)
		case "Mobile":
			a.Mobile = param.AsString(vv)
		case "Password":
			a.Password = param.AsString(vv)
		case "Salt":
			a.Salt = param.AsString(vv)
		case "SafePwd":
			a.SafePwd = param.AsString(vv)
		case "Avatar":
			a.Avatar = param.AsString(vv)
		case "Gender":
			a.Gender = param.AsString(vv)
		case "LastLogin":
			a.LastLogin = param.AsUint(vv)
		case "LastIp":
			a.LastIp = param.AsString(vv)
		case "Disabled":
			a.Disabled = param.AsString(vv)
		case "Online":
			a.Online = param.AsString(vv)
		case "RoleIds":
			a.RoleIds = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		case "FileSize":
			a.FileSize = param.AsUint64(vv)
		case "FileNum":
			a.FileNum = param.AsUint64(vv)
		}
	}
}

func (a *NgingUser) AsRow() param.Store {
	r := param.Store{}
	r["id"] = a.Id
	r["username"] = a.Username
	r["email"] = a.Email
	r["mobile"] = a.Mobile
	r["password"] = a.Password
	r["salt"] = a.Salt
	r["safe_pwd"] = a.SafePwd
	r["avatar"] = a.Avatar
	r["gender"] = a.Gender
	r["last_login"] = a.LastLogin
	r["last_ip"] = a.LastIp
	r["disabled"] = a.Disabled
	r["online"] = a.Online
	r["role_ids"] = a.RoleIds
	r["created"] = a.Created
	r["updated"] = a.Updated
	r["file_size"] = a.FileSize
	r["file_num"] = a.FileNum
	return r
}

func (a *NgingUser) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *NgingUser) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
