import React from 'react';

import ProjectList from '../component/ProjectList';

class FavouriteProject extends React.PureComponent {
  render() {
    return (
      <ProjectList mode="favourite"/>
    );
  }
}

export default FavouriteProject;
