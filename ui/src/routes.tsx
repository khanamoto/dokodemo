import React from 'react'
import Signin from './components/Signin'
import Signout from './components/Signout'
import Signup from './components/Signup'
import Organization from './components/Organization'

const routes = [
  { name: 'Signin', path: '/signin', exact: true, main: () => <Signin /> },
  { name: 'Signout', path: '/signout', exact: true, main: () => <Signout /> },
  { name: 'Signup', path: '/signup', exact: true, main: () => <Signup /> },
  { name: 'Organization', path: '/organizations', exact: true, main: () => <Organization /> },
]

export default routes
