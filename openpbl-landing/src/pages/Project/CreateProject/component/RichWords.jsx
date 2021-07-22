import React, {useEffect, useState} from "react";
import {Button, Card, message} from "antd";
import LzEditor from "react-lz-editor"
import uniqBy from 'lodash/uniqBy';
import findIndex from 'lodash/findIndex';

import OSSUploaderFile from "./oss";
import ResourceApi from "../../../../api/ResourceApi";
import "../section-edit.less"

function RichWords(obj) {
  const sid = obj.section.id
  const [fileList, setFileList] = useState([])
  const [content, setContent] = useState(obj.content)

  useEffect(() => {

  }, [])
  const receiveHtml = (content) => {
    console.log("recieved markdown content", content);
    setContent(content)
  }

  const saveContent = e => {
    ResourceApi.updateResourceContent(obj.section.resource.id, content)
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

  const customRequest = ({
                           file,
                           onError,
                           onSuccess,
                         }) => {
    OSSUploaderFile(file, '/openct/' + sid + '/content/', onSuccess, onError);
  };

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
