import React, {useEffect, useState} from 'react';
import QueueAnim from 'rc-queue-anim';
import project from "../../../../api/project";
import {Table} from "antd";

const columns = [
  {
    title: '学生',
    dataIndex: 'name',
    key: 'name'
  },
  {
    title: '邮箱',
    dataIndex: 'studentEmail',
    key: 'studentEmail'
  },
  {
    title: '操作',
    dataIndex: 'action',
    key: 'action'
  },
]

function StudentAdmin(obj) {
  const pid = obj.project.id
  const [students, setStudents] = useState([])
  useEffect(() => {
    project.getProjectStudents(pid)
      .then((res) => {
        if (res.data.students) {
          setStudents(res.data.students);
        }
      })
  }, []);

  return (
    <QueueAnim>
      <div style={{textAlign: 'left'}} key="1">
        <Table dataSource={students} columns={columns}/>
      </div>
    </QueueAnim>
  );
}

export default StudentAdmin;
