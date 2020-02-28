import React from 'react'
import {useState,useEffect} from 'react'
import { Grid, TextField, Button, Typography } from '@material-ui/core'
import {Redirect} from 'react-router-dom'
import axios from 'axios'
export default function Login(props) {
    const [redirect,setRedirect] = useState({
        status:false,
        token:''
    })
    const [toDashboard,changeRedirect] = useState(false)
    const [userInfo,setUserInfo] = useState({
        email:'',
        username:'',
        password:'',
        token:'',
    });
    const handleChange = (input) => event =>{
        setUserInfo({...userInfo,[input]:event.target.value})
    }
    const login = () => {
        axios.post("/api/v1/login",userInfo)
            .then((res) => {
                // Succesfull login
                const token = res.data.account["Token"]
                setRedirect({...redirect,"token":token})
                changeRedirect(true)
            })
            .catch((err) => {
                console.log(err)
                console.log(err.response.data.message)
            })
    }

    useEffect(() => {
        // Handle redirect
        props.setState({...props.globalState,"token":redirect.token})
    },[toDashboard])
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
                    label="Password"
                    variant="outlined"
                    style={{width:'50vw',marginBlockEnd:'1vh'}}
                    onChange={handleChange('password')}/>
            </Grid>
            <Grid item xs={12}>
                <Button variant="filled" onClick={login}> LOGIN </Button>
            </Grid>
            <Grid item xs={12}>
                <a href="/changePassword">
                    <Typography variant="subtitle1"> Change your password </Typography>
                </a>
            </Grid>  
            {toDashboard ? 
                <Redirect to="/dashboard"/>
                : <></>
            }
        </Grid>
    )
}