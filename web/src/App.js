import {Route, BrowserRouter, Switch, Redirect} from "react-router-dom";
import './App.less';
import AuthCallback from "./pages/User/Auth/AuthCallback";
import HeaderLayout from "./pages/component/Layout/HeaderLayout";
import StudentEvidenceContent from "./pages/Project/Evidence/component/StudentEvidenceContent";

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Route exact path="/" render={() => (
          <Redirect to="/home"/>
        )}/>
        <Route path="/home" component={HeaderLayout} />
        <Route exact path="/export/project/:projectId/student/:studentId/evidence" component={StudentEvidenceContent} />
        <Route exact path="/callback" component={AuthCallback} />
      </BrowserRouter>
    </div>
  );
}

export default App;
