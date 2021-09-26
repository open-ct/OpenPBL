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

import React, {useEffect, useState} from 'react';
import QueueAnim from 'rc-queue-anim';
import {Avatar, Button, message, Pagination, Popconfirm, Table} from "antd";
import {DeleteOutlined} from "@ant-design/icons"
import {Link} from "react-router-dom";

import ProjectApi from "../../../../api/ProjectApi";
import util from "../../../component/Util"

function StudentAdmin(obj) {
  const pid = obj.project.id
  const [students, setStudents] = useState([])
  const [loadStudents, setLoadStudents] = useState(false)

  const [total, setTotal] = useState(0)
  const [page, setPage] = useState(1)

  useEffect(() => {
    if (pid !== undefined) {
      updateStudentList()
    }
  }, [pid]);

  const updateStudentList = () => {
    setLoadStudents(true)
    ProjectApi.getProjectStudents(pid)
      .then((res) => {
        setLoadStudents(false)
        if (res.data.code === 200) {
          setStudents(res.data.students);
          setTotal(res.data.count)
        }
      })
      .catch(e => {
        console.log(e)
      })
  }
  const removeStudent = (action, record) => {
    ProjectApi.removeStudent(record.projectId, record.studentId)
      .then(res => {
        if (res.data.code === 200) {
          message.success(res.data.msg)
          updateStudentList()
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e => {
        console.log(e)
      })
  }

  return (
    <QueueAnim>
      <div style={{textAlign: 'left'}} key="1">
        <Table
          loading={loadStudents}
          dataSource={students}
          columns={[
            {
              title: '头像',
              dataIndex: 'avatar',
              key: 'avatar',
              render: avatar => (
                <Avatar src={avatar}/>
              )
            },
            {
              title: '学生',
              dataIndex: 'name',
              key: 'name'
            },
            {
              title: '加入时间',
              dataIndex: 'joinTime',
              key: 'joinTime',
              render: joinTime => (
                <span>{util.FilterTime(joinTime)}</span>
              )
            },
            {
              title: '',
              dataIndex: 'showEvidence',
              render: (text, item, index) => (
                <Link to={`/home/project/${pid}/student/${item.studentId}/evidence`}>
                  <Button type="text">查看学习证据</Button>
                </Link>
              )
            },
            {
              title: '操作',
              dataIndex: 'action',
              key: 'action',
              render: (action, item) => (
                <Popconfirm title="确定移除学生？" onConfirm={e => removeStudent(action, item)}>
                  <Button shape="circle" type="text" style={{color: 'red'}} icon={<DeleteOutlined/>}/>
                </Popconfirm>
              )
            },
          ]}
          pagination={false}
        />
        <Pagination
          total={total}
          showTotal={t => `共${total}名学生参加`}
          current={page}
          onChange={updateStudentList}
          onShowSizeChange={updateStudentList}
          style={{marginTop: '20px', textAlign: 'right'}}
        />
      </div>
    </QueueAnim>
  );
}

export default StudentAdmin;
