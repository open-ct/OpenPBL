import React from "react";
import {Button, Input, message, Upload} from "antd";
import {InboxOutlined} from "@ant-design/icons";
import FillSurvey from "./FillSurvey";
import SubmitApi from "../../../../api/SubmitApi";


function TaskCard(obj) {

  const updateComment = (item, index) => {
    SubmitApi.updateSubmit(obj.pid, item.id, item.submit.id, item.submit)
      .then(res=>{
        if (res.data.code === 200) {
          message.success(res.data.msg)
          obj.getTasks()
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e=>{console.log(e)})
  }
  const submitComment = (item, index) => {
    item.submit.submitType = item.taskType
    SubmitApi.createSubmit(obj.pid, item.id, item.submit)
      .then(res=>{
        if (res.data.code === 200) {
          message.success(res.data.msg)
          obj.getTasks()
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e=>{console.log(e)})
  }
  const changeComment = (v, index) => {
    obj.item.submit.submitContent = v.target.value
    obj.setTaskItem(obj.item, index)
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
    <>
      <p className="task-title">{obj.item.taskTitle}</p>
      <p>{obj.item.taskIntroduce}</p>
      {obj.item.taskType === 'file' ?
        <div>
          <Upload.Dragger {...props} disabled={!obj.editable}>
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
      {obj.item.taskType === 'comment' ?
        <div>
          <Input.TextArea value={obj.item.submit.submitContent} onChange={v=>changeComment(v, obj.index)} />
          <div style={{textAlign: 'right', marginTop: '10px'}}>
            {obj.item.submitted ?
              <Button disabled={!obj.editable} type="primary" onClick={e=>updateComment(obj.item, obj.index)}>更新</Button>
              :
              <Button disabled={!obj.editable} type="primary" onClick={e => submitComment(obj.item, obj.index)}>提交</Button>
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