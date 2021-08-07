import React from 'react';
import DocumentTitle from 'react-document-title';
import {enquireScreen} from 'enquire-js';
import {Avatar, Button, Card, Form, Input, Tabs} from 'antd';
import {UserOutlined} from '@ant-design/icons';
import localStorage from 'localStorage';
import QueueAnim from 'rc-queue-anim';

import './register.less';
import GlobalHeader from '../../component/GlobalHeader/GlobalHeader';

const {TabPane} = Tabs;

class Register extends React.PureComponent {
  state = {
    type: this.props.match.params.type,
  };

  componentDidMount() {
    enquireScreen((b) => {
      this.setState({
        isMobile: !!b,
      });
    });
  }

  onFinish = (v) => {
    console.log(v, this.state.type);
    localStorage.setItem('type', this.state.type);
    window.location = '#/project';
  }
  onFinishFailed = () => {

  }
  changeTab = (k) => {
    this.setState({
      type: k,
    });
  }

  render() {
    const {type} = this.state;
    return (
      <DocumentTitle title="Login">
        <div style={{height: '100vh'}}>
          <GlobalHeader/>
          <QueueAnim delay={100}>
            <Card className="login-card" key="1">
              <Tabs defaultActiveKey={type} onChange={this.changeTab}>
                <TabPane tab="学生登录" key="student"/>
                <TabPane tab="教师登录" key="teacher"/>
              </Tabs>
              <Avatar
                icon={<UserOutlined/>}
                size={64}
              />
              <Form
                name="basic"
                initialValues={{username: 'username', password: '123456'}}
                onFinish={this.onFinish}
                onFinishFailed={this.onFinishFailed}
                style={{marginTop: '20px'}}
              >
                <Form.Item
                  label=""
                  name="username"
                  rules={[{required: true, message: '请输入用户名'}]}
                >
                  <Input placeholder="请输入用户名" size="large" name="username"/>
                </Form.Item>
                <Form.Item
                  label=""
                  name="password"
                  rules={[{required: true, message: '清输入密码'}]}
                >
                  <Input.Password placeholder="清输入密码" size="large" name="password"/>
                </Form.Item>
                <Form.Item>
                  <Button type="primary" htmlType="submit" shape="round" size="large" block>
                    登录
                  </Button>
                </Form.Item>
                <Form.Item>
                  <Button size="large" shape="round" block>
                    注册
                  </Button>
                </Form.Item>
              </Form>
            </Card>
          </QueueAnim>
        </div>
      </DocumentTitle>
    );
  }
}


export default Register;
