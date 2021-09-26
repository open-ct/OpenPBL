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

package util

import "github.com/casdoor/casdoor-go-sdk/auth"

func GetUserId(claims *auth.Claims) (id string) {
	return claims.Name
}

func IsStudent(claims *auth.Claims) (b bool) {
	return claims.Tag != "老师"
}
func IsTeacher(claims *auth.Claims) (b bool) {
	return claims.Tag == "老师"
}