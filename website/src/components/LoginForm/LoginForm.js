import React, { useState } from 'react'
import { Grid,TextField, Button } from '@material-ui/core'
import axios from 'axios'
import { useHistory } from 'react-router-dom'
function LoginForm(props) {
    const history = useHistory()
    const { redirect } = props
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
        let url = ""
        if (process.env.NODE_ENV == "production") {
            url = '/api/v1/login'
        }else{
            url = 'http://localhost:8004/api/v1/login'
        }
        axios.post(url,user)
            .then(res => {
                const token = res.data.account.Token
                props.cookieHandler(token)
                history.push(redirect)
            })
            .catch(err => {
                console.log(err)
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
                onChange={handleChange('password')}
                style={{marginBottom:'1vh',minWidth:'80vw'}} 
                variant="outlined"  
                label="Password" />
            </Grid>
            <Grid item xs={12}>
                <Button variant="outlined" onClick={() => login()}> LOGIN </Button>
            </Grid>
        </>
    )    
}
export default LoginForm;