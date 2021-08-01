import React, {useEffect, useState} from 'react';
import 'echarts/lib/chart/pie';
import 'echarts/lib/component/tooltip';
import 'echarts/lib/component/title';
import 'echarts/lib/component/legend';
import 'echarts/lib/component/markPoint';
import ReactEcharts from 'echarts-for-react';
import QueueAnim from 'rc-queue-anim';
import {Button, Card, Col, Divider, InputNumber, Row, Table, message} from "antd";

import TaskApi from "../../../../api/TaskApi";
import ProjectApi from "../../../../api/ProjectApi";

function ProjectEvaluation(obj) {
  const pid = obj.project.id
  const published = obj.project.published

  const [data, setData] = useState([])
  const [tasks, setTasks] = useState([])

  const [learnMinute, setLearnMinute] = useState(obj.project.learnMinute)
  const [learnMinuteWeight, setLearnMinuteWeight] = useState(obj.project.learnMinuteWeight)

  const [editWeight, setEditWeight] = useState(false)

  useEffect(() => {
    getTasks()
  }, []);
  const getTasks = () => {
    TaskApi.getProjectTasks(pid)
      .then(res=>{
        if (res.data.code === 200) {
          setTasks(res.data.data)
          setEchartsData(res.data.data)
        }
      })
      .catch(e=>{console.log(e)})
  }

  const checkWeight = () => {
    let total = learnMinuteWeight
    for (let i=0; i<tasks.length; i++) {
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
  const edit = () => {
    setEditWeight(true)
  }
  const save = () => {
    if (checkWeight()) {
      setEditWeight(false)
      let data = {
        learnMinute: learnMinute,
        learnMinuteWeight: learnMinuteWeight,
        tasks: JSON.stringify(tasks)
      }
      ProjectApi.updateWeight(pid, data)
        .then(res=>{
          if (res.data.code === 200) {
            getTasks()
            obj.loadProjectDetail()
          }
        })
        .catch(e=>{console.log(e)})

    }
  }
  const cancel = () => {
    setEditWeight(false)
  }
  const changeLearnMinute = (v) => {
    setLearnMinute(v)
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
    for (let i=0; i<tasks.length; i++) {
      d.push({
        value: tasks[i].taskWeight,
        name: tasks[i].taskTitle,
      })
    }
    setData(d)
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
    <QueueAnim>
      <div style={{textAlign: 'left', marginBottom: '30px'}} key="1">
        <ReactEcharts option={getOptions()}/>

        {!published ?
          <div style={{textAlign: 'right'}}>
            {editWeight ?
              <>
                <Button onClick={cancel} style={{marginRight: '20px'}}>取消</Button>
                <Button onClick={save} type="primary">保存</Button>
              </>
              :
              <Button onClick={edit}>编辑权重</Button>
            }
          </div>
          : null
        }
        <p style={{textAlign: 'center', fontWeight: 'bold', fontSize: '20px', marginTop: '20px'}}>权重占比</p>
        <Divider />
        <Row style={{padding: '15px'}} gutter={[10, 10]}>
          <Col span={8}>
            <span style={{fontWeight: 'bold'}}>章节学习时长</span></Col>
          <Col span={10}>
            <span>每小节学习时长不少于&nbsp;&nbsp;</span>
            {editWeight ?
              <span><InputNumber value={learnMinute} onChange={changeLearnMinute} min={0}/></span>
              :
              <span>{obj.project.learnMinute}</span>
            }
            <span>&nbsp;&nbsp;分钟</span>
          </Col>
          <Col span={6}>
            <span>权重&nbsp;&nbsp;</span>
            {editWeight ?
              <span><InputNumber value={learnMinuteWeight} onChange={changeLearnMinuteWeight} min={0}/></span>
              :
              <span>{obj.project.learnMinuteWeight}</span>
            }
            &nbsp;&nbsp;%
          </Col>
        </Row>
        <Divider />
        <Table
          columns={[
            {
              title: '任务标题',
              dataIndex: 'taskTitle',
              key: 'taskTitle'
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
                      <InputNumber onChange={v=>changeTaskWeight(v, index)} value={item.taskWeight} min={0} max={100}/>&nbsp;&nbsp;%
                    </>
                    :
                    <span>{text}&nbsp;&nbsp;%</span>
                  }
                </>
              )
            }
          ]}
          dataSource={tasks}
          pagination={false}
        />
      </div>
    </QueueAnim>
  );
}

export default ProjectEvaluation;
