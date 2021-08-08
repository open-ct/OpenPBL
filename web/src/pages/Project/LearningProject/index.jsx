import React from 'react';

import ProjectList from '../component/ProjectList';

class LearningProject extends React.PureComponent {
  render() {
    return (
      <ProjectList mode="learning" style={{height: '100%'}}/>
    );
  }
}

export default LearningProject;
