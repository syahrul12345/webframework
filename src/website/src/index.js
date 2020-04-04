import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import Home from './screens/home';
import * as serviceWorker from './serviceWorker';

// Redux stuff
import { Provider } from 'react-redux';
import { createStore } from 'redux';
import userReducer from  './redux-modules/user'

const store = createStore(userReducer);
const component = 
    <Provider store={store}>
        <Home/>
    </Provider>

ReactDOM.render(component, document.getElementById('root'));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
