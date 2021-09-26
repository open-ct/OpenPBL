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

import axios from 'axios'
import {message} from 'antd'

function getServerUrl() {
  const hostname = window.location.hostname
  if (hostname === 'localhost') {
    return `http://${hostname}:5000/api`
  }
  return '/api'
}

const request = axios.create({
  baseURL: getServerUrl(),
  timeout: 10000,
  withCredentials: true
})

request.interceptors.response.use(res=>{
  if (res.data.code === 401) {
    message.error(res.data.msg)
    setTimeout(()=>{
      window.location.href = '/'
    }, 1000)
  } else if (res.data.code === 403) {
    message.error(res.data.msg)
    setTimeout(()=>{
      window.location.href = '/'
    }, 1000)
  } else {
    return res
  }
})

export default request
