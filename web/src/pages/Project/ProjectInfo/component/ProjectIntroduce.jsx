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

import React, {useEffect} from 'react';
import {RiseOutlined, UnorderedListOutlined} from '@ant-design/icons';
import QueueAnim from 'rc-queue-anim';


function ProjectIntroduce(obj) {
  useEffect(() => {
  }, []);
  return (
    <QueueAnim>
      <div style={{textAlign: 'left'}} key="1">
        <p>学科：{obj.project.subjects}</p>
        <p>技能：{obj.project.skills}</p><br/>
        <h2><UnorderedListOutlined/>&nbsp;&nbsp;项目简介</h2>
        <p>{obj.project.projectIntroduce}</p><br/>
        <h2><RiseOutlined/>&nbsp;&nbsp;项目目标</h2>
        <p>{obj.project.projectGoal}</p><br/>
      </div>
    </QueueAnim>
  );
}

export default ProjectIntroduce
