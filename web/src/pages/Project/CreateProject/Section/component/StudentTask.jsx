import React, {useEffect, useState} from "react";
import {Button, Card, Divider, Dropdown, Input, Menu, message, Popconfirm} from "antd";
import {ArrowDownOutlined, ArrowUpOutlined, DeleteOutlined, UpOutlined} from "@ant-design/icons";

import TaskApi from "../../../../../api/TaskApi";
import "./section-edit.less"


function StudentTask(obj) {
  const pid = obj.pid
  const sid = obj.sid
  const [tasks, setTasks] = useState([])
  useEffect(() => {
    if (sid !== undefined) {
      getTasks()
    }
  }, [sid])
  const getTasks = () => {
    TaskApi.getSectionTasks(sid, pid)
      .then(res => {
        if (res.data.tasks === null) {
          setTasks([])
        } else {
          setTasks(res.data.tasks)
        }
      })
  }

  const saveContent = (item, index) => {
    TaskApi.updateTask(pid, item)
      .then(res => {
        if (res.data.code === 200) {
          message.success(res.data.msg)
        } else {
          message.error(res.data.msg)
        }
      })
  }
  const addTask = (tp) => {
    let len = tasks.length
    let o = 0
    if (len > 0) {
      o = tasks[len - 1].taskOrder + 1
    }
    // taskType: survey file comment
    let t = {
      sectionId: sid,
      taskOrder: o,
      taskType: tp.key,
      chapterNumber: obj.section.chapterNumber,
      sectionNumber: obj.section.sectionNumber
    }
    TaskApi.createTask(pid, t)
      .then(res => {
        if (res.data.code === 200) {
          t.id = res.data.data
          tasks.push(t)
          setTasks([...tasks])
        }
      })
  }

  const changeTitle = (value, index) => {
    tasks[index].taskTitle = value.target.value
    setTasks([...tasks])
  }
  const changeIntroduce = (value, index) => {
    tasks[index].taskIntroduce = value.target.value
    setTasks([...tasks])
  }
  const deleteTask = (item, index) => {
    TaskApi.deleteTask(pid, item.id)
      .then(res => {
        if (res.data.code === 200) {
          tasks.splice(index, 1)
          setTasks([...tasks])
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e => {
        console.log(e)
      })
  }
  const exchangeTask = (index1, index2) => {
    if (index1 < 0 || index2 >= tasks.length) {
    } else {
      let id1 = tasks[index1].id
      let id2 = tasks[index2].id
      TaskApi.exchangeTask(id1, id2)
        .then(res => {
          if (res.data.code === 200) {
            let t1 = tasks[index1]
            tasks[index1] = tasks[index2]
            tasks[index2] = t1
            setTasks([...tasks])
          } else {
            message.error(res.data.msg)
          }
        })
    }
  }
  const gotoSurvey = (item, index) => {
    saveContent(item, index)
    window.location.href = `/home/project/${pid}/section/${obj.section.id}/task/${item.id}/survey/edit`
  }

  return (
    <Card className="resource-card">
      <p className="card-title">????????????</p>
      {tasks.map((item, index) => (
        <div key={index.toString()}>
          <Divider/>
          <p className="task-title">
            {item.taskType === 'file' ? '??????????????????' : null}
            {item.taskType === 'comment' ? '????????????' : null}
            {item.taskType === 'survey' ? '??????????????????' : null}
            <span style={{float: 'right'}}>
              <Button shape="circle" type="text" icon={<ArrowUpOutlined/>}
                      onClick={e => exchangeTask(index - 1, index)}/>
              <Button shape="circle" type="text" icon={<ArrowDownOutlined/>}
                      onClick={e => exchangeTask(index, index + 1)}/>
              &nbsp;&nbsp;
              <Popconfirm title="?????????????????????" onConfirm={e => deleteTask(item, index)} placement="topRight">
                <Button shape="circle" type="text" icon={<DeleteOutlined/>} style={{color: 'red'}}/>
              </Popconfirm>
            </span>
          </p>
          <Input placeholder="????????????" value={item.taskTitle} onChange={e => changeTitle(e, index)}/>
          <Input.TextArea placeholder="????????????" value={item.taskIntroduce} onChange={e => changeIntroduce(e, index)}
                          style={{marginTop: '20px'}}/>
          {item.taskType === 'survey' ?
            <Button style={{marginTop: '10px'}} onClick={e => gotoSurvey(item, index)}>????????????</Button>
            : null
          }

          <div style={{marginTop: '10px', textAlign: 'right'}}>
            <span>
              <Button type="primary" onClick={e => saveContent(item, index)}>??????</Button>
            </span>
          </div>
        </div>
      ))}
      <div style={{marginTop: '20px', textAlign: 'right'}}>
        <Dropdown
          trigger="click"
          overlay={
            <Menu onClick={addTask}>
              <Menu.Item key="file">????????????</Menu.Item>
              <Menu.Item key="comment">????????????</Menu.Item>
              <Menu.Item key="survey">????????????</Menu.Item>
            </Menu>
          }
        >
          <Button>????????????<UpOutlined/></Button>
        </Dropdown>
      </div>
    </Card>
  )
}

export default StudentTask
