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

import React, {useEffect, useState} from "react";
import {Card, Col, Divider, Image, Row, Empty} from "antd";
import {Link} from "react-router-dom";
import {EyeOutlined, TeamOutlined} from "@ant-design/icons";
import HomeApi from "../../../api/HomeApi";
import util from "../../component/Util";

const topColResponsiveProps = {
  xs: 12,
  sm: 12,
  md: 8,
  lg: 8,
  xl: 6,
};

function LatestProject(obj) {
  const [projects, setProjects] = useState([])

  useEffect(() => {
    getLatestProject()
  }, [])

  const getLatestProject = () => {
    HomeApi.getLatestProjects(4)
      .then(res => {
        if (res.data.code === 200) {
          if (res.data.projects !== null) {
            setProjects(res.data.projects)
          }
        }
      })
      .catch(e => {
        console.log(e)
      })
  }

  return (
    <div style={{maxWidth: '1200px', margin: 'auto'}}>
      <Row gutter={[20, 20]}>
        {
          projects.map((item, index) => (
            <Col key={index.toString()} {...topColResponsiveProps}>
              <Link to={`/home/project/${item.id}/info`}>
                <Card
                  hoverable
                  bordered={false}
                  style={{
                    borderRadius: '10px',
                  }}
                  cover={
                    <Image
                      src={item.image}
                      preview={false}
                      style={{
                        borderTopLeftRadius: '10px',
                        borderTopRightRadius: '10px',
                      }}
                      fallback={"https://cdn.open-ct.com/task-resources//openpbl-empty-project.png"}
                    />
                  }
                >
                  <Card.Meta
                    title={
                      item.projectTitle === '' ? '无':item.projectTitle
                    }
                    description={
                      <div>
                        {item.subjects === '' ? '无' : item.subjects}
                      </div>
                    }
                  />

                  <span
                    style={{
                      color: 'gray',
                      fontSize: 'small',
                    }}
                  >
                      <EyeOutlined/>&nbsp;
                    {item.readNum}
                    </span>
                  <Divider type="vertical"/>
                  <span
                    style={{
                      color: 'gray',
                      fontSize: 'small',
                    }}
                  >
                      <TeamOutlined/>&nbsp;
                    {item.joinNum}
                    </span>
                  <span
                    style={{
                      color: 'gray',
                      fontSize: 'small',
                      float: 'right',
                    }}
                  >
                      {util.FilterMoment(item.createAt)}
                    </span>
                </Card>
              </Link>
            </Col>
          ))
        }
      </Row>
      {projects.length === 0 ?
        <Empty /> : null
      }
    </div>
  )
}

export default LatestProject