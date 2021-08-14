import React from "react";
import MessageList from "../component/MessageList";

function UnreadMessage() {
  return (
    <div>
      <MessageList readType="unread"/>
    </div>
  )
}

export default UnreadMessage