import React from 'react';
import {Link} from 'react-router-dom';
import {Col, Menu, Popover, Row} from 'antd';
import {enquireScreen} from 'enquire-js';

// const LOGO_URL = 'https://cdn.open-ct.com/logo/openct.png';
// const LOGO_URL = require('../../assets/logo.svg');

class GlobalHeader extends React.Component {
  state = {
    menuVisible: false,
    menuMode: 'horizontal',
    current: 'landing',
  };

  componentDidMount() {
    enquireScreen((b) => {
      console.log(b);
      this.setState({menuMode: b ? 'inline' : 'horizontal'});
    });
  }

  handleClick = e => {
    this.setState({current: e.key});
  };

  render() {
    const {menuMode, menuVisible, current} = this.state;

    const menu = (
      <Menu
        mode={menuMode}
        id="nav"
        key="nav"
        defaultSelectedKeys={[current]}
        onClick={this.handleClick}
      >
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
        <Menu.Item key="student">
          <Link to="/user/login">
            学生登录
          </Link>
        </Menu.Item>
        <Menu.Item key="teacher">
          <Link to="/user/login">
            教师登录
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
    );

    return (
      <div id="header" className="header">
        {menuMode === 'inline' ? (
          <Popover
            overlayClassName="popover-menu"
            placement="bottomRight"
            content={menu}
            trigger="click"
            visible={menuVisible}
            arrowPointAtCenter
            onVisibleChange={this.onMenuVisibleChange}
          >


          </Popover>
        ) : null}
        <Row>
          <Col xxl={4} xl={5} lg={8} md={8} sm={24} xs={24}>
            <div>
              <Link to="/landing">
                <img src="assets/logo.svg" style={{height: '64px', marginLeft: '20px'}} alt="logo"/>
                <span style={{fontWeight: "bold", fontSize: '30px', top: '50%'}}>OpenCT</span>
              </Link>
            </div>
          </Col>
          <Col xxl={20} xl={19} lg={16} md={16} sm={0} xs={0}>
            <div className="header-meta">
              {/*<div id="preview">*/}
              {/*  <a*/}
              {/*    id="preview-button"*/}
              {/*    target="_blank"*/}
              {/*    href="/register"*/}
              {/*    rel="noopener noreferrer"*/}
              {/*  >*/}
              {/*    <Button icon="eye-o">*/}
              {/*      注册*/}
              {/*    </Button>*/}
              {/*  </a>*/}
              {/*</div>*/}
              {menuMode === 'horizontal' ? <div id="menu">{menu}</div> : null}
            </div>
          </Col>
        </Row>
      </div>
    );
  }
}

export default GlobalHeader;
