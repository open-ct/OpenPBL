// Copyright 2021 The casbin Authors. All Rights Reserved.
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
// /login/oauth/authorize?client_id=9145af9a92d8194b5ede&response_type=code&redirect_uri=http://localhost:3000/project&scope=read&state=openpbl
// /login/oauth/authorize?client_id=0ba528121ea87b3eb54d&response_type=code&redirect_uri=https://oa.casbin.com/callback&scope=read&state=app-casbin-oa
export let authConfig = {
    category: "email",
    serverUrl: "http://121.196.149.183:8000", // your Casdoor URL, like the official one: https://door.casbin.com
    clientId: "9145af9a92d8194b5ede", // your Casdoor OAuth Client ID
    appName: "openpbl", // your Casdoor application name, like: "app-built-in"
    organizationName: "openct", // your Casdoor organization name, like: "built-in"
}

export function initAuthWithConfig(config) {
    authConfig = config;
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
