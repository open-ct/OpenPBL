import React, {useEffect, useState} from "react";
import {Button, Pagination, Popconfirm, Table} from "antd";
import {DeleteOutlined} from "@ant-design/icons"

import util from "../../component/Util"
import MessageApi from "../../../api/MessageApi";

function MessageList(obj) {
  const [page, setPage] = useState(1)
  const [pageSize, setPageSize] = useState(10)
  const [messages, setMessages] = useState([])
  const [loading, setLoading] = useState(false)
  const [btLoading, setBtLoading] = useState(false)
  const [total, setTotal] = useState(0)

  useEffect(()=>{
    updateMessageList((page-1)*pageSize, pageSize);
  }, [])

  const updateMessageList = (from, size) => {
    setLoading(true)
    const query = {
      readType: obj.readType,
      messageType: 'all',
      from: from,
      size: size,
      orderType: 'desc'
    }
    MessageApi.getUserMessages(query)
      .then(res=>{
        setLoading(false)
        if (res.data.code === 200) {

          setMessages(res.data.messages)
          setTotal(res.data.count)
        }
      })
      .catch(e=>{console.log(e)})
  }

  const changePage = p => {
    setPage(p)
    updateMessageList((p-1)*pageSize, pageSize)
  }

  const deleteMessage = item => {
    setLoading(true)
    MessageApi.deleteMessage(item)
      .then(res=>{
        setLoading(false)
        if (res.data.code === 200) {
          updateMessageList((page-1)*pageSize, pageSize)
        }
      })
      .catch(e=>{console.log(e)})
  }
  const setReadMessage = (expanded, record) => {
    if (expanded && !record.readMessage) {
      MessageApi.readMessage(record)
        .then(res => {
          if (res.data.code === 200) {
            updateMessageList((page - 1) * pageSize, pageSize)
          }
        })
        .catch(e => {
          console.log(e)
        })
    }
  }
  const setReadAll = () => {
    setBtLoading(true)
    MessageApi.readAllMessage()
      .then(res=>{
        setBtLoading(false)
        if (res.data.code === 200) {
          updateMessageList((page - 1) * pageSize, pageSize)
        }
      })
      .catch(e=>{console.log(e)})
  }

  return (
    <div>
      {obj.readType === "all" || obj.readType === "unread" ?
        <div style={{float: 'right', padding: '20px'}}>
          <Button onClick={setReadAll} loading={btLoading}>全部设为已读</Button>
        </div>
        : null
      }
      <Table
        loading={loading}
        rowKey="id"
        dataSource={messages}
        columns={[
          {
            title: '状态',
            dataIndex: 'readMessage',
            key: 'readMessage',
            render: (text, item, index) =>(
              <>
                {!text ?
                  <span style={{color: 'green'}}>未读</span>
                  :
                  <span style={{color: 'gray'}}>已读</span>
                }
              </>
            )
          },
          {
            title: '标题',
            dataIndex: 'messageTitle',
            key: 'messageTitle'
          },
          {
            title: '时间',
            dataIndex: 'createAt',
            key: 'createAt',
            render: (text, item, index) => (
              <>
                {util.FilterMoment(text)}
              </>
            )
          },
          {
            title: '类型',
            dataIndex: 'messageType',
            key: 'messageType',
            render: (text, item, index) => (
              <>
                {text === 'warning' ? '警告': null}
                {text === 'info' ? '信息': null}
                {text === 'error' ? '错误': null}
              </>
            )
          },
          {
            title: '内容',
            dataIndex: 'content',
            key: 'content'
          },
          {
            title: '操作',
            dataIndex: 'action',
            key: 'action',
            render: (text, item, index) => (
              <>
                <Popconfirm title="确定删除消息？" onConfirm={e=>deleteMessage(item)}>
                  <Button shape="circle" type="text" icon={<DeleteOutlined />}/>
                </Popconfirm>
              </>
            )
          }
        ]}
        expandable={{
          expandedRowRender: record => <p style={{ margin: 10 }}>{record.content}</p>,
          rowExpandable: record => true,
          onExpand: setReadMessage
        }}

        pagination={false}
      />
      <Pagination
        total={total}
        showTotal={t => `共${total}条消息`}
        current={page}
        onChange={changePage}
        onShowSizeChange={()=>updateMessageList(page-1, pageSize)}
        style={{margin: '20px', textAlign: 'right'}}
      />
    </div>
  )
}

export default MessageList