import React from 'react';
import { Grid } from "@material-ui/core";
import LoginForm from '../../components/LoginForm';

const LoginPage = (props) => {
    const { redirect } = props

    return(
      <Grid
      container
      spacing={0}
      direction="column"
      alignItems="center"
      justify="center"
      style={{ minHeight: '100vh' }}
      >
          <LoginForm redirect= {redirect}/>
      </Grid>   
    )
}

export default LoginPage