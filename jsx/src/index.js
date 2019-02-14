// Import React and ReactDOM Libs
import React from 'react';
import ReactDOM from 'react-dom';

// Create a react component
const App = () => {
  const button = {text: 'Click me!'};
  const labelText = 'Example:';

  return (
    <div>
      <label className="label" htmlFor="name">
        {labelText}
      </label>
      <input id="name" type="text" />
      <button style={{backgroundColor: 'blue', color: 'white'}}>
        {button.text}
      </button>
    </div>
  );
};

// Take the react component and show it on the screen
ReactDOM.render(<App />, document.querySelector('#root'));
