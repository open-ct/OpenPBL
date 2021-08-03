import React, {useEffect, useState} from "react";
import {Menu, Progress} from 'antd'
import QueueAnim from 'rc-queue-anim';

import ChapterApi from "../../../../api/ChapterApi";
import SectionApi from "../../../../api/SectionApi";

const {SubMenu} = Menu;

function ProjectOutline(obj) {
  const pid = obj.project.id
  const [chapters, setChapters] = useState([])

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
      .catch(e=>{console.log(e)})
  }
  const gotoLearning = (item, subItem) => {
    window.open(`/project/${pid}/section/${subItem.id}/preview?back=/project/${pid}/info`)
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
          <SubMenu style={{fontSize: '2.7vh'}} key={index.toString()} title={item.chapterName}>
            {(item.sections === null || item.sections === undefined) ? null :
              item.sections.map((subItem, subIndex) => (
                <Menu.Item key={index.toString() + subIndex.toString()} onClick={e => gotoLearning(item, subItem)}>
                  {subItem.sectionName}
                  {obj.project.learning ?
                    <>
                      <span style={{float: 'right'}}>
                        <Progress
                          trailColor="lightgray"
                          width={35}
                          strokeWidth={10}
                          type="circle"
                          percent={((subItem.learnMinute+subItem.learnSecond/60)/subItem.sectionMinute*100).toFixed(1)}
                        />
                      </span>
                      <span style={{float: 'right', marginRight: '20px'}}>
                        学习进度：
                        {subItem.learnMinute}&nbsp;:&nbsp;{subItem.learnSecond}&nbsp;/&nbsp;
                        {subItem.sectionMinute}
                      </span>
                    </>
                    : null }
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
