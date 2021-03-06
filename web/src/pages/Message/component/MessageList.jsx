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
          <Button onClick={setReadAll} loading={btLoading}>??????????????????</Button>
        </div>
        : null
      }
      <Table
        loading={loading}
        rowKey="id"
        dataSource={messages}
        columns={[
          {
            title: '??????',
            dataIndex: 'readMessage',
            key: 'readMessage',
            render: (text, item, index) =>(
              <>
                {!text ?
                  <span style={{color: 'green'}}>??????</span>
                  :
                  <span style={{color: 'gray'}}>??????</span>
                }
              </>
            )
          },
          {
            title: '??????',
            dataIndex: 'messageTitle',
            key: 'messageTitle'
          },
          {
            title: '??????',
            dataIndex: 'createAt',
            key: 'createAt',
            render: (text, item, index) => (
              <>
                {util.FilterMoment(text)}
              </>
            )
          },
          {
            title: '??????',
            dataIndex: 'messageType',
            key: 'messageType',
            render: (text, item, index) => (
              <>
                {text === 'warning' ? '??????': null}
                {text === 'info' ? '??????': null}
                {text === 'error' ? '??????': null}
              </>
            )
          },
          {
            title: '??????',
            dataIndex: 'content',
            key: 'content'
          },
          {
            title: '??????',
            dataIndex: 'action',
            key: 'action',
            render: (text, item, index) => (
              <>
                <Popconfirm title="?????????????????????" onConfirm={e=>deleteMessage(item)}>
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
        showTotal={t => `???${total}?????????`}
        current={page}
        onChange={changePage}
        onShowSizeChange={()=>updateMessageList(page-1, pageSize)}
        style={{margin: '20px', textAlign: 'right'}}
      />
    </div>
  )
}

export default MessageList