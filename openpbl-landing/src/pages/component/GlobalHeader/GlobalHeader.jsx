import React from "react";
import {Avatar, Col, Dropdown, Layout, Menu, Row} from "antd";
import {Link} from "react-router-dom";
import {LogoutOutlined, SettingOutlined} from '@ant-design/icons';

import * as Auth from "../../User/Auth/Auth"
import AuthApi from "../../../api/AuthApi"

import './global-header.less'

const {Header} = Layout;

class GlobalHeader extends React.PureComponent {
  state = {
    current: this.props.current,
    account: null,
  }

  componentDidMount() {
    AuthApi.getAccount()
      .then((res) => {
        if (res.data.code === 200) {
          this.setState({
            account: res.data.data
          })
        }
      })
      .catch((e) => {
        console.log(e)
      })
  }

  handleRightDropdownClick(e) {
    let account = this.state.account;
    console.log(account)
    if (e.key === 'my-account') {
      window.open(Auth.getMyProfileUrl(account));
    } else if (e.key === 'logout') {
      AuthApi.logout()
        .then((res) => {
          if (res.data.code === 200) {
            this.setState({
              account: null
            })
          }
        })
        .catch(e => {
          console.log(e)
        })
    }
  }

  renderRightDropdown() {
    const menu = (
      <Menu onClick={this.handleRightDropdownClick.bind(this)} style={{width: '150px', padding: '5px'}}>
        <Menu.Item key='my-account'>
          <SettingOutlined/>
          我的账户
        </Menu.Item>
        <Menu.Item key='logout'>
          <LogoutOutlined/>
          退出账号
        </Menu.Item>
      </Menu>
    )
    return (
      <Dropdown overlay={menu} placement="bottomRight">
        <div style={{cursor: 'pointer'}}>
          <Avatar size="large" src={this.state.account.avatar}/>&nbsp;
          <span>{this.state.account.username}</span>
        </div>
      </Dropdown>
    );
  }

  renderAccount() {
    if (this.state.account === undefined || this.state.account === null) {
      return (
        <a href={Auth.getAuthorizeUrl()}>
          登录注册
        </a>
      );
    } else {
      return (
        this.renderRightDropdown()
      )
    }
  }

  render() {
    const {current} = this.state;
    return (
      <div>
        <Header style={{backgroundColor: 'white'}}>
          <Row>
            <Col xxl={16} xl={12} lg={10} md={8} sm={10} xs={16}>
              <Link to="/landing">
                <div className="logo">
                  <span style={{fontSize: '25px', color: 'black', float: 'left', marginLeft: '80px'}}>OpenCT</span>
                </div>
              </Link>
            </Col>
            <Col xxl={6} xl={10} lg={12} md={14} sm={12} xs={6}>
              <Menu theme="light" mode="horizontal" defaultSelectedKeys={[current]} style={{border: 0}}>
                <Menu.Item key="landing">
                  <Link to="/landing">
                    首页
                  </Link>
                </Menu.Item>
                <Menu.Item key="project">
                  <Link to="/project">
                    项目学习
                  </Link>
                </Menu.Item>
                <Menu.Item key="bbs">
                  <a href="https://bbs.open-ct.com">
                    在线论坛
                  </a>
                </Menu.Item>
              </Menu>
            </Col>
            <Col xxl={2} xl={2} lg={2} md={2} sm={2} xs={2}>
              {
                this.renderAccount()
              }
            </Col>
          </Row>
        </Header>
      </div>
    );
  }
}

export default GlobalHeader;
