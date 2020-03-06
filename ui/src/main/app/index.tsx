import "regenerator-runtime/runtime";

import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import store from "./store";
import App from "./components/app";
import "./scss/index.scss";

ReactDOM.render((
    <Provider store={store}>
        <App />
    </Provider>
), document.getElementById("main"));