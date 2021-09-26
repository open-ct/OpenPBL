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
import {Button, Checkbox, Divider, List, Radio} from "antd";

import TaskApi from "../../../../api/TaskApi";
import ChapterApi from "../../../../api/ChapterApi";
import util from "../../../component/Util"
import htmlDocx from 'html-docx-js/dist/html-docx';
import saveAs from 'file-saver';
import Question from "../../CreateProject/Survey/component/Question";

const blank = Question.blank


function StudentEvidenceContent(obj) {
  const pid = obj.match.params.projectId
  const studentId = obj.match.params.studentId
  const [tasks, setTasks] = useState([])
  const [learning, setLearning] = useState(false)
  const [editable, setEditable] = useState(false)
  const [teacherScore, setTeacherScore] = useState(false)
  const [showCount, setShowCount] = useState(false)
  const [chapters, setChapters] = useState([])
  const [showMinute, setShowMinute] = useState(false)

  useEffect(() => {
    getChapters()
    getTasks()
    setTimeout(() => {
      saveFile()
    }, 500)
  }, []);
  const getChapters = () => {
    ChapterApi.getProjectChapters(pid, studentId)
      .then((res) => {
        if (res.data.chapters === null) {
          setChapters([])
        } else {
          setChapters(res.data.chapters)
          setShowMinute(res.data.showMinute)
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
                for (let j = 0; j < t[i].questions.length; j++) {
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
            setTeacherScore(res.data.teacherScore)
            setShowCount(res.data.showCount)
          }
        }
      })
      .catch(e => {
        console.log(e)
      })
  }

  const getScore = (score, weight) => {
    return (score * weight / 100).toFixed(2)
  }
  const getPercent = (item) => {
    let p = ((item.learnMinute + item.learnSecond / 60) / item.sectionMinute * 100).toFixed(1)
    if (p > 100) {
      return 100
    }
    return p
  }
  const saveFile = () => {
    let content = document.getElementById("evidence").innerHTML
    let page = `<!DOCTYPE html><html><head><meta charset="UTF-8"></head><body>${content}</body></html>`
    const converted = htmlDocx.asBlob(page)
    saveAs(converted, `${studentId}.docx`)
  }

  return (
    <div style={{padding: '40px'}}>
      <div style={{float: 'right'}}><Button onClick={saveFile}>保存</Button></div>

      <div id="evidence">
        <h2>一、章节学习时长</h2>
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
                  locale={{emptyText: '--'}}
                  renderItem={
                    item => (
                      <List.Item>
                        <p>{util.FormatSectionName(item.sectionName, item.chapterNumber, item.sectionNumber)}</p>
                        要求学习时长：{item.sectionMinute}分钟,&nbsp;
                        学生学习时长：{item.learnMinute}分{item.learnSecond}秒,&nbsp;
                        学习进度：{getPercent(item)}%
                      </List.Item>
                    )
                  }
                /><br/>
              </>
            }
          </div>
        ))}

        <h2>二、学生任务</h2>

        <List
          size="large"
          dataSource={tasks}
          locale={{emptyText: '--'}}
          renderItem={
            item => (
              <List.Item>
                <div style={{textAlign: 'left'}}>
                  <div>
                    <p
                      style={{fontWeight: 'bold'}}>{util.FormatSectionName(item.taskTitle, item.chapterNumber, item.sectionNumber)}</p>
                    {item.submitted ?
                      <span style={{color: 'green'}}>已提交</span>
                      :
                      <span>未提交</span>
                    },&nbsp;
                    {item.submit.scored ?
                      <span style={{color: 'green'}}>已打分</span>
                      :
                      <span>未打分</span>
                    }&nbsp;&nbsp;
                    任务权重：{item.taskWeight}%,&nbsp;
                    学生得分：{getScore(item.submit.score, item.taskWeight)}/{item.taskWeight}
                  </div>
                  <div>
                    标题：
                    <p>{item.taskTitle}</p>
                    描述：
                    <p>{item.taskIntroduce}</p>
                    {item.taskType === 'file' ?
                      <div>
                        <p>{item.submit.submitContent}</p>
                      </div>
                      : null
                    }
                    {item.taskType === 'comment' ?
                      <div>
                        学生提交：<br/>{item.submit.submitContent}
                      </div>
                      : null
                    }
                    {item.taskType === 'survey' ?
                      <>学生提交：<br/>
                        <h2 style={{textAlign: 'center'}}>{item.survey.surveyTitle}</h2>
                        <div style={{marginLeft: '20px', marginRight: '20px'}}>
                          {item.questions.map((subItem, subIndex) => (
                            <div key={subIndex.toString()}>
                              {subItem.questionType === 'singleChoice' || subItem.questionType === 'scale5' || subItem.questionType === 'scale7' ?
                                <div style={{textAlign: "left", marginTop: '10px'}}>
                                  <p>{subItem.questionTitle}</p>
                                  <Radio.Group value={item.choices[subIndex].choiceOptions[0]}>
                                    {subItem.questionOptions.map((optItem, optIndex) => (
                                      <div style={{marginTop: '10px'}} key={optIndex.toString()}>
                                        <Radio value={optIndex.toString()}>
                                          {optItem}
                                        </Radio>
                                      </div>
                                    ))}
                                  </Radio.Group>
                                </div>
                                : null}
                              {subItem.questionType === 'multipleChoice' ?
                                <div style={{textAlign: "left", marginTop: '10px'}}>
                                  <p>{subItem.questionTitle}</p>
                                  <Checkbox.Group value={item.choices[subIndex].choiceOptions}>
                                    {subItem.questionOptions.map((optItem, optIndex) => (
                                      <div key={optIndex.toString()} style={{marginTop: '10px'}}>
                                        <Checkbox value={optIndex.toString()}>
                                          {optItem}
                                        </Checkbox>
                                      </div>
                                    ))}
                                  </Checkbox.Group>
                                </div>
                                : null}
                              {subItem.questionType === 'blankFill' ?
                                <div style={{textAlign: "left", marginTop: '10px'}}>
                                  <p>{subItem.questionTitle}</p>
                                  {subItem.questionOptions.map((optItem, optIndex) => (
                                    <div style={{marginTop: '10px'}} key={optIndex.toString()}>
                                      {optItem === blank ?
                                        <span style={{float: 'left', margin: '5px'}}>
                                          <u>&nbsp;{item.choices[subIndex].choiceOptions[optIndex]}&nbsp;</u>
                                        </span>
                                        :
                                        <span style={{float: 'left', margin: '5px'}}>{optItem}</span>
                                      }
                                    </div>
                                  ))}
                                </div>
                                : null}

                              {subItem.questionType === 'briefAnswer' ?
                                <div style={{textAlign: "left", marginTop: '10px'}}>
                                  <p>{subItem.questionOptions[0]}</p>
                                  {item.choices[subIndex].choiceOptions[0]}
                                </div>
                                : null}

                              <Divider/>
                            </div>
                          ))}
                        </div>

                      </>
                      : null
                    }
                  </div>
                </div>
              </List.Item>
            )
          }
        />
      </div>
    </div>
  )
}

export default StudentEvidenceContent