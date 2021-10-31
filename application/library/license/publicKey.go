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

package license

import (
	"io/ioutil"
	"path/filepath"

	"github.com/admpub/log"
	"github.com/webx-top/echo"
)

var publicKey string

func PublicKey() string {
	return publicKey
}

func GetOrLoadPublicKey() string {
	if len(publicKey) == 0 {
		LoadPublicKey()
	}
	return publicKey
}

func LoadPublicKey() {
	pubKeyFile := filepath.Join(echo.Wd(), `data`, `nging.pem.pub`)
	b, err := ioutil.ReadFile(pubKeyFile)
	if err != nil {
		log.Error(`Failed to reading public key file [ ` + pubKeyFile + ` ]: ` + err.Error())
		return
	}
	publicKey = string(b)
}

func SetPublicKey(pubkey string) {
	publicKey = pubkey
}
