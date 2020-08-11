import React from 'react'
import Signin from './components/Signin'
import Signout from './components/Signout'
import Signup from './components/Signup'
import Registration from './components/Registration'

const routes = [
  { name: 'Signin', path: '/signin', exact: true, main: () => <Signin /> },
  { name: 'Signout', path: '/signout', exact: true, main: () => <Signout /> },
  { name: 'Signup', path: '/signup', exact: true, main: () => <Signup /> },
  { name: 'Registration', path: '/registrations', exact: true, main: () => <Registration /> },
]

export default routes
