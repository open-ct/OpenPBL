import { Route, BrowserRouter, Switch } from "react-router-dom";
import './App.less';
import AuthCallback from "./pages/User/Auth/AuthCallback";
import HeaderLayout from "./pages/component/Layout/HeaderLayout";

function App() {
  return (
    <div className="App">
      <BrowserRouter>
          <Route path="/" component={HeaderLayout} />
          <Route exact path="/callback" component={AuthCallback} />
      </BrowserRouter>
    </div>
  );
}

export default App;
