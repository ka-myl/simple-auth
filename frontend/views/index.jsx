import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

import Login from './Login';
import Register from './Register';
import Dashboard from './Dashboard';

const AppViews = () => (
  <Router>
    <Switch>
      <Route path="/login">
        <Login />
      </Route>
      <Route path="/register">
        <Register />
      </Route>
      <Route path="/">
        <Dashboard />
      </Route>
    </Switch>
  </Router>
);

export default AppViews;
