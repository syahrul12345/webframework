import React, { useState,useEffect } from 'react'
import { Grid,TextField, Button } from '@material-ui/core'
import axios from 'axios'
import { useHistory } from 'react-router-dom'

// Redux stuff
import { CreateAccountAction } from '../../redux-modules/user/actions'
import { connect } from 'react-redux';
// Cookie stuff
import { useCookies } from 'react-cookie';
// Helper functions
import { getCreateAccountUrl } from '../../utils'

function CreateAccountForm(props) {
    const history = useHistory()
    const [cookies, setCookie, removeCookie] = useCookies(['cookie-name']);
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
                // console.log(res.data.account)
                const account = res.data.account
                setCookie('x-token',`bearer ${account.Token}`)
                props.dispatch(CreateAccountAction(account,account.Token))
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
function mapStateToProps(state) {
    return {
      account: state.account,
      token: state.token,
    };

}
export default connect(mapStateToProps)(CreateAccountForm);