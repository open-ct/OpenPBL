import React from "react";
import {Button, Checkbox, Divider, Input, Radio} from "antd";

import Question from "../../CreateProject/Survey/component/Question"
import "../preview.less"
import "../../CreateProject/Section/component/section-edit.less"

const blank = Question.blank

function FillSurvey(obj) {

  const submitSurvey = e => {

  }
  return (
    <div className="survey">
      <h2 style={{textAlign: 'center'}}>{obj.item.survey.surveyTitle}</h2>
      <Divider />
      <div style={{marginLeft: '20px', marginRight: '20px'}}>
        {obj.item.questions.map((subItem, subIndex)=>(
          <div key={subIndex.toString()}>
            {subItem.questionType==='singleChoice' ?
              <div style={{textAlign: "left", marginTop: '10px'}}>
                <p>{subItem.questionTitle}</p>
                <Radio.Group>
                  {subItem.questionOptions.map((optItem, optIndex) => (
                    <div style={{marginTop: '10px'}} key={optIndex.toString()}>
                      <Radio value={optIndex}>
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
                  {subItem.questionOptions.map((optItem, optIndex) => (
                    <div style={{marginTop: '10px'}} key={optIndex.toString()}>
                      <Checkbox>
                        {optItem}
                      </Checkbox>
                    </div>
                  ))}
              </div>
              : null}
            {subItem.questionType==='blankFill' ?
              <div style={{textAlign: "left", marginTop: '10px'}}>
                <p>{subItem.questionTitle}</p>
                {subItem.questionOptions.map((optItem, optIndex) => (
                  <div style={{marginTop: '10px'}} key={optIndex.toString()}>
                    {optItem === blank ?
                      <span style={{float: 'left', margin: '5px'}}>
                        <Input style={{borderBottom: '2px solid black'}} bordered={false}/>
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
                <Input.TextArea/>
              </div>
              : null}

            <Divider/>
          </div>
        ))}
      </div>
      <Button type="primary" onClick={{submitSurvey}} style={{float: 'right', marginTop: '10px'}}>提交</Button>
    </div>
  )
}

export default FillSurvey