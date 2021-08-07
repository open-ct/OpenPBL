import React from 'react';
import ProjectList from '../component/ProjectList';
import CarouselPBL from './component/CarouselPBL';

class PublicProject extends React.PureComponent {
  render() {
    return (
      <div style={{ minHeight: '100vh',margin: 'auto', maxWidth: '1200px' }}>
        <CarouselPBL/><br />
        <div >
          <ProjectList mode="public"/>
        </div>
      </div>
    );
  }
}

export default PublicProject;
