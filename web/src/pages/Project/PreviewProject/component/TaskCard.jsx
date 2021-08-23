import React, {useEffect, useState} from "react";
import {Button, Input, message, Upload} from "antd";
import {InboxOutlined} from "@ant-design/icons";
import FillSurvey from "./FillSurvey";
import SubmitApi from "../../../../api/SubmitApi";
import FileApi from "../../../../api/FileApi";
import SectionApi from "../../../../api/SectionApi";


function TaskCard(obj) {
  const [fileList, setFileList] = useState([])

  useEffect(()=>{
    if (obj.item.taskType === "file") {
      getSubmitFiles()
    }
  }, [])

  const getSubmitFiles = () => {
    SubmitApi.getSubmitFiles(obj.pid, obj.item.id, obj.item.submit.id)
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

  const updateComment = (item, index) => {
    SubmitApi.updateSubmit(obj.pid, item.id, item.submit.id, item.submit)
      .then(res => {
        if (res.data.code === 200) {
          message.success(res.data.msg)
          obj.getTasks()
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e => {
        console.log(e)
      })
  }
  const submitComment = (item, index) => {
    item.submit.submitType = item.taskType
    SubmitApi.createSubmit(obj.pid, item.id, item.submit)
      .then(res => {
        if (res.data.code === 200) {
          message.success(res.data.msg)
          obj.getTasks()
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e => {
        console.log(e)
      })
  }
  const changeComment = (v, index) => {
    obj.item.submit.submitContent = v.target.value
    obj.setTaskItem(obj.item, index)
  }

  const removeFile = file => {
    let data = {
      objectKey: file.filePath,
    }
    FileApi.deleteFile(JSON.stringify(data))
      .then(res=>{
        if (res.data.status === 'ok') {
          SubmitApi.deleteSubmitFile(obj.pid, obj.item.id, obj.item.submit.id, file.id)
            .then(res=>{
              if (res.data.code === 200) {
                getSubmitFiles()
              }
            })
            .catch(e=>{console.log(e)})
        }
      })
      .catch(e=>{console.log(e)})
  }
  const onUploadFile = file => {
    file = file.file
    const index = file.name.lastIndexOf('.');
    if (index === -1) {
      message.error('不能识别文件类型');
      return
    }
    if (file.size > 1024 * 1024 * 1024) {
      message.error('文件不能大于1GB');
      return
    }
    let filePath = `/openpbl/project/${obj.pid}/task/${obj.item.id}/${obj.studentId}/${file.name}`
    FileApi.uploadFile("admin", "openpbl", obj.studentId, filePath, file)
      .then(res=>{
        if (res.data.status === 'ok') {
          let e = false
          for (let i=0; i<fileList.length; i++) {
            if (fileList[i].name === file.name) {
              let f = fileList[i]
              f.url = res.data.data
              updateFile(f)
              e = true
            }
          }
          if (!e) {
            uploadFile(filePath, file.name, res.data.data)
          }
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
    SubmitApi.uploadSubmitFile(obj.pid, obj.item.id, obj.item.submit.id, d)
      .then(res=>{
        if (res.data.code === 200) {
          message.success(res.data.msg)
          obj.getTasks()
          getSubmitFiles()
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e=>{console.log(e)})
  }
  const updateFile = (f) => {
    SubmitApi.updateSubmitFile(obj.pid, obj.item.id, obj.item.submit.id, f)
      .then(res=>{
        if (res.data.code === 200) {
          message.success(res.data.msg)
          obj.getTasks()
          getSubmitFiles()
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
    <>
      <p className="task-title">{obj.item.taskTitle}</p>
      <p>{obj.item.taskIntroduce}</p>
      {obj.item.taskType === 'file' ?
        <div>
          <Upload.Dragger {...props} disabled={!obj.editable}>
            <p className="ant-upload-drag-icon">
              <InboxOutlined/>
            </p>
            <p className="ant-upload-text">点击或拖动文件上传</p>
            <p className="ant-upload-hint">hint
            </p>
          </Upload.Dragger>
        </div>
        : null
      }
      {obj.item.taskType === 'comment' ?
        <div>
          <Input.TextArea value={obj.item.submit.submitContent} onChange={v => changeComment(v, obj.index)}/>
          <div style={{textAlign: 'right', marginTop: '10px'}}>
            {obj.item.submitted ?
              <Button disabled={!obj.editable} type="primary"
                      onClick={e => updateComment(obj.item, obj.index)}>更新</Button>
              :
              <Button disabled={!obj.editable} type="primary"
                      onClick={e => submitComment(obj.item, obj.index)}>提交</Button>
            }
          </div>
        </div>
        : null
      }
      {obj.item.taskType === 'survey' ?
        <FillSurvey
          pid={obj.pid}
          item={obj.item}
          index={obj.index}
          getTasks={obj.getTasks}
          learning={obj.learning}
          editable={obj.editable}
          showCount={obj.showCount}
          setTaskItem={obj.setTaskItem}
        />
        : null
      }
    </>
  )
}

export default TaskCard