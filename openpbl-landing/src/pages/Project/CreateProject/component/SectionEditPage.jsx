import React, {useEffect, useState} from "react";
import {Col, Row, Input, PageHeader, Card} from "antd";
import DocumentTitle from 'react-document-title';

import SectionApi from "../../../../api/SectionApi";
import "./section-edit.less"

function SectionEditPage(obj) {
  const pid = obj.match.params.pid
  const sid = obj.match.params.sid;
  const [section, setSection] = useState({})
  const [resources, setResources] = useState([])

  useEffect(()=>{
    SectionApi.getSection(sid)
      .then(res=>{
        setSection(res.data.section)
      })
      .catch(e=>{console.log(e)})

    setResources([
      {type: 'words', title: '介绍', content: ''},
      {type: 'file', title: '这是一个视频', content: '/../..'},
      {type: 'homeWork', title: '这是一个作业', content: ''}
    ])

  }, [])

  const addRichWords = e => {

  }
  const addFile = e => {

  }
  const addHomeWork = e => {

  }

  const back = e => {
    window.location.href = '/project/edit/outline/' + pid
  }

  return (
    <DocumentTitle title="Project">
      <div style={{backgroundColor: '#f2f4f5', minHeight: '100vh',}}>
        <PageHeader
          className="site-page-header"
          onBack={back}
          title="返回"
          subTitle="我的项目"
        />
        <div style={{ padding: '20px', margin: 'auto' }}>
          <Card>
            <h2 style={{ fontWeight: 'bold' }}>{section.sectionName}</h2>
          </Card>
          {
            resources.map((item, index)=>(
              <Card className="resource-card" key={index.toString()}>
                <h3 style={{textAlign: 'left'}}>{item.title}</h3>
              </Card>
            ))
          }
          <Row gutter={[20, 20]} style={{ marginTop: '20px' }}>
            <Col span={8}>
              <Card hoverable bordered={false} onClick={addRichWords}>添加富文本框</Card>
            </Col>
            <Col span={8}>
              <Card hoverable bordered={false} onClick={addFile}>添加文件资源</Card>
            </Col>
            <Col span={8}>
              <Card hoverable bordered={false} onClick={addHomeWork}>添加学生任务</Card>
            </Col>
          </Row>
        </div>
      </div>
    </DocumentTitle>
  )
}

export default SectionEditPage
