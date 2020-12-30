import React from "react";
import { browserHistory } from "utils/history";
import { Router, Route, Switch, Redirect } from "react-router-dom";
import Post from "components/post";
import Home from "components/home";
import Notfound from "components/404";

const App = () => {
  return ( <Router history={browserHistory}>
    <Switch>
      <Route exact path="/posts/:id" component={Post} />
      <Route exact path="/404" component={Notfound} />
      <Route exact path="/" component={Home} />
      <Redirect to="/404" />
    </Switch>
  </Router>);
}

export default App;
