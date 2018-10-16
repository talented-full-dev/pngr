import React, { Component } from 'react'
import { Menu, Container } from 'semantic-ui-react'
import { Subscribe } from 'unstated'
import { NavLink } from 'react-router-dom'

import UserContainer from '../Containers/User'
import API from '../Api'

// helper for semanticUI + react-router
const Link = props => (
  <NavLink
    exact
    {...props}
    activeClassName="active"
  />
);

class Nav extends Component {

  loggedOutMenu = (
    <Menu.Menu position="right">
      <Menu.Item as={Link} to="/login" name="Log In" />
      <Menu.Item as={Link} to="/signup" name="Sign Up" />
    </Menu.Menu>
  );

  loggedInMenu = (
    <Menu.Menu position="right">
      <Menu.Item as={Link} to="/dashboard" name="Dashboard" />
      <Menu.Item link={true} onClick={API.logout} content="Log Out"/>
    </Menu.Menu>
  );

  render() {
    return (
      <Menu fixed="top" inverted>
        <Container>
          <Menu.Item as={Link} to="/" name="Home" />
          <Subscribe to={[UserContainer]}>
            {userContainer => {
              return userContainer.state.user.id === 0 ? this.loggedOutMenu : this.loggedInMenu;
            }}
          </Subscribe>
        </Container>
      </Menu>
    );
  }
}

export default Nav;