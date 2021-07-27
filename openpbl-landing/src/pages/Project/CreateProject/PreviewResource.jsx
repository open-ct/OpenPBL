import React, {useEffect, useState} from "react";
import {Card, PageHeader, Input, Upload, message, Button, Collapse, Spin, Radio, Divider} from "antd";
import DocumentTitle from 'react-document-title';
import {InboxOutlined} from '@ant-design/icons'
import SectionApi from "../../../api/SectionApi";
import "./section-edit.less"
import "./preview.less"
import TaskApi from "../../../api/TaskApi";
import FillSurvey from "./component/FillSurvey";


function PreviewResource(obj) {
  const sid = obj.match.params.sid
  const pid = obj.match.params.pid
  const [section, setSection] = useState({resource:{}})
  const [tasks, setTasks] = useState([])

  useEffect(()=>{
    getSectionResource()
    getTasks()
  }, [])
  const getSectionResource = () => {
    SectionApi.getSectionDetail(sid)
      .then(res=>{
        setSection(res.data.section)
      })
      .catch(e=>{console.log(e)})
  }
  const getTasks = () => {
    TaskApi.getSectionTasks(sid)
      .then(res=>{
        if (res.data.tasks === null) {
          setTasks([])
        } else {
          let t = res.data.tasks
          for (let i = 0; i < t.length; i++) {
            if (t[i].questions !== undefined && t[i].questions != null) {
              for (let j = 0; j < t[i].questions.length; j++) {
                t[i].questions[j].questionOptions = t[i].questions[j].questionOptions.split(",")
              }
            }
          }
          setTasks(t)
        }
      })
      .catch(e=>{console.log(e)})
  }
  const back = e => {
    window.location.href = `/project/${pid}/section/${sid}/edit`
  }
  const props = {
    name: 'file',
    multiple: true,
    action: '',
    onChange(info) {
      const { status } = info.file;
      if (status !== 'uploading') {
        console.log(info.file, info.fileList);
      }
      if (status === 'done') {
        message.success(`${info.file.name} 上传成功`);
      } else if (status === 'error') {
        message.error(`${info.file.name} 上传失败`);
      }
    },
    onDrop(e) {
      console.log('Dropped files', e.dataTransfer.files);
    },
  };

  return (
  <DocumentTitle title="Project">
    <div style={{backgroundColor: '#f2f4f5', minHeight: '100vh'}}>
      <PageHeader
        className="site-page-header"
        onBack={back}
        title="返回"
        subTitle="我的项目"
      />
      <div style={{ padding: '20px', margin: 'auto'}}>
        <Card>
          <h2 style={{ fontWeight: 'bold' }}>{section.sectionName}</h2>
        </Card>
        <Card className="resource-card">
          <div dangerouslySetInnerHTML={{__html: section.resource.content}}/>
        </Card>
        <Card className="resource-card">
          <p className="card-title">文件资源</p>
          <p>{section.resource.fileTitle}</p>
          <p>{section.resource.fileIntroduce}</p>
        </Card>
        {tasks.map((item, index)=>(
          <Card className="resource-card" key={index.toString()}>
            <p className="card-title">学生任务</p>
            <p className="task-title">{item.taskTitle}</p>
            <p>{item.taskIntroduce}</p>
            {item.taskType === 'file' ?
              <div>
                <Upload.Dragger {...props}>
                  <p className="ant-upload-drag-icon">
                    <InboxOutlined />
                  </p>
                  <p className="ant-upload-text">点击或拖动文件上传</p>
                  <p className="ant-upload-hint">hint
                  </p>
                </Upload.Dragger>
              </div>
              : null
            }
            {item.taskType === 'comment' ?
              <div>
                <Input.TextArea />
                <Button type="primary" style={{float: 'right', marginTop: '10px'}}>提交</Button>
              </div>
              : null
            }
            {item.taskType === 'survey' ?
              <FillSurvey
                item={item}
              />
              : null
            }
          </Card>
        ))
        }
      </div><br/>
    </div>
  </DocumentTitle>
)
}

export default PreviewResource
