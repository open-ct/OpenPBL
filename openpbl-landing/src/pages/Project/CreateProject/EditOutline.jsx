import React from "react";
import DocumentTitle from 'react-document-title';
import {Card, Divider, PageHeader, Steps} from "antd";

import './index.less'
import OutlineEditPage from "./component/OutlineEditPage";

const {Step} = Steps;


class EditInfo extends React.PureComponent {
  state = {
    current: 0,
    pid: this.props.match.params.id
  }

  componentDidMount() {

  }

  back = e => {
    this.props.history.push('/project/info/' + this.state.pid)
  }

  render() {
    const {pid} = this.state
    return (
      <DocumentTitle title="Project">
        <div style={{backgroundColor: '#f2f4f5', minHeight: '100vh',}}>
          <PageHeader
            className="site-page-header"
            onBack={() => this.back()}
            title="返回"
            subTitle="我的项目"
          />
          <div
            style={{
              paddingTop: '20px',
              paddingLeft: '20px',
              paddingRight: '20px',
              paddingBottom: '20px'
            }}
          >
            <Card className="edit-card">
              <Steps current={1} style={{marginLeft: '20px'}}>
                <Step title="填写信息"/>
                <Step title="添加章节" style={{marginRight: '20px'}}/>
              </Steps>
              <Divider/>
              <OutlineEditPage pid={pid}/>
            </Card>
          </div>
        </div>
      </DocumentTitle>
    )
  }
}

export default EditInfo
