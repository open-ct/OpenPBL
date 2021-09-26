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

import {trim} from "./Util";
import axios from "axios";

export let authConfig = {
  serverUrl: process.env.REACT_APP_CASDOOR_ENDPOINT,
  clientId: process.env.REACT_APP_CLIENT_ID,
  appName: process.env.REACT_APP_APP_NAME,
  organizationName: process.env.REACT_APP_CASDOOR_ORGANIZATION,
}

export function initAuthWithConfig(config) {
  authConfig = config;
}
export function getCasdoorService() {
  return axios.create({
    baseURL: authConfig.serverUrl,
    timeout: 10000,
    withCredentials: true
  })
}

export function getUser(teacherId) {
  let request = getCasdoorService()
  return request({
    url: '/api/get-user',
    method: 'get',
    params: {id: authConfig.organizationName + '/' + teacherId}
  })
}

export function getAccount(account) {
  let request = getCasdoorService()
  return request({
    url: `/account?access_token=${account.accessToken}`,
    method: 'get',
  })
}

export function getAuthorizeUrl() {
  const redirectUri = `${window.location.origin}/callback/`;
  const scope = "read";
  const state = authConfig.appName;
  return `${trim(authConfig.serverUrl)}/login/oauth/authorize?client_id=${authConfig.clientId}&response_type=code&redirect_uri=${redirectUri}&scope=${scope}&state=${state}`;
}

export function getUserProfileUrl(userName, account) {
  let param = "";
  if (account !== undefined && account !== null) {
    param = `?access_token=${account.accessToken}`;
  }
  return `${trim(authConfig.serverUrl)}/users/${authConfig.organizationName}/${userName}${param}`;
}

export function getMyProfileUrl(account) {
  let param = "";
  if (account !== undefined && account !== null) {
    param = `?access_token=${account.accessToken}`;
  }
  return `${trim(authConfig.serverUrl)}/account${param}`;
}
