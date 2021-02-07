const isProduction = process.env.NODE_ENV === 'production'
const isDocker = process.env.NODE_ENV === 'docker'
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
        '~/plugins/element-ui',
        '~/plugins/axios'
    ],
    modules:[
        '@nuxtjs/axios'
    ],
    /*
   ** Axios module configuration
   ** See https://axios.nuxtjs.org/options
   */
    axios: {
        proxy: true,
        credentials: true,
    },
    proxy: {
        '/api':'http://127.0.0.1:8002'
        // '/api/': isProduction
        //     ? 'https://mlog.club'
        //     : isDocker
        //         ? 'http://bbs-go-server:8082'
        //         : 'http://127.0.0.1:8082',
    },
}