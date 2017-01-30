var App = React.createClass({
  componentWillMount: function() {
    this.setState({idToken: null})
  },
  render: function() {
    if (this.state.idToken) {
      return (<LoggedIn idToken={this.state.idToken} />);
    } else {
      return (<Home />);
    }
  }
});

var Home = React.createClass({
  render: function() {
    return (
    <div className="container">
      <div className="col-xs-12 jumbotron text-center">
        <h1>Expense Tracker</h1>
        <p>Save all your expenses securely online!</p>
        <a className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>
      </div>
    </div>);
  }
});


ReactDOM.render(
  <App />,
  document.getElementById('app')
);