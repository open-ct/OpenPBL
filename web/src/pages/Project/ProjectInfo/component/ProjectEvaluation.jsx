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

import React, {useEffect, useState} from 'react';
import 'echarts/lib/chart/pie';
import 'echarts/lib/component/tooltip';
import 'echarts/lib/component/title';
import 'echarts/lib/component/legend';
import 'echarts/lib/component/markPoint';
import ReactEcharts from 'echarts-for-react';
import {Button, Col, Divider, InputNumber, List, Menu, message, Row, Table} from "antd";

import TaskApi from "../../../../api/TaskApi";
import ProjectApi from "../../../../api/ProjectApi";
import ChapterApi from "../../../../api/ChapterApi";
import SectionApi from "../../../../api/SectionApi";
import util from "../../../component/Util"
import "./index.less"

const {SubMenu} = Menu

function ProjectEvaluation(obj) {
  const pid = obj.project.id
  const published = obj.project.published
  const type = localStorage.getItem("type")

  const [chapters, setChapters] = useState([])
  const [loadChapters, setLoadChapters] = useState(false)

  const [defaultOpenedKeys, setDefaultOpenedKeys] = useState([])
  const [openedKeys, setOpenedKeys] = useState([])
  const [data, setData] = useState([])
  const [tasks, setTasks] = useState([])
  const [loadTasks, setLoadTasks] = useState(false)

  const [learnMinuteWeight, setLearnMinuteWeight] = useState(obj.project.learnMinuteWeight)

  const [editWeight, setEditWeight] = useState(false)
  const [editMinute, setEditMinute] = useState(false)

  useEffect(() => {
    getChapters()
    getTasks()
  }, []);
  const getTasks = () => {
    setLoadTasks(true)
    TaskApi.getProjectTasks(pid)
      .then(res => {
        setLoadTasks(false)
        if (res.data.code === 200) {
          if (res.data.data != null) {
            setTasks(res.data.data)
            setEchartsData(res.data.data)
          } else {
            setEchartsData([])
          }
        }
      })
      .catch(e => {
        console.log(e)
      })
  }
  const getChapters = () => {
    setLoadChapters(true)
    ChapterApi.getProjectChapters(pid)
      .then((res) => {
        setLoadChapters(false)
        if (res.data.chapters === null) {
          setChapters([])
          setDefaultOpenedKeys(0)
        } else {
          setChapters(res.data.chapters)
          setOpenAllKeys(res.data.chapters.length)
        }
      })
      .catch(e => {
        console.log(e)
      })
  }
  const setOpenAllKeys = n => {
    let keys = []
    for (let i = 0; i < n; i++) {
      keys.push(i.toString())
    }
    setDefaultOpenedKeys(keys)
  }
  const checkWeight = () => {
    let total = learnMinuteWeight
    for (let i = 0; i < tasks.length; i++) {
      total = total + tasks[i].taskWeight
    }
    if (total === 100) {
      return true
    }
    if (total > 100) {
      message.error('权重之和不能大于100')
    } else if (total < 100) {
      message.error('权重之和不能小于100')
    }
    return false
  }
  const edit1 = () => {
    setEditMinute(true)
  }
  const edit2 = () => {
    setEditWeight(true)
  }
  const saveMinute = () => {
    let sections = []
    for (let i = 0; i < chapters.length; i++) {
      sections = sections.concat(chapters[i].sections)
    }
    let data = JSON.stringify(sections)
    SectionApi.updateSectionsMinute(pid, {sections: data})
      .then(res => {
        setEditMinute(false)
        if (res.data.code === 200) {
          getChapters()
        }
      })
      .catch(e => {
        console.log(e)
      })
  }
  const saveWeight = () => {
    if (checkWeight()) {
      let data = {
        learnMinuteWeight: learnMinuteWeight,
        tasks: JSON.stringify(tasks)
      }
      ProjectApi.updateWeight(pid, data)
        .then(res => {
          setEditWeight(false)
          if (res.data.code === 200) {
            getTasks()
            obj.loadProjectDetail()
          }
        })
        .catch(e => {
          console.log(e)
        })
    }
  }
  const cancel1 = () => {
    setEditMinute(false)
  }
  const cancel2 = () => {
    setEditWeight(false)
  }
  const changeLearnMinute = (v, index, subIndex) => {
    console.log(index, subIndex)

    chapters[index].sections[subIndex].sectionMinute = v
    setChapters([...chapters])
  }
  const changeLearnMinuteWeight = (v) => {
    setLearnMinuteWeight(v)
  }
  const changeTaskWeight = (v, index) => {
    tasks[index].taskWeight = v
    setTasks([...tasks])
  }

  const setEchartsData = (tasks) => {
    let d = [
      {value: learnMinuteWeight, name: '学习时长'}
    ]
    for (let i = 0; i < tasks.length; i++) {
      d.push({
        value: tasks[i].taskWeight,
        name: tasks[i].taskTitle,
      })
    }
    setData(d)
  }

  const getColumns = () => {
    return [
      {
        title: '任务标题',
        dataIndex: 'taskTitle',
        key: 'taskTitle',
        render: (text, item, index) => (
          <>
            {util.FormatSectionName(item.taskTitle, item.chapterNumber, item.sectionNumber)}
          </>
        )
      },
      {
        title: '任务描述',
        dataIndex: 'taskIntroduce',
        key: 'taskIntroduce',
        ellipsis: true,
      },
      {
        title: '权重',
        dataIndex: 'taskWeight',
        key: 'taskWeight',
        render: (text, item, index) => (
          <>
            {editWeight ?
              <>
                <InputNumber onChange={v => changeTaskWeight(v, index)} value={item.taskWeight} min={0}
                             max={100}/>&nbsp;&nbsp;%
              </>
              :
              <span>{text}&nbsp;&nbsp;%</span>
            }
          </>
        )
      }]
  }

  const getOptions = () => ({
    title: {
      x: 'center',
    },
    tooltip: {
      trigger: 'item',
      formatter: '{a} <br/>{b} : {c} ({d}%)',
    },
    legend: {
      orient: 'vertical',
      top: 20,
      left: 5,
    },
    series: [
      {
        name: '评价方案',
        type: 'pie',
        data,
      },
    ],
  });
  return (
    <div style={{textAlign: 'left', marginBottom: '30px'}} key="1">
      <ReactEcharts option={getOptions()} showLoading={loadTasks}/>
      <div>
        <Divider orientation="left">
          <p className="evidence-title">章节学习时长</p>
        </Divider>
        {!published ?
          <div style={{float: 'right'}}>
            {editMinute ?
              <>
                <Button onClick={cancel1} style={{marginRight: '20px'}}>取消</Button>
                <Button onClick={saveMinute} type="primary">保存</Button>
              </>
              :
              <Button onClick={edit1}>编辑时长</Button>
            }
          </div>
          : null
        }
      </div>


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
                locale={{emptyText: '空的'}}
                loading={loadChapters}
                renderItem={
                  (item, subIndex) => (
                    <List.Item>
                      {util.FormatSectionName(item.sectionName, item.chapterNumber, item.sectionNumber)}
                      {editMinute ?
                        <div style={{float: 'right'}}>
                          学习时长不少于&nbsp;&nbsp;
                          <InputNumber value={item.sectionMinute} onChange={v => changeLearnMinute(v, index, subIndex)}
                                       min={0}/>
                          &nbsp;&nbsp;分钟
                        </div>
                        :
                        <div style={{float: 'right'}}>
                          学习时长不少于&nbsp;&nbsp;{item.sectionMinute}&nbsp;&nbsp;分钟
                        </div>
                      }

                    </List.Item>
                  )
                }
              /><br/>
            </>
          }
        </div>
      ))}
      <div>
        <Divider orientation="left">
          <p className="evidence-title">权重占比</p>
        </Divider>
        {!published ?
          <div style={{float: 'right'}}>
            {editWeight ?
              <>
                <Button onClick={cancel2} style={{marginRight: '20px'}}>取消</Button>
                <Button onClick={saveWeight} type="primary">保存</Button>
              </>
              :
              <Button onClick={edit2}>编辑权重</Button>
            }
          </div>
          : null
        }
      </div>
      <Divider/>
      <Row style={{padding: '15px'}} gutter={[10, 10]}>
        <Col span={8}>
          <span style={{fontWeight: 'bold'}}>学习时长</span></Col>
        <Col span={8}>
        </Col>
        <Col span={8}>
          <span>权重&nbsp;&nbsp;</span>
          {editWeight ?
            <span><InputNumber value={learnMinuteWeight} onChange={changeLearnMinuteWeight} min={0}/></span>
            :
            <span>{obj.project.learnMinuteWeight}</span>
          }
          &nbsp;&nbsp;%
        </Col>
      </Row>
      <Divider/>
      <Table
        loading={loadTasks}
        columns={getColumns()}
        dataSource={tasks}
        pagination={false}
      />
    </div>
  );
}

export default ProjectEvaluation;
