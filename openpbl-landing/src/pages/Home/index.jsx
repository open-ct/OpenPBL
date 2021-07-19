import React from 'react';
import DocumentTitle from 'react-document-title';
import {enquireScreen} from 'enquire-js';

import GlobalHeader from '../component/GlobalHeader/GlobalHeader';
import Banner from './Banner';
import Page1 from './Page1';
import Page2 from './Page2';
import Footer from './Footer';
import './static/style';
import TestPage from "./TestPage";
import Teams2 from "./Teams2";
import {Teams20DataSource} from './data.source';


let isMobile;

enquireScreen((b) => {
  isMobile = b;
});

class Home extends React.PureComponent {
  state = {
    isMobile,
  }

  componentDidMount() {
    enquireScreen((b) => {
      this.setState({
        isMobile: !!b,
      });
    });
  }

  render() {
    return (
      <DocumentTitle title="OpenCT">
        <div style={{textAlign: 'left'}}>
          <GlobalHeader current="landing"/>
          <div className="home-wrapper">
            <Banner isMobile={this.state.isMobile}/>
            <TestPage isMobile={this.state.isMobile}/>
            <Page1 isMobile={this.state.isMobile}/>
            <Page2/>
            <Teams2 dataSource={Teams20DataSource} isMobile={this.state.isMobile}/>
          </div>
          <Footer/>
        </div>
      </DocumentTitle>
    );
  }
}

export default Home;