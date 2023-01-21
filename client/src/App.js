import React from 'react'
import {container} from 'semantic-ui-react'
import "./App.css"
import ToDOList from "./To-DO-List"

function App(){
  return(
    <div>
      <container>
        <ToDOList/>
      </container>
    </div>
  );
}

export default App;
