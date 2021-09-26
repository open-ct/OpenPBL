// Copyright 2021 The OpenPBL Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React, {useEffect, useState} from "react";
import {Button, Card, PageHeader} from "antd";
import DocumentTitle from 'react-document-title';
import {Link} from "react-router-dom";


import SectionApi from "../../../../api/SectionApi";
import RichWords from "./component/RichWords";
import FileResource from "./component/FileResource";
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
        <div style={{padding: '20px', margin: 'auto', maxWidth: '1400px'}}>
          <Card>
            <h2 style={{fontWeight: 'bold'}}>{section.sectionName}</h2>
          </Card>
          <RichWords section={section} pid={pid} sid={sid} account={obj.account} content={section.resource.content}/>
          <FileResource section={section} pid={pid} sid={sid} account={obj.account}/>
          <StudentTask section={section} pid={pid} sid={sid}/>
        </div>
        <div style={{textAlign: 'center'}}>
          <Link to={`/home/project/${pid}/section/${sid}/preview?back=/home/project/${pid}/section/${sid}/edit`}>
            <Button style={{marginBottom: '20px'}}>预览</Button>
          </Link>
        </div>
      </div>
    </DocumentTitle>
  )
}

export default SectionEditPage
