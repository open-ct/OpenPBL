import React, {useEffect, useState} from "react";
import {Col, Collapse, Divider, List, Progress, Row, Table} from "antd";

import TaskApi from "../../../../api/TaskApi";
import ChapterApi from "../../../../api/ChapterApi";
import TaskCard from "../../PreviewProject/component/TaskCard";
import util from "../../component/Util"

function StudentEvidence(obj) {
  const pid = obj.project === undefined ? obj.pid : obj.project.id
  const studentId = obj.studentId
  const [tasks, setTasks] = useState([])
  const [learning, setLearning] = useState(false)
  const [editable, setEditable] = useState(false)
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
    TaskApi.getProjectTasksDetail(pid, studentId)
      .then(res => {
        if (res.data.code === 200) {
          if (res.data.tasks === null) {
            setTasks([])
          } else {
            let t = res.data.tasks
            for (let i = 0; i < t.length; i++) {
              if (t[i].questions !== undefined && t[i].questions != null) {
                for (let j = 0; j < t[i].questions.length; j++) {
                  t[i].questions[j].questionOptions = t[i].questions[j].questionOptions.split(",")
                }
              } else {
                t[i].questions = []
              }
              if (t[i].choices !== undefined && t[i].choices != null) {
                for (let j = 0; j < t[i].choices.length; j++) {
                  t[i].choices[j].choiceOptions = t[i].choices[j].choiceOptions.split(",")
                }
              } else {
                t[i].choices = []
                for (let j=0; j<t[i].questions.length; j++) {
                  t[i].choices.push({
                    choiceOptions: [],
                    choiceOrder: j
                  })
                }
              }
            }
            setTasks(t)
            setLearning(res.data.learning)
            setEditable(res.data.editable)
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
    if (learning) {
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
  const setTaskItem = (item, index) => {
    tasks[index] = item
    setTasks([...tasks])
  }
  return (
    <div>
      <Divider orientation="left">
        <p className="evidence-title">章节学习时长</p>
      </Divider>
      {chapters.map((item, index) => (
        <div key={index.toString()} style={{textAlign: 'left'}}>
          <p style={{fontWeight: 'bold', fontSize: '16px'}}>
            {util.FormatChapterName(item.chapterName, item.chapterNumber)}
          </p>
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
                      {util.FormatSectionName(item.sectionName, item.chapterNumber, item.sectionNumber)}
                      {learning ?
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
                {util.FormatSectionName(item.taskTitle, item.chapterNumber, item.sectionNumber)}
                <span style={{float: 'right', marginRight: '20px'}}>
                  {item.submit.scored ?
                    <span style={{color: 'green'}}>已打分</span>
                    :
                    <span>未打分</span>
                  }
                </span>
                <span style={{float: 'right', marginRight: '20px'}}>
                  {item.submitted ?
                    <span style={{color: 'green'}}>已提交</span>
                    :
                    <span>未提交</span>
                  }
                </span>
                <span style={{float: 'right', marginRight: '20px'}}>
                  权重：{item.taskWeight}
                </span>
                <span style={{float: 'right', marginRight: '20px'}}>
                  得分：{getScore(item.submit.score, item.taskWeight)}&nbsp;/&nbsp;{item.taskWeight}
                </span>
              </>
            }
          >
            <TaskCard
              pid={pid}
              item={item}
              index={index}
              learning={learning}
              editable={editable}
              setTaskItem={setTaskItem}
              getTasks={getTasks}
            />
          </Collapse.Panel>
        ))}
      </Collapse>

{/*      <Table
        columns={getColumns()}
        dataSource={tasks}
        pagination={false}
      />*/}
    </div>
  )
}

export default StudentEvidence