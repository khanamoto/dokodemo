import React, { useState, useContext, FormEvent } from 'react'
import axios from 'axios'
import Button from '@material-ui/core/Button'
import TextField from '@material-ui/core/TextField'
import Typography from '@material-ui/core/Typography'

const Signup = () => {
  const [name, setName] = useState('')
  const [userName, setUserName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setErrors] = useState('')

  const handleForm = (e: FormEvent) => {
    e.preventDefault()

    // const params = {
    //   name: name,
    //   userName: userName,
    //   email: email,
    //   password: password,
    // }

    const params = new URLSearchParams()
    params.append('name', name)
    params.append('userName', userName)
    params.append('email', email)
    params.append('password', password)

    axios
      .post('http://localhost:8000/signup', params)
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
        ユーザー登録
      </Typography>
      <form onSubmit={e => handleForm(e)}>
        <div style={{ display: 'flex', flexDirection: 'column' }}>
          <TextField
            label="name"
            value={name}
            onChange={e => setName(e.target.value)}
            name="name"
            type="name"
            required
            style={styles.field}
          />
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
            label="email"
            value={email}
            onChange={e => setEmail(e.target.value)}
            name="email"
            type="email"
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
              登録する
            </Button>
          </div>

          <div>{error}</div>
        </div>
      </form>
    </div>
  )
}

export default Signup
