import React, {useEffect, useState} from "react";
import {Menu} from 'antd'
import QueueAnim from 'rc-queue-anim';

import ChapterApi from "../../../../api/ChapterApi";
import SectionApi from "../../../../api/SectionApi";

const {SubMenu} = Menu;

function ProjectOutline(obj) {
  const pid = obj.project.id
  const [chapters, setChapters] = useState([])

  useEffect(() => {
    ChapterApi.getProjectChapters(pid)
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
      SectionApi.getChapterSections(item.id)
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
  const gotoLearning = (item, subItem) => {
    window.open('/project/learning/' + pid + '/' + item.id + '/' + subItem.id)
  }

  return (
    <QueueAnim>
      <div key="1">
        <Menu
          style={{width: '100%'}}
          defaultSelectedKeys={[]}
          defaultOpenKeys={[]}
          mode="inline"
        >{chapters.map((item, index) => (
          <SubMenu style={{fontSize: '2.7vh'}} key={index.toString()} title={item.chapterName}
                   onTitleClick={e => handleClick(item, index)}>
            {(item.sections === null || item.sections === undefined) ? null :
              item.sections.map((subItem, subIndex) => (
                <Menu.Item key={index.toString() + subIndex.toString()} onClick={e => gotoLearning(item, subItem)}>
                  {subItem.sectionName}
                </Menu.Item>
              ))
            }
          </SubMenu>
        ))}
        </Menu>
      </div>
    </QueueAnim>
  )
}

export default ProjectOutline
