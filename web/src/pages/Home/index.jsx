import React from "react";
import CarouselPBL from "./component/CarouselPBL";
import DocumentTitle from 'react-document-title'
import LatestProject from "./component/LatestProject";

class Home extends React.PureComponent {
  render() {
    return (
      <DocumentTitle title="Home">
        <div
          style={{
            margin: 'auto',
            padding: '20px',
          }}
        >
          <CarouselPBL/><br/>
          <div>
            <LatestProject />
          </div>
        </div>
      </DocumentTitle>
    );
  }
}

export default Home
