import React, {useEffect, useState} from "react";
import {Button, Pagination, Table} from "antd";

import util from "../../component/Util"

function MessageList(obj) {
  const [page, setPage] = useState(1)
  const [pageSize, setPageSize] = useState(10)

  useEffect(()=>{
    obj.updateMessageList(page - 1, pageSize);
  }, [])

  const changePage = p => {
    setPage(p)
    obj.updateMessageList((p-1)*pageSize, pageSize)
  }

  return (
    <div>
      <Table
        loading={obj.loading}
        dataSource={obj.messages}
        columns={[
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
            title: '操作',
            dataIndex: 'active',
            key: 'active',
            render: (text, item, index) => (
              <>
                <Button>未读</Button>
              </>
            )
          }
        ]}
        pagination={false}
      />
      <Pagination
        total={obj.total}
        showTotal={t => `共${obj.total}条消息`}
        current={page}
        onChange={changePage}
        onShowSizeChange={()=>obj.updateMessageList(page-1, pageSize)}
        style={{margin: '20px', textAlign: 'right'}}
      />
    </div>
  )
}

export default MessageList