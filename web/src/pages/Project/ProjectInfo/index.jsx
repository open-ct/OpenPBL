import React from 'react';
import DocumentTitle from 'react-document-title';
import {
  Avatar,
  BackTop,
  Button,
  Card,
  Col,
  Divider,
  Image,
  Menu,
  message,
  PageHeader,
  Popconfirm,
  Row,
  Tag,
  Tooltip
} from 'antd';
import QueueAnim from 'rc-queue-anim';
import {Link} from 'react-router-dom';
import {DeleteOutlined, StarFilled, StarOutlined} from "@ant-design/icons";


import ProjectIntroduce from './component/ProjectIntroduce';
import ProjectOutline from './component/ProjectOutline';
import ProjectEvaluation from './component/ProjectEvaluation';
import ProjectComment from './component/ProjectComment';
import ProjectApi from "../../../api/ProjectApi";
import StudentApi from "../../../api/StudentApi";
import StudentAdmin from "./component/StudentAdmin";
import {getUser} from "../../User/Auth/Auth";
import util from "../../component/Util"
import StudentEvidence from "../Evidence/component/StudentEvidence";

class ProjectInfo extends React.PureComponent {
  constructor(props) {
    super(props);
    const pid = this.props.match.params.id;
    let url = new URLSearchParams(this.props.location.search)
    let menu = url.get("menu")
    if (menu === undefined || menu === null) {
      menu = 'project-introduce'
    }
    this.state = {
      pid: pid,
      project: {},
      teacher: {},
      menu: menu,
      lastLearn: {},
      favBtLoading: false,
      learnBtLoading: false,
      exitBtLoading: false,
      closeBtLoading: false,
      cloneBtLoading: false,
      deleteBtLoading: false
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
      .then(res => {
        if (res.data.code === 200) {
          this.setState({
            lastLearn: res.data.data
          })
        }
      })
      .catch(e => {
        console.log(e)
      })
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
      .catch((e) => {
        console.log(e)
      })
  }

  handleClick = (e) => {
    if (this.props.account.tag === '??????' && e.key === "student-evidence" && !this.state.project.learning) {
      message.warn("??????????????????")
      return
    }
    this.setState({
      menu: e.key,
    });
  }
  back = e => {
    if (this.props.account.tag === '??????') {
      window.location.href = '/home/my-project/published'
    } else {
      window.location.href = '/home/my-project/learning'
    }
  }
  learnProject = e => {
    this.setState({
      learnBtLoading: true
    })
    StudentApi.learnProject(this.state.pid)
      .then((res) => {
        this.setState({
          learnBtLoading: true
        })
        if (res.data.code === 200) {
          this.loadProjectDetail()
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
          this.loadProjectDetail()
        }
      })
      .catch((e) => {
        console.log(e)
      })
  }
  closeProject = e => {
    this.setState({
      closeBtLoading: true
    })
    ProjectApi.closeProject(this.state.pid)
      .then(res => {
        this.setState({
          closeBtLoading: false
        })
        if (res.data.code === 200) {
          this.loadProjectDetail()
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e => {
        console.log(e)
      })
  }
  exitProject = e => {
    this.setState({
      exitBtLoading: true
    })
    StudentApi.exitProject(this.state.pid)
      .then(res => {
        this.setState({
          exitBtLoading: false
        })
        if (res.data.code === 200) {
          if (this.props.account.tag === '??????') {
            window.location.href = '/home/my-project/published'
          } else {
            window.location.href = '/home/my-project/learning'
          }
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e => {
        console.log(e)
      })
  }
  deleteProject = e => {
    this.setState({deleteBtLoading: true})
    ProjectApi.deleteProject(this.state.pid)
      .then(res => {
        this.setState({deleteBtLoading: false})
        if (res.data.code === 200) {
          if (this.props.account.tag === '??????') {
            window.location.href = '/home/my-project/published'
          } else {
            window.location.href = '/home/my-project/learning'
          }
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e => {
        console.log(e)
      })
  }
  addFavourite = e => {
    this.setState({
      favBtLoading: true
    })
    ProjectApi.addFavourite(this.state.pid)
      .then(res => {
        this.setState({
          favBtLoading: false
        })
        if (res.data.code === 200) {
          this.loadProjectDetail()
          message.success(res.data.msg)
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e => {
        console.log(e)
      })
  }
  removeFavourite = e => {
    this.setState({
      favBtLoading: true
    })
    ProjectApi.removeFavourite(this.state.pid)
      .then(res => {
        this.setState({
          favBtLoading: false
        })
        if (res.data.code === 200) {
          this.loadProjectDetail()
          message.success(res.data.msg)
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e => {
        console.log(e)
      })
  }

  setProject = project => {
    this.setState({
      project: project
    })
  }
  projectState = () => {
    if (this.state.project.published) {
      if (this.state.project.closed) {
        return "( ????????? )"
      } else {
        return "( ????????? )"
      }
    } else {
      return "( ????????? )"
    }
  }
  cloneProject = () => {
    this.setState({
      cloneBtLoading: true
    })
    ProjectApi.cloneProject(this.state.pid)
      .then(res => {
        this.setState({
          cloneBtLoading: false
        })
        if (res.data.code === 200) {
          message.success(res.data.msg)
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e => {
        console.log(e)
      })
  }

  render() {
    const {
      project,
      teacher,
      menu,
      pid,
      lastLearn,
      favBtLoading,
      learnBtLoading,
      exitBtLoading,
      closeBtLoading,
      cloneBtLoading,
      deleteBtLoading
    } = this.state;

    const teacherBt = (
      <div style={{float: 'right'}}>
        {project.published && !project.closed ?
          <>
            <Popconfirm
              title="???????????????????"
              placement="topRight"
              onConfirm={this.closeProject}
              okText="??????"
              cancelText="??????"
            >
              <Button
                shape="round"
                size="middle"
                danger
                style={{margin: '5px'}}
                loading={closeBtLoading}
              >
                ????????????
              </Button>
            </Popconfirm>
            <Popconfirm
              title="???????????????????"
              placement="topRight"
              onConfirm={this.cloneProject}
              okText="??????"
              cancelText="??????"
            >
              <Button
                shape="round"
                size="middle"
                style={{margin: '5px'}}
                loading={cloneBtLoading}
              >
                ????????????
              </Button>
            </Popconfirm>
          </>
          : null}
        {!project.published ?
          <div>
            <Popconfirm
              title="???????????????????"
              placement="topRight"
              onConfirm={this.publishProject}
              okText="??????"
              cancelText="??????"
            >
              <Button
                shape="round"
                size="middle"
                style={{margin: '5px'}}
              >
                ????????????
              </Button>
            </Popconfirm>
            <Link to={`/home/project/${pid}/info/edit`} target="_blank">
              <Button
                shape="round"
                size="middle"
                style={{margin: '5px'}}
              >
                ????????????
              </Button>
            </Link>
            <Popconfirm
              title="???????????????????"
              placement="topRight"
              onConfirm={this.cloneProject}
              okText="??????"
              cancelText="??????"
            >
              <Button
                shape="round"
                size="middle"
                style={{margin: '5px'}}
                loading={cloneBtLoading}
              >
                ????????????
              </Button>
            </Popconfirm>
            <Popconfirm
              title="??????????????????????????????????????????"
              onConfirm={this.deleteProject}
              placement="topRight"
            >
              <Button
                shape="circle"
                size="middle"
                type="danger"
                loading={deleteBtLoading}
                style={{marginTop: '5px', marginLeft: '20px', marginBottom: '5px', marginRight: '5px'}}
                icon={<DeleteOutlined/>}
              />
            </Popconfirm>
          </div> : null}
        {project.closed ?
          <div>
            <Popconfirm
              title="???????????????????"
              placement="topRight"
              onConfirm={this.cloneProject}
              okText="??????"
              cancelText="??????"
            >
              <Button
                shape="round"
                size="middle"
                style={{margin: '5px'}}
                loading={cloneBtLoading}
              >
                ????????????
              </Button>
            </Popconfirm>
            <Popconfirm
              title="??????????????????????????????????????????"
              onConfirm={this.deleteProject}
              placement="topRight"
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
          : null}
      </div>
    )
    const studentBt = (
      <div style={{float: 'right'}}>
        {project.learning ?
          <>
            <Link to={`/home/project/${project.id}/section/${lastLearn.id}/preview?back=/project/${project.id}/info`}>

              {lastLearn.last ?
                <span>
                  ???????????????{util.FormatSectionName(lastLearn.sectionName, lastLearn.chapterNumber, lastLearn.sectionNumber)}&nbsp;&nbsp;
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
              >????????????
              </Button>
            </Link>
            <Popconfirm title="?????????????????????" onConfirm={this.exitProject} placement="topRight">
              <Button
                type="danger"
                shape="round"
                size="middle"
                style={{margin: '5px'}}
                loading={exitBtLoading}
              >????????????
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
            loading={learnBtLoading}
          >????????????
          </Button>
        }
      </div>
    )

    return (
      <DocumentTitle title="Project">
        <div style={{backgroundColor: '#f2f4f5'}}>
          <BackTop/>
          <div
            style={{
              width: '100%',
              height: '100%',
              paddingRight: '20px',
              paddingLeft: '20px',
              paddingBottom: '20px',
              margin: 'auto',
              maxWidth: '1200px',
            }}
          >
            <PageHeader
              className="site-page-header"
              onBack={() => this.back()}
              title="??????"
              subTitle="????????????"
            />
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
                        fallback={"https://cdn.open-ct.com/task-resources//openpbl-empty-project.png"}
                      />
                    </Col>
                    <Col span={1}>&nbsp;</Col>
                    <Col flex="auto">
                      <p style={{fontSize: '20px'}}>
                        {project.projectTitle}&nbsp;&nbsp;
                        {project.created ?
                          <>
                            <Tag color="geekblue">??????????????????</Tag>
                            <span style={{color: 'gray', fontSize: '15px'}}>
                              {this.projectState()}
                            </span>
                          </>
                          : null
                        }
                        {project.learning ? <Tag color="geekblue">????????????</Tag> : null}
                        <span style={{float: 'right'}}>
                        {project.favourite ?
                          <Tooltip title="??????????????????">
                            <Button shape="circle" type="text" loading={favBtLoading} icon={<StarFilled/>}
                                    onClick={this.removeFavourite}/>
                          </Tooltip>
                          :
                          <Tooltip title="????????????">
                            <Button shape="circle" type="text" loading={favBtLoading} icon={<StarOutlined/>}
                                    onClick={this.addFavourite}/>
                          </Tooltip>
                        }
                        </span>
                      </p>
                      <p
                        style={{fontSize: '14px', color: 'gray'}}
                      >
                        ???????????????{util.FilterTime(project.createAt)}&nbsp;&nbsp;&nbsp;&nbsp;
                        ???????????????{project.published ? util.FilterTime(project.publishedAt) : "?????????"}&nbsp;&nbsp;&nbsp;&nbsp;
                        ???????????????
                        {project.closed ?
                          util.FilterTime(project.closedAt)
                          :
                          <>
                            {project.timedEnd ? util.FilterTime(project.endTime) : "??????"}
                          </>
                        }
                      </p>
                      <div>
                        <span>{project.readNum}&nbsp;?????????</span>
                        <Divider type="vertical"/>
                        <span>{project.joinNum}&nbsp;???????????????</span>
                        <Divider type="vertical"/>
                        <span>???????????????&nbsp;
                          <Avatar src={teacher.avatar}/>&nbsp;&nbsp;
                          {teacher.displayName}
                          </span>
                      </div>
                      <br/>
                      {this.props.account.tag === '??????' ?
                        studentBt
                        :
                        <>
                          {project.created ?
                            teacherBt
                            : null
                          }
                        </>
                      }
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
                  <Menu.Item key="project-introduce">????????????</Menu.Item>
                  <Menu.Item key="project-outline">????????????</Menu.Item>
                  <Menu.Item key="project-comment">????????????</Menu.Item>
                  <Menu.Item key="project-evaluation">????????????</Menu.Item>

                  {project.created ? <Menu.Item key="student-admin">????????????</Menu.Item>
                    : null}
                  {project.learning ? <Menu.Item key="student-evidence">????????????</Menu.Item>
                    : null}
                </Menu>
                <div style={{
                  backgroundColor: 'white',
                  padding: '20px',
                }}
                >
                  {menu === 'project-introduce' ? <ProjectIntroduce project={project}/> : null}
                  {menu === 'project-outline' ? <ProjectOutline project={project}/> : null}
                  {menu === 'project-comment' ? <ProjectComment project={project} account={this.props.account}/> : null}
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
      </DocumentTitle>
    );
  }
}

export default ProjectInfo
