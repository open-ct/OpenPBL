import React from "react";
import {Button, message, Result, Spin} from "antd";

import AuthApi from "../../../api/AuthApi";


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
          localStorage.setItem("type", res.data.data.tag)
          message.success("登录成功，跳转主页");
          window.location.href = "/project"
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
