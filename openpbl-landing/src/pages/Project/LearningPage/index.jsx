import React from 'react';
import {Card, message, PageHeader, Progress, Upload} from 'antd';
import {Document, Page} from 'react-pdf';
import ProjectApi from "../../../api/ProjectApi";


class Learning extends React.PureComponent {
  state = {
    pid: this.props.match.params.pid,
    cid: this.props.match.params.cid,
    sid: this.props.match.params.sid,
  }

  componentDidMount() {

  }

  onDocumentLoadSuccess = (numPages) => {
    this.setState({
      numPages,
    });
  };
  onUploadFile = (info) => {
    if (info.file.status !== 'uploading') {
      console.log(info.file, info.fileList);
    }
    if (info.file.status === 'done') {
      message.success(`${info.file.name} 文件上传成功`);
    } else if (info.file.status === 'error') {
      message.error(`${info.file.name} 文件上传失败`);
    }
  }
  back = e => {
    console.log('back')
    this.props.history.push('/project/info/' + this.state.pid)
  }

  render() {
    const { pid, cid, sid } = this.state
    return (
      <div style={{ minHeight: '100vh', backgroundColor: '#f2f4f5' }}>
        <PageHeader
          className="site-page-header"
          onBack={() => this.back()}
          title="返回"
          subTitle="项目信息"
        />
        <div style={{ paddingLeft: '20px', paddingRight: '20px' }}>
          <Card style={{ maxWidth: '1200px', margin: 'auto' }} >
            <h2>介绍</h2>
          </Card>
        </div>
      </div>
    );
  }
}

export default Learning;
