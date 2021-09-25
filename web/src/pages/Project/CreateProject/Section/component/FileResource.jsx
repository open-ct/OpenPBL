import React, {useEffect, useState} from "react";
import {Button, Card, Divider, Input, message, Upload} from "antd";
import {InboxOutlined} from "@ant-design/icons";

import "./section-edit.less"
import FileApi from "../../../../../api/FileApi";
import SectionApi from "../../../../../api/SectionApi";

const {Dragger} = Upload;

function FileResource(obj) {
  const [fileList, setFileList] = useState([])

  useEffect(() => {
    getSectionFiles()
  }, [])

  const getSectionFiles = () => {
    SectionApi.getSectionFiles(obj.sid, obj.pid)
      .then(res=>{
        if (res.data.code === 200) {
          if (res.data.data !== null) {
            setFileList(res.data.data)
          } else {
            setFileList([])
          }
        }
      })
      .catch(e=>{console.log(e)})
  }
  const removeFile = file => {
    SectionApi.deleteSectionFile(obj.sid, obj.pid, file.id)
      .then(res=>{
        if (res.data.code === 200) {
          message.success(res.data.msg)
          getSectionFiles()
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e=>{console.log(e)})
  }
  const onUploadFile = file => {
    file = file.file
    let f = fileList
    f.push({name: file.name, status: 'uploading'})
    setFileList([...f])

    const index = file.name.lastIndexOf('.');
    if (index === -1) {
      message.error('不能识别文件类型');
      return
    }
    if (file.size > 1024 * 1024 * 1024 * 10) {
      message.error('文件不能大于10GB');
      return
    }
    let name = file.name.substr(0, index)
    const postfix = file.name.substr(index);
    let ts = new Date().getTime()
    let filePath = `/openpbl/resources/${name}${ts}${postfix}`
    FileApi.uploadFile("admin", "openpbl", obj.account.name, filePath, file)
      .then(res=>{
        if (res.data.status === 'ok') {
          uploadFile(filePath, file.name, res.data.data)
        } else {
          message.error('文件名过长，上传失败')
          getSectionFiles()
        }
      })
      .catch(e=>{console.log(e)})
  }
  const uploadFile = (filePath, name, url) => {
    let d = {
      filePath: filePath,
      name: name,
      url: url
    }
    SectionApi.uploadSectionFile(obj.sid, obj.pid, d)
      .then(res=>{
        if (res.data.code === 200) {
          message.success(res.data.msg)
          getSectionFiles()
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e=>{console.log(e)})
  }
  const updateFile = (f) => {
    SectionApi.updateSectionFile(obj.sid, obj.pid, f)
      .then(res=>{
        if (res.data.code === 200) {
          message.success(res.data.msg)
          getSectionFiles()
        }
      })
      .catch(e=>{console.log(e)})
  }

  const props = {
    name: 'file',
    multiple: false,
    customRequest: onUploadFile,
    fileList: fileList,
    onRemove: removeFile,
  };

  return (
    <Card className="resource-card">
      <p className="card-title">上传文件资源</p>
      <Divider/>
      <Dragger {...props} style={{marginTop: '20px'}} >
        <p className="ant-upload-drag-icon">
          <InboxOutlined/>
        </p>
        <p className="ant-upload-text">点击或拖动文件上传</p>
        <p className="ant-upload-hint">
          多个文件打包上传
        </p>
      </Dragger>

    </Card>
  )
}

export default FileResource
