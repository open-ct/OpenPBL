import React, {useEffect, useState} from "react";
import {Button, Input, Menu, Modal, Row, Col, message} from 'antd'
import {EditOutlined} from '@ant-design/icons'
import {Link} from 'react-router-dom'

import ProjectApi from "../../../../api/ProjectApi";
import SectionEditPage from "./SectionEditPage";

const {SubMenu} = Menu;

function OutlineEditPage(obj) {
  const pid = obj.pid

  const [chapters, setChapters] = useState([])
  const [chapterModalVisible, setChapterModalVisible] = useState(false)
  const [sectionModalVisible, setSectionModalVisible] = useState(false)

  const [exist, setExist] = useState(false)
  const [chapter, setChapter] = useState({})
  const [section, setSection] = useState({})

  const [index, setIndex] = useState('')
  const [subIndex, setSubIndex] = useState('')

  const [chapterName, setChapterName] = useState('')
  const [sectionName, setSectionName] = useState('')

  useEffect(() => {
    ProjectApi.getProjectChapters(pid)
      .then((res) => {
        if (res.data.chapters === null) {
          setChapters([])
        } else {
          setChapters(res.data.chapters)
        }
      })
      .catch((e)=>{console.log(e)})
  }, [])

  const handleClick = (item, key) => {
    if (item.sections === undefined) {
      ProjectApi.getChapterSections(item.id)
        .then((res) => {
          if (res.data.sections == null) {
            item.sections = []
          } else {
            item.sections = res.data.sections
          }
          chapters[key] = item
          setChapters([...chapters])
        })
        .catch((e) => {
          console.log(e)
        })
    }
  }
  const addChapter = e => {
    setExist(false)
    setChapterModalVisible(true)
  }
  const modifyChapter = (c, i) => {
    setChapter(c)
    setIndex(i)
    setExist(true)
    setChapterModalVisible(true)
  }
  const addSection = (index) => {
    setExist(false)
    setIndex(index)
    setSectionModalVisible(true)
  }
  const modifySection = (s, index, subIndex) => {
    setSection(s)
    setIndex(index)
    setSubIndex(subIndex)
    setExist(true)
    setSectionModalVisible(true)
  }
  const createChapter = e => {
    if (chapterName === '') {
      message.error('请输入章名')
      return
    }
    if (exist) {
      let c = {
        id: chapter.id,
        projectId: chapter.projectId,
        chapterName: chapterName,
        chapterNumber: chapter.chapterNumber,
      }
      ProjectApi.updateProjectChapter(c)
        .then((res)=>{
          if (res.data.code === 200) {
            let s = chapters[index].sections
            c.sections = s
            chapters[index] = c
            setChapters([...chapters])
            message.success(res.data.msg)
            setChapterModalVisible(false)
            setChapterName('')
          }
        })
        .catch((e)=>{console.log(e)})
    } else {
      let l = chapters.length
      let cp = {chapterName: chapterName, chapterNumber: l, projectId: pid}
      ProjectApi.createProjectChapter(cp)
        .then((res) => {
          setChapterModalVisible(false)
          setChapterName('')
          if (res.data.id) {
            cp.id = res.data.id
            chapters.push(cp)
            setChapters([...chapters])
          }
        })
        .catch((e) => {
          console.log(e)
        })
    }
  }
  const createSection = () => {
    if (sectionName === '') {
      message.error('请输入小节名')
      return
    }
    if (exist) {
      let s = {
        id: section.id,
        chapterId: section.chapterId,
        sectionName: sectionName,
        sectionNumber: section.sectionNumber
      }
      ProjectApi.updateChapterSection(s)
        .then((res)=>{
          if (res.data.code === 200) {
            chapters[index].sections[subIndex] = s
            setChapters([...chapters])
            message.success(res.data.msg)
            setSectionModalVisible(false)
            setSectionName('')
          }
        })
        .catch((e)=>{console.log(e)})
    } else {
      let l = 0
      if (chapters[index].sections !== null) {
        l = chapters[index].sections.length
      }
      let sec = {sectionName: sectionName, sectionNumber: l, chapterId: chapters[index].id}
      ProjectApi.createChapterSection(sec)
        .then((res) => {
          setSectionModalVisible(false)
          setSectionName('')
          if (res.data.id) {
            sec.id = res.data.id;
            if (chapters[index].sections === null) {
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
  const cancelCreateChapter = e => {
    setChapterModalVisible(false)
  }
  const cancelCreateSection = e => {
    setSectionModalVisible(false)
  }

  const changeChapterName = value => {
    setChapterName(value.target.value)
  }
  const changeSectionName = value => {
    setSectionName(value.target.value)
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
              <span style={{ float: 'right', marginRight: '20px' }}>
                <Button shape="circle" type="text" onClick={e => modifyChapter(item, index)} icon={<EditOutlined/>} />
              </span>
            </div>
          }
          onTitleClick={e => handleClick(item, index)}
        >
          {(item.sections === null || item.sections === undefined) ? null :
            item.sections.map((subItem, subIndex) => (
              <Menu.Item key={index.toString() + subIndex.toString()}>
                {subItem.sectionName}
                <span style={{ float: 'right', marginRight: '20px' }}>
                  <Button shape="circle" type="text" onClick={e => modifySection(subItem, index, subIndex)} icon={<EditOutlined/>} />
                </span>
              </Menu.Item>
            ))
          }
          <Button type="round" style={{float: 'right', margin: '3px'}} onClick={e => addSection(index)}>添加小节</Button>
        </SubMenu>
      ))}
      </Menu>
      <div style={{textAlign: 'right'}}>
        <Button type="round" style={{margin: '5px'}} onClick={addChapter}>添加章节</Button>
      </div>
      <div style={{textAlign: 'right', marginTop: '20px', marginRight: '20px'}}>
        <Link to={"/project/info/" + pid}>
          <Button type="primary" size="middle" style={{float: 'right'}}>完成</Button>
        </Link>
        <Link to={"/project/edit/info/" + pid}>
          <Button size="middle" style={{float: 'right', marginRight: '20px'}}>上一页</Button>
        </Link>
      </div>

      <Modal title="" visible={chapterModalVisible} onOk={createChapter} onCancel={cancelCreateChapter}>
        <br/>
        <Row>
          <Col span={3}>
            <p>章名：</p>
          </Col>
          <Col span={20}>
            <Input value={chapterName} onChange={changeChapterName} />
          </Col>
        </Row>
      </Modal>

      <Modal visible={sectionModalVisible} onOk={createSection} onCancel={cancelCreateSection}>
        <br/>
        <Row>
          <Col span={3}>
            <p>小节名：</p>
          </Col>
          <Col span={20}>
            <Input value={sectionName} onChange={changeSectionName} />
          </Col>
        </Row>
      </Modal>
    </div>
  )
}

export default OutlineEditPage
