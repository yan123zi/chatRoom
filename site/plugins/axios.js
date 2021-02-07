import qs from 'qs'

export default function ({$axios, app}) {
    $axios.onRequest((reqConfig => {
        reqConfig.transformRequest = [
            function (data) {
                if (process.client && data instanceof FormData) {
                    // 如果是FormData就不转换
                    return data
                }
                data = qs.stringify(data)
                return data
            }
        ]
    }))
    $axios.onResponse((response => {
        if (response.status !== 200) {
            return Promise.reject(response)
        }
        const respObj = response.data
        if (respObj.success) {
            return Promise.resolve(respObj.data)
        } else {
            return Promise.reject(respObj)
        }
    }))
}
