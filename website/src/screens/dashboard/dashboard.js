import React from 'react'
import { connect } from 'react-redux';
// Cookie stuff
import { useCookies } from 'react-cookie';
import jwt from 'jwt-decode'

const Dashboard = (props) => {
    // Pull account from the store
    const { account } = props
    const [cookies, setCookie, removeCookie] = useCookies(['cookie-name']);
    if (account.id == undefined) {
        // send this token to the backend to verify
        const tokenHeader = cookies["x-token"].split(" ")
        const token = tokenHeader[1]
        // send the token to the backend
        
    }
    return(
        <div> THIS IS THE DASHBAORD</div>
    )
}
function mapStateToProps(state) {
    return {
      account: state.account
    };
  }
export default connect(mapStateToProps)(Dashboard);