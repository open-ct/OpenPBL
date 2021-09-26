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

import React, {useEffect, useState} from "react";
import {Menu, message, Progress} from 'antd'
import QueueAnim from 'rc-queue-anim';

import ChapterApi from "../../../../api/ChapterApi";
import util from "../../../component/Util"

const {SubMenu} = Menu;

function ProjectOutline(obj) {
  const pid = obj.project.id
  const [chapters, setChapters] = useState([])

  useEffect(() => {
    if (pid !== undefined) {
      getChapters()
    }
  }, [pid])
  const getChapters = () => {
    ChapterApi.getProjectChapters(pid)
      .then((res) => {
        if (res.data.chapters === null) {
          setChapters([])
        } else {
          setChapters(res.data.chapters)
        }
      })
      .catch(e => {
        console.log(e)
      })
  }
  const gotoLearning = (item, subItem) => {
    if (!obj.project.learning && obj.account.tag === '学生') {
      message.warn("请先加入学习")
      return
    }
    window.location.href = (`/home/project/${pid}/section/${subItem.id}/preview?back=/home/project/${pid}/info`)
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
          <SubMenu style={{fontSize: '2.7vh'}} key={index.toString()}
                   title={util.FormatChapterName(item.chapterName, item.chapterNumber)}>
            {(item.sections === null || item.sections === undefined) ? null :
              item.sections.map((subItem, subIndex) => (
                <Menu.Item key={index.toString() + subIndex.toString()} onClick={e => gotoLearning(item, subItem)}>
                  {util.FormatSectionName(subItem.sectionName, subItem.chapterNumber, subItem.sectionNumber)}
                  {obj.project.learning ?
                    <>
                      <span style={{float: 'right'}}>
                        <Progress
                          trailColor="lightgray"
                          width={35}
                          strokeWidth={10}
                          type="circle"
                          percent={((subItem.learnMinute + subItem.learnSecond / 60) / subItem.sectionMinute * 100).toFixed(1)}
                        />
                      </span>
                      <span style={{float: 'right', marginRight: '20px'}}>
                        学习进度：
                        {subItem.learnMinute}&nbsp;:&nbsp;{subItem.learnSecond}&nbsp;/&nbsp;
                        {subItem.sectionMinute}
                      </span>
                    </>
                    : null}
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
