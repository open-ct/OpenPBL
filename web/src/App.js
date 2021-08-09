import { Route, BrowserRouter, Switch } from "react-router-dom";
import {Provider} from "react-redux";

import store from './store'
import './App.less';
import AuthCallback from "./pages/User/Auth/AuthCallback";
import HeaderLayout from "./pages/component/Layout/HeaderLayout";

function App() {
  return (
    <Provider store={store}>
      <div className="App">
        <BrowserRouter>
          <Route path="/" component={HeaderLayout} />
          <Route exact path="/callback" component={AuthCallback} />
        </BrowserRouter>
      </div>
    </Provider>
  );
}

export default App;
