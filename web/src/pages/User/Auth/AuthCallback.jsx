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

import React from "react";
import {Button, message, Result, Spin} from "antd";

import AuthApi from "../../../api/AuthApi";
import {getAccount} from "./Auth"
import * as Auth from "./Auth";

class AuthCallback extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      msg: null,
    };
  }

  componentDidMount() {
    this.login();
  }

  login() {
    const params = new URLSearchParams(this.props.location.search);
    AuthApi.login(params.get('code'), params.get('state'))
      .then((res) => {
        if (res.data.code === 200) {
          AuthApi.getAccount()
            .then((resp) => {
              message.success("登录成功，跳转主页");
              window.location.href = "/home"
              window.open(Auth.getMyProfileUrl(resp.data.data));
            })
            .catch((e) => {
              console.log(e)
            })
        } else {
          this.setState({
            msg: res.msg,
          });
        }
      })
      .catch((e) => {
        console.log(e)
      })

  }

  render() {
    return (
      <div style={{textAlign: "center"}}>
        {
          (this.state.msg === null) ? (
            <Spin size="large" tip="登录中..." style={{paddingTop: "10%"}}/>
          ) : (
            <div style={{display: "inline"}}>
              <Result
                status="error"
                title="登录失败"
                subTitle={this.state.msg}
                extra={[
                  <Button type="primary" key="details">
                    信息
                  </Button>,
                  <Button key="help">帮助</Button>,
                ]}
              >
              </Result>
            </div>
          )
        }
      </div>
    )
  }
}

export default AuthCallback
