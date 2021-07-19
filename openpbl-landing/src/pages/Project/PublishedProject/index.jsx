import React from 'react';

import ProjectList from '../component/ProjectList';

class PublishedProject extends React.PureComponent {
  render() {
    return (
      <ProjectList mode="published"/>
    );
  }
}

export default PublishedProject;
