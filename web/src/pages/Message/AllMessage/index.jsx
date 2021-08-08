import React, {useState} from "react";


import MessageList from "../component/MessageList";
import MessageApi from "../../../api/MessageApi"

function AllMessage() {
  const [messages, setMessages] = useState([])
  const [loading, setLoading] = useState(false)
  const [total, setTotal] = useState(0)

  const updateMessageList = (from, size) => {
    setLoading(true)
    const query = {
      readType: 'all',
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
  return (
    <div>
      <MessageList
        loading={loading}
        total={total}
        updateMessageList={updateMessageList}
        messages={messages}
      />
    </div>
  )
}

export default AllMessage