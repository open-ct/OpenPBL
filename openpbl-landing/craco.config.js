const CracoLessPlugin = require('craco-less');
const path = require("path");
const resolve = dir => path.resolve(__dirname, dir);

module.exports = {
    plugins: [
        {
            plugin: CracoLessPlugin,
            options: {
                lessLoaderOptions: {
                    lessOptions: {
                        javascriptEnabled: true,
                    },
                },
            },
        },
    ],
    babel: {
        plugins: [
            ['import', { libraryName: 'antd', libraryDirectory: 'es', style: 'css' }],
            ['@babel/plugin-proposal-decorators', { legacy: true }]
        ]
    },
    webpack: {
        alias: {
            '@': resolve("src"),
        }
    }
};
