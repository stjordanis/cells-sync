import React from 'react'
import { Route } from 'react-router-dom'
import Servers from '../components/Servers'
import External from './External'
import CallbackPage from './Callback'

export default function Routes(props) {
    const {match, socket} = props;
  return (
    <React.Fragment>
      <Route path={match.path} exact={true} render={(p)=><Servers {...p} socket={socket}/>}/>
      <Route path={`${match.path}/callback`}  render={(p)=><CallbackPage {...p} socket={socket}/>} />
      <Route path={`${match.path}/external`}  render={(p)=><External {...p} socket={socket}/>} />
    </React.Fragment>
  );
}