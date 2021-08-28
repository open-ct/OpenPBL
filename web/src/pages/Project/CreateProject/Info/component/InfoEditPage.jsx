import React, {useEffect, useState} from "react";
import {Button, Col, DatePicker, Input, message, Row, Select, Upload} from "antd";
import ImgCrop from 'antd-img-crop';
import {LoadingOutlined, PlusOutlined} from "@ant-design/icons";

import '../../Outline/index.less'
import ProjectApi from "../../../../../api/ProjectApi";
import util from "../../../../component/Util"

function InfoEditPage(obj) {
  const pid = obj.pid

  const iniSubjects = ['语文', '数学', '英语', '科学', '政治', '历史', '地理', '化学', '生物', '美术', '音乐']
  const iniSkills = ['信息、媒体与技术技能', '生活与职业技能', '文化理解与传承素养', '审辨思维', '创新素养', '沟通素养', '合作素养']
  const [subjects, setSubjects] = useState(iniSubjects)
  const [skills, setSkills] = useState(iniSkills)

  const [nextPageLoading, setNextPageLoading] = useState(false)

  const [change, setChange] = useState(false)

  const [selectedSubjects, setSelectedSubjects] = useState([])
  const [selectedSkills, setSelectedSkills] = useState([])
  const [loading, setLoading] = useState(false)
  const [imageUrl, setImageUrl] = useState()
  const [projectTitle, setProjectTitle] = useState('')
  const [projectIntroduce, setProjectIntroduce] = useState('')
  const [projectGoal, setProjectGoal] = useState('')

  const [endTimeStr, setEndTimestr] = useState('')
  const [endTime, setEndTime] = useState()

  const stringToList = str => {
    if (str === '') {
      return []
    }
    return str.split(',')
  }
  useEffect(() => {
    loadSubjectsAndSkills()
    ProjectApi.getProjectDetail(pid)
      .then((res) => {
        if (res.data.code === 200) {
          setImageUrl(res.data.project.image)
          setProjectTitle(res.data.project.projectTitle)
          setProjectIntroduce(res.data.project.projectIntroduce)
          setProjectGoal(res.data.project.projectGoal)
          setSelectedSkills(stringToList(res.data.project.skills))
          setSelectedSubjects(stringToList(res.data.project.subjects))
          if (res.data.project.timedEnd) {
            setEndTime(util.GetMoment(res.data.project.endTime.slice(0, 10)))
          }
        }
      })
      .catch((e) => {
        console.log(e)
      })
  }, [])

  const loadSubjectsAndSkills = () => {
    ProjectApi.getSubjectsAndSkills(pid)
      .then(res => {
        if (res.data.code === 200) {
          if (res.data.subjects !== null) {
            let s = res.data.subjects
            for (let i = 0; i < iniSubjects.length; i++) {
              if (s.indexOf(iniSubjects[i]) < 0) {
                s.push(iniSubjects[i])
              }
            }
            setSubjects(s)
          }
          if (res.data.skills !== null) {
            let s = res.data.skills
            for (let i = 0; i < iniSkills.length; i++) {
              if (s.indexOf(iniSkills[i]) < 0) {
                s.push(iniSkills[i])
              }
            }
            setSkills(s)
          }
        }
      })
      .catch(e => {
        console.log(e)
      })
  }

  const changeTitle = value => {
    setChange(true)
    setProjectTitle(value.target.value)
  }
  const changeIntroduce = value => {
    setChange(true)
    setProjectIntroduce(value.target.value)
  }
  const changeGoal = value => {
    setChange(true)
    setProjectGoal(value.target.value)
  }

  const onUploadImage = (file) => {
    setChange(true)

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
      endTime: endTimeStr + " 23:59:59 +0800 CST",
    }
    ProjectApi.updateProject(data, pid)
      .then((res) => {
        setNextPageLoading(false)
        if (res.data.code === 200) {
          message.success(res.data.msg)
          setTimeout(() => {
            window.location.href = `/home/project/${pid}/outline/edit`
          }, 200)
        }
      })
      .catch((e) => {
        console.log(e)
      })
  }
  const nextPage = () => {
    setNextPageLoading(true)
    if (checkInput()) {
      if (change) {
        onFinish();
      } else {
        setNextPageLoading(false)
        window.location.href = `/home/project/${pid}/outline/edit`
      }
    }
  }
  const checkInput = () => {
    if (imageUrl === "") {
      message.error("请上传图片")
      return false
    }
    if (projectTitle === "") {
      message.error("请输入标题")
      return false
    }
    if (projectIntroduce === "") {
      message.error("请输入简介")
      return false
    }
    return true
  }

  const handleSubjectsChange = selected => {
    setChange(true)
    setSelectedSubjects(selected)
  }
  const handleSkillsChange = selected => {
    setChange(true)
    setSelectedSkills(selected)
  }

  const chooseEndTime = (date, dateString) => {
    setChange(true)
    setEndTimestr(dateString)
    setEndTime(date)

    console.log(date)
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
            mode="tags"
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
        <Col span={4} offset={2}>结束时间：</Col>
        <Col span={16}>
          <DatePicker value={endTime} onChange={chooseEndTime} placeholder="选择结束日期"/>
        </Col>
      </Row>
      <Row style={{marginTop: '20px'}}>
        <Col span={4} offset={20}>
          <Button
            size="middle"
            style={{marginRight: '20px'}}
            onClick={nextPage}
            loading={nextPageLoading}
          >下一页
          </Button>
        </Col>
      </Row>
    </div>
  )
}

export default InfoEditPage
