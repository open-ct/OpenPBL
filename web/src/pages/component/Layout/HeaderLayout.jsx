import React from "react";
import {Avatar, Badge, Button, Col, Dropdown, Layout, Menu, Row} from "antd";
import {Link, Redirect, Route, Switch} from "react-router-dom";
import {BellOutlined, LogoutOutlined, SettingOutlined} from '@ant-design/icons';

import './index.less'

import * as Auth from "../../User/Auth/Auth"
import AuthApi from "../../../api/AuthApi"
import Home from "../../Home";
import Project from "../../Project";
import MyProject from "../../Project/MyProject";
import Message from "../../Message";
import ProjectInfo from "../../Project/ProjectInfo";
import EditInfo from "../../Project/CreateProject/Info";
import EditOutlined from "../../Project/CreateProject/Outline";
import Evidence from "../../Project/Evidence";
import SectionEditPage from "../../Project/CreateProject/Section/SectionEditPage";
import SurveyEditPage from "../../Project/CreateProject/Survey/SurveyEditPage";
import PreviewSection from "../../Project/PreviewProject/PreviewSection";


class HeaderLayout extends React.Component {
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
          localStorage.setItem("type", res.data.data.tag)
        } else {
          localStorage.setItem("type", "")
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
            localStorage.setItem("type", "")
            window.location.href = '/'
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
          <SettingOutlined/>&nbsp;&nbsp;
          我的账户
        </Menu.Item>
        <Menu.Item key='logout'>
          <LogoutOutlined/>&nbsp;&nbsp;
          退出账号
        </Menu.Item>
      </Menu>
    )
    return (
      <Dropdown overlay={menu} placement="bottomRight">
        <div style={{cursor: 'pointer'}}>
          <Avatar size="large" src={this.state.account.avatar}/>&nbsp;
          <span>{this.state.account.name}</span>
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
      <Layout style={{minHeight: '100vh', textAlign: 'left'}}>
        <Layout.Header style={{backgroundColor: 'white'}}>
          <Row>
            <Col xxl={15} xl={11} lg={8} md={6} sm={6} xs={10}>
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
                <Menu.Item key="public-project">
                  <Link to="/public-project">
                    公共项目
                  </Link>
                </Menu.Item>
                <Menu.Item key="my-project">
                  <Link to="/my-project">
                    我的空间
                  </Link>
                </Menu.Item>
                <Menu.Item key="bbs">
                  <a href="https://bbs.open-ct.com">
                    在线论坛
                  </a>
                </Menu.Item>
              </Menu>
            </Col>
            <Col xxl={3} xl={3} lg={4} md={4} sm={6} xs={8}>
              {
                <>
                  <span style={{float: 'left', marginRight: '20px'}}>
                    <Link to="/message">
                      <Button
                        shape="circle"
                        type="text"
                        icon={
                          <Badge count={999} overflowCount={99} size="small">
                            <BellOutlined/>
                          </Badge>
                        }
                        size="large"
                      />
                    </Link>
                  </span>
                  {this.renderAccount()}
                </>
              }
            </Col>
          </Row>
        </Layout.Header>
        <Layout.Content>
          <Switch>
            <Route exact path="/" render={() => (
              <Redirect to="/landing"/>
            )}/>

            <Route exact path="/landing" component={Home}/>
            <Route exact path="/public-project" component={Project}/>
            <Route path="/my-project" component={MyProject}/>
            <Route path="/message" component={Message}/>

            <Route exact path="/project/:id/info" component={ProjectInfo}/>
            <Route exact path="/project/:id/info/edit" component={EditInfo}/>
            <Route exact path="/project/:id/outline/edit" component={EditOutlined}/>

            <Route exact path="/project/:pid/student/:sid/evidence" component={Evidence}/>

            <Route exact path="/project/:pid/section/:sid/edit" component={SectionEditPage}/>
            <Route exact path="/project/:pid/section/:sid/task/:tid/survey/edit" component={SurveyEditPage}/>
            <Route exact path="/project/:pid/section/:sid/preview" component={PreviewSection}/>
          </Switch>
        </Layout.Content>
        <Layout.Footer style={{textAlign: 'center'}}>OpenPBL</Layout.Footer>
      </Layout>
    );
  }
}

export default HeaderLayout;