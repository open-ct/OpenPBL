import React from 'react';
import ProjectList from '../component/ProjectList';

class PublicProject extends React.PureComponent {
  render() {
    return (
      <div
        style={{maxWidth: '1200px', margin: 'auto'}}>
        <ProjectList mode="public"/>
      </div>
    );
  }
}

export default PublicProject;
