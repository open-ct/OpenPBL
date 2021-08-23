import qs from 'qs'
import {getCasdoorService} from "../pages/User/Auth/Auth";

const FileApi = {
  uploadFile(owner, tag, parent, fullFilePath, file, provider="") {
    let request = getCasdoorService()
    let formData = new FormData();
    formData.append("file", file);
    return request({
      url: `/api/upload-resource?owner=${owner}&tag=${tag}&parent=${parent}&fullFilePath=${encodeURIComponent(fullFilePath)}&provider=${provider}`,
      method: 'post',
      data: formData
    })
  },
  deleteFile(data, provider="") {
    let request = getCasdoorService()
    return request({
      url: `/api/delete-resource?provider=${provider}`,
      method: 'post',
      data: data
    })
  }
}

export default FileApi