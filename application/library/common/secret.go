/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present  Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package common

import (
	"github.com/webx-top/codec"
	"github.com/webx-top/com"
	"github.com/webx-top/echo"
)

// SetRandomSecret 设置随机密码
func SetRandomSecret(ctx echo.Context, sessionKey string, storeKey ...string) {
	secret := com.RandomAlphanumeric(32)
	ctx.Session().Set(`secrect_`+sessionKey, secret)
	if len(storeKey) > 0 {
		ctx.Set(storeKey[0], secret)
	} else {
		ctx.Set(sessionKey, secret)
	}
}

// DeleteRandomSecret 删除随机密码
func DeleteRandomSecret(ctx echo.Context, sessionKey string) {
	ctx.Session().Delete(`secrect_` + sessionKey)
}

// DecryptedByRandomSecret 用上次设置的随机密码解密
func DecryptedByRandomSecret(ctx echo.Context, sessionKey string, datas ...*string) {
	secret := ctx.Session().Get(`secrect_` + sessionKey)
	if secret != nil {
		secrets, _ := secret.(string)
		Decrypt(secrets, datas...)
	}
}

// Decrypt 数据解密
func Decrypt(secret string, datas ...*string) {
	crypto := codec.NewDesECBCrypto()
	for _, data := range datas {
		if len(*data) == 0 {
			continue
		}
		*data = crypto.Decode(*data, secret)
	}
}

// Encrypt 数据加密
func Encrypt(secret string, datas ...*string) {
	crypto := codec.NewDesECBCrypto()
	for _, data := range datas {
		if len(*data) == 0 {
			continue
		}
		*data = crypto.Encode(*data, secret)
	}
}

// GenSecret 生成随机密钥
func GenSecret(sizes ...int) string {
	var size int
	if len(sizes) > 0 {
		size = sizes[0]
	}
	if size < 1 {
		size = 32
	}
	secret := com.RandomAlphanumeric(32)
	return secret
}
