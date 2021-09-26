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
    window.location.href = `/home/project/${this.state.pid}/info`
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
            <InfoEditPage pid={pid} account={this.props.account}/>
          </Card>
        </div>
      </DocumentTitle>
    )
  }
}

export default EditInfo
