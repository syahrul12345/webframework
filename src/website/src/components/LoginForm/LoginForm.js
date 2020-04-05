import React, { useState } from 'react'
import { Grid,TextField, Button, Typography, Link } from '@material-ui/core'
import { useHistory } from 'react-router-dom'
import { getLoginUrl } from '../../utils'

import axios from 'axios'

// Redux stuff
import { LoginAction } from '../../redux-modules/user/actions'
import { connect } from 'react-redux';

// Cookie stuff
import { useCookies } from 'react-cookie';
// DialogBox
import DialogBox from '../Dialog'

function LoginForm(props) {
    const history = useHistory()
    const { redirect } = props
    const [setCookie] = useCookies(['cookie-name']);
    
    const [errorMessage,setErrorMessage] = useState('')
    const [openDialog,setOpenDialog] = useState(false)
    const handleDialogClose = () => {
        setOpenDialog(false);
    };
  
    const [user,setUser] = useState({
        email:'',
        password:'',
    })
    const handleChange = (input) => event => {
        setUser({
            ...user,[input]:event.target.value
        })
    }
    const login = () => {
        const url = getLoginUrl()
        axios.post(url,user)
            .then(res => {
                const account = res.data.account
                setCookie('x-token',`bearer ${account.Token}`)
                props.dispatch(LoginAction(account,account.Token))
                history.push(redirect)
            })
            .catch(err => {
                setErrorMessage(err.response.data.message)
                setOpenDialog(true)
            })
    }
    return(
        <>
            <Grid item xs={12} md={12}>
                <TextField 
                onChange={handleChange('email')}
                style={{marginBottom:'1vh',minWidth:'80vw'}} 
                variant="outlined" 
                label="Email" />
            </Grid>
            <Grid item xs={12} md={12}>
                <TextField 
                type="password"
                onChange={handleChange('password')}
                style={{marginBottom:'1vh',minWidth:'80vw'}} 
                variant="outlined"  
                label="Password" />
            </Grid>
            <Typography style={{marginBottom:'1vh'}} variant="subtitle2"> 
                <Link href="/forgetPassword"> Forget Password </Link> 
            </Typography>
            <Typography style={{marginBottom:'1vh'}} variant="subtitle2"> No Account? Register 
                <Link href="/create"> here </Link> 
            </Typography>
            <Grid item xs={12}>
                <Button variant="outlined" onClick={() => login()}> LOGIN </Button>
            </Grid>
            <DialogBox errorMessage={errorMessage} handler={handleDialogClose} openDialog={openDialog}/>
        </>
    )    
}

function mapStateToProps(state) {
    return {
      token: state.token
    };

}
export default connect(mapStateToProps)(LoginForm);