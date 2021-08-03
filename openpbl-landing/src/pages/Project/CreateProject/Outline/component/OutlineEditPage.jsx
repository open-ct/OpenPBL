import React, {useEffect, useState} from "react";
import {Button, Col, Input, Menu, message, Modal, Popconfirm, Row} from 'antd'
import {ArrowDownOutlined, ArrowUpOutlined, DeleteOutlined, EditOutlined} from '@ant-design/icons'
import {Link} from 'react-router-dom'

import ChapterApi from "../../../../../api/ChapterApi"
import SectionApi from "../../../../../api/SectionApi"

const {SubMenu} = Menu;

function OutlineEditPage(obj) {
  const pid = obj.pid

  const [chapters, setChapters] = useState([])
  const [chapterModalVisible, setChapterModalVisible] = useState(false)
  const [sectionModalVisible, setSectionModalVisible] = useState(false)

  const [opt, setOpt] = useState('add')
  const [chapter, setChapter] = useState({})
  const [section, setSection] = useState({})

  const [index, setIndex] = useState('')
  const [subIndex, setSubIndex] = useState('')

  const [chapterName, setChapterName] = useState('')
  const [sectionName, setSectionName] = useState('')

  useEffect(() => {
    getChapters()
  }, [])
  const getChapters = () => {
    ChapterApi.getProjectChapters(pid)
      .then((res) => {
        if (res.data.chapters === null) {
          setChapters([])
        } else {
          setChapters(res.data.chapters)
        }
      })
      .catch((e) => {
        console.log(e)
      })
  }

  const addChapter = e => {
    setOpt('add')
    setChapterModalVisible(true)
  }
  const modifyChapter = (e, c, index) => {
    e.stopPropagation()

    setChapter(c)
    setChapterName(c.chapterName)
    setIndex(index)
    setOpt('modify')
    setChapterModalVisible(true)
  }
  const addSection = (c, index) => {
    setOpt('add')
    setChapter(c)
    setIndex(index)
    setSectionModalVisible(true)
  }
  const modifySection = (s, index, subIndex) => {
    setSection(s)
    setSectionName(s.sectionName)
    setIndex(index)
    setSubIndex(subIndex)
    setOpt('modify')
    setSectionModalVisible(true)
  }
  const del = e => {
    e.stopPropagation()
  }
  const cancel = e => {
    e.stopPropagation()
  }
  const deleteChapter = (e, c, index) => {
    e.stopPropagation()
    ChapterApi.deleteProjectChapter(c)
      .then((res) => {
        if (res.data.code === 200) {
          chapters.splice(index, 1)
          setChapters([...chapters])
          message.success(res.data.msg)
        }
      })
      .catch((e) => {
        console.log(e)
      })
  }
  const deleteSection = (s, index, subIndex) => {
    SectionApi.deleteChapterSection(pid, s)
      .then((res) => {
        if (res.data.code === 200) {
          let sections = chapters[index].sections
          sections.splice(subIndex, 1)
          chapters[index].sections = sections
          setChapters([...chapters])
          message.success(res.data.msg)
        }
      })
  }
  const doChapter = e => {
    if (chapterName === '') {
      message.error('请输入章名')
      return
    }
    if (opt === 'modify') {
      let c = {
        id: chapter.id,
        projectId: chapter.projectId,
        chapterName: chapterName,
        chapterNumber: chapter.chapterNumber,
      }
      ChapterApi.updateProjectChapter(c)
        .then((res) => {
          if (res.data.code === 200) {
            setChapterModalVisible(false)
            message.success(res.data.msg)

            chapters[index].chapterName = chapterName
            setChapters([...chapters])
            setChapterName('')
          }
        })
        .catch((e) => {
          console.log(e)
        })
    } else if (opt === 'add') {
      let len = chapters.length
      let l = 0
      if (len > 0) {
        l = chapters[len - 1].chapterNumber + 1
      }
      let cp = {chapterName: chapterName, chapterNumber: l, projectId: pid}
      ChapterApi.createProjectChapter(cp)
        .then((res) => {
          setChapterModalVisible(false)
          setChapterName('')
          if (res.data.code === 200) {
            cp.id = res.data.data
            chapters.push(cp)
            setChapters([...chapters])
          }
        })
        .catch((e) => {
          console.log(e)
        })
    }
  }
  const doSection = () => {
    if (sectionName === '') {
      message.error('请输入小节名')
      return
    }
    if (opt === 'modify') {
      let s = {
        id: section.id,
        chapterId: section.chapterId,
        sectionName: sectionName,
        sectionNumber: section.sectionNumber,
        chapterNumber: section.chapterNumber
      }
      SectionApi.updateChapterSection(pid, s)
        .then((res) => {
          if (res.data.code === 200) {
            setSectionModalVisible(false)

            chapters[index].sections[subIndex].sectionName = sectionName
            setChapters([...chapters])
            setSectionName('')
          }
        })
        .catch((e) => {
          console.log(e)
        })
    } else if (opt === 'add') {
      let l = 0
      if (chapters[index].sections !== undefined && chapters[index].sections !== null) {
        let len = chapters[index].sections.length
        if (len > 0) {
          l = chapters[index].sections[len - 1].sectionNumber + 1
        }
      }
      let sec = {
        sectionName: sectionName,
        sectionNumber: l,
        chapterId: chapters[index].id,
        chapterNumber: chapters[index].chapterNumber
      }
      SectionApi.createChapterSection(pid, sec)
        .then((res) => {
          setSectionModalVisible(false)
          setSectionName('')
          if (res.data.code === 200) {
            sec.id = res.data.data;
            if (chapters[index].sections === undefined || chapters[index].sections === null) {
              chapters[index].sections = [sec]
            } else {
              chapters[index].sections.push(sec)
            }
            setChapters([...chapters])
            setChapterName('')
          }
        })
        .catch((e) => {
          console.log(e)
        })
    }
  }
  const cancelDoChapter = e => {
    setChapterModalVisible(false)
    setChapterName('')
  }
  const cancelDoSection = e => {
    setSectionModalVisible(false)
    setSectionName('')
  }

  const changeChapterName = value => {
    setChapterName(value.target.value)
  }
  const changeSectionName = value => {
    setSectionName(value.target.value)
  }
  const exchangeChapter = (e, index, index2) => {
    e.stopPropagation()

    if (index < 0 || index2 >= chapters.length) {
    } else {
      let id1 = chapters[index].id
      let id2 = chapters[index2].id
      ChapterApi.exchangeProjectChapter(pid, id1, id2)
        .then((res) => {
          if (res.data.code === 200) {
            let c1 = chapters[index]
            chapters[index] = chapters[index2]
            chapters[index2] = c1
            setChapters([...chapters])
          }
        })
        .catch(e => {
          console.log(e)
        })
    }
  }
  const exchangeSection = (index, subIndex, subIndex2) => {
    if (subIndex < 0 || subIndex2 >= chapters[index].sections.length) {
    } else {
      let id1 = chapters[index].sections[subIndex].id
      let id2 = chapters[index].sections[subIndex2].id
      SectionApi.exchangeChapterSection(chapters[index], id1, id2)
        .then(res => {
          if (res.data.code === 200) {
            let s1 = chapters[index].sections[subIndex]
            chapters[index].sections[subIndex] = chapters[index].sections[subIndex2]
            chapters[index].sections[subIndex2] = s1
            setChapters([...chapters])
          }
        })
        .catch(e => {
          console.log(e)
        })
    }
  }

  return (
    <div>
      <Menu
        style={{width: '100%'}}
        defaultSelectedKeys={[]}
        defaultOpenKeys={[]}
        mode="inline"
      >{chapters.map((item, index) => (
        <SubMenu
          style={{fontSize: '2.7vh'}}
          key={index.toString()}
          title={
            <div>
              {item.chapterName}
              <span style={{float: 'right', marginRight: '20px'}}>
                <Button shape="circle" type="text" onClick={e => modifyChapter(e, item, index)} icon={<EditOutlined/>}/>
                <Button shape="circle" type="text" icon={<ArrowUpOutlined/>} onClick={e => exchangeChapter(e, index - 1, index)}/>
                <Button shape="circle" type="text" icon={<ArrowDownOutlined/>} onClick={e => exchangeChapter(e, index, index + 1)}/>
                <Popconfirm title="确定删除章节？" onConfirm={e => deleteChapter(e, item, index)} onCancel={cancel} placement="topLeft">
                  <Button onClick={del} shape="circle" type="text" icon={<DeleteOutlined/>} style={{color: 'red', marginLeft: '20px'}}/>
                </Popconfirm>
              </span>
            </div>
          }
        >
          {(item.sections === null || item.sections === undefined) ? null :
            item.sections.map((subItem, subIndex) => (
              <Menu.Item key={index.toString() + subIndex.toString()}>
                {subItem.sectionName}
                <span style={{float: 'right', marginRight: '20px'}}>
                  <Link to={`/project/${pid}/section/${subItem.id}/edit`}>
                    <Button type="text">编辑资源</Button>
                  </Link>

                  <Button shape="circle" type="text" onClick={e => modifySection(subItem, index, subIndex)} icon={<EditOutlined/>}/>
                  <Button shape="circle" type="text" icon={<ArrowUpOutlined/>} onClick={e => exchangeSection(index, subIndex - 1, subIndex)}/>
                  <Button shape="circle" type="text" icon={<ArrowDownOutlined/>} onClick={e => exchangeSection(index, subIndex, subIndex + 1)}/>

                  <Popconfirm title="确定删除小节？" onConfirm={e => deleteSection(subItem, index, subIndex)} placement="topLeft">
                    <Button shape="circle" type="text" icon={<DeleteOutlined/>} style={{color: 'red', marginLeft: '20px'}}/>
                  </Popconfirm>

                </span>
              </Menu.Item>
            ))
          }
          <Button type="round" style={{float: 'right', margin: '3px'}}
                  onClick={e => addSection(item, index)}>添加小节</Button>
        </SubMenu>
      ))}
      </Menu>
      <div style={{textAlign: 'right'}}>
        <Button type="round" style={{margin: '5px'}} onClick={addChapter}>添加章节</Button>
      </div>
      <div style={{textAlign: 'right', marginTop: '20px', marginRight: '20px'}}>
        <Link to={`/project/${pid}/info`}>
          <Button type="primary" size="middle" style={{float: 'right'}}>完成</Button>
        </Link>
        <Link to={`/project/${pid}/info/edit`}>
          <Button size="middle" style={{float: 'right', marginRight: '20px'}}>上一页</Button>
        </Link>
      </div>

      <Modal title="" visible={chapterModalVisible} onOk={doChapter} onCancel={cancelDoChapter}>
        <br/>
        <Row>
          <Col span={3}>
            <p>章名：</p>
          </Col>
          <Col span={20}>
            <Input value={chapterName} onChange={changeChapterName} onPressEnter={doChapter}/>
          </Col>
        </Row>
      </Modal>

      <Modal visible={sectionModalVisible} onOk={doSection} onCancel={cancelDoSection}>
        <br/>
        <Row>
          <Col span={3}>
            <p>小节名：</p>
          </Col>
          <Col span={20}>
            <Input value={sectionName} onChange={changeSectionName} onPressEnter={doSection}/>
          </Col>
        </Row>
      </Modal>
    </div>
  )
}

export default OutlineEditPage
