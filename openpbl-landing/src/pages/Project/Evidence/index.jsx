import React, {useEffect} from "react";
import {Card, PageHeader} from "antd";
import DocumentTitle from 'react-document-title';

function Evidence(obj) {
  const pid = obj.match.params.pid

  useEffect(()=>{
    console.log(pid)
  }, [])

  const back = e => {
    window.location.href = `/project/${pid}/info`
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
            <Card>
              evidence
            </Card>
          </div>
        </div>
      </div>
    </DocumentTitle>
  )
}

export default Evidence