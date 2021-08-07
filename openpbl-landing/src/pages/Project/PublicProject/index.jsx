import React from 'react';
import ProjectList from '../component/ProjectList';
import CarouselPBL from './component/CarouselPBL';

class PublicProject extends React.PureComponent {
  render() {
    return (
      <div style={{maxWidth: '1200px', margin: 'auto'}}>
        <CarouselPBL/><br />
        <div >
          <ProjectList mode="public"/>
        </div>
      </div>
    );
  }
}

export default PublicProject;
