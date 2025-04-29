import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import SignIn from "./components/loginForm/loginForm";
import Register from "./components/registerForm/registerForm";
import CreateProject from "./components/projectCreationForm/projectCreation";
import CreateIssue from "./components/issueCreationForm/issueCreation";

import LandingPage from "./pages/LandingPage";
import Projects from "./pages/Projects";
import StandUp from "./pages/StandUp";
import ProjectDetails from "./pages/Project";
import Issue from "./pages/Issue";
import EditIssue from "./pages/EditIssue";
import ProjectAssignment from "./pages/ProjectAssignment";
import Scope from "./pages/Scope";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<LandingPage />} />
        <Route path="/projects" element={<Projects />} />
        <Route path="/login" element={<SignIn />} />
        <Route path="/register" element={<Register />} />
        <Route path="/create-project" element={<CreateProject />} />
        <Route path="/create-issue" element={<CreateIssue />} />
        <Route path="/create-issue" element={<CreateIssue />} />
        <Route path="/projects/:key" element={<ProjectDetails />} /> 
        <Route path="/standup/:key" element={<StandUp />} /> 
        <Route path="/issues/:id" element={<Issue />} /> 
        <Route path="/issues/:id/edit" element={<EditIssue />} />
        <Route path="/assign-user" element={<ProjectAssignment />} />
        <Route path="/scope/:id" element={<Scope />} />        
      </Routes>
    </Router>
  );
}

export default App;
