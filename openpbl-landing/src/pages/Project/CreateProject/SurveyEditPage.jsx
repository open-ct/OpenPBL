import React, {useEffect, useState} from "react";
import {PageHeader, Card, Input, Divider, Button} from "antd";
import DocumentTitle from 'react-document-title';

import SurveyApi from "../../../api/SurveyApi";

function SurveyEditPage(obj) {
  const pid = obj.match.params.pid
  const sid = obj.match.params.sid
  const tid = obj.match.params.tid
  const [survey, setSurvey] = useState({questions:[]})

  useEffect(()=>{
    if (tid !== undefined) {
      SurveyApi.getSurveyDetailByTaskId(tid)
        .then(res => {
          if (res.data.questions === null) {
            res.data.questions = []
          }
          setSurvey(res.data.survey)
        })
    }
  }, [tid])

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
          subTitle="编辑资源"
        />
        <div style={{ padding: '20px', margin: 'auto', maxWidth: '1200px' }}>
          <Card>
            <Input />
            {survey.questions.map((item, index) => (
              <div key={index.toString()}>
                <Divider />
                <p>{item.questionTitle}</p>
              </div>
            ))}
            <Button>添加</Button>
          </Card>
        </div>
      </div>
    </DocumentTitle>
  )
}

export default SurveyEditPage
