import React, {useEffect, useState} from "react";
import {Button, Input, Menu, Modal} from 'antd'
import {Link} from 'react-router-dom'

import ProjectApi from "../../../../api/ProjectApi";

const {SubMenu} = Menu;

function OutlineEditPage(obj) {
  const pid = obj.pid

  const [chapters, setChapters] = useState([])
  const [isModalVisible, setIsModalVisible] = useState(false)

  const [add, setAdd] = useState('')
  const [index, setIndex] = useState()

  const [name, setName] = useState('')
  const [introduce, setIntroduce] = useState('')

  useEffect(() => {
    ProjectApi.getProjectChapters(pid)
      .then((res) => {
        if (res.data.chapters === null) {
          setChapters([])
        } else {
          setChapters(res.data.chapters)
        }
      })
  }, [])

  const handleClick = (item, key) => {
    if (item.sections === undefined) {
      ProjectApi.getChapterSections(item.id)
        .then((res) => {
          item.sections = res.data.sections
          chapters[key] = item
          setChapters([...chapters])
        })
        .catch((e) => {
          console.log(e)
        })
    }
  }
  const createChapter = e => {
    setIsModalVisible(true)
    setAdd('chapter')
  }
  const createSection = (index) => {
    setIsModalVisible(true)
    setAdd('section')
    setIndex(index)
  }
  const handleOk = e => {
    if (name === '') {

    } else {
      if (add === 'chapter') {
        let l = chapters.length
        let cp = {chapterName: name, chapterNumber: l, projectId: pid}
        ProjectApi.createProjectChapter(cp)
          .then((res) => {
            setIsModalVisible(false)
            if (res.data.id) {
              cp.id = res.data.id
              chapters.push(cp)
              setChapters([...chapters])
              setName('')
              setIntroduce('')
            }
          })
          .catch((e) => {
            console.log(e)
          })
      } else {
        let l = 0
        if (chapters[index].sections !== null) {
          l = chapters[index].sections.length
        }
        let sec = {sectionName: name, sectionNumber: l, chapterId: chapters[index].id}
        ProjectApi.createChapterSection(sec)
          .then((res) => {
            setIsModalVisible(false)
            if (res.data.id) {
              sec.id = res.data.id;
              if (chapters[index].sections === null) {
                chapters[index].sections = [sec]
              } else {
                chapters[index].sections.push(sec)
              }
              setChapters([...chapters])
              setName('')
              setIntroduce('')
            }
          })
          .catch((e) => {
            console.log(e)
          })
      }
    }
  }
  const handleCancel = e => {
    setIsModalVisible(false)
  }

  const changeName = value => {
    setName(value.target.value)
  }
  const changeIntroduce = i => {
    setIntroduce(i.target.value)
  }

  return (
    <div>
      <Menu
        style={{width: '100%'}}
        defaultSelectedKeys={[]}
        defaultOpenKeys={[]}
        mode="inline"
      >{chapters.map((item, index) => (
        <SubMenu key={index.toString()} title={item.chapterName} onTitleClick={e => handleClick(item, index)}>
          {(item.sections === null || item.sections === undefined) ? null :
            item.sections.map((subItem, subIndex) => (
              <Menu.Item key={index.toString() + subIndex.toString()}>
                {subItem.sectionName}
              </Menu.Item>
            ))
          }
          <Button type="round" style={{float: 'right', margin: '3px'}} onClick={e => createSection(index)}>添加小节</Button>
        </SubMenu>
      ))}
      </Menu>
      <div style={{textAlign: 'right'}}>
        <Button type="round" style={{margin: '5px'}} onClick={createChapter}>添加章节</Button>
      </div>
      <div style={{textAlign: 'right', marginTop: '20px', marginRight: '20px'}}>
        <Link to={"/project/info/" + pid}>
          <Button type="primary" size="middle" style={{float: 'right'}}>完成</Button>
        </Link>
        <Link to={"/project/edit/info/" + pid}>
          <Button size="middle" style={{float: 'right', marginRight: '20px'}}>上一页</Button>
        </Link>
      </div>

      <Modal title="" visible={isModalVisible} onOk={handleOk} onCancel={handleCancel}>
        <br/>
        <div>
          {add === 'chapter' ? '章名：' : '小节名：'}
          <Input value={name} maxLength={20} onChange={changeName}/>
        </div>
        <div style={{marginTop: '10px'}}>介绍：
          <Input.TextArea row={3} maxLength={200} value={introduce} onChange={changeIntroduce}/>
        </div>
      </Modal>
    </div>
  )
}

export default OutlineEditPage
