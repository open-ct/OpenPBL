// deprecated

import React from 'react';
import {Affix, Button, Layout, Menu} from 'antd';
import {
  CheckCircleOutlined,
  CheckOutlined,
  CopyOutlined,
  HighlightOutlined,
  SyncOutlined,
} from '@ant-design/icons';
import {Switch, Route, Link} from 'react-router-dom'

import LearningProject from '../LearningProject';
import './menu-bar.less';

import FinishedProject from '../FinishedProject';
import PublishedProject from "../PublishedProject";
import EditingProject from "../EditingProject";
import localStorage from "localStorage";
import ProjectApi from "../../../api/ProjectApi";


class MenuBar extends React.PureComponent {
  state = {
    menu: this.props.menu,
    top: 10,
    type: localStorage.getItem('type'),
  }

  handleClick = (e) => {
    let key = e.key;
    this.setState({menu: key});
  }

  createProject = e => {
    ProjectApi.createProject()
      .then((res)=>{
        if (res.data.code === 200) {
          window.open(`/project/${res.data.data}/info/edit`)
        }
      })
      .catch((e)=>{console.log(e)})
  }

  render() {
    const {menu, top, type} = this.state;

    return (
      <Layout style={{minHeight: '90.5vh'}}>
        <Affix offsetTop={top}>
          <Layout.Sider
            breakpoint="lg"
            collapsedWidth="0"
            style={{height: '100%', backgroundColor: '#f2f4f5'}}
          >
            <Menu
              defaultSelectedKeys={['published']}
              className="menu-bar"
              mode="inline"
            >
              {type === 'teacher' ?
                <>
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
                </>
                :
                <>
                  <Menu.Item key="learning" icon={<SyncOutlined/>}>
                    进行中
                  </Menu.Item>
                  <Menu.Item key="finished" icon={<CheckOutlined/>}>
                    已完成
                  </Menu.Item>
                </>
              }
            </Menu>
            {type === 'teacher' ?
              <Button
                style={{marginLeft: '5px', marginTop: '10px'}}
                size='large'
                type='primary'
                block
                shape='round'
                onClick={this.createProject}
              >创建项目</Button>
              :
              null
            }
          </Layout.Sider>
        </Affix>

          {/*{type === 'teacher' ?
            <Content className="project-content">
              {menu === 'published' ? <PublishedProject/> : null}
              {menu === 'editing' ? <EditingProject/> : null}
              {menu === 'finished' ? <FinishedProject/> : null}
            </Content>
            :
            <Content className="project-content">
              {menu === 'learning' ? <LearningProject/> : null}
              {menu === 'finished' ? <FinishedProject/> : null}
            </Content>
          }*/}
      </Layout>
    );
  }
}

export default MenuBar;
