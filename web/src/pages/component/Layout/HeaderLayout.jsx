import React from "react";
import {Avatar, Badge, Button, Col, Dropdown, Layout, Menu, Row, Tag, Image} from "antd";
import {Link, Redirect, Route, Switch} from "react-router-dom";
import {BellOutlined, LogoutOutlined, SettingOutlined} from '@ant-design/icons';
import {connect} from "react-redux";

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


const logo = "https://cdn.open-ct.com/logo/openct_logo_1082x328.png"

class HeaderLayout extends React.Component {
  state = {
    current: 'home',
    account: null,
    messageCount: 0,
    menu: 'home'
  }

  componentDidMount() {
    AuthApi.getAccount()
      .then((res) => {
        if (res.data.code === 200) {
          this.setState({
            account: res.data.data
          })
          this.props.setUserType(res.data.data.tag)
        } else {
          this.props.setUserType("")
        }
      })
      .catch((e) => {console.log(e)})

    this.changeMenu()
  }

  changeMenu = (e) => {
    if (e !== undefined ) {
      this.setState({menu: e.key})
      return
    }
    const p = this.props.location.pathname
    console.log(p)
    if (p.startsWith('/home')) {
      this.setState({menu: 'home'})
    } else if (p.startsWith("/my-project")) {
      this.setState({menu: 'my-project'})
    } else if (p.startsWith("/public-project")) {
      this.setState({menu: 'public-project'})
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
            this.props.setUserType("")
            window.location.href = '/'
          }
        })
        .catch(e => {console.log(e)})
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
          {this.props.userType === 'student' ? <Tag>学生</Tag> :
            <>
              {this.props.userType === 'teacher' ? <Tag>教师</Tag> :
                <Tag>未知角色</Tag>
              }
            </>
          }
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
            <Col xxl={14} xl={10} lg={8} md={6} sm={6} xs={10}>
              <Link to="/home">
                <Image height={60} style={{margin: '2px'}} src={logo} preview={false}/>
              </Link>
            </Col>
            <Col xxl={6} xl={10} lg={10} md={10} sm={10} xs={3}>
              <Menu
                theme="light"
                mode="horizontal"
                style={{border: 0}}
                defaultSelectedKeys={['home']}
                selectedKeys={[menu]}
                onClick={e=>this.changeMenu(e)}
              >
                <Menu.Item key="home">
                  <Link to="/home">
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
            <Col xxl={4} xl={4} lg={6} md={8} sm={8} xs={11}>
              {
                <>
                  <span style={{float: 'left', marginRight: '20px'}}>
                    <Link to="/message">
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
            <Route exact path="/" render={() => (
              <Redirect to="/home"/>
            )}/>

            <Route exact path="/home" component={Home}/>
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

function mapStateToProps(state) {
  return {
    userType: state.get("userType").get("userType")
  }
}

const setType = (userType) => {
  return {
    type: 'set',
    userType: userType
  }
}

function mapDispatchToProps(dispatch) {
  return {
    setUserType: (userType)=>dispatch(setType(userType))
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(HeaderLayout);
