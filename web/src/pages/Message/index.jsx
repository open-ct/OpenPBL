import React from "react";
import DocumentTitle from "react-document-title";
import {Affix, Layout, Menu} from "antd";
import {Link, Redirect, Route, Switch} from "react-router-dom";

import AllMessage from "./AllMessage";
import UnreadMessage from "./UnreadMessage";
import ReadMessage from "./ReadMessage";
import Communication from "./Communication";

class Message extends React.Component {
  state = {
    menu: 'all-message'
  }

  render() {
    const {menu} = this.state
    return (
      <DocumentTitle title="Project">
        <div style={{maxWidth: '1200px', margin: 'auto', padding: '20px'}}>
          <Layout>
            <Affix offsetTop={0}>
              <Layout.Sider breakpoint="lg" collapsedWidth="0">
                <Menu theme="light" mode="inline" defaultSelectedKeys={['all-message']}>
                  <Menu.Item key="all-message">
                    <Link to="/message/all">
                      全部消息
                    </Link>
                  </Menu.Item>
                  <Menu.Item key="unread-message">
                    <Link to="/message/unread">
                      未读消息
                    </Link>
                  </Menu.Item>
                  <Menu.Item key="read-message">
                    <Link to="/message/read">
                      已读消息
                    </Link>
                  </Menu.Item>
                  <Menu.Item key="remind">
                    <Link to="/message/communication">
                      留言沟通
                    </Link>
                  </Menu.Item>
                </Menu>
              </Layout.Sider>
            </Affix>
            <Layout.Content style={{backgroundColor: 'white', marginLeft: '10px'}}>
              <Switch>
                <Route exact path="/message" render={() => (
                  <Redirect to="/message/all"/>
                )}/>
                <Route exact path="/message/all" component={AllMessage}/>
                <Route exact path="/message/unread" component={UnreadMessage}/>
                <Route exact path="/message/read" component={ReadMessage}/>
                <Route exact path="/message/communication" component={Communication}/>
              </Switch>
            </Layout.Content>
          </Layout>
        </div>
      </DocumentTitle>
    )
  }
}

export default Message