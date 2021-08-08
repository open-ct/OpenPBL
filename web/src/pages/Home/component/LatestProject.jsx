import React, {useEffect, useState} from "react";
import HomeApi from "../../../api/HomeApi";
import {Card, Col, Divider, Image, Row} from "antd";
import {Link} from "react-router-dom";
import {EyeOutlined, TeamOutlined} from "@ant-design/icons";
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
          setProjects(res.data.projects)
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
              <Link to={`/project/${item.id}/info`}>
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
                    />
                  }
                >
                  <Card.Meta
                    title={item.projectTitle}
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
    </div>
  )
}

export default LatestProject