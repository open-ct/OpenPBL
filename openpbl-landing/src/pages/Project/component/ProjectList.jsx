import React, {useEffect, useState} from 'react';
import QueueAnim from 'rc-queue-anim';
import {Card, Col, Divider, Image, Input, Pagination, Row, Select, Tag} from 'antd';
import {EyeOutlined, TeamOutlined} from '@ant-design/icons';
import {Link} from 'react-router-dom';

import './project-list.less';

import ProjectListApi from '../../../api/ProjectListApi'

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
  const [loading, setLoading] = useState(true);
  const [page, setPage] = useState(1);
  const [size, setSize] = useState(10);

  const [total, setTotal] = useState(0);

  const [subjects, setSubjects] = useState([]);
  const [selectedSubjects, setSelectedSubjects] = useState([]);

  const [value, setValue] = useState('')

  const [skills, setSkills] = useState([]);
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
    let q = {
      page: p-1,
      size: size,
      orderBy: 'create_at',
      orderType: 'desc',
      subject: subject,
      skill: skill,
      text: text
    }
    ProjectListApi.getUserProjectList(mode, q)
      .then((res) => {
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

    setTimeout(() => {
      setLoading(false);
    }, 100);

    setSubjects(['语文', '数学', '英语', '科学', '信息技术']);
    setSkills(['学习与创新技能', '信息、媒体与技术技能', '生活与职业技能']);

  }, []);

  const handleSubjectsChange = selected => {
    setSelectedSubjects(selected)
    updateProjectList(page, size, selected.toString(), selectedSkills.toString(), value)
  }
  const handleSkillsChange = selected => {
    setSelectedSkills(selected)
    updateProjectList(page, size, selectedSubjects.toString(), selected.toString(), value)
  }

  const onSearch = (v) => {
    setValue(v)
    updateProjectList(page, size, selectedSubjects.toString(), selectedSkills.toString(), v)
  };
  return (
    <QueueAnim delay={100} className="queue-simple">
      <Card style={{backgroundColor: 'white', borderRadius: '4px', textAlign: 'left'}}>
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

      <div key="1" style={{marginTop: '10px'}}>
        <Row gutter={[15, 15]}>
          {
            learningProjectList.map((item, index) => (
              <Col key={index.toString()} {...topColResponsiveProps}>
                <Link to={`/project/info/${item.id}`}>
                  <Card
                    hoverable
                    bordered={false}
                    className="item-card"
                    cover={
                      <Image
                        alt="example"
                        src={item.image}
                        preview={false}
                      />
                    }
                  >
                    <Meta
                      title={item.projectTitle}
                      description={
                        <div>
                          {item.subjects === '' ? '无' : item.subjects}
                          {item.learning ? <Tag color="geekblue" style={{zIndex: '999', float: 'right'}}>学习中</Tag> : null}
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
                      {item.createTime}
                    </span>
                  </Card>
                </Link>
              </Col>
            ))
          }
        </Row>
        <div
          style={{
            margin: 'auto',
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
    </QueueAnim>
  );
}

export default ProjectList;
