import React, { useState, useContext, FormEvent } from 'react'
import axios from 'axios'
import Button from '@material-ui/core/Button'
import TextField from '@material-ui/core/TextField'
import Typography from '@material-ui/core/Typography'

const Registration = () => {
  const [organizationName, setOrganizationName] = useState('')
  const [url, setUrl] = useState('')
  const [fullName, setFullName] = useState('')
  const [userName, setUserName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setErrors] = useState('')

  const handleForm = (e: FormEvent) => {
    e.preventDefault()

    const params = new URLSearchParams()
    params.append('organizationName', organizationName)
    params.append('url', url)
    params.append('fullName', fullName)
    params.append('userName', userName)
    params.append('email', email)
    params.append('password', password)

    // axios.defaults.withCredentials = true
    axios
      .post('http://localhost:8000/registrations', params, { withCredentials: true })
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
        新規登録して始める
      </Typography>
      <form onSubmit={e => handleForm(e)}>
        <div style={{ display: 'flex', flexDirection: 'column' }}>
          <TextField
            label="会社・団体名"
            value={organizationName}
            onChange={e => setOrganizationName(e.target.value)}
            name="organizationName"
            type="organizationName"
            required
            style={styles.field}
          />
          <TextField
            label="URL"
            value={url}
            onChange={e => setUrl(e.target.value)}
            name="url"
            type="url"
            required
            style={styles.field}
          />
          <TextField
            label="お名前"
            value={fullName}
            onChange={e => setFullName(e.target.value)}
            name="fullName"
            type="fullName"
            required
            style={styles.field}
          />
          <TextField
            label="アカウント名"
            value={userName}
            onChange={e => setUserName(e.target.value)}
            name="userName"
            type="userName"
            required
            style={styles.field}
          />
          <TextField
            label="メールアドレス"
            value={email}
            onChange={e => setEmail(e.target.value)}
            name="email"
            type="email"
            required
            style={styles.field}
          />
          <TextField
            label="パスワード"
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

export default Registration
