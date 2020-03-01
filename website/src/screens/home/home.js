import React from 'react'
import Create from '../createaccount'
import Login from '../login'
import Dashboard from '../dashboard'
import { BrowserRouter as Router,
        Switch ,
        Route} from 'react-router-dom'
function App() {
  // Set the global state

  return (
    <Router>
      <Switch>
        <Route exact path="/create">
          <Create redirect="/dashboard"/>
        </Route>
        <Route exact path="/">
          <Login redirect="/dashboard"/>
        </Route>
        <Route exact path="/dashboard">
          <Dashboard/>
        </Route>
      </Switch>
    </Router>
  );
}

export default App;