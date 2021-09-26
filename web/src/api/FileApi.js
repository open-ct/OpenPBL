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