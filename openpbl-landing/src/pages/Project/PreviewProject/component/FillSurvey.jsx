import React from "react";
import {Button, Checkbox, Divider, Input, Radio, message} from "antd";
import qs from 'qs'

import Question from "../../CreateProject/Survey/component/Question"
import "../preview.less"
import "../../CreateProject/Section/component/section-edit.less"
import SubmitApi from "../../../../api/SubmitApi";

const blank = Question.blank

function FillSurvey(obj) {
  const submitSurvey = e => {
    obj.item.submit.submitType = 'survey'
    let data = Object.assign({}, obj.item.submit)
    let c = [...obj.item.choices]
    for (let i=0; i<c.length; i++) {
      c[i].choiceOptions = c[i].choiceOptions.toString()
    }
    data['choices'] = JSON.stringify(c)
    SubmitApi.createSubmit(obj.pid, obj.item.id, data)
      .then(res=>{
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
    obj.item.submit.submitType = 'survey'
    let data = Object.assign({}, obj.item.submit)
    let c = [...obj.item.choices]
    for (let i=0; i<c.length; i++) {
      c[i].choiceOptions = c[i].choiceOptions.toString()
    }
    data['choices'] = JSON.stringify(c)
    SubmitApi.updateSubmit(obj.pid, obj.item.id, obj.item.submit.id, data)
      .then(res=>{
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
  }
  const changeBriefAnswer = (v, subIndex) => {
    obj.item.choices[subIndex].choiceOptions[0] = v.target.value
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
                        <Input value={obj.item.choices[subIndex].choiceOptions[optIndex]} onChange={v=>changeBlankFill(v, subIndex, optIndex)} style={{borderBottom: '2px solid black'}} bordered={false}/>
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
          <Button type="primary" onClick={updateSurvey}>更新</Button>
          :
          <Button disabled={!obj.learning} type="primary" onClick={submitSurvey}>提交</Button>
        }
      </div>
    </div>
  )
}

export default FillSurvey