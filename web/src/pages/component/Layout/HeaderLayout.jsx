import React from "react";
import {Avatar, Badge, Button, Col, Dropdown, Image, Layout, Menu, message, Row, Tag} from "antd";
import {Link, Redirect, Route, Switch} from "react-router-dom";
import {BellOutlined, LogoutOutlined, SettingOutlined} from '@ant-design/icons';

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
import {getAccount} from "../../User/Auth/Auth";


const logo = "https://cdn.open-ct.com/logo/openct_logo_1082x328.png"

class HeaderLayout extends React.Component {
  state = {
    current: 'home',
    account: undefined,
    messageCount: 0,
    menu: 'home'
  }

  componentDidMount() {
    AuthApi.getAccount()
      .then((res) => {
        this.setState({
          account: res.data.data
        })
      })
      .catch((e) => {
        console.log(e)
      })

    this.changeMenu()
  }

  changeMenu = (e) => {
    if (e !== undefined) {
      this.setState({menu: e.key})
      return
    }
    const p = this.props.location.pathname
    if (p.startsWith("/home/my-project")) {
      this.setState({menu: 'my-project'})
    } else if (p.startsWith("/home/public-project")) {
      this.setState({menu: 'public-project'})
    } else {
      this.setState({menu: 'home'})
    }
  }

  renderHomeIfLoggedIn(component) {
    if (this.state.account === null) {
      message.warn('请先登录')
      return <Redirect to={'/home'}/>
    } else if (this.state.account === undefined) {
      return null
    } else {
      return component
    }
  }

  handleRightDropdownClick(e) {
    let account = this.state.account;
    if (e.key === 'my-account') {
      window.open(Auth.getMyProfileUrl(account));
    } else if (e.key === 'logout') {
      AuthApi.logout()
        .then((res) => {
          if (res.data.code === 200) {
            this.setState({
              account: null
            })
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
      <Dropdown overlay={menu} placement="bottomRight" trigger="click">
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
        <>
        <span style={{float: 'left'}}>
          <Tag>{this.state.account.tag}</Tag>
        </span>
          {this.renderRightDropdown()}
        </>
      )
    }
  }

  render() {
    const {menu, messageCount} = this.state;
    return (
      <Layout style={{minHeight: '100vh', textAlign: 'left'}}>
        <Layout.Header style={{backgroundColor: 'white', paddingLeft: '4px', paddingRight: '4px'}}>
          <Row>
            <Col xxl={16} xl={14} lg={10} md={6} sm={8} xs={0}>
              <Link to="/home">
                <Image height={50} width={165} style={{margin: '7px'}} src={logo} preview={false}/>
              </Link>
            </Col>
            <Col xxl={4} xl={6} lg={8} md={10} sm={8} xs={6}>
              <Menu
                theme="light"
                mode="horizontal"
                style={{border: 0}}
                defaultSelectedKeys={['home']}
                selectedKeys={[menu]}
                onClick={e => this.changeMenu(e)}
              >
                <Menu.Item key="home">
                  <Link to="/home">
                    首页
                  </Link>
                </Menu.Item>
                <Menu.Item key="public-project">
                  <Link to="/home/public-project">
                    公共项目
                  </Link>
                </Menu.Item>
                <Menu.Item key="my-project">
                  <Link to="/home/my-project">
                    我的空间
                  </Link>
                </Menu.Item>
              </Menu>
            </Col>
            <Col xxl={4} xl={4} lg={6} md={8} sm={8} xs={18}>
              {
                <>
                  <span style={{float: 'left', marginRight: '20px'}}>
                    <Link to="/home/message">
                      <Button
                        shape="circle"
                        type="text"
                        icon={
                          <Badge count={messageCount} overflowCount={99} size="small">
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
            <Route exact path="/home" render={(props) => <Home account={this.state.account} {...props} />}/>
            <Route exact path="/home/public-project"
                   render={(props) => this.renderHomeIfLoggedIn(<Project account={this.state.account} {...props} />)}/>
            <Route path="/home/my-project" render={(props) => this.renderHomeIfLoggedIn(<MyProject
              account={this.state.account} {...props} />)}/>
            <Route path="/home/message"
                   render={(props) => this.renderHomeIfLoggedIn(<Message account={this.state.account} {...props} />)}/>

            <Route exact path="/home/project/:id/info" render={(props) => this.renderHomeIfLoggedIn(<ProjectInfo
              account={this.state.account} {...props} />)}/>
            <Route exact path="/home/project/:id/info/edit"
                   render={(props) => this.renderHomeIfLoggedIn(<EditInfo account={this.state.account} {...props} />)}/>
            <Route exact path="/home/project/:id/outline/edit"
                   render={(props) => this.renderHomeIfLoggedIn(<EditOutlined
                     account={this.state.account} {...props} />)}/>

            <Route exact path="/home/project/:pid/student/:sid/evidence"
                   render={(props) => this.renderHomeIfLoggedIn(<Evidence account={this.state.account} {...props} />)}/>

            <Route exact path="/home/project/:pid/section/:sid/edit"
                   render={(props) => this.renderHomeIfLoggedIn(<SectionEditPage
                     account={this.state.account} {...props} />)}/>
            <Route exact path="/home/project/:pid/section/:sid/task/:tid/survey/edit"
                   render={(props) => this.renderHomeIfLoggedIn(<SurveyEditPage
                     account={this.state.account} {...props} />)}/>
            <Route exact path="/home/project/:pid/section/:sid/preview"
                   render={(props) => this.renderHomeIfLoggedIn(<PreviewSection
                     account={this.state.account} {...props} />)}/>
          </Switch>
        </Layout.Content>
        <Layout.Footer style={{textAlign: 'center'}}>OpenPBL</Layout.Footer>
      </Layout>
    );
  }
}

export default HeaderLayout
