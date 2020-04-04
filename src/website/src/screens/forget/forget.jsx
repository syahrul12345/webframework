import React from 'react';

import { Grid } from '@material-ui/core';
import ForgetPasswordForm from '../../components/ForgetPasswordForm';
import Header from '../../components/Header';

const ForgetScreen = () => {
  return (
    <>
      <Header title="Forget Password"/>
      <Grid
      container
      spacing={0}
      direction="column"
      alignItems="center"
      justify="center"
      style={{ minHeight: '100vh' }}
      >
          <ForgetPasswordForm/>
      </Grid>  
    </>
  )
}
export default ForgetScreen;