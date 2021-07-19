import React from 'react';

import ProjectList from '../component/ProjectList';

class EditingProject extends React.PureComponent {
  render() {
    return (
      <ProjectList mode="editing"/>
    );
  }
}

export default EditingProject;
