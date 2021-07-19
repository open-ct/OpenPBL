import React from 'react';
import DocumentTitle from 'react-document-title';

import GlobalHeader from '../component/GlobalHeader/GlobalHeader';
import MenuBar from './component/MenuBar';

class Project extends React.PureComponent {
  state = {}

  componentDidMount() {
  }

  render() {
    return (
      <DocumentTitle title="Project">
        <div style={{ minHeight: '100vh' }}>
          <GlobalHeader current="project"/>
          <MenuBar menu="public" />
        </div>
      </DocumentTitle>
    );
  }
}


export default Project;
