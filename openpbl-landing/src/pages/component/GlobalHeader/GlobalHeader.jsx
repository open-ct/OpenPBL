import React from "react";
import {Col, Layout, Menu, Row} from "antd";
import {Link} from "react-router-dom";

import './global-header.less'

const {Header} = Layout;

class GlobalHeader extends React.PureComponent {
  state = {
    current: this.props.current
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
            <Col xxl={8} xl={12} lg={14} md={16} sm={14} xs={8}>
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
                <Menu.Item key="login">
                  <Link to="/user/login">
                    用户登录
                  </Link>
                </Menu.Item>
                <Menu.Item key="register">
                  <Link to="/user/register">
                    用户注册
                  </Link>
                </Menu.Item>
                <Menu.Item key="bbs">
                  <a href="https://bbs.open-ct.com">
                    在线论坛
                  </a>
                </Menu.Item>
              </Menu>
            </Col>
          </Row>
        </Header>
      </div>
    );
  }
}

export default GlobalHeader;
