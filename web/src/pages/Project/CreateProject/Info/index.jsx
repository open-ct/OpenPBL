import React from "react";
import DocumentTitle from 'react-document-title';
import {Card, Divider, PageHeader, Steps} from "antd";

import '../Outline/index.less'
import InfoEditPage from "./component/InfoEditPage";

const {Step} = Steps;


class EditInfo extends React.PureComponent {
  state = {
    current: 0,
    pid: this.props.match.params.id
  }

  componentDidMount() {
  }

  back = e => {
    window.location.href = `/project/${this.state.pid}/info`
  }

  render() {
    const {pid} = this.state
    return (
      <DocumentTitle title="Project">
        <div
          style={{
            paddingLeft: '20px',
            paddingRight: '20px',
            paddingBottom: '20px',
            maxWidth: '1200px',
            margin: 'auto',
          }}
        >
          <PageHeader
            className="site-page-header"
            onBack={() => this.back()}
            title="返回"
            subTitle="我的项目"
          />
          <Card className="edit-card">
            <Steps current={0} style={{marginLeft: '20px'}}>
              <Step title="填写信息"/>
              <Step title="添加章节" style={{marginRight: '20px'}}/>
            </Steps>
            <Divider/>
            <InfoEditPage pid={pid}/>
          </Card>
        </div>
      </DocumentTitle>
    )
  }
}

export default EditInfo
