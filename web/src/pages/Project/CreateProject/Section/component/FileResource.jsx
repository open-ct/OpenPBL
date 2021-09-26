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
          message.error(res.data.msg)
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
