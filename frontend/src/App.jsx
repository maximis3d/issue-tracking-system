import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import SignIn from "./components/loginForm/loginForm";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<SignIn />} />
      </Routes>
    </Router>
  );
}

export default App;
