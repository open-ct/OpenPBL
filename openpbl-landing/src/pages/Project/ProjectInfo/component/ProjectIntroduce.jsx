import React, {useEffect, useState} from 'react';
import {RiseOutlined, UnorderedListOutlined} from '@ant-design/icons';
import QueueAnim from 'rc-queue-anim';
import { connect } from "react-redux";


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
