/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present Wenhui Shen <swh@admpub.com>

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

package notice

import (
	"sync/atomic"
)

func NewProgress() *Progress {
	return &Progress{
		Total:   -1,
		Finish:  -1,
		Percent: 0,
	}
}

type Progress struct {
	Total        int64   `json:"total" xml:"total"`
	Finish       int64   `json:"finish" xml:"finish"`
	Percent      float64 `json:"percent" xml:"percent"`
	Complete     bool    `json:"complete" xml:"complete"`
	control      IsExited
	autoComplete bool
}

func (p *Progress) IsExited() bool {
	if p.control == nil {
		return false
	}
	return p.control.IsExited()
}

func (p *Progress) SetControl(control IsExited) *Progress {
	p.control = control
	return p
}

func (p *Progress) CalcPercent() *Progress {
	if p.Total > 0 {
		p.Percent = (float64(p.Finish) / float64(p.Total)) * 100
		if p.Percent < 0 {
			p.Percent = 0
		}
	} else if p.Total == 0 {
		p.Percent = 100
	} else {
		p.Percent = 0
	}
	return p
}

func (p *Progress) Add(n int64) *Progress {
	if atomic.LoadInt64(&p.Finish) > 0 {
		atomic.StoreInt64(&p.Finish, 0)
	}
	if atomic.LoadInt64(&p.Total) == -1 {
		n++
	}
	atomic.AddInt64(&p.Total, n)
	return p
}

func (p *Progress) Done(n int64) int64 {
	if atomic.LoadInt64(&p.Finish) == -1 {
		n++
	}
	newN := atomic.AddInt64(&p.Finish, n)
	if p.autoComplete && newN >= atomic.LoadInt64(&p.Total) {
		p.SetComplete()
	}
	return newN
}

func (p *Progress) AutoComplete(on bool) *Progress {
	p.autoComplete = on
	return p
}

func (p *Progress) SetComplete() *Progress {
	p.Complete = true
	return p
}

func (p *Progress) Callback(total int64, exec func(callback func(strLen int)) error) error {
	var remains int64 = 100
	var partPercent float64
	var perByteVal float64
	if total > 0 {
		perByteVal = float64(remains) / float64(total)
	}
	var callback = func(strLen int) {
		if perByteVal <= 0 {
			return
		}
		partPercent += perByteVal * float64(strLen)
		if partPercent < 1 {
			return
		}
		percent := int64(partPercent)
		remains -= percent
		p.Done(percent)
		partPercent = partPercent - float64(percent)
	}
	err := exec(callback)
	if remains > 0 {
		p.Done(remains)
	}
	return err
}
