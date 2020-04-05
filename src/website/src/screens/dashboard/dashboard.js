import React from 'react'
import { connect } from 'react-redux';
// Cookie stuff
import { useCookies } from 'react-cookie';
import jwt from 'jwt-decode'

const Dashboard = (props) => {
    // Pull account from the store
    const { account } = props
    console.log(account)
    const [cookies, setCookie, removeCookie] = useCookies(['cookie-name']);
    // Redux is non persistent, so account in redux store is deleted upon refresh. Cookies are still stored by browser,so we send it to the backend.
    if (account.ID == undefined) {
        // send this token to the backend to verify
        const tokenHeader = cookies["x-token"].split(" ")
        const token = tokenHeader[1]
        // send the token to the backend
        console.log(token)
    }
    return(
        <div> THIS IS THE DASHBAORD</div>
    )
}
function mapStateToProps(state) {
    return {
      account: state.user.account
    };
  }
export default connect(mapStateToProps)(Dashboard);