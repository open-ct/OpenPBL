import React, {useEffect, useState} from "react";
import {Button, Card, Checkbox, Divider, Layout, Menu, message, PageHeader, Popconfirm, Radio} from "antd";
import DocumentTitle from 'react-document-title';

import SurveyApi from "../../../api/SurveyApi";
import {ArrowDownOutlined, ArrowUpOutlined, DeleteOutlined, EditOutlined} from "@ant-design/icons";

const qType = {
  'singleChoice': '单选',
  'multipleChoice': '多选',
  'blankFill': '填空',
  'briefAnswer': '简答',
  'scale5': '五点量表',
  'scale7': '七点量表'
}

function SurveyEditPage(obj) {
  const pid = obj.match.params.pid
  const sid = obj.match.params.sid
  const tid = obj.match.params.tid
  const [survey, setSurvey] = useState({})
  const [questions, setQuestions] = useState([])

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

            console.log(qs)
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
      opt = ""
    } else {

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

  const back = e => {
    window.location.href = '/project/edit/section/' + pid + '/' + sid
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
              <Layout.Content style={{backgroundColor: 'white'}}>
                <div style={{margin: '30px'}}>
                  <h2>{survey.surveyTitle}</h2>
                  {questions.map((item, index) => (
                    <div key={index.toString()}>
                      <Divider/>
                      <div>
                        <p style={{float: 'left'}}>{item.questionTitle}
                          <span style={{color: 'gray'}}>
                          &nbsp;[{getType(item.questionType)}]
                        </span>
                        </p>
                        <p style={{float: 'right'}}>
                          <Button shape="circle" type="text" icon={<EditOutlined/>}/>
                          <Button shape="circle" type="text" icon={<ArrowUpOutlined/>}
                                  onClick={e => exchangeQuestion(index - 1, index)}/>
                          <Button shape="circle" type="text" icon={<ArrowDownOutlined/>}
                                  onClick={e => exchangeQuestion(index, index + 1)}/>
                          <Popconfirm title="确定删除小节？" onConfirm={e => deleteQuestion(item, index)} placement="topLeft">
                            &nbsp;&nbsp;<Button shape="circle" type="text" style={{color: 'red'}}
                                                icon={<DeleteOutlined/>}/>
                          </Popconfirm>
                        </p>
                      </div>
                      <br/><br/>
                      {item.questionType === 'singleChoice' ?
                        <div style={{textAlign: "left", marginTop: '10px'}}>
                          <Radio.Group>
                            {item.questionOptions.map((subItem, subIndex) => (
                              <div style={{marginTop: '10px'}} key={subIndex.toString()}>
                                <Radio value={subIndex}>
                                  {subItem}
                                </Radio>
                              </div>
                            ))}
                          </Radio.Group>
                        </div>
                        : null}
                      {item.questionType === 'multipleChoice' ?
                        <div style={{textAlign: "left", marginTop: '10px'}}>
                          {item.questionOptions.map((subItem, subIndex) => (
                            <div style={{marginTop: '10px'}} key={subIndex.toString()}>
                              <Checkbox>
                                {subItem}
                              </Checkbox>
                            </div>
                          ))}
                        </div> : null}
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
