import React from 'react';

import TodoList from './components/TodoList'

class App extends React.Component {
  render() {
    return (
      <div className="App">
        <p>Todo!</p>
        <TodoList/>
      </div>
    );  
  }
}

export default App;
