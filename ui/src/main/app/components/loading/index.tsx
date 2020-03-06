import React from "react";
import Icon from "../icon";
import "./style.scss";

export default () => (
    <div className="loading">
        <div className="loading__inner">
            <Icon name="spinner" type="light" />
        </div>
    </div>
);