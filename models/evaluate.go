// Copyright 2021 The OpenPBL Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

type Evaluate struct {
	Chapter        `xorm:"extends"`
	SectionOutline `xorm:"extends"`
}


func GetProjectEvaluate(pid string) (evaluate []Evaluate, err error) {
	var c []Chapter
	err = (&Chapter{}).GetEngine().
		Where("project_id = ?", pid).
		Asc("chapter_number").
		Find(&c)
	evaluate = make([]Evaluate, len(c))
	var outline SectionOutline
	for i:=0; i< len(c); i++ {
		evaluate[i].Chapter = c[i]
		err = (&Section{}).GetEngine().
			Where("chapter_id = ?", c[i].Id).
			Find(&outline)
		evaluate[i].SectionOutline = outline
	}
	return
}