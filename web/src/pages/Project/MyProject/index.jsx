import React from 'react';
import DocumentTitle from 'react-document-title';
import {Link, Redirect, Route, Switch} from 'react-router-dom'
import {CheckCircleOutlined, CheckOutlined, CopyOutlined, HighlightOutlined, SyncOutlined} from "@ant-design/icons";
import {connect} from "react-redux";

import PublishedProject from "../PublishedProject";
import EditingProject from "../EditingProject";
import FinishedProject from "../FinishedProject";
import LearningProject from "../LearningProject";
import {Affix, Button, Layout, Menu} from "antd";
import ProjectApi from "../../../api/ProjectApi";

class MyProject extends React.PureComponent {

  createProject = e => {
    ProjectApi.createProject()
      .then((res) => {
        if (res.data.code === 200) {
          window.open(`/project/${res.data.data}/info/edit`)
        }
      })
      .catch((e) => {
        console.log(e)
      })
  }

  render() {
    return (
      <DocumentTitle title="My Project">
        <Layout style={{margin: '20px'}}>
          <Affix offsetTop={0}>
            <Layout.Sider
              breakpoint="lg"
              collapsedWidth="0"
              theme="light"
              style={{backgroundColor: '#f2f4f5'}}
            >
              {this.props.userType === 'teacher' ?
                <Menu
                  defaultSelectedKeys={['published']}
                  className="menu-bar"
                  mode="inline"
                >
                  <Menu.Item key="published" icon={<CheckCircleOutlined/>}>
                    <Link to="/my-project/published">
                      已发布
                    </Link>
                  </Menu.Item>
                  <Menu.Item key="editing" icon={<HighlightOutlined/>}>
                    <Link to="/my-project/editing">
                      未发布
                    </Link>
                  </Menu.Item>
                  <Menu.Item key="finished" icon={<CopyOutlined/>}>
                    <Link to="/my-project/finished">
                      已结束
                    </Link>
                  </Menu.Item>
                </Menu>
                :
                <Menu
                  defaultSelectedKeys={['learning']}
                  className="menu-bar"
                  mode="inline"
                >
                  <Menu.Item key="learning" icon={<SyncOutlined/>}>
                    <Link to="/my-project/learning">
                      进行中
                    </Link>
                  </Menu.Item>
                  <Menu.Item key="finished" icon={<CheckOutlined/>}>
                    <Link to="/my-project/finished">
                      已完成
                    </Link>
                  </Menu.Item>
                </Menu>
              }
              {this.props.userType === 'teacher' ?
                <Button
                  type='primary'
                  shape='round'
                  size="large"
                  onClick={this.createProject}
                  style={{
                    width: '180px',
                    margin: '10px',
                  }}
                >创建项目</Button>
                :
                null
              }
            </Layout.Sider>
          </Affix>
          <Layout.Content style={{marginLeft: '10px'}}>
            {this.props.userType === 'teacher' ?
              <Switch>
                <Route exact path="/my-project" render={() => (
                  <Redirect to="/my-project/published"/>
                )}/>
                <Route exact path="/my-project/published" component={PublishedProject}/>
                <Route exact path="/my-project/editing" component={EditingProject}/>
                <Route exact path="/my-project/finished" component={FinishedProject}/>
              </Switch>
              :
              <Switch>
                <Route exact path="/my-project" render={() => (
                  <Redirect to="/my-project/learning"/>
                )}/>
                <Route exact path="/my-project/learning" component={LearningProject}/>
                <Route exact path="/my-project/finished" component={FinishedProject}/>
              </Switch>
            }
          </Layout.Content>
        </Layout>
      </DocumentTitle>
    );
  }
}

function mapStateToProps(state) {
  return {
    userType: state.get("userType").get("userType")
  }
}

export default connect(mapStateToProps, null)(MyProject);
