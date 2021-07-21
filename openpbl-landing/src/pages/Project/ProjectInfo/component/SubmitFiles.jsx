import React, {useEffect, useState} from 'react';
import {Button, Dropdown, Input, Menu, message, Modal, Table, Upload} from 'antd'
import {SettingOutlined, UploadOutlined} from '@ant-design/icons';
import QueueAnim from 'rc-queue-anim';
import ProjectApi from "../../../../api/ProjectApi";


const menu = (
  <Menu style={{width: '200px'}}>
    <Menu.Item>
      选项
    </Menu.Item>
    <Menu.Item danger>删除</Menu.Item>
  </Menu>
);
const columns = [
  {
    title: '文件名',
    dataIndex: 'fileName',
    key: 'fileName'
  },
  {
    title: '介绍',
    dataIndex: 'submitIntroduce',
    key: 'submitIntroduce',
  },
  {
    title: '提交时间',
    dataIndex: 'createAt',
    key: 'createAt',
  },
  {
    title: '操作',
    dataIndex: 'action',
    render: () => (
      <Dropdown overlay={menu} trigger="click">
        <SettingOutlined style={{cursor: 'pointer'}}/>
      </Dropdown>
    ),
  }
];

const props = {
  name: 'file',
  action: '',
  headers: {
    authorization: 'authorization-text',
  },
  onChange(info) {
    if (info.file.status !== 'uploading') {
      console.log(info.file, info.fileList);
    }
    if (info.file.status === 'done') {
      message.success(`${info.file.name} file uploaded successfully`);
    } else if (info.file.status === 'error') {
      message.error(`${info.file.name} file upload failed.`);
    }
  },
};

function SubmitFiles(obj) {
  const pid = obj.project.id
  const sid = localStorage.getItem('uid')
  const [isModalVisible, setIsModalVisible] = useState(false)
  const [submitFiles, setSubmitFiles] = useState([])
  const [introduce, setIntroduce] = useState('')

  useEffect(() => {
    ProjectApi.getSubmitFiles(sid, pid)
      .then((res) => {
        if (res.data.files === null) {
          setSubmitFiles([])
        } else {
          setSubmitFiles(res.data.files)
        }
      })
      .catch((e) => {
        console.log(e)
      })
  }, []);

  const handleOk = e => {
    let s = {
      studentId: sid,
      projectId: pid,
      submitIntroduce: introduce,
      filePath: '',
      fileName: ''
    }
    ProjectApi.createSubmitFiles(s)
      .then((res) => {
        setIntroduce('')
        if (res.data.id) {
          s.id = res.data.id
          submitFiles.push(s)
          setSubmitFiles([...submitFiles])
        }
      })
      .catch((e) => {
        console.log(e)
      })

    setIsModalVisible(false)
  }
  const handleCancel = e => {
    setIsModalVisible(false)
  }
  const sub = e => {
    setIsModalVisible(true)
  }
  const changeIntroduce = e => {
    setIntroduce(e.target.value)
  }

  return (
    <QueueAnim>
      <div style={{textAlign: 'left'}} key="1">
        <Table dataSource={submitFiles} columns={columns}/>
        <Button type="text" onClick={sub}>提交材料</Button>
      </div>
      <Modal title="提交材料" visible={isModalVisible} onOk={handleOk} onCancel={handleCancel}>
        介绍：
        <Input.TextArea maxLength={200} value={introduce} onChange={changeIntroduce}/>
        <Upload {...props}>
          <Button icon={<UploadOutlined/>} style={{marginTop: '20px'}}>
            上传文件
          </Button>
        </Upload>
      </Modal>
    </QueueAnim>
  );
}

export default SubmitFiles;
