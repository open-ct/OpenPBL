import React from 'react';
import {Affix, Button, Layout, Menu} from 'antd';
import {
  AlignLeftOutlined,
  CheckCircleOutlined,
  CheckOutlined,
  CopyOutlined,
  HighlightOutlined,
  SyncOutlined,
  UserOutlined
} from '@ant-design/icons';

import LearningProject from '../LearningProject';
import './menu-bar.less';

import PublicProject from '../PublicProject';
import FinishedProject from '../FinishedProject';
import PublishedProject from "../PublishedProject";
import EditingProject from "../EditingProject";
import localStorage from "localStorage";
import project from "../../../api/project";

const {SubMenu} = Menu;
const {Sider, Content} = Layout;

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
  gotoEditPage = e => {
    this.props.history.push('/project/edit/')
  }

  createProject = e => {
    project.createProject()
      .then((res)=>{
        if (res.data.id) {
          window.open('/project/edit/info/' + res.data.id)
        }
      })
      .catch((e)=>{console.log(e)})
  }

  render() {
    const {menu, top, type} = this.state;

    return (
      <Layout style={{minHeight: '90vh'}}>
        <Affix offsetTop={top}>
          <Sider
            breakpoint="lg"
            collapsedWidth="0"
            style={{height: '100%', backgroundColor: '#f2f4f5'}}
          >
            <Menu
              onClick={this.handleClick}
              className="menu-bar"
              defaultOpenKeys={['my-project']}
              defaultSelectedKeys={[menu]}
              mode="inline"
            >
              <Menu.Item className="menu-item" key="public" icon={<AlignLeftOutlined/>}>
                公共项目
              </Menu.Item>
              <SubMenu className="menu-item" key="my-project" title="我的项目" icon={<UserOutlined/>}>
                {type === 'teacher' ?
                  <Menu.Item key="published" icon={<CheckCircleOutlined/>}>
                    已发布
                  </Menu.Item>
                  :
                  <Menu.Item key="learning" icon={<SyncOutlined/>}>
                    进行中
                  </Menu.Item>
                }
                {type === 'teacher' ?
                  <Menu.Item key="editing" icon={<HighlightOutlined/>}>
                    未发布
                  </Menu.Item>
                  :
                  <Menu.Item key="finished" icon={<CheckOutlined/>}>
                    已完成
                  </Menu.Item>
                }
                {type === 'teacher' ?
                  <Menu.Item key="finished" icon={<CopyOutlined/>}>
                    已结束
                  </Menu.Item>
                  :
                  null
                }
              </SubMenu>
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
          </Sider>
        </Affix>

        <Layout>
          {type === 'teacher' ?
            <Content className="project-content">
              {menu === 'public' ? <PublicProject/> : null}
              {menu === 'published' ? <PublishedProject/> : null}
              {menu === 'editing' ? <EditingProject/> : null}
              {menu === 'finished' ? <FinishedProject/> : null}
            </Content>
            :
            <Content className="project-content">
              {menu === 'public' ? <PublicProject/> : null}
              {menu === 'learning' ? <LearningProject/> : null}
              {menu === 'finished' ? <FinishedProject/> : null}
            </Content>
          }
        </Layout>
      </Layout>
    );
  }
}

export default MenuBar;
