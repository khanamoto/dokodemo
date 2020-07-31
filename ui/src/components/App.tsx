import React from 'react'
import { Switch, BrowserRouter as Router, Route, Link } from 'react-router-dom'
import routes from '../routes'
// import { Link } from 'react-router-dom'

const App: React.FC = () => {
  return (
    <Router>
      <ul>
        <li>
          <Link to="/">ホーム</Link>
        </li>
        <li>
          <Link to="/signup">ユーザー 新規登録</Link>
        </li>
        <li>
          <Link to="/group">グループ 新規登録</Link>
        </li>
      </ul>

      <Switch>
        {routes.map(route => (
          <Route key={route.path} path={route.path} exact={route.exact} component={route.main} />
        ))}
      </Switch>
    </Router>
  )
}

export default App
