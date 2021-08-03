import React, {useEffect, useState} from "react";
import {InputNumber, Table} from "antd";
import TaskApi from "../../../../api/TaskApi";

function StudentEvidence(obj) {
  const pid = obj.project.id
  const [tasks, setTasks] = useState([])

  useEffect(() => {
    getTasks()
  }, []);
  const getTasks = () => {
    TaskApi.getProjectTasks(pid)
      .then(res=>{
        if (res.data.code === 200) {
          if (res.data.data != null) {
            setTasks(res.data.data)
          }
        }
      })
      .catch(e=>{console.log(e)})
  }
  const getColumns = () => {
    let c = [
      {
        title: '任务标题',
        dataIndex: 'taskTitle',
        key: 'taskTitle'
      },
      {
        title: '任务描述',
        dataIndex: 'taskIntroduce',
        key: 'taskIntroduce',
        ellipsis: true,
      },
      {
        title: '权重',
        dataIndex: 'taskWeight',
        key: 'taskWeight',
        render: (text, item, index) => (
          <span>{text}&nbsp;&nbsp;%</span>
        )
      }]
    if (obj.project.learning) {
      c.push({
        title: '得分',
        dataIndex: 'score',
        key: 'score',
        render: (text, item, index) => (
          <>
            <span>{getScore(text, item.taskWeight)}&nbsp;/&nbsp;{item.taskWeight}</span>
          </>
        )
      })
      c.push({
        title: '状态',
        dataIndex: 'score',
        key: 'score',
        render: (text, item, index) => (
          <>
            {item.scored ?
              <span>已打分</span>
              :
              <span>未打分</span>
            }
          </>
        )
      })
    }
    return c
  }
  const getScore = (score, weight) => {
    return (score * weight / 100).toFixed(2)
  }
  return (
    <div>
      <Table
        columns={getColumns()}
        dataSource={tasks}
        pagination={false}
      />
    </div>
  )
}

export default StudentEvidence