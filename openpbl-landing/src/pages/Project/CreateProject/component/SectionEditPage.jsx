import React, {useEffect} from "react";
import {Col, Input, Row} from "antd";

function SectionEditPage(obj) {
  const pid = obj.pid;

  useEffect(()=>{

  }, [])

  return (
    <div>
      <Row>
        <Col span={3}>
          小节名：
        </Col>
        <Col span={20}>
          <Input />
        </Col>
      </Row>
    </div>
  )
}

export default SectionEditPage
