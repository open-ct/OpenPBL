import React, {useEffect} from "react";
import {Button, Card, PageHeader} from "antd";
import DocumentTitle from 'react-document-title';
import StudentEvidence from "./component/StudentEvidence";


function Evidence(obj) {
  const pid = obj.match.params.pid
  const sid = obj.match.params.sid

  useEffect(() => {
  }, [])

  const back = e => {
    window.location.href = `/home/project/${pid}/info?menu=student-admin`
  }

  const exportFile = e => {
    window.open(`/export/project/${pid}/student/${sid}/evidence`)
  }
  return (
    <DocumentTitle title="Project">
      <div style={{backgroundColor: '#f2f4f5', minHeight: '100vh'}}>
        <PageHeader
          className="site-page-header"
          onBack={back}
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
            }}
          >
            <div style={{padding: '10px', textAlign: 'center'}}>
              <Button onClick={exportFile}>导出</Button>
            </div>
            <div>
              <Card>
                <StudentEvidence
                  pid={pid}
                  studentId={sid}
                />
              </Card>
            </div>
          </div>
        </div>
      </div>
    </DocumentTitle>
  )
}

export default Evidence