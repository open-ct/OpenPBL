import React from 'react';
import DocumentTitle from 'react-document-title';

import GlobalHeader from '../component/GlobalHeader/GlobalHeader';
import PublicProject from "./PublicProject";

class Project extends React.PureComponent {
  state = {}

  componentDidMount() {
  }

  render() {
    return (
      <DocumentTitle title="Project">
        <div style={{ minHeight: '100vh', backgroundColor: '#f2f4f5' }}>
          <GlobalHeader current="project"/>
          <div style={{marginLeft: '10px', marginRight: '10px'}}>
            <PublicProject/>
          </div>
        </div>
      </DocumentTitle>
    );
  }
}


export default Project;
