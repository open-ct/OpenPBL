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
import {Avatar, Button, Comment, Form, Input, message, Pagination, Popconfirm, Switch, Tooltip} from 'antd';
import {DeleteOutlined} from "@ant-design/icons"
import QueueAnim from 'rc-queue-anim';

import CommentApi from "../../../../api/CommentApi"
import util from "../../../component/Util"

const {TextArea} = Input;


function ProjectComment(obj) {
  const [commentList, setCommentList] = useState([]);
  const [actions, setActions] = useState([]);
  const [text, setText] = useState([])
  const [isTeacher, setIsTeacher] = useState(false)
  const [page, setPage] = useState(1);
  const [pageSize, setPageSize] = useState(10)
  const [total, setTotal] = useState(0);
  const [value, setValue] = useState('');
  const [submitting, setSubmitting] = useState(false);

  const updateCommentList = (p, size, isTeacher) => {
    let query = {
      from: (p - 1) * size,
      size: size,
      text: '',
      isTeacher: isTeacher,
    }
    CommentApi.getProjectComments(obj.project.id, query)
      .then(res => {
        if (res.data.code === 200) {
          if (res.data.comments !== null) {
            setCommentList(res.data.comments)
            setTotal(res.data.count)
          }
        }
      })
      .catch(e => {
        console.log(e)
      })

    setPage(p);
  };
  useEffect(() => {
    updateCommentList(1, pageSize, isTeacher);
  }, []);

  const onSubmit = () => {
    if (value === "") {
      message.warn("请输入留言内容")
      return
    }
    setSubmitting(true);
    let c = {
      content: value,
      projectId: obj.project.id,
    }
    CommentApi.createComment(obj.project.id, c)
      .then(res => {
        setSubmitting(false);
        if (res.data.code === 200) {
          message.success(res.data.msg)
          updateCommentList(page, pageSize, isTeacher)
          setValue('');
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e => {
        console.log(e)
      })
  };
  const deleteComment = (item) => {
    CommentApi.deleteComment(obj.project.id, item.id)
      .then(res => {
        if (res.data.code === 200) {
          message.success(res.data.msg)
          updateCommentList(page, pageSize, isTeacher)
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e => {
        console.log(e)
      })
  }

  const onChange = (v) => {
    setValue(v.target.value);
  };
  const onChangeSwitch = e => {
    setIsTeacher(e)
    updateCommentList(page, pageSize, e)
  }

  return (
    <QueueAnim>
      <div style={{textAlign: 'left'}} key="1">
        <Form.Item>
          <TextArea rows={4} onChange={onChange} value={value}/>
        </Form.Item>
        <Form.Item>
          <Button htmlType="submit" loading={submitting} onClick={onSubmit} type="primary">
            提交
          </Button>
        </Form.Item>

        <Switch
          checkedChildren="只看老师"
          unCheckedChildren="所有"
          defaultChecked={false}
          onChange={onChangeSwitch}
          value={isTeacher}
        />
        {commentList.map((item, index) => (
          <Comment
            key={index.toString()}
            actions={
              [
                <>
                  {item.isTeacher ?
                    <span key="0" style={{color: 'green'}}>老师</span>
                    :
                    <span key="0">学生</span>
                  }
                </>,
                <>
                  {obj.account.name === item.userId ?
                    <Popconfirm title="确定删除留言？" onConfirm={e => deleteComment(item)}>
                      <Button type="text" shape="circle" icon={<DeleteOutlined/>}/>
                    </Popconfirm>
                    : null
                  }
                </>
              ]
            }
            author={item.userName}
            avatar={
              <Avatar
                src={item.userAvatar}
                alt={item.userName}
              />
            }
            content={item.content}
            datetime={
              <Tooltip title={util.FilterTime(item.createAt)}>
                <span>{util.FilterMoment(item.createAt)}</span>
              </Tooltip>
            }
          />
        ))
        }<br/>
        <div style={{textAlign: 'right', marginRight: '40px'}}>
          <Pagination
            total={total}
            showTotal={t => `共${total}项内容`}
            current={page}
            onChange={updateCommentList}
            onShowSizeChange={updateCommentList}
          />
        </div>
      </div>
    </QueueAnim>
  );
}

export default ProjectComment;
