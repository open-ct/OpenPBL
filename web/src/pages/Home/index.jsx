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

import React from "react";
import CarouselPBL from "./component/CarouselPBL";
import DocumentTitle from 'react-document-title'
import LatestProject from "./component/LatestProject";
import {Divider} from "antd";

class Home extends React.PureComponent {
  render() {
    return (
      <DocumentTitle title="Home">
        <div
          style={{
            margin: 'auto',
            padding: '20px',
          }}
        >
          <CarouselPBL/><br/>
          <Divider orientation="center">最近的项目</Divider>
          <div>
            <LatestProject />
          </div>
        </div>
      </DocumentTitle>
    );
  }
}

export default Home
