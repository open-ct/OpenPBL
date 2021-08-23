import React, {useEffect, useState} from "react";
import {Button, Card, Divider, message} from "antd";
import LzEditor from "react-lz-editor"
import uniqBy from 'lodash/uniqBy';
import findIndex from 'lodash/findIndex';

import ResourceApi from "../../../../../api/ResourceApi";
import "./section-edit.less"
import FileApi from "../../../../../api/FileApi";
import SectionApi from "../../../../../api/SectionApi";

function RichWords(obj) {
  const [fileList, setFileList] = useState([])
  const [content, setContent] = useState(obj.content)

  useEffect(() => {
  }, [])
  const receiveHtml = (content) => {
    setContent(content)
  }

  const saveContent = () => {
    ResourceApi.updateResourceContent(obj.pid, obj.section.resource.id, content)
      .then(res => {
        if (res.data.code === 200) {
          obj.section.resource.content = content
          message.success(res.data.msg)
        } else {
          message.error(res.data.msg)
        }
      })
      .catch(e => {
        console.log(e)
      })
  }

  const handleChange = (changedValue) => {
    let currFileList = changedValue.fileList;
    console.error(JSON.stringify(changedValue));
    currFileList = currFileList.filter(f => (!f.length));
    currFileList = currFileList.map((file) => {
      if (file.response) {
        file.url = file.response.url;
      }
      if (!file.length) {
        return file;
      }
    });
    currFileList = currFileList.filter((file) => {
      const hasNoExistCurrFileInUploadedList = !~findIndex(
        fileList, item => item.name === file.name,
      );
      if (hasNoExistCurrFileInUploadedList) {
        fileList.push(file);
      }
      return !!file.response || (!!file.url && file.status === 'done') || file.status === 'uploading';
    });
    currFileList = uniqBy(currFileList, 'name');
    if (!!currFileList && currFileList.length !== 0) {
      setFileList(currFileList);
    }
  };

  const customRequest = ({file, onError, onSuccess}) => {
    const index = file.name.lastIndexOf('.');
    if (index === -1) {
      return onError('不能识别文件类型');
    }
    if (file.size > 1024 * 1024 * 6) {
      return onError('文件不能大于6MB');
    }

    const postfix = file.name.substr(index);
    const r = new FileReader();
    r.addEventListener('load', () => upload(r.result, onSuccess, onError));
    r.readAsDataURL(file);

    // OSSUploaderFile(file, `/openpbl/${obj.pid}/`, onSuccess, onError);
  };
  const upload = (file, onSuccess, onError) => {
    let data = new FormData()
    let fileName = new Date().getTime()
    data.append("fileName", fileName)
    data.append("filePath", `/openpbl/project/${obj.pid}/resource/${fileName}`)
    data.append("file", file)
    FileApi.uploadFile(data)
      .then(res=>{
        if (res.data.code === 200) {
          let url = res.data.data
          onSuccess({url: url}, file);
        } else {
          message.error(res.data.msg)
          onError(res.data.msg)
        }
      })
      .catch(e=>{
        onError(e)
      })
  }

  const uploadProps = {
    onChange: handleChange,
    listType: 'picture',
    fileList: fileList,
    customRequest: customRequest,
    multiple: true,
    showUploadList: true
  }

  return (
    <Card className="resource-card">
      <p className="card-title">文本内容</p>
      <Divider/>
      <div style={{textAlign: 'left'}}>
        <LzEditor
          readOnly={true}
          active={true}
          importContent={obj.section.resource.content}
          cbReceiver={receiveHtml}
          uploadProps={uploadProps}
        />
      </div>
      <div style={{marginTop: '10px'}}>
        <span style={{float: 'right'}}>
          <Button type="primary" onClick={saveContent}>保存</Button>
        </span>
      </div>
    </Card>
  )
}

export default RichWords
