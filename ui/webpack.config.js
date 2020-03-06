const webpack = require("webpack");
const path = require("path");
const MiniCSSExtractPlugin = require("mini-css-extract-plugin");
const CopyWebpackPlugin = require("copy-webpack-plugin");

const PRODUCTION = process.env.NODE_ENV === "production";

const config = {
    mode: PRODUCTION ? "production" : "development",
    entry: path.resolve(__dirname, `src/main/app/index.tsx`),
    output: {
        path: path.resolve(__dirname, "dist"),
        filename: "bundle.js"
    },
    module: {
        rules: [
            {
                test: /\.ts[x]?$/,
                use: [
                    {
                        loader: "babel-loader"
                    },
                    {
                        loader: "ts-loader"
                    }
                ]
            },
            {
                test: /\.[s]?css$/,
                use: [
                    {
                        loader: PRODUCTION ? MiniCSSExtractPlugin.loader : "style-loader"
                    },
                    {
                        loader: "css-loader"
                    },
                    {
                        loader: "sass-loader"
                    }
                ]
            },
            {
                test: /\.(png|jpg|woff2?|ttf|otf|eot|svg)$/,
                loader: "file-loader"
            }
        ]
    },
    resolve: {
        extensions: [".ts", ".tsx", ".js"]
    },
    plugins: [
        new webpack.EnvironmentPlugin(["NODE_ENV"])
    ]
};

if (process.env.NODE_ENV === "production") {
    config.devtool = "source-map";
    config.module.rules[1].use[0].options = {
        hmr: !PRODUCTION
    };
    config.plugins.push(new MiniCSSExtractPlugin({
        filename: "bundle.css"
    }));
    config.plugins.push(new CopyWebpackPlugin([
        { from: "src/main/app/index.html", to: "index.html" },
        { from: "src/main/app/webfonts/*", to: "webfonts/[name].[ext]" }
    ]));
} else {
    config.devServer = {
        historyApiFallback: true,
        hot: true,
        inline: true,
        contentBase: "./src/main/app",
        port: 8000,
        publicPath: "/",
        // proxy: {
        //     "/webfonts": {
        //         target: "http://localhost:8000",
        //         pathRewrite: {
        //             "^": ""
        //         }
        //     }
        // }
    };
    config.devtool = "inline-source-map";
    config.plugins.push(new webpack.HotModuleReplacementPlugin());
}

module.exports = config;