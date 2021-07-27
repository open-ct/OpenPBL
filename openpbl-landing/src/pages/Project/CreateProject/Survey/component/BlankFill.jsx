import React, {useState} from "react";
import {Button, Dropdown, Input, Menu, message, Popconfirm} from "antd";
import {
  ArrowDownOutlined,
  ArrowUpOutlined,
  DeleteOutlined,
  EditOutlined,
  MinusOutlined,
  PlusOutlined,
  SaveOutlined
} from "@ant-design/icons";
import Question from "./Question";


const blank = Question.blank

function BlankFill(obj) {
  const [opt, setOpt] = useState(obj.item.questionOptions)
  const [title, setTitle] = useState(obj.item.questionTitle)

  const changeTitle = value => {
    setTitle(value.target.value)
  }
  const changeOpt = (value, index) => {
    opt[index] = value.target.value
    setOpt([...opt])
  }
  const addOpt = (key, index) => {
    let v = '描述'
    if (key.key === '2') {
      v = '&nbsp;'
    }
    opt.splice(index + 1, 0, v)
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
          {opt.map((item, index) => (
            <div style={{marginTop: '10px'}} key={index.toString()}>
              {item === blank ?
                <span style={{float: 'left', margin: '5px'}}>
                  <Input style={{borderBottom: '2px solid black'}} bordered={false}/>
                </span>
                :
                <span style={{float: 'left', margin: '5px'}}>
                  <Input value={item} onChange={value => changeOpt(value, index)}/>
                </span>
              }
              <span style={{float: 'left', margin: '5px'}}>
                <Button type="circle" size="small" icon={<MinusOutlined/>} onClick={e => delOpt(index)}/>
                <Dropdown
                  overlay={
                    <Menu onClick={key => addOpt(key, index)}>
                      <Menu.Item key="1">文字</Menu.Item>
                      <Menu.Item key="2">填空</Menu.Item>
                    </Menu>
                  }
                  trigger={['click']}>
                  <Button type="circle" size="small" className="ant-dropdown-link" icon={<PlusOutlined/>}
                          onClick={e => e.preventDefault()} style={{marginLeft: '4px'}}/>
                </Dropdown>
              </span>
            </div>
          ))}
        </div>
        :
        <div style={{textAlign: "left", marginTop: '10px'}}>
          {obj.item.questionOptions.map((subItem, subIndex) => (
            <div style={{marginTop: '10px'}} key={subIndex.toString()}>
              {subItem === blank ?
                <span style={{float: 'left', margin: '5px'}}>
                  <Input style={{borderBottom: '2px solid black'}} bordered={false}/>
                </span>
                :
                <span style={{float: 'left', margin: '5px'}}>{subItem}</span>}
            </div>
          ))}
        </div>
      }
    </div>
  )
}

export default BlankFill