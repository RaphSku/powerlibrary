import React from 'react';

import './App.css';

import Header from "./components/Header";
import View   from "./components/View";
import Form   from "./components/Form";

function App() {
  return (
    <div className="App">
      <Header />
      <main>
        <View />
        <Form />
      </main>
    </div>
  );
}

export default App;
