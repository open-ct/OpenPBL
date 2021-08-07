import React from 'react';

import ProjectList from '../component/ProjectList';

class FinishedProject extends React.PureComponent {
  render() {
    return (
      <ProjectList mode="finished"/>
    );
  }
}

export default FinishedProject;
