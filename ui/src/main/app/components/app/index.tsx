import React, { useState, useEffect, FunctionComponent } from "react";
import { Router } from "@reach/router";
import { useDispatch } from "react-redux";
import Favicon from "react-favicon";

import HomePage from "../home-page";
import Loading from "../loading";

import faviconURL from "../../img/favicon.png";
import "./style.scss";

const App: FunctionComponent = () => {
    const [ loading, setLoading ] = useState(true);

    const dispatch = useDispatch();

    useEffect(() => {
        const load = async () => {
            setTimeout(() => setLoading(false), 1000); // TODO: Fetch Server data
        };

        load();
    }, []);

    return (
        <div className="app">
            <Favicon url={faviconURL} />
            {loading ? (
                <Loading />
            ) : (
                <Router basepath="">
                    <HomePage path="/" />
                </Router>
            )}
        </div>
    );
}

export default App;