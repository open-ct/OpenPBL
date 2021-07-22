import React, {useEffect, useState} from "react";
import {Card, PageHeader} from "antd";
import DocumentTitle from 'react-document-title';

import SectionApi from "../../../api/SectionApi";
import "./section-edit.less"

function PreviewResource(obj) {
  const sid = obj.match.params.sid
  const pid = obj.match.params.pid
  const [section, setSection] = useState({resource:{}})
  useEffect(()=>{
    SectionApi.getSectionDetail(sid)
      .then(res=>{
        setSection(res.data.section)
      })
      .catch(e=>{console.log(e)})
  }, [])

  const back = e => {
    window.location.href = '/project/edit/section/' + pid + '/' + sid
  }


  return (
  <DocumentTitle title="Project">
    <div style={{backgroundColor: '#f2f4f5', minHeight: '100vh'}}>
      <PageHeader
        className="site-page-header"
        onBack={back}
        title="返回"
        subTitle="我的项目"
      />
      <div style={{ padding: '20px', margin: 'auto', maxWidth: '1200px'}}>
        <Card>
          <h2 style={{ fontWeight: 'bold' }}>{section.sectionName}</h2>
        </Card>
        <Card className="resource-card">
          <div dangerouslySetInnerHTML={{__html: section.resource.content}}/>
        </Card>
        <Card className="resource-card">
          文件
        </Card>
        <Card className="resource-card">
          学生任务
        </Card>
      </div>
    </div>
  </DocumentTitle>
)
}

export default PreviewResource
