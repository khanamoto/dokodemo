import React from 'react'
import Signup from './components/Signup'
import Group from './components/Group'

const routes = [
  { name: 'Signup', path: '/signup', exact: true, main: () => <Signup /> },
  { name: 'Group', path: '/group', exact: true, main: () => <Group /> },
]

export default routes
