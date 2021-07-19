import React from 'react';
import ProjectList from '../component/ProjectList';
import CarouselPBL from './component/CarouselPBL';

class PublicProject extends React.PureComponent {
  render() {
    return (
      <div style={{ minHeight: '100vh' }}>
        <CarouselPBL/><br />
        <ProjectList mode="public"/>
      </div>
    );
  }
}

export default PublicProject;
