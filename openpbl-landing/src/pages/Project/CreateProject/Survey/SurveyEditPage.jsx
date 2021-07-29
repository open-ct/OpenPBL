import React, {useEffect, useState} from "react";
import {Affix, Card, Divider, Layout, Menu, message, PageHeader, Radio} from "antd";
import DocumentTitle from 'react-document-title';

import SurveyApi from "../../../../api/SurveyApi";

import SingleChoice from "./component/SingleChoice";
import MultipleChoice from "./component/MultipleChoice";
import BlankFill from "./component/BlankFill";
import BriefAnswer from "./component/BriefAnswer";
import Question from "./component/Question";
import Scale from "./component/Scale";

const qType = {
  'singleChoice': '单选',
  'multipleChoice': '多选',
  'blankFill': '填空',
  'briefAnswer': '简答',
  'scale5': '五点量表',
  'scale7': '七点量表'
}

const blank = Question.blank

function SurveyEditPage(obj) {
  const pid = obj.match.params.pid
  const sid = obj.match.params.sid
  const tid = obj.match.params.tid

  const [survey, setSurvey] = useState({})
  const [questions, setQuestions] = useState([])
  const [editing, setEditing] = useState([])

  useEffect(() => {
    if (tid !== undefined) {
      SurveyApi.getSurveyDetailByTaskId(tid)
        .then(res => {
          if (res.data.questions === null) {
            setQuestions([])
          } else {
            let qs = res.data.questions
            for (let i = 0; i < qs.length; i++) {
              qs[i].questionOptions = qs[i].questionOptions.split(",")
            }
            setQuestions(qs)
            setEditing(new Array(qs.length).fill(false))
          }
          setSurvey(res.data.survey)
        })
    }
  }, [tid])

  const createQuestion = (e) => {
    let len = questions.length
    let l = 0
    if (len > 0) {
      l = questions[len - 1].questionOrder + 1
    }
    let opt = ""
    if (e.key === 'singleChoice' || e.key === 'multipleChoice') {
      opt = ['选项1', '选项2']
    } else if (e.key === 'blankFill') {
      opt = ['题目描述', blank]
    } else if (e.key === 'briefAnswer') {
      opt = ['题目描述']
    } else if (e.key === 'scale5') {
      opt = ['选项1', '选项2', '选项3', '选项4', '选项5']
    } else {
      opt = ['选项1', '选项2', '选项3', '选项4', '选项5', '选项6', '选项7']
    }
    let q = {
      surveyId: survey.id,
      questionType: e.key,
      questionOrder: l,
      questionTitle: '标题',
      questionOptions: opt
    }
    q.questionOptions = q.questionOptions.toString()
    SurveyApi.createQuestion(q)
      .then(res => {
        if (res.data.code === 200) {
          q.id = res.data.data
          q.questionOptions = q.questionOptions.split(",")
          questions.push(q)
          setQuestions([...questions])
        } else {
          message.error(res.data.msg)
        }
      })
  }
  const saveQuestion = (item, title, opt, index) => {
    let q = Object.assign({}, item)
    q.questionOptions = opt.toString()
    q.questionTitle = title
    SurveyApi.updateQuestion(q)
      .then(res => {
        editing[index] = false
        setEditing([...editing])
        if (res.data.code === 200) {
          q.questionOptions = q.questionOptions.split(",")
          questions[index] = q
          setQuestions([...questions])
          message.success(res.data.msg)
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e => {
        console.log(e)
      })
  }

  const deleteQuestion = (item, index) => {
    SurveyApi.deleteQuestion(item.id)
      .then(res => {
        if (res.data.code === 200) {
          questions.splice(index, 1)
          setQuestions([...questions])
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e => {
        console.log(e)
      })
  }
  const exchangeQuestion = (index1, index2) => {
    if (index1 < 0 || index2 >= questions.length) {
    } else {
      SurveyApi.exchangeQuestion(questions[index1].id, questions[index2].id)
        .then(res => {
          if (res.data.code === 200) {
            let q1 = questions[index1]
            questions[index1] = questions[index2]
            questions[index2] = q1
            setQuestions([...questions])
          } else {
            message.error(res.data.msg)
          }
        })
    }
  }

  const editQuestion = (item, index) => {
    editing[index] = true
    setEditing([...editing])
  }
  const back = e => {
    window.location.href = `/project/${pid}/section/${sid}/edit`
  }

  const getType = t => {
    return qType[t]
  }

  return (
    <DocumentTitle title="Project">
      <div style={{backgroundColor: '#f2f4f5', minHeight: '100vh'}}>
        <PageHeader
          className="site-page-header"
          onBack={back}
          title="返回"
          subTitle="编辑资源"
        />
        <div style={{padding: '20px', margin: 'auto', maxWidth: '1200px'}}>
          <Card>
            <Layout>
              <Affix offsetTop={0}>
                <Layout.Sider breakpoint="lg" collapsedWidth="0" style={{backgroundColor: '#f2f4f5'}}>
                  <Menu theme="light" mode="inline" onClick={createQuestion}>
                    <Menu.SubMenu title="选择" key="1">
                      <Menu.Item key="singleChoice">单选</Menu.Item>
                      <Menu.Item key="multipleChoice">多选</Menu.Item>
                    </Menu.SubMenu>
                    <Menu.SubMenu title="问答" key="2">
                      <Menu.Item key="blankFill">填空</Menu.Item>
                      <Menu.Item key="briefAnswer">简答</Menu.Item>
                    </Menu.SubMenu>
                    <Menu.SubMenu title="量表" key="3">
                      <Menu.Item key="scale5">五点量表</Menu.Item>
                      <Menu.Item key="scale7">七点量表</Menu.Item>
                    </Menu.SubMenu>
                  </Menu>
                </Layout.Sider>
              </Affix>
              <Layout.Content style={{backgroundColor: 'white'}}>
                <div style={{margin: '30px'}}>
                  <h2>{survey.surveyTitle}</h2>
                  {questions.map((item, index) => (
                    <div key={index.toString()}>
                      <Divider/>
                      {item.questionType === 'singleChoice' ?
                        <SingleChoice
                          item={item}
                          index={index}
                          editing={editing[index]}
                          getType={getType}
                          editQuestion={editQuestion}
                          saveQuestion={saveQuestion}
                          deleteQuestion={deleteQuestion}
                          exchangeQuestion={exchangeQuestion}
                        />
                        : null}

                      {item.questionType === 'scale5' || item.questionType === 'scale7' ?
                        <Scale
                          item={item}
                          index={index}
                          editing={editing[index]}
                          getType={getType}
                          editQuestion={editQuestion}
                          saveQuestion={saveQuestion}
                          deleteQuestion={deleteQuestion}
                          exchangeQuestion={exchangeQuestion}
                        />
                        : null}

                      {item.questionType === 'multipleChoice' ?
                        <MultipleChoice
                          item={item}
                          index={index}
                          editing={editing[index]}
                          getType={getType}
                          editQuestion={editQuestion}
                          saveQuestion={saveQuestion}
                          deleteQuestion={deleteQuestion}
                          exchangeQuestion={exchangeQuestion}
                        />
                        : null}

                      {item.questionType === 'blankFill' ?
                        <BlankFill
                          item={item}
                          index={index}
                          editing={editing[index]}
                          getType={getType}
                          editQuestion={editQuestion}
                          saveQuestion={saveQuestion}
                          deleteQuestion={deleteQuestion}
                          exchangeQuestion={exchangeQuestion}
                        />
                        : null}

                      {item.questionType === 'briefAnswer' ?
                        <BriefAnswer
                          item={item}
                          index={index}
                          editing={editing[index]}
                          getType={getType}
                          editQuestion={editQuestion}
                          saveQuestion={saveQuestion}
                          deleteQuestion={deleteQuestion}
                          exchangeQuestion={exchangeQuestion}
                        />
                        : null}
                    </div>
                  ))}
                </div>
              </Layout.Content>
            </Layout>
          </Card>
        </div>
      </div>
    </DocumentTitle>
  )
}

export default SurveyEditPage
