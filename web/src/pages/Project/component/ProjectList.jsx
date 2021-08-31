import React, {useEffect, useState} from 'react';
import {Card, Col, Divider, Image, Input, Pagination, Row, Select, Spin, Tag} from 'antd';
import {EyeOutlined, TeamOutlined} from '@ant-design/icons';
import {Link} from "react-router-dom"

import './project-list.less';
import ProjectListApi from '../../../api/ProjectListApi'
import util from '../../component/Util'
import ProjectApi from "../../../api/ProjectApi";

const {Meta} = Card;
const {Search} = Input;
const topColResponsiveProps = {
  xs: 24,
  sm: 12,
  md: 8,
  lg: 8,
  xl: 6,
};

function ProjectList(obj) {
  const [learningProjectList, setLearningProjectList] = useState([]);

  const [loadProjects, setLoadProjects] = useState(false)

  const [page, setPage] = useState(1);
  const [size, setSize] = useState(10);

  const iniSubjects = ['语文', '数学', '英语', '科学', '政治', '历史', '地理', '化学', '生物', '美术', '音乐']
  const iniSkills = ['信息、媒体与技术技能', '生活与职业技能', '文化理解与传承素养', '审辨思维', '创新素养', '沟通素养', '合作素养']
  const [subjects, setSubjects] = useState(iniSubjects)
  const [skills, setSkills] = useState(iniSkills)
  const [total, setTotal] = useState(0);
  const [selectedSubjects, setSelectedSubjects] = useState([]);
  const [value, setValue] = useState('')
  const [selectedSkills, setSelectedSkills] = useState([])

  const type = localStorage.getItem('type')
  const uid = localStorage.getItem('uid')
  const mode = obj.mode

  const loadPage = (p, size) => {
    setSize(size)
    setPage(p)
    updateProjectList(p, size, selectedSubjects.toString(), selectedSkills.toString(), value)
  }

  const updateProjectList = (p, size, subject, skill, text) => {
    setLoadProjects(true)
    let q = {
      from: (p - 1) * size,
      size: size,
      orderBy: 'create_at',
      orderType: 'desc',
      subject: subject,
      skill: skill,
      text: text
    }
    ProjectListApi.getUserProjectList(mode, q)
      .then((res) => {
        setLoadProjects(false)
        if (res.data.projects === null) {
          setLearningProjectList([])
        } else {
          setLearningProjectList(res.data.projects)
        }
        setTotal(res.data.count);
      })
      .catch((e) => {
        console.log(e)
      })
    setPage(p);
  }

  useEffect(() => {
    loadPage(1, 10);
    loadSubjectsAndSkills()
  }, []);

  const loadSubjectsAndSkills = () => {
    ProjectApi.getSubjectsAndSkills(0)
      .then(res => {
        if (res.data.code === 200) {
          if (res.data.subjects !== null) {
            let s = res.data.subjects
            for (let i = 0; i < iniSubjects.length; i++) {
              if (s.indexOf(iniSubjects[i]) < 0) {
                s.push(iniSubjects[i])
              }
            }
            setSubjects(s)
          }
          if (res.data.skills !== null) {
            let s = res.data.skills
            for (let i = 0; i < iniSkills.length; i++) {
              if (s.indexOf(iniSkills[i]) < 0) {
                s.push(iniSkills[i])
              }
            }
            setSkills(s)
          }
        }
      })
      .catch(e => {
        console.log(e)
      })
  }

  const handleSubjectsChange = selected => {
    setSelectedSubjects(selected)
    updateProjectList(page, size, selected.toString(), selectedSkills.toString(), value)
  }
  const handleSkillsChange = selected => {
    setSelectedSkills(selected)
    updateProjectList(page, size, selectedSubjects.toString(), selected.toString(), value)
  }
  const viewProject = (item) => {
    if (item.published && !item.closed) {
      ProjectApi.viewProject(item.id)
        .then(res => {
        })
        .catch(e => {
        })
    }
  }

  const onSearch = (v) => {
    setValue(v)
    updateProjectList(page, size, selectedSubjects.toString(), selectedSkills.toString(), v)
  };
  return (
    <div>
      <Card bordered={false} style={{backgroundColor: 'white', borderRadius: '4px', textAlign: 'left'}}>
        <Search
          size="large"
          placeholder="搜索"
          onSearch={onSearch}
        />

        <Row>
          <Col span={2}>
            <h3 className="options-text">学&nbsp;&nbsp;科</h3>
          </Col>
          <Col span={22}>
            <Select
              showArrow
              mode="multiple"
              placeholder="选择学科"
              value={selectedSubjects}
              onChange={handleSubjectsChange}
              style={{width: '100%', marginTop: '20px'}}
            >
              {subjects.map(item => (
                <Select.Option key={item} value={item}>
                  {item}
                </Select.Option>
              ))}
            </Select>
          </Col>
        </Row>
        <Row>
          <Col span={2}>
            <h3 className="options-text">技&nbsp;&nbsp;能</h3>
          </Col>
          <Col span={22}>
            <Select
              showArrow
              mode="multiple"
              placeholder="选择技能"
              value={selectedSkills}
              onChange={handleSkillsChange}
              style={{width: '100%', marginTop: '10px'}}
            >
              {skills.map(item => (
                <Select.Option key={item} value={item}>
                  {item}
                </Select.Option>
              ))}
            </Select>
          </Col>
        </Row>
      </Card>

      <div style={{marginTop: '10px', textAlign: 'center'}}>
        <Spin spinning={loadProjects}/>
        <Row gutter={[20, 20]} style={{textAlign: 'left'}}>
          {
            learningProjectList.map((item, index) => (
              <Col key={index.toString()} {...topColResponsiveProps}>
                <Link to={`/home/project/${item.id}/info`}>
                <Card
                  hoverable
                  bordered={false}
                  onClick={e => viewProject(item)}
                  style={{
                    borderRadius: '10px',
                  }}
                  cover={
                    <Image
                      src={item.image}
                      preview={false}
                      style={{
                        borderTopLeftRadius: '10px',
                        borderTopRightRadius: '10px',
                      }}
                      fallback={"https://cdn.open-ct.com/task-resources//openpbl-empty-project.png"}
                    />
                  }
                >
                  <Meta
                    title={
                      item.projectTitle === '' ? "无" : item.projectTitle
                    }
                    description={
                      <div>
                        <span className="des-text">{item.subjects === '' ? '--' : item.subjects}</span>
                        {item.learning ?
                          <Tag color="geekblue" style={{zIndex: '999', float: 'right'}}>学习中</Tag> : null}
                        {item.teacherId === uid && type === 'teacher' ?
                          <Tag color="geekblue" style={{zIndex: '999', float: 'right'}}>我的项目</Tag> : null}
                      </div>
                    }
                  />

                  <span
                    style={{
                      color: 'gray',
                      fontSize: 'small',
                    }}
                  >
                      <EyeOutlined/>&nbsp;
                    {item.readNum}
                    </span>
                  <Divider type="vertical"/>
                  <span
                    style={{
                      color: 'gray',
                      fontSize: 'small',
                    }}
                  >
                      <TeamOutlined/>&nbsp;
                    {item.joinNum}
                    </span>
                  <span
                    style={{
                      color: 'gray',
                      fontSize: 'small',
                      float: 'right',
                    }}
                  >
                      {util.FilterMoment(item.createAt)}
                    </span>
                </Card>
                </Link>
              </Col>
            ))
          }
        </Row>
        <div
          style={{
            textAlign: 'center',
            marginTop: '20px',
          }}
        >
          <Pagination
            total={total}
            showTotal={t => `共${total}项内容`}
            current={page}
            onChange={loadPage}
            onShowSizeChange={loadPage}
          />
        </div>
        <br/>
      </div>
    </div>
  );
}

export default ProjectList;
