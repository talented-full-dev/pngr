import React, { Component } from 'react'
import { Link } from 'react-router-dom'
import { Segment, Message, Header, Button } from 'semantic-ui-react'
import API from '../../../Api'

export default class Posts extends Component {
  state = {
    loading: true,
    posts: [],
    error: ''
  }
  
  componentDidMount() {
    API.getPosts()
      .then(posts => {
        this.setState({posts: posts});
      })
      .catch(error => {
        this.setState({error: error});
      })
      .finally(() => {
        this.setState({loading: false})
      });
  }

  render() {
    const { loading, posts, error } = this.state;
    return (
      <div>
        <Header attached='top' as='h3'>My Posts</Header>
        {error ? <Message attached danger>{error}</Message> : ''}
        {posts.length === 0 && !loading ? <Message attached warning>No posts to show</Message> : ''}
        {posts.map(({id, title, body}, i) => (
          <Segment key={id} attached>
            <h4>{title}</h4>
            <p>{body}</p>
          </Segment>
        ))}
        
        <Segment attached='bottom'>
          <Link to='/posts/create'>
            <Button primary>New Post</Button>
          </Link>
        </Segment>
      </div>
    );
  }
}