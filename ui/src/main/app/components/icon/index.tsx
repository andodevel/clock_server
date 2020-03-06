import React, { FunctionComponent } from "react";
import "./style.scss";

interface IconProps {
    name: string;
    type?: "regular" | "light" | "solid" | "duotone";
}

const Icon: FunctionComponent<IconProps> = ({ name, type = "regular" }) => {
    let className = "";

    switch (type) {
        case "regular":
            className += "far";
            break;
        
        case "light":
            className += "fal";
            break;

        case "solid":
            className += "fas";
            break;

        case "duotone":
            className += "fad";
            break;

        default:
            break;
    }

    className += ` fa-${name}`;

    return <i className={className}></i>;
};

export default Icon;