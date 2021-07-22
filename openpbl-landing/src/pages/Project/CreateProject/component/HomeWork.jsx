import React, {useEffect, useState} from "react";
import {Button, Card, Divider, Input, message} from "antd";

import "../section-edit.less"
import ResourceApi from "../../../../api/ResourceApi";

function HomeWork(obj) {

  const [homeWorkTitle, setHomeWorkTitle] = useState('')
  const [homeWorkIntroduce, setHomeWorkIntroduce] = useState('')

  useEffect(()=>{
    setHomeWorkTitle(obj.section.resource.homeWorkTitle)
    setHomeWorkIntroduce(obj.section.resource.homeWorkIntroduce)
  }, [obj.section.resource.homeWorkIntroduce, obj.section.resource.homeWorkTitle])

  const saveContent = e => {
    let q = {
      id: obj.section.resource.id,
      hasHomeWork: true,
      homeWorkTitle: homeWorkTitle,
      homeWorkIntroduce: homeWorkIntroduce,
    }
    ResourceApi.updateResource(q)
      .then(res=>{
        if (res.data.code === 200){
          message.success(res.data.msg)
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e=>{console.log(e)})
  }

  const changeHomeWorkTitle = value => {
    setHomeWorkTitle(value.target.value)
  }
  const changeHomeWorkIntroduce = value => {
    setHomeWorkIntroduce(value.target.value)
  }

  return (
    <Card className="resource-card">
      <p className="card-title">学生任务</p>
      <Divider />
      <Input placeholder="任务标题" value={homeWorkTitle} onChange={changeHomeWorkTitle}/>
      <Input.TextArea placeholder="任务描述" value={homeWorkIntroduce} onChange={changeHomeWorkIntroduce} style={{marginTop: '20px'}} />
      <div style={{marginTop: '10px'}}>
        <span style={{float: 'right'}}>
          <Button type="primary" onClick={saveContent}>保存</Button>
        </span>
      </div>
    </Card>
  )
}

export default HomeWork
