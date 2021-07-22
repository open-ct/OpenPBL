import React, {useEffect, useState} from "react";
import {Button, Card, Divider, Input, message, Upload} from "antd";
import {LoadingOutlined, InboxOutlined} from "@ant-design/icons";


import "../section-edit.less"
import ResourceApi from "../../../../api/ResourceApi";

const { Dragger } = Upload;

function FileResource(obj) {

  const [fileTitle, setFileTitle] = useState('')
  const [fileIntroduce, setFileIntroduce] = useState('')
  const [filePath, setFilePath] = useState('')

  useEffect(()=>{
    setFileTitle(obj.section.resource.fileTitle)
    setFileIntroduce(obj.section.resource.fileIntroduce)
    setFilePath(obj.section.resource.filePath)
  }, [fileTitle, obj.section.resource.fileIntroduce, obj.section.resource.filePath, obj.section.resource.fileTitle])

  const onUploadFile = e => {
  }
  const saveContent = e => {
    let q = {
      id: obj.section.resource.id,
      fileTitle: fileTitle,
      fileIntroduce: fileIntroduce,
      filePath: filePath
    }
    ResourceApi.updateResource(q)
      .then(res=>{
        if (res.data.code === 200){
          message.success(res.data.msg)
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e=>{console.log(e)})
  }
  const changeTitle = value => {
    setFileTitle(value.target.value)
  }
  const changeIntroduce = value => {
    setFileIntroduce(value.target.value)
  }

  const props = {
    name: 'file',
    multiple: false,
    action: 'https://www.mocky.io/v2/5cc8019d300000980a055e76',
    onChange(info) {
      const { status } = info.file;
      if (status !== 'uploading') {
        console.log(info.file, info.fileList);
      }
      if (status === 'done') {
        setFilePath(info.file.name)
        message.success(`${info.file.name} 上传成功`);
      } else if (status === 'error') {
        message.error(`${info.file.name} 上传失败`);
      }
    },
    onDrop(e) {
      console.log('删除文件', e.dataTransfer.files);
      setFilePath('')
    },
  };

  return (
    <Card className="resource-card">
      <p className="card-title">上传文件资源</p>
      <Divider />
      <Input value={fileTitle} placeholder="文件标题" maxLength={50} onChange={changeTitle} />
      <Input.TextArea value={fileIntroduce} placeholder="文件描述" onChange={changeIntroduce} style={{marginTop: '20px'}} />

      <Dragger {...props} style={{marginTop: '20px'}}>
        <p className="ant-upload-drag-icon">
          <InboxOutlined />
        </p>
        <p className="ant-upload-text">点击或拖动文件上传</p>
        <p className="ant-upload-hint">
          多个文件打包上传
        </p>
      </Dragger>
      <div style={{marginTop: '10px'}}>
        <span style={{float: 'right'}}>
          <Button type="primary" onClick={saveContent}>保存</Button>
        </span>
      </div>
    </Card>
  )
}

export default FileResource
