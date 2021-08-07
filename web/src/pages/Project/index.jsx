import React from 'react';
import DocumentTitle from 'react-document-title'

import PublicProject from "./PublicProject";

class Project extends React.PureComponent {
  render() {
    return (
      <DocumentTitle title="Public Project">
        <div
          style={{
            backgroundColor: '#f2f4f5',
            padding: '20px'
          }}
        >
          <PublicProject/>
        </div>
      </DocumentTitle>
    );
  }
}


export default Project;
