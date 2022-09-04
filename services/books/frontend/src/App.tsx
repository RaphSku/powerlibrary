import React from 'react';

import './App.css';

import Header    from "./components/Header";
import View      from "./components/View";
import Form      from "./components/Form";
import FormShelf from "./components/FormShelf";

function App() {
  return (
    <div className="App">
      <Header />
      <main>
        <View />
        <FormShelf />
        <Form />
      </main>
    </div>
  );
}

export default App;
