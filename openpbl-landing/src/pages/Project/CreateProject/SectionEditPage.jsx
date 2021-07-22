import React, {useEffect, useState} from "react";
import {PageHeader, Card, Button} from "antd";
import DocumentTitle from 'react-document-title';

import SectionApi from "../../../api/SectionApi";
import RichWords from "./component/RichWords";
import FileResource from "./component/FileResource";
import HomeWork from "./component/HomeWork";
import {Link} from "react-router-dom";

function SectionEditPage(obj) {
  const pid = obj.match.params.pid
  const sid = obj.match.params.sid;
  const [section, setSection] = useState({resource:{}})

  useEffect(()=>{
    SectionApi.getSectionDetail(sid)
      .then(res=>{
        setSection(res.data.section)
      })
      .catch(e=>{console.log(e)})

  }, [])


  const back = e => {
    window.location.href = '/project/edit/outline/' + pid
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
        <div style={{ padding: '20px', margin: 'auto', maxWidth: '1200px' }}>
          <Card>
            <h2 style={{ fontWeight: 'bold' }}>{section.sectionName}</h2>
          </Card>
          <RichWords section={section} />
          <FileResource section={section} />
          <HomeWork section={section}/>
        </div>
        <Link to={"/project/preview/" + pid + '/' + sid}>
          <Button style={{marginBottom: '20px'}}>预览</Button>
        </Link>
      </div>
    </DocumentTitle>
  )
}

export default SectionEditPage
