import React, { FunctionComponent } from "react";
import "./style.scss";

interface PageProps {
    header?: any; // FIXME: Header Object
}

const Page: FunctionComponent<PageProps> = ({ children, header = null }) => {
    return (
        <div className="page">
            {header && (
                <header className="page__header">
                    <img className="page__header-logo" src={header.logo} />
                </header>
            )}
            <div className="page__content">
                {children}
            </div>
        </div>
    );
};

export default Page;