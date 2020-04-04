import React, {useState} from 'react'
import axios from 'axios';

import { getResetPasswordUrl } from '../../utils'
import { Grid,TextField, Button,Typography,Link } from '@material-ui/core'
import DialogBox from '../Dialog'
const ForgetPasswordForm = () => {
  const [email,setEmail] = useState()
  const [errorMessage,setErrorMessage] = useState('')
  const [openDialog,setOpenDialog] = useState(false)

  const handleChange = () => event => {
    setEmail(event.target.value)
  }
  const handleDialogClose = () => {
    setOpenDialog(false);
};

  const reset = () => {
    const url = getResetPasswordUrl()
    axios.post(url,{
      email,
    }).then(res => {
      setErrorMessage(res.data.message)
      setOpenDialog(true)
    }).catch(err => {
      setErrorMessage(err.toString())
      setOpenDialog(true)
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
      <Grid item xs={12}>
          <Button variant="outlined" onClick={() => reset()}> Reset Password </Button>
      </Grid>
      <DialogBox errorMessage={errorMessage} handler={handleDialogClose} openDialog={openDialog}/>
  </>
  )
}

export default ForgetPasswordForm;