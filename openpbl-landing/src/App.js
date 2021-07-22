import { Route, BrowserRouter, Switch } from "react-router-dom";

import './App.less';
import Home from './pages/Home/index'
import Project from "./pages/Project";
import ProjectInfo from "./pages/Project/ProjectInfo/index";
import PublicProject from "./pages/Project/PublicProject";
import LearningProject from "./pages/Project/LearningProject";
import FinishedProject from "./pages/Project/FinishedProject";

import AuthCallback from "./pages/User/Auth/AuthCallback";
import Login from "./pages/User/Login";
import Register from "./pages/User/Register";
import Learning from "./pages/Project/LearningPage";
import EditInfo from "./pages/Project/CreateProject/EditInfo";
import EditOutlined from "./pages/Project/CreateProject/EditOutline"
import SectionEditPage from "./pages/Project/CreateProject/component/SectionEditPage";

function App() {
  return (
      <div className="App">
        <BrowserRouter>
          <Switch>
            <Route exact path="/" component={Home} />
            <Route exact path="/callback" component={AuthCallback} />

            <Route exact path="/landing" component={Home} />
            <Route exact path="/project" component={Project} />

            <Route exact path="/project/public" component={PublicProject} />
            <Route exact path="/project/learning" component={LearningProject} />
            <Route exact path="/project/finished" component={FinishedProject} />
            <Route exact path="/project/info/:id" component={ProjectInfo} />

            <Route exact path="/project/edit/info/:id" component={EditInfo} />
            <Route exact path="/project/edit/outline/:id" component={EditOutlined} />
            <Route exact path="/project/edit/section/:pid/:sid" component={SectionEditPage} />

            <Route exact path="/project/learning/:pid/:cid/:sid" component={Learning} />

            <Route exact path="/user/login" component={Login} />
            <Route exact path="/user/register" component={Register} />

          </Switch>
        </BrowserRouter>
      </div>
  );
}

export default App;
