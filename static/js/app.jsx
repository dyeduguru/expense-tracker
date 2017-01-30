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
        <h1>We R VR</h1>
        <p>Provide valuable feedback to VR experience developers.</p>
        <a className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>
      </div>
    </div>);
  }
});