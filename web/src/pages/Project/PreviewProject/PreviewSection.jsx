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

import React from "react";
import {BackTop, Card, PageHeader, Divider} from "antd";
import DocumentTitle from 'react-document-title';

import SectionApi from "../../../api/SectionApi";
import "../CreateProject/Section/component/section-edit.less"
import "./preview.less"
import TaskApi from "../../../api/TaskApi";
import util from "../../component/Util"
import StudentApi from "../../../api/StudentApi";
import TaskCard from "./component/TaskCard";

class PreviewSection extends React.Component {
  constructor(props) {
    super(props);
    let url = new URLSearchParams(this.props.location.search)
    const backUrl = url.get('back')
    this.state = {
      backUrl: backUrl,
      sid: this.props.match.params.sid,
      pid: this.props.match.params.pid,
      section: {resource: {}},
      sectionFiles: [],
      tasks: [],
      learning: false,
      editable: false,
      showCount: false,
      timer: null,
      minute: 0,
      second: 0,
    }
  }

  componentDidMount() {
    this.getSectionDetail()
    this.getTasks()
    this.getSectionFiles()
    window.addEventListener("beforeunload", this.saveTime)
  }

  componentWillUnmount() {
    this.saveTime()
  }

  saveTime = () => {
    if (this.state.learning && this.state.editable) {
      if (this.state.timer != null) {
        clearTimeout(this.state.timer)
      }
      let data = {
        learnMinute: this.state.minute,
        learnSecond: this.state.second
      }
      StudentApi.updateLearnSection(this.state.pid, this.state.sid, data)
        .then(res => {
        })
        .catch(e => {
          console.log(e)
        })
    }
  }

  getSectionDetail = () => {
    SectionApi.getSectionDetail(this.state.sid, this.state.pid)
      .then(res => {
        this.setState({
          section: res.data.section
        })
      })
      .catch(e => {
        console.log(e)
      })
  }
  getTasks = () => {
    TaskApi.getSectionTasks(this.state.sid, this.state.pid)
      .then(res => {
        if (res.data.tasks === null) {
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
              for (let j = 0; j < t[i].questions.length; j++) {
                t[i].choices.push({
                  choiceOptions: [],
                  choiceOrder: j
                })
              }
            }
          }
          this.setState({
            tasks: t
          })
        }
        this.setState({
          learning: res.data.learning,
          editable: res.data.editable,
          showCount: res.data.showCount
        })

        if (res.data.learning && res.data.editable && this.state.timer === null) {
          this.setTimer()
        }
      })
      .catch(e => {
        console.log(e)
      })
  }
  setTimer = () => {
    StudentApi.getLearnSection(this.state.pid, this.state.sid)
      .then(res => {
        if (res.data.code === 200) {
          let m = res.data.data.learnMinute
          let s = res.data.data.learnSecond
          this.setState({
            minute: m,
            second: s
          })
          if (this.state.timer !== null) {
            clearTimeout(this.state.timer)
          }
          this.state.timer = setTimeout(this.count, 1000)
        }
      })
      .catch(e => {
        console.log(e)
      })
  }
  count = () => {
    if (this.state.second >= 59) {
      this.setState({
        second: 0,
        minute: ++this.state.minute
      })
    } else {
      this.setState({
        second: ++this.state.second
      })
    }
    this.state.timer = setTimeout(this.count, 1000)
  }
  back = e => {
    if (this.state.backUrl === undefined || this.state.backUrl === null) {
      this.props.history.push(`/home/project/${this.state.pid}/section/${this.state.sid}/edit`)
    } else {
      this.props.history.push(this.state.backUrl)
    }
  }
  setTaskItem = (item, index) => {
    this.state.tasks[index] = item
    this.setState({
      tasks: [...this.state.tasks]
    })
  }
  getSectionFiles = () => {
    SectionApi.getSectionFiles(this.state.sid, this.state.pid)
      .then(res=>{
        if (res.data.code === 200) {
          if (res.data.data !== null) {
            this.setState({
              sectionFiles: res.data.data
            })
          }
        }
      })
      .catch(e=>{console.log(e)})
  }

  render() {
    const {section, tasks, learning, pid, editable, showCount, minute, second, sectionFiles} = this.state
    return (
      <DocumentTitle title="Project">
        <div style={{backgroundColor: '#f2f4f5', minHeight: '100vh'}}>
          <BackTop/>
          <PageHeader
            className="site-page-header"
            onBack={this.back}
            title="返回"
            subTitle="我的项目"
          />
          <div style={{padding: '20px', margin: 'auto', maxWidth: '1400px'}}>
            <Card>
              <h2 style={{fontWeight: 'bold'}}>
                {util.FormatSectionName(section.sectionName, section.chapterNumber, section.sectionNumber)}
              </h2>
              {learning && editable ?
                <span style={{float: 'right'}}>{minute}&nbsp;分&nbsp;{second}&nbsp;秒</span>
                : null}
            </Card>
            <Card className="resource-card">
              <div dangerouslySetInnerHTML={{__html: section.resource.content}}/>
            </Card>
            <Card className="resource-card">
              <p className="card-title">文件资源</p>
              <Divider />
              {sectionFiles.map((item, index)=>(
                <div>
                  <a target="_blank" href={item.url}>{item.name}</a>
                  <Divider />
                </div>
              ))}
            </Card>
            {tasks.map((item, index) => (
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
                  account={this.props.account}
                  item={item}
                  index={index}
                  learning={learning}
                  editable={editable}
                  showCount={showCount}
                  setTaskItem={this.setTaskItem}
                  getTasks={this.getTasks}
                />
              </Card>
            ))
            }
          </div>
          <br/>
        </div>
      </DocumentTitle>
    )
  }
}

export default PreviewSection
