import React from 'react';
import DocumentTitle from 'react-document-title';
import {Button, Card, Col, Divider, Image, Menu, PageHeader, Popconfirm, Row} from 'antd';
import QueueAnim from 'rc-queue-anim';
import localStorage from 'localStorage';
import {Link} from 'react-router-dom';

import ProjectIntroduce from './component/ProjectIntroduce';
import ProjectOutline from './component/ProjectOutline';
import ProjectEvaluation from './component/ProjectEvaluation';
import ProjectComment from './component/ProjectComment';
import project from "../../../api/project";
import student from "../../../api/student";
import teacher from "../../../api/teacher";
import StudentAdmin from "./component/StudentAdmin";
import SubmitFiles from "./component/SubmitFiles";


class ProjectInfo extends React.PureComponent {
  constructor(props) {
    super(props);
    const pid = this.props.match.params.id;

    // TODO axios get project info
    const projectInfo = {
      id: pid,
      image: '',
      projectIntroduce: '',
      projectTitle: '',
      name: '',
      teacherId: '',
      subjects: '',
      skills: '',
      createAt: '',
      readNum: 0,
      joinNum: 0,
      learning: null,
    };
    this.state = {
      pid: pid,
      project: projectInfo,
      menu: 'project-introduce',
      type: localStorage.getItem('type'),
      uid: localStorage.getItem('uid')
    };
  }

  componentDidMount() {
    if (this.state.type === 'teacher') {
      project.getProjectDetailForTeacher(this.state.pid)
        .then((res) => {
          this.setState({
            project: res.data.project
          })
        })
        .catch((e) => {
          console.log(e)
        })
    } else if (this.state.type === 'student') {
      project.getProjectDetailForStudent(this.state.pid)
        .then((res) => {
          this.setState({
            project: res.data.project
          })
        })
        .catch((e) => {
          console.log(e)
        })
    }
  }

  handleClick = (e) => {
    this.setState({
      menu: e.key,
    });
  }
  back = e => {
    console.log('back')
    this.props.history.push('/project')
  }
  learnProject = e => {
    student.learnProject(this.state.uid, this.state.pid)
      .then((res) => {
        if (res.data.result) {
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
    teacher.publishProject(this.state.pid)
      .then((res) => {
        if (res.data.result) {
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

  confirmCloseProject = e => {
    // TODO axios close project
  }
  cancelCloseProject = e => {
  }

  filterTime = t => {
    let d = new Date(t)
    return d.getFullYear().toString() + '-' +
      (d.getMonth() + 1).toString() + '-' +
      d.getDate().toString() + ' ' +
      d.getHours().toString() + ':' +
      d.getMinutes().toString() + ':' +
      d.getSeconds().toString()
  }


  render() {
    const {project, menu, type, pid } = this.state;

    const teacherBt = (
      <div style={{float: 'right'}}>
        {project.published ?
          <Popconfirm
            title="确定结束项目?"
            placement="topRight"
            onConfirm={this.confirmCloseProject}
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
          :
          <div>
            <Button
              shape="round"
              size="middle"
              style={{margin: '5px'}}
              onClick={this.publishProject}
            >
              发布项目
            </Button>
            <Link to={`/project/edit/info/${pid}`} target="_blank">
              <Button
                shape="round"
                size="middle"
                style={{margin: '5px'}}
              >
              编辑项目
              </Button>
            </Link>
          </div>
        }
      </div>
    )
    const studentBt = (
      <div style={{float: 'right'}}>
        {project.learning ?
          <Link to={`/project/learning/${project.id}`}>
            <Button
              shape="round"
              size="middle"
              style={{margin: '5px'}}
            >继续学习
            </Button>
          </Link>
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
                        <p style={{
                          fontSize: '13px',
                          color: 'gray',
                        }}
                        >发布时间：{this.filterTime(project.createAt)}
                        </p>
                        <div>
                          <span>{project.readNum}&nbsp;人看过</span>
                          <Divider type="vertical"/>
                          <span>{project.readNum}&nbsp;人加入学习</span>
                          <Divider type="vertical"/>
                          <span>授课教师：{project.name}</span>
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
                    {type === 'student' ? <Menu.Item key="submit-files">证据收集</Menu.Item>
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
                    {menu === 'project-evaluation' ? <ProjectEvaluation project={project}/> : null}

                    {menu === 'student-admin' ? <StudentAdmin project={project}/> : null}

                    {menu === 'submit-files' ? <SubmitFiles project={project}/> : null}

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

export default ProjectInfo;
