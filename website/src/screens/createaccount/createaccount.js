import React, { useState,useEffect } from 'react';
import { Grid } from "@material-ui/core";
import { useCookies } from 'react-cookie';
import CreateAccountForm from '../../components/CreateAccountForm';

export default function CreateAccountPage(props) {
    const { redirect } = props
    const [myCookie,setMyCookie] = useState('')
    const [cookies, setCookie, removeCookie] = useCookies(['cookie-name']);
    useEffect(()=> {
      // Prevent routing on the first load
      // The cookie will be set from the login form.
      if (myCookie !== '') {
        setCookie('x-token',`bearer ${myCookie}`)
      }
    },[cookies,myCookie])

    return(
      <Grid
      container
      spacing={0}
      direction="column"
      alignItems="center"
      justify="center"
      style={{ minHeight: '100vh' }}
      >
          <CreateAccountForm redirect= {redirect} cookieHandler={setMyCookie}/>
      </Grid>   
    )
}