import React, {useEffect, useState} from "react";
import {Button, Col, Input, message, Row, Select, Upload} from "antd";
import {Link} from 'react-router-dom'
import ImgCrop from 'antd-img-crop';
import '../../Outline/index.less'
import {LoadingOutlined, PlusOutlined} from "@ant-design/icons";
import ProjectApi from "../../../../../api/ProjectApi";


function InfoEditPage(obj) {
  const pid = obj.pid

  const subjects = ['语文', '数学', '英语', '科学']
  const skills = ['学习与创新技能', '信息、媒体与技术技能', '生活与职业技能']

  const [p, setP] = useState({})

  const [selectedSubjects, setSelectedSubjects] = useState([])
  const [selectedSkills, setSelectedSkills] = useState([])
  const [loading, setLoading] = useState(false)
  const [imageUrl, setImageUrl] = useState()
  const [projectTitle, setProjectTitle] = useState('')
  const [projectIntroduce, setProjectIntroduce] = useState('')
  const [projectGoal, setProjectGoal] = useState('')

  const stringToList = str => {
    if (str === '') {
      return []
    }
    return str.split(',')
  }
  useEffect(() => {
    ProjectApi.getProjectDetail(pid)
      .then((res) => {
        if (res.data.code === 200) {
          console.log(res.data.project)
          setP(res.data.project)
          setImageUrl(res.data.project.image)
          setProjectTitle(res.data.project.projectTitle)
          setProjectIntroduce(res.data.project.projectIntroduce)
          setProjectGoal(res.data.project.projectGoal)
          setSelectedSkills(stringToList(res.data.project.skills))
          setSelectedSubjects(stringToList(res.data.project.subjects))
        }
      })
      .catch((e) => {
        console.log(e)
      })
  }, [])

  const changeTitle = value => {
    setProjectTitle(value.target.value)
  }
  const changeIntroduce = value => {
    setProjectIntroduce(value.target.value)
  }
  const changeGoal = value => {
    setProjectGoal(value.target.value)
  }

  const onUploadImage = (file) => {
    setLoading(true)
    file = file.file;
    const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png';
    if (!isJpgOrPng) {
      message.error('只能上传 JPG/PNG 格式文件');
    }
    const isLt2M = file.size / 1024 / 1024 < 2;
    if (!isLt2M) {
      message.error('图片应当小于 2MB');
    }

    const r = new FileReader();
    r.addEventListener('load', () => upload(r.result));
    r.readAsDataURL(file);
  }

  const upload = (file) => {
    setImageUrl(file)
    setLoading(true)

  }

  const onPreview = file => {
    let src = file.url;
    if (!src) {
      src = new Promise(resolve => {
        const reader = new FileReader();
        reader.readAsDataURL(file.originFileObj);
        reader.onload = () => resolve(reader.result);
      });
    }
    const image = new Image();
    image.src = src;
    const imgWindow = window.open(src);
    imgWindow.document.write(image.outerHTML);
  }

  const onFinish = () => {
    let data = {
      id: pid,
      image: imageUrl,
      projectTitle: projectTitle,
      projectIntroduce: projectIntroduce,
      projectGoal: projectGoal,
      subjects: selectedSubjects.toString(),
      skills: selectedSkills.toString(),
    }
    ProjectApi.updateProject(data, pid)
      .then((res) => {
        if (res.data.code === 200) {
          window.location.href = `/project/${pid}/outline/edit`
        }
      })
      .catch((e) => {
        console.log(e)
      })
  }
  const handleSubjectsChange = selected => {
    setSelectedSubjects(selected)
  }
  const handleSkillsChange = selected => {
    setSelectedSkills(selected)
  }

  return (
    <div>
      <Row>
        <Col span={4} offset={2}>上传图片：</Col>
        <Col span={16}>
          <ImgCrop rotate aspect="1.2">
            <Upload
              name="avatar"
              listType="picture-card"
              className="avatar-uploader"
              showUploadList={false}
              customRequest={onUploadImage}
              onPreview={onPreview}
            >
              {
                imageUrl ? <img src={imageUrl} alt="avatar" style={{width: '100%'}}/>
                  :
                  <div>
                    {loading ? <LoadingOutlined/> : <PlusOutlined/>}
                    <div style={{marginTop: 8}}>上传图片</div>
                  </div>
              }
            </Upload>
          </ImgCrop>
        </Col>
      </Row>
      <Row style={{marginTop: '20px'}}>
        <Col span={4} offset={2}>输入标题：</Col>
        <Col span={16}><Input value={projectTitle} maxLength={40} onChange={changeTitle}/></Col>
      </Row>
      <Row style={{marginTop: '20px'}}>
        <Col span={4} offset={2}>输入简介：</Col>
        <Col span={16}>
          <Input.TextArea rows={5} value={projectIntroduce} maxLength={1000} onChange={changeIntroduce}/>
        </Col>
      </Row>
      <Row style={{marginTop: '20px'}}>
        <Col span={4} offset={2}>输入目标：</Col>
        <Col span={16}>
          <Input.TextArea rows={3} value={projectGoal} maxLength={500} onChange={changeGoal}/>
        </Col>
      </Row>

      <Row style={{marginTop: '20px'}}>
        <Col span={4} offset={2}>选择学科：</Col>
        <Col span={16}>
          <Select
            showArrow
            mode="multiple"
            placeholder="选择学科"
            value={selectedSubjects}
            onChange={handleSubjectsChange}
            style={{width: '100%'}}
          >
            {subjects.map(item => (
              <Select.Option key={item} value={item}>
                {item}
              </Select.Option>
            ))}
          </Select>
        </Col>
      </Row>
      <Row style={{marginTop: '20px'}}>
        <Col span={4} offset={2}>选择技能：</Col>
        <Col span={16}>
          <Select
            showArrow
            mode="multiple"
            placeholder="选择技能"
            value={selectedSkills}
            onChange={handleSkillsChange}
            style={{width: '100%'}}
          >
            {skills.map(item => (
              <Select.Option key={item} value={item}>
                {item}
              </Select.Option>
            ))}
          </Select>
        </Col>
      </Row>
      <Row style={{marginTop: '20px'}}>
        <Col span={4} offset={20}>
          <Link to={`/project/${pid}/outline/edit`}>
            <Button
              size="middle"
              style={{marginRight: '20px'}}
            >下一页</Button>
          </Link>
          <Button
            type="primary"
            size="middle"
            onClick={onFinish}
          >提交</Button>
        </Col>
      </Row>
    </div>
  )
}

export default InfoEditPage
