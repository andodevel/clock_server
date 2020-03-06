import React from "react";
import { RouteComponentProps } from "@reach/router";
import Page from "../page";
import "./style.scss";

interface HomePageProps extends RouteComponentProps, JSX.IntrinsicAttributes {

}

export default ({}: HomePageProps) => {
    return (
        <Page>
            <div className="home-page">
                <header>
                    <h1><mark>CLOCK</mark> SERVER</h1>
                    <span>0.0.1</span>
                </header>
                <main>
                    <div className="container">
                    </div>
                </main>
            </div>
        </Page>
    )
};