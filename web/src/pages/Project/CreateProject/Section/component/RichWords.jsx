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
import {Button, Card, Divider, message, Upload} from "antd";
import {FileImageOutlined, VideoCameraAddOutlined, AudioOutlined} from "@ant-design/icons"
import BraftEditor from 'braft-editor'
import { ContentUtils } from 'braft-utils'

import 'braft-editor/dist/index.css'
import './rich-words.less'
import ResourceApi from "../../../../../api/ResourceApi";
import "./section-edit.less"
import FileApi from "../../../../../api/FileApi";

function RichWords(obj) {
  const [editorState, setEditorState] = useState()
  const [content, setContent] = useState('')
  const [saveBtLoading, setSaveBtLoading] = useState(false)

  useEffect(() => {
    if (obj.content !== undefined) {
      initEditor()
    }
  }, [obj.content])

  const initEditor = () => {
    setEditorState(BraftEditor.createEditorState(obj.content))
    setContent(obj.content)
  }

  const receiveHtml = (content) => {
    setContent(content)
  }

  const changeContent = (e) => {
    setEditorState(e)
    setContent(e.toHTML())
  }

  const saveContent = () => {
    setSaveBtLoading(true)
    ResourceApi.updateResourceContent(obj.pid, obj.section.resource.id, content)
      .then(res => {
        setSaveBtLoading(false)
        if (res.data.code === 200) {
          obj.section.resource.content = content
          message.success(res.data.msg)
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e => {
        console.log(e)
      })
  }

  const uploadFile = (param, type, size) => {
    let file = param.file
    const index = file.name.lastIndexOf('.');
    if (index === -1) {
      message.error("不能识别文件类型")
      return
    }
    if (file.size > 1024 * 1024 * size) {
      message.error(`文件不能大于${size}MB`)
      return
    }
    const postfix = file.name.substr(index);
    let fileName = new Date().getTime()
    let filePath = `/openpbl/resources/${fileName}${postfix}`
    FileApi.uploadFile("admin", "openpbl", obj.account.name, filePath, file)
      .then(res=>{
        if (res.data.status === 'ok') {
          let url = res.data.data
          setEditorState(ContentUtils.insertMedias(editorState, [{
            type: type,
            url: url
          }]))
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e=>{console.log(e)})
  };

  const controls = [
    'undo', 'redo', 'separator',
    'font-size', 'line-height', 'letter-spacing', 'separator',
    'text-color', 'bold', 'italic', 'underline', 'strike-through', 'separator',
    'superscript', 'subscript', 'remove-styles', 'emoji',  'separator', 'text-indent', 'text-align', 'separator',
    'headings', 'list-ul', 'list-ol', 'blockquote', 'code', 'separator',
    'link', 'separator', 'hr', 'separator',
    'separator', 'clear'
  ]
  const extendControls = [
    'separator',
    {
      key: 'insert-image',
      type: 'component',
      component: (
        <div style={{float: 'left'}}>
          <Upload
            accept="image/*"
            showUploadList={false}
            customRequest={param=>uploadFile(param, 'IMAGE', '5')}
          >
            <Button size="large" type="text" icon={<FileImageOutlined />} data-title="插入图片"/>
          </Upload>
        </div>
      )
    }, {
      key: 'insert-audio',
      type: 'component',
      component: (
        <div style={{float: 'left'}}>
          <Upload
            accept="audio/*"
            showUploadList={false}
            customRequest={param=>uploadFile(param, 'AUDIO', '10')}
          >
            <Button size="large" type="text" icon={<AudioOutlined />} data-title="插入音频"/>
          </Upload>
        </div>
      )
    }, {
      key: 'insert-video',
      type: 'component',
      component: (
        <div style={{float: 'left'}}>
          <Upload
            accept="video/*"
            showUploadList={false}
            customRequest={param=>uploadFile(param, 'VIDEO', '1024')}
          >
            <Button size="large" type="text" icon={<VideoCameraAddOutlined />} data-title="插入视频"/>
          </Upload>
        </div>
      )
    }
  ]

  return (
    <Card className="resource-card">
      <p className="card-title">文本内容</p>
      <Divider/>
      <div>
        <BraftEditor
          style={{height: '100%'}}
          value={editorState}
          onChange={changeContent}
          controls={controls}
          extendControls={extendControls}
          onBlur={saveContent}
        />
      </div>
      <div style={{marginTop: '10px'}}>
        <span style={{float: 'left'}}>
          <Button type="primary" onClick={saveContent} loading={saveBtLoading}>保存</Button>
        </span>
      </div>
    </Card>
  )
}

export default RichWords
