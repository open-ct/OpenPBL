import React, {useState} from "react";
import {Button, Input, Popconfirm} from "antd";
import {ArrowDownOutlined, ArrowUpOutlined, DeleteOutlined, EditOutlined, SaveOutlined} from "@ant-design/icons";


function BriefAnswer(obj) {
  const [opt, setOpt] = useState(obj.item.questionOptions)
  const [title, setTitle] = useState(obj.item.questionTitle)

  const changeTitle = value => {
    setTitle(value.target.value)
  }
  const changeQ = v => {
    opt[0] = v.target.value
    setOpt([...opt])
  }
  return (
    <div>
      <div>
        {obj.editing ?
          <p style={{float: 'left'}}>
            <Input value={title} onChange={changeTitle}/>
          </p>
          :
          <p style={{float: 'left'}}>{obj.item.questionTitle}
            <span style={{color: 'gray'}}>&nbsp;
              [{obj.getType(obj.item.questionType)}]
          </span>
          </p>
        }
        <p style={{float: 'right'}}>
          {obj.editing ?
            <Button shape="circle" type="text" onClick={e => obj.saveQuestion(obj.item, title, opt, obj.index)}
                    icon={<SaveOutlined/>}/>
            :
            <>
              <Button shape="circle" type="text" onClick={e => obj.editQuestion(obj.item, obj.index)}
                      icon={<EditOutlined/>}/>
              <Button shape="circle" type="text" icon={<ArrowUpOutlined/>}
                      onClick={e => obj.exchangeQuestion(obj.index - 1, obj.index)}/>
              <Button shape="circle" type="text" icon={<ArrowDownOutlined/>}
                      onClick={e => obj.exchangeQuestion(obj.index, obj.index + 1)}/>
            </>
          }
          <Popconfirm title="确定删除题目？"
                      onConfirm={e => obj.deleteQuestion(obj.item, obj.index)}
                      placement="topLeft">
            &nbsp;&nbsp;<Button shape="circle" type="text"
                                style={{color: 'red'}}
                                icon={<DeleteOutlined/>}/>
          </Popconfirm>
        </p>
      </div>
      <br/><br/>

      {obj.editing ?
        <div style={{textAlign: "left", marginTop: '10px'}}>
          <Input value={opt[0]} onChange={changeQ}/>
          <Input.TextArea style={{marginTop: '5px'}}/>
        </div>
        :
        <div style={{textAlign: "left", marginTop: '10px'}}>
          <p>{obj.item.questionOptions[0]}</p>
          <Input.TextArea/>
        </div>
      }

    </div>
  )
}

export default BriefAnswer