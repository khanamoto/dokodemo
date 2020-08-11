import React, { useState, useContext, FormEvent } from 'react'
import axios from 'axios'
import Button from '@material-ui/core/Button'
import TextField from '@material-ui/core/TextField'
import Typography from '@material-ui/core/Typography'

const Signin = () => {
  const [userName, setUserName] = useState('')
  const [password, setPassword] = useState('')
  const [error, setErrors] = useState('')

  const handleForm = (e: FormEvent) => {
    e.preventDefault()

    const params = new URLSearchParams()
    params.append('userName', userName)
    params.append('password', password)

    // axios.defaults.withCredentials = true
    axios
      .post('http://localhost:8000/signin', params, { withCredentials: true })
      .then(res => {
        console.log(res)
      })
      .catch(err => {
        setErrors(err.message)
      })
  }

  const styles = {
    field: {
      marginBottom: 20,
    },
  }

  return (
    <div>
      <Typography variant="subtitle1" component="h2">
        ログイン
      </Typography>
      <form onSubmit={e => handleForm(e)}>
        <div style={{ display: 'flex', flexDirection: 'column' }}>
          <TextField
            label="userName"
            value={userName}
            onChange={e => setUserName(e.target.value)}
            name="userName"
            type="userName"
            required
            style={styles.field}
          />
          <TextField
            label="password"
            value={password}
            onChange={e => setPassword(e.target.value)}
            name="password"
            type="password"
            required
            style={styles.field}
          />

          <div>
            <Button type="submit" variant="contained" color="primary">
              ログインする
            </Button>
          </div>

          <div>{error}</div>
        </div>
      </form>
    </div>
  )
}

export default Signin
