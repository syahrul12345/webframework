import React, { useState } from 'react'
import { Grid,TextField, Button } from '@material-ui/core'
import { getCreateAccountUrl } from '../../utils'
import { useHistory } from 'react-router-dom'
import axios from 'axios'

function CreateAccountForm(props) {
    const history = useHistory()
    const { redirect } = props
    const [userInfo,setUserInfo] = useState({
        email:'',
        username:'',
        password:'',
    });
    const handleChange = (input) => event =>{
        setUserInfo({...userInfo,[input]:event.target.value})
    }
    const createAccount = () => {
        const url = getCreateAccountUrl()
        axios.post(url,userInfo)
            .then((res) => {
                // Set the returned cookie from the backend
                const token = res.data.account.Token
                props.cookieHandler(token)
                history.push(redirect)
            })
            .catch((err) => {
                console.log(err.response.data.message)
            })
    }
    return(
        <Grid
        container
        direction="column"
        justify="center"
        alignItems="center"
        alignContent="center"
        style={{minHeight:'60vh'}}
        >
            <Grid item xs={12}>
                <TextField
                    label="Email"
                    variant="outlined"
                    style={{width:'50vw',marginBlockEnd:'1vh'}}
                    onChange={handleChange('email')}/>
            </Grid>
            <Grid item xs={12}>
                <TextField
                    label="Username"
                    variant="outlined"
                    style={{width:'50vw',marginBlockEnd:'1vh'}}
                    onChange={handleChange('username')}/>
            </Grid>
            <Grid item xs={12}>
                <TextField
                    label="Password"
                    variant="outlined"
                    style={{width:'50vw',marginBlockEnd:'1vh'}}
                    onChange={handleChange('password')}/>
            </Grid>
            <Grid item xs={12}>
                <Button variant="outlined" onClick={createAccount}> Create Account </Button>
            </Grid>
        </Grid>
    )
}
export default CreateAccountForm;