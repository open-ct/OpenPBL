import React, {useEffect, useState} from "react";
import {Col, Collapse, Divider, List, Progress, Row, Table} from "antd";

import TaskApi from "../../../../api/TaskApi";
import ChapterApi from "../../../../api/ChapterApi";


function StudentEvidence(obj) {
  const pid = obj.project.id
  const [tasks, setTasks] = useState([])
  const [chapters, setChapters] = useState([])

  useEffect(() => {
    getChapters()

    getTasks()
  }, []);
  const getChapters = () => {
    ChapterApi.getProjectChapters(pid)
      .then((res) => {
        if (res.data.chapters === null) {
          setChapters([])
        } else {
          setChapters(res.data.chapters)
        }
      })
      .catch(e => {
        console.log(e)
      })
  }
  const getTasks = () => {
    TaskApi.getProjectTasksAndSubmits(pid)
      .then(res => {
        if (res.data.code === 200) {
          if (res.data.data != null) {
            setTasks(res.data.data)
          }
        }
      })
      .catch(e => {
        console.log(e)
      })
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
              <span style={{color: 'green'}}>已打分</span>
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
      <Divider orientation="left">
        <p className="evidence-title">章节学习时长</p>
      </Divider>
      {chapters.map((item, index) => (
        <div key={index.toString()} style={{textAlign: 'left'}}>
          <p style={{fontWeight: 'bold', fontSize: '16px'}}>{item.chapterName}</p>
          {(item.sections === null || item.sections === undefined) ?
            null
            :
            <>
              <List
                size="large"
                dataSource={item.sections}
                renderItem={
                  item => (
                    <List.Item>
                      {item.sectionName}
                      {obj.project.learning ?
                        <>
                      <span style={{float: 'right'}}>
                        <Progress
                          trailColor="lightgray"
                          width={30}
                          strokeWidth={10}
                          type="circle"
                          percent={((item.learnMinute + item.learnSecond / 60) / item.sectionMinute * 100).toFixed(1)}
                        />
                      </span>
                          <span style={{float: 'right', marginRight: '20px'}}>
                        学习进度：
                            {item.learnMinute}&nbsp;:&nbsp;{item.learnSecond}&nbsp;/&nbsp;
                            {item.sectionMinute}
                      </span>
                        </>
                        : null}
                    </List.Item>
                  )
                }
              /><br/>
            </>
          }
        </div>
      ))}

      <Divider orientation="left">
        <p className="evidence-title">学生任务</p>
      </Divider>

      <Collapse style={{textAlign: 'left'}}>
        {tasks.map((item, index) => (
          <Collapse.Panel
            key={index.toString()}
            header={
              <>
                {item.taskTitle}
                <span style={{float: 'right', marginRight: '20px'}}>
                  {item.scored ?
                    <span style={{color: 'green'}}>已打分</span>
                    :
                    <span>未打分</span>
                  }
                </span>
                <span style={{float: 'right', marginRight: '20px'}}>
                  权重：{item.taskWeight}
                </span>
                <span style={{float: 'right', marginRight: '20px'}}>
                  得分：{getScore(item.score, item.taskWeight)}&nbsp;/&nbsp;{item.taskWeight}
                </span>
              </>
            }
          >
            neironog
          </Collapse.Panel>
        ))}
      </Collapse>

      <Table
        columns={getColumns()}
        dataSource={tasks}
        pagination={false}
      />
    </div>
  )
}

export default StudentEvidence