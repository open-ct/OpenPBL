import React from 'react';
import DocumentTitle from 'react-document-title';
import {Avatar, Button, Card, Col, Divider, Image, Menu, PageHeader, Popconfirm, Row, message, BackTop} from 'antd';
import QueueAnim from 'rc-queue-anim';
import localStorage from 'localStorage';
import {Link} from 'react-router-dom';

import ProjectIntroduce from './component/ProjectIntroduce';
import ProjectOutline from './component/ProjectOutline';
import ProjectEvaluation from './component/ProjectEvaluation';
import ProjectComment from './component/ProjectComment';
import ProjectApi from "../../../api/ProjectApi";
import StudentApi from "../../../api/StudentApi";
import StudentAdmin from "./component/StudentAdmin";
import {getUser} from "../../User/Auth/Auth";
import {DeleteOutlined} from "@ant-design/icons";
import util from "../component/Util"
import StudentEvidence from "../Evidence/component/StudentEvidence";

class ProjectInfo extends React.PureComponent {
  constructor(props) {
    super(props);
    const pid = this.props.match.params.id;
    let url = new URLSearchParams(this.props.location.search)
    let menu = url.get("menu")
    console.log(menu)
    if (menu === undefined || menu === null) {
      menu = 'project-introduce'
    }
    this.state = {
      pid: pid,
      project: {},
      teacher: {},
      menu: menu,
      type: localStorage.getItem('type'),
      lastLearn: {}
    };
  }

  componentDidMount() {
    this.loadProjectDetail()
    this.loadLastLearn()
  }
  loadProjectDetail = () => {
    ProjectApi.getProjectDetail(this.state.pid)
      .then((res) => {
        this.setState({
          project: res.data.project
        })
        this.loadTeacherInfo(res.data.project.teacherId)
      })
      .catch((e) => {
        console.log(e)
      })
  }
  loadLastLearn = () => {
    StudentApi.getLastLearnSection(this.state.pid)
      .then(res=>{
        if (res.data.code === 200) {
          this.setState({
            lastLearn: res.data.data
          })
        }
      })
      .catch(e=>{console.log(e)})
  }
  loadTeacherInfo = (teacherId) => {
    getUser(teacherId)
      .then((res) => {
        if (res.status === 200) {
          this.setState({
            teacher: res.data
          })
        }
      })
      .catch((e)=>{console.log(e)})
  }

  handleClick = (e) => {
    if (this.state.type==='student' && e.key === "student-evidence" && !this.state.project.learning) {
      message.warn("请先加入学习")
      return
    }
    this.setState({
      menu: e.key,
    });
  }
  back = e => {
    window.location.href = '/my-project'
  }
  learnProject = e => {
    StudentApi.learnProject(this.state.pid)
      .then((res) => {
        if (res.data.code === 200) {
          let p = Object.assign({}, this.state.project)
          p.learning = true
          this.setState({
            project: p
          })
        }
      })
      .catch((e) => {
        console.log(e)
      })
  }
  publishProject = e => {
    ProjectApi.publishProject(this.state.pid)
      .then((res) => {
        if (res.data.code === 200) {
          let p = Object.assign({}, this.state.project)
          p.published = true
          this.setState({
            project: p
          })
        }
      })
      .catch((e) => {
        console.log(e)
      })
  }
  closeProject = e => {
    ProjectApi.closeProject(this.state.pid)
      .then(res=>{
        if (res.data.code === 200) {
          let p = Object.assign({}, this.state.project)
          p.closed = true
          this.setState({
            project: p
          })
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e=>{console.log(e)})
  }
  exitProject = e => {
    StudentApi.exitProject(this.state.pid)
      .then(res => {
        if (res.data.code === 200) {
          window.location.href = "/my-project"
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e=>{console.log(e)})
  }
  deleteProject = e => {
    ProjectApi.deleteProject(this.state.pid)
      .then(res=>{
        if (res.data.code === 200) {
          window.location.href = "/my-project"
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e=>{console.log(e)})
  }

  setProject = project => {
    this.setState({
      project: project
    })
  }

  render() {
    const {project, teacher, menu, type, pid, lastLearn} = this.state;

    const teacherBt = (
      <div style={{float: 'right'}}>
        {project.published && !project.closed ?
          <Popconfirm
            title="确定结束项目?"
            placement="topRight"
            onConfirm={this.closeProject}
            onCancel={this.cancelCloseProject}
            okText="确定"
            cancelText="取消"
          >
            <Button
              shape="round"
              size="middle"
              danger
              style={{margin: '5px'}}
            >
              结束项目
            </Button>
          </Popconfirm>
          : null }
        {!project.published ?
          <div>
            <Popconfirm
                title="确定发布项目?"
                placement="topRight"
                onConfirm={this.publishProject}
                okText="确定"
                cancelText="取消"
            >
              <Button
                shape="round"
                size="middle"
                style={{margin: '5px'}}
              >
                发布项目
              </Button>
            </Popconfirm>
            <Link to={`/project/${pid}/info/edit`} target="_blank">
              <Button
                shape="round"
                size="middle"
                style={{margin: '5px'}}
              >
              编辑项目
              </Button>
            </Link>
            <Popconfirm
              title="确定删除项目？删除后不能恢复"
              onConfirm={this.deleteProject}
              placement="topLeft"
            >
              <Button
                shape="circle"
                size="middle"
                type="danger"
                style={{marginTop: '5px', marginLeft: '20px', marginBottom: '5px', marginRight: '5px'}}
                icon={<DeleteOutlined/>}
              />
            </Popconfirm>
          </div> : null }
        {project.closed ?
          <div>
            <Popconfirm
              title="确定删除项目？删除后不能恢复"
              onConfirm={this.deleteProject}
              placement="topLeft"
            >
              <Button
                shape="circle"
                size="middle"
                type="danger"
                style={{marginTop: '5px', marginLeft: '20px', marginBottom: '5px', marginRight: '5px'}}
                icon={<DeleteOutlined/>}
              />
            </Popconfirm>
          </div>
          : null }
      </div>
    )
    const studentBt = (
      <div style={{float: 'right'}}>
        {project.learning ?
          <>
            <Link to={`/project/${project.id}/section/${lastLearn.id}/preview?back=/project/${project.id}/info`}>

              {lastLearn.last ?
                <span>
                  上次学到：{util.FormatSectionName(lastLearn.sectionName, lastLearn.chapterNumber, lastLearn.sectionNumber)}&nbsp;&nbsp;
                  <span style={{color: 'gray'}}>
                    {util.FilterMoment(lastLearn.exitAt)}&nbsp;&nbsp;
                  </span>
                </span>
                : null
              }
              <Button
                shape="round"
                size="middle"
                style={{margin: '5px'}}
              >继续学习
              </Button>
            </Link>
            <Popconfirm title="确认退出项目？" onConfirm={this.exitProject} placement="topRight" >
              <Button
                type="danger"
                shape="round"
                size="middle"
                style={{margin: '5px'}}
              >退出项目
              </Button>
            </Popconfirm>
          </>
          :
          <Button
            type="primary"
            shape="round"
            size="middle"
            style={{margin: '5px'}}
            onClick={this.learnProject}
          >加入学习
          </Button>
        }
      </div>
    )

    return (
      <DocumentTitle title="Project">
        <div style={{backgroundColor: '#f2f4f5'}}>
          <BackTop />

          <PageHeader
            className="site-page-header"
            onBack={() => this.back()}
            title="返回"
            subTitle="项目详情"
          />
          <div
            style={{
              width: '100%',
              height: '100%',
              display: 'flex',
              paddingRight: '20px',
              paddingLeft: '20px',
              paddingBottom: '20px',
              marginTop: '20px'
            }}
          >
            <div
              style={{
                margin: 'auto',
                width: '100%',
                maxWidth: '1200px',
                minHeight: '100vh',
              }}
            >
              <QueueAnim>
                <div key="1">
                  <Card hoverable style={{textAlign: 'left'}}>
                    <Row>
                      <Col span={6}>
                        <Image
                          alt="example"
                          src={project.image}
                          placeholder
                          preview={false}
                        />
                      </Col>
                      <Col span={1}>&nbsp;</Col>
                      <Col flex="auto">
                        <p style={{fontSize: '20px'}}>{project.projectTitle}</p>
                        <p
                          style={{fontSize: '14px', color: 'gray'}}
                        >发布时间：{util.FilterTime(project.createAt)}
                        </p>
                        <div>
                          <span>{project.readNum}&nbsp;人看过</span>
                          <Divider type="vertical"/>
                          <span>{project.joinNum}&nbsp;人加入学习</span>
                          <Divider type="vertical"/>
                          <span>授课教师：&nbsp;
                            <Avatar src={teacher.avatar} />&nbsp;&nbsp;
                            {teacher.displayName}
                          </span>
                        </div>
                        <br/>
                        {type === 'student' ? studentBt : teacherBt}
                      </Col>
                    </Row>
                  </Card>
                </div>
                <div key="2">
                  <Menu
                    onClick={this.handleClick}
                    selectedKeys={menu}
                    mode="horizontal"
                    style={{marginTop: '20px'}}
                  >
                    <Menu.Item key="project-introduce">项目信息</Menu.Item>
                    <Menu.Item key="project-outline">项目大纲</Menu.Item>
{/*
                    <Menu.Item key="project-comment">评论区</Menu.Item>
*/}
                    <Menu.Item key="project-evaluation">评价方案</Menu.Item>

                    {type === 'teacher' ? <Menu.Item key="student-admin">学生管理</Menu.Item>
                      : null}
                    {type === 'student' ? <Menu.Item key="student-evidence">证据收集</Menu.Item>
                      : null}
                  </Menu>
                  <div style={{
                    backgroundColor: 'white',
                    padding: '20px',
                  }}
                  >
                    {menu === 'project-introduce' ? <ProjectIntroduce project={project}/> : null}
                    {menu === 'project-outline' ? <ProjectOutline project={project}/> : null}
                    {menu === 'project-comment' ? <ProjectComment project={project}/> : null}
                    {menu === 'project-evaluation' ?
                      <ProjectEvaluation
                        project={project}
                        setProject={this.setProject}
                        loadProjectDetail={this.loadProjectDetail}
                      />
                      : null
                    }

                    {menu === 'student-admin' ? <StudentAdmin project={project}/> : null}

                    {menu === 'student-evidence' ? <StudentEvidence project={project}/> : null}

                  </div>
                </div>
              </QueueAnim>
            </div>
          </div>
        </div>
      </DocumentTitle>
    );
  }
}

export default ProjectInfo
