import React from "react";
import { Page } from "./Componentes/Page";
import {BrowserRouter as Router} from "react-router-dom";
import {Header} from './Componentes/header/index';
function App() {
  return (
    <div className="App">
      <Router>
      <Header/>
        <Page/>
      </Router>
    </div>
  );
}

export default App;
