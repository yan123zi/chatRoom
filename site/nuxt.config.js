export default {
    server:{
        port: 8001,
        host: '127.0.0.1'
    },
    head:{
        meta: [
            { charset: 'utf-8' },
            {
                name: 'viewport',
                content:
                    'width=device-width, initial-scale=1, maximum-scale=1, minimum-scale=1, user-scalable=no, minimal-ui',
            },
            { name: 'window-target', content: '_top' },
        ],
    },
    plugins:[
        '~/plugins/element-ui'
    ]
}