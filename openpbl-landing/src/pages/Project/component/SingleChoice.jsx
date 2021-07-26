import React, {useState} from "react";
import {Button, Input, Popconfirm, Radio, message} from "antd";
import {ArrowDownOutlined, ArrowUpOutlined, DeleteOutlined, MinusOutlined, PlusOutlined, EditOutlined, SaveOutlined} from "@ant-design/icons";

function SingleChoice(obj) {
  const [opt, setOpt] = useState(obj.item.questionOptions)
  const [title, setTitle] = useState(obj.item.questionTitle)

  const changeTitle = value => {
    setTitle(value.target.value)
  }
  const changeOpt = (value, index) => {
    opt[index] = value.target.value
    setOpt([...opt])
  }
  const addOpt = index => {
    opt.splice(index+1, 0, '')
    setOpt([...opt])
  }
  const delOpt = index => {
    if (opt.length === 1) {
      message.error("不能删除最后一个选项")
    } else {
      opt.splice(index, 1)
      setOpt([...opt])
    }
  }
  return (
    <div>
      <div>
        {obj.editing ?
          <p style={{float: 'left'}}>
            <Input value={title} onChange={changeTitle} />
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
                    icon={<SaveOutlined />}/>
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
          <Radio.Group>
            {opt.map((item, index) => (
              <div style={{marginTop: '10px'}} key={index.toString()}>
                <Radio value={index}>
                  <Input value={item} onChange={value=>changeOpt(value, index)}/>
                </Radio>
                <Button type="circle" size="small" icon={<MinusOutlined />} onClick={e=>delOpt(index)}/>
                <Button type="circle" size="small" icon={<PlusOutlined />} onClick={e=>addOpt(index)} style={{marginLeft: '10px'}}/>
              </div>
            ))}
          </Radio.Group>
        </div>
        :
        <div style={{textAlign: "left", marginTop: '10px'}}>
          <Radio.Group>
            {obj.item.questionOptions.map((subItem, subIndex) => (
              <div style={{marginTop: '10px'}} key={subIndex.toString()}>
                <Radio value={subIndex}>
                  {subItem}
                </Radio>
              </div>
            ))}
          </Radio.Group>
        </div>
      }
    </div>
  )

}

export default SingleChoice