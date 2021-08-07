import React from 'react';
import DocumentTitle from 'react-document-title';

import GlobalHeader from '../../component/GlobalHeader/GlobalHeader';
import MenuBar from "../component/MenuBar";
import localStorage from "localStorage";

class MyProject extends React.PureComponent {
  state = {
    type: localStorage.getItem('type'),
  }
  render() {
    const {type} = this.state
    return (
      <DocumentTitle title="Project">
        <div style={{minHeight: '100vh'}}>
          <GlobalHeader current="my-project"/>
          <div>
            {type === 'teacher' ?
              <MenuBar menu="editing"/>
              :
              <MenuBar menu="learning"/>
            }
          </div>
        </div>
      </DocumentTitle>
    );
  }
}


export default MyProject;
