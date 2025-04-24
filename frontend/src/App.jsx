import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import SignIn from "./components/loginForm/loginForm";
import Register from "./components/registerForm/registerForm";
import CreateProject from "./components/projectCreationForm/projectCreation";
import CreateIssue from "./components/issueCreationForm/issueCreation";

import LandingPage from "./pages/LandingPage";
import Projects from "./pages/Projects";

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
      </Routes>
    </Router>
  );
}

export default App;
