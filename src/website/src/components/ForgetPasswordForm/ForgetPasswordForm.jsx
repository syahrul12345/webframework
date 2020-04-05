import React, {useState} from 'react'
import axios from 'axios';

import { getResetPasswordUrl } from '../../utils'
import { Grid,TextField, Button,Typography } from '@material-ui/core'

const ForgetPasswordForm = () => {
  const [email,setEmail] = useState()
  const [resetMessage,setResetMessage] = useState('')

  const handleChange = () => event => {
    setEmail(event.target.value)
  }

  const reset = () => {
    const url = getResetPasswordUrl()
    axios.post(url,{
      email,
    }).then(res => {
      setResetMessage(res.data.message)
    }).catch(err => {
      setResetMessage(err.response.data.message)
    })
  }
  return (
    <>
      <Grid item xs={12} md={12}>
          <TextField 
          onChange={handleChange('email')}
          style={{marginBottom:'1vh',minWidth:'80vw'}} 
          variant="outlined" 
          label="Email" />
      </Grid>
      {resetMessage !== "" ?
        <Typography variant="subtitle2" style={{textAlign:'center',color:'red',paddingBottom:'1vh'}}> {resetMessage} </Typography>
      :
        <></>
      }
      <Grid item xs={12}>
          <Button variant="outlined" onClick={() => reset()}> Reset Password </Button>
      </Grid>
  </>
  )
}

export default ForgetPasswordForm;