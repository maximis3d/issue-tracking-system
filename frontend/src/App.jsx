import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import SignIn from "./components/loginForm/loginForm";
import Register from "./components/registerForm/registerForm";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<SignIn />} />
        <Route path="/register" element={<Register />} />
      </Routes>
    </Router>
  );
}

export default App;
