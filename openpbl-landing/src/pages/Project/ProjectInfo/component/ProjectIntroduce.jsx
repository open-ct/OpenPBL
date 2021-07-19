import React, {useEffect, useState} from 'react';
import {RiseOutlined, UnorderedListOutlined} from '@ant-design/icons';
import QueueAnim from 'rc-queue-anim';

function ProjectIntroduce(obj) {
  const [project, setProject] = useState({});
  useEffect(() => {
    setProject(obj.project);
  }, []);
  return (
    <QueueAnim>
      <div style={{textAlign: 'left'}} key="1">
        <p>学科：{project.subjects}</p>
        <p>技能：{project.skills}</p><br/>
        <h2><UnorderedListOutlined/>&nbsp;&nbsp;项目简介</h2>
        <p>{project.projectIntroduce}</p><br/>
        <h2><RiseOutlined/>&nbsp;&nbsp;项目目标</h2>
        <p>{project.projectGoal}</p><br/>
      </div>
    </QueueAnim>
  );
}

export default ProjectIntroduce;
