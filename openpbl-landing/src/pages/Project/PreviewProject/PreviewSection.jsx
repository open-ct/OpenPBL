import React, {useEffect, useState} from "react";
import {Card, PageHeader, Input, Upload, message, Button, BackTop} from "antd";
import DocumentTitle from 'react-document-title';
import {InboxOutlined} from '@ant-design/icons'

import SectionApi from "../../../api/SectionApi";
import "../CreateProject/Section/component/section-edit.less"
import "./preview.less"
import TaskApi from "../../../api/TaskApi";
import util from "../component/Util"
import StudentApi from "../../../api/StudentApi";
import TaskCard from "./component/TaskCard";

function PreviewSection(obj) {
  let url = new URLSearchParams(obj.location.search)
  const backUrl = url.get('back')

  const sid = obj.match.params.sid
  const pid = obj.match.params.pid
  const [section, setSection] = useState({resource:{}})
  const [tasks, setTasks] = useState([])
  const [learning, setLearning] = useState(false)
  const [editable, setEditable] = useState(false)

  const [minute, setMinute] = useState(0)
  const [second, setSecond] = useState(0)
  const [timer, setTimer] = useState(null)
  let s = 0
  let m = 0

  useEffect(()=>{
    getSectionDetail()
    getTasks()
  }, [])
  useEffect(()=>{
    window.onbeforeunload = leave
  }, [])
  const leave = () => {
    if (timer != null) {
      clearTimeout(timer)
    }
    let data = {
      learnMinute: m,
      learnSecond: s
    }
    StudentApi.updateLearnSection(pid, sid, data)
      .then(res=>{
      })
      .catch(e=>{console.log(e)})
    window.removeEventListener("onbeforeunload", leave)
  }

  const getSectionDetail = () => {
    SectionApi.getSectionDetail(sid, pid)
      .then(res=>{
        setSection(res.data.section)
      })
      .catch(e=>{console.log(e)})
  }
  const getTasks = () => {
    TaskApi.getSectionTasks(sid, pid)
      .then(res=>{
        if (res.data.tasks === null) {
          setTasks([])
        } else {
          let t = res.data.tasks
          for (let i = 0; i < t.length; i++) {
            if (t[i].questions !== undefined && t[i].questions != null) {
              for (let j = 0; j < t[i].questions.length; j++) {
                t[i].questions[j].questionOptions = t[i].questions[j].questionOptions.split(",")
                t[i].questions[j].questionCount = t[i].questions[j].questionCount.split(",")
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
        }
        setLearning(res.data.learning)
        setEditable(res.data.editable)

        if (res.data.learning) {
          getTimer()
        }
      })
      .catch(e=>{console.log(e)})
  }
  const getTimer = () => {
    StudentApi.getLearnSection(pid, sid)
      .then(res=>{
        if (res.data.code === 200) {
          setSecond(res.data.data.learnSecond)
          setMinute(res.data.data.learnMinute)
          s = res.data.data.learnSecond
          m = res.data.data.learnMinute
        }
      })
      .catch(e=>{console.log(e)})

    setTimeout(count, 1000)
  }
  const count = () => {
    if (s === 30) {
      s = 0
      m ++
      setMinute(m)
    } else {
      s++
    }
    setSecond(s)
    setTimeout(count, 1000)
  }
  const back = e => {
    if (backUrl === undefined || backUrl === null) {
      window.location.href = `/project/${pid}/section/${sid}/edit`
    } else {
      window.location.href = backUrl
    }
  }
  const setTaskItem = (item, index) => {
    tasks[index] = item
    setTasks([...tasks])
  }

  return (
  <DocumentTitle title="Project">
    <div style={{backgroundColor: '#f2f4f5', minHeight: '100vh'}}>
      <BackTop />
      <PageHeader
        className="site-page-header"
        onBack={back}
        title="返回"
        subTitle="我的项目"
      />
      <div style={{ padding: '20px', margin: 'auto'}}>
        <Card>
          <h2 style={{ fontWeight: 'bold' }}>
            {util.FormatSectionName(section.sectionName, section.chapterNumber, section.sectionNumber)}
          </h2>
          {learning ?
            <span style={{float: 'right'}}>{minute}&nbsp;:&nbsp;{second}</span>
            : null}
        </Card>
        <Card className="resource-card">
          <div dangerouslySetInnerHTML={{__html: section.resource.content}}/>
        </Card>
        <Card className="resource-card">
          <p className="card-title">文件资源</p>
          <p>{section.resource.fileTitle}</p>
          <p>{section.resource.fileIntroduce}</p>
        </Card>
        {tasks.map((item, index)=>(
          <Card className="resource-card" key={index.toString()}>
            <p className="card-title">学生任务
              {item.submitted ?
                <span className="submit-status" style={{color: 'green'}}>
                  权重占比&nbsp;{item.taskWeight}&nbsp;%&nbsp;&nbsp;已提交&nbsp;&nbsp;{util.FilterTime(item.submit.createAt)}
                </span>
                :
                <span className="submit-status" style={{color: 'gray'}}>
                  权重占比&nbsp;{item.taskWeight}&nbsp;%&nbsp;&nbsp;未提交
                </span>
              }
            </p>

            <TaskCard
              pid={pid}
              item={item}
              index={index}
              learning={learning}
              editable={editable}

              setTaskItem={setTaskItem}
              getTasks={getTasks}
            />
          </Card>
        ))
        }
      </div><br/>
    </div>
  </DocumentTitle>
)
}

export default PreviewSection
