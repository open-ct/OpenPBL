import React, {useEffect, useState} from "react";
import {Button, Checkbox, Divider, Input, Radio, message, Collapse} from "antd";

import Question from "../../CreateProject/Survey/component/Question"
import "../preview.less"
import "../../CreateProject/Section/component/section-edit.less"
import SubmitApi from "../../../../api/SubmitApi";
import ReactEcharts from "echarts-for-react";

const blank = Question.blank

function FillSurvey(obj) {
  const [submitLoading, setSubmitLoading] = useState(false)

  const submitSurvey = e => {
    setSubmitLoading(true)
    obj.item.submit.submitType = 'survey'
    let data = Object.assign({}, obj.item.submit)
    let c = [...obj.item.choices]
    for (let i=0; i<c.length; i++) {
      c[i].questionId = obj.item.questions[i].id
      c[i].choiceOptions = c[i].choiceOptions.toString()
    }
    data['choices'] = JSON.stringify(c)
    SubmitApi.createSubmit(obj.pid, obj.item.id, data)
      .then(res=>{
        setSubmitLoading(false)
        if (res.data.code === 200) {
          message.success(res.data.msg)
          obj.getTasks()
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e=>{console.log(e)})
  }
  const updateSurvey = e => {
    setSubmitLoading(true)
    obj.item.submit.submitType = 'survey'

    let data = Object.assign({}, obj.item.submit)
    let c = [...obj.item.choices]
    for (let i=0; i<c.length; i++) {
      c[i].choiceOptions = c[i].choiceOptions.toString()
    }
    data['choices'] = JSON.stringify(c)
    SubmitApi.updateSubmit(obj.pid, obj.item.id, obj.item.submit.id, data)
      .then(res=>{
        setSubmitLoading(false)
        if (res.data.code === 200) {
          message.success(res.data.msg)
          obj.getTasks()
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e=>{console.log(e)})

  }
  const changeCheckBox = (v, subIndex) => {
    obj.item.choices[subIndex].choiceOptions = v
    obj.setTaskItem(obj.item, obj.index)
  }
  const changeRadio = (v, subIndex) => {
    obj.item.choices[subIndex].choiceOptions[0] = v.target.value
    obj.setTaskItem(obj.item, obj.index)
  }
  const changeBlankFill = (v, subIndex, optIndex) => {
    obj.item.choices[subIndex].choiceOptions[optIndex] = v.target.value
    obj.setTaskItem(obj.item, obj.index)
  }
  const changeBriefAnswer = (v, subIndex) => {
    obj.item.choices[subIndex].choiceOptions[0] = v.target.value
    obj.setTaskItem(obj.item, obj.index)
  }


  const showCount = e => {
    if (e.length > 0) {
      console.log('draw')
    }
  }
  const getOption = subIndex => {
    return {
      title: {
        x: 'center',
      },
      tooltip: {
        trigger: 'item',
      },
      xAxis: {
        type: 'category',
        data: obj.item.questions[subIndex].questionOptions
      },
      yAxis: {
        minInterval: 1,
      },
      series: [
        {
          type: 'bar',
          data: obj.item.questions[subIndex].questionCount
        },
      ],
    }
  }

  return (
    <div className="survey">
      <h2 style={{textAlign: 'center'}}>{obj.item.survey.surveyTitle}</h2>
      <Divider />
      <div style={{marginLeft: '20px', marginRight: '20px'}}>
        {obj.item.questions.map((subItem, subIndex)=>(
          <div key={subIndex.toString()}>
            {subItem.questionType==='singleChoice' || subItem.questionType==='scale5' || subItem.questionType==='scale7' ?
              <div style={{textAlign: "left", marginTop: '10px'}}>
                <p>{subItem.questionTitle}</p>
                <Radio.Group value={obj.item.choices[subIndex].choiceOptions[0]} onChange={v=>changeRadio(v, subIndex)}>
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
            {subItem.questionType==='multipleChoice' ?
              <div style={{textAlign: "left", marginTop: '10px'}}>
                <p>{subItem.questionTitle}</p>
                <Checkbox.Group value={obj.item.choices[subIndex].choiceOptions} onChange={v=>changeCheckBox(v, subIndex)}>
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
            {subItem.questionType==='blankFill' ?
              <div style={{textAlign: "left", marginTop: '10px'}}>
                <p>{subItem.questionTitle}</p>
                {subItem.questionOptions.map((optItem, optIndex) => (
                  <div style={{marginTop: '10px'}} key={optIndex.toString()}>
                    {optItem === blank ?
                      <span style={{float: 'left', margin: '5px'}}>
                        <Input value={obj.item.choices[subIndex].choiceOptions[optIndex]} onChange={v=>changeBlankFill(v, subIndex, optIndex)} style={{borderBottom: '1.5px solid black'}} bordered={false}/>
                      </span>
                      :
                      <span style={{float: 'left', margin: '5px'}}>{optItem}</span>}
                  </div>
                ))}
              </div>
              : null}

            {subItem.questionType==='briefAnswer' ?
              <div style={{textAlign: "left", marginTop: '10px'}}>
                <p>{subItem.questionOptions[0]}</p>
                <Input.TextArea value={obj.item.choices[subIndex].choiceOptions[0]} onChange={v=>changeBriefAnswer(v, subIndex)}/>
              </div>
              : null}

            <Divider/>
          </div>
        ))}
      </div>
      <div style={{textAlign: 'right', marginTop: '10px'}}>
        {obj.item.submitted ?
          <Button disabled={!obj.editable} type="primary" onClick={updateSurvey} loading={submitLoading}>更新</Button>
          :
          <Button disabled={!obj.editable} type="primary" onClick={submitSurvey} loading={submitLoading}>提交</Button>
        }
      </div>
      {obj.showCount || obj.item.submitted ?
        <div style={{textAlign: 'left', marginTop: '20px'}}>
          <Collapse onChange={showCount}>
            <Collapse.Panel key={1} header="查看统计结果">
              {obj.item.questions.map((subItem, subIndex)=>(
                <div key={subIndex.toString()}>
                  {subItem.questionType==='singleChoice' || subItem.questionType==='multipleChoice'
                  || subItem.questionType==='scale5' || subItem.questionType==='scale7' ?
                    <div style={{textAlign: "left", marginTop: '10px'}}>
                      <p>{subItem.questionTitle}</p>
                      <ReactEcharts option={getOption(subIndex)}/>
                      <Divider />
                    </div>
                    :
                    null
                  }
                </div>
              ))}
            </Collapse.Panel>
          </Collapse>
        </div>
        : null
      }
    </div>
  )
}

export default FillSurvey