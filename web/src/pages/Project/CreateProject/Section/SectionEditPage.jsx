import React, {useEffect, useState} from "react";
import {Button, Card, PageHeader} from "antd";
import DocumentTitle from 'react-document-title';
import {Link} from "react-router-dom";


import SectionApi from "../../../../api/SectionApi";
import RichWords from "./component/RichWords";
import FileResource from "../component/FileResource";
import StudentTask from "./component/StudentTask";

function SectionEditPage(obj) {
  const pid = obj.match.params.pid
  const sid = obj.match.params.sid
  const [section, setSection] = useState({resource: {}})

  useEffect(() => {
    SectionApi.getSectionDetail(sid, pid)
      .then(res => {
        setSection(res.data.section)
      })
      .catch(e => {
        console.log(e)
      })

  }, [])


  const back = e => {
    window.location.href = `/home/project/${pid}/outline/edit`
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
        <div style={{padding: '20px', margin: 'auto'}}>
          <Card>
            <h2 style={{fontWeight: 'bold'}}>{section.sectionName}</h2>
          </Card>
          <RichWords section={section}/>
          <FileResource section={section}/>
          <StudentTask section={section} pid={pid}/>
        </div>
        <div style={{textAlign: 'center'}}>
          <Link to={`/home/project/${pid}/section/${sid}/preview?back=/project/${pid}/section/${sid}/edit`}>
            <Button style={{marginBottom: '20px'}}>预览</Button>
          </Link>
        </div>
      </div>
    </DocumentTitle>
  )
}

export default SectionEditPage
