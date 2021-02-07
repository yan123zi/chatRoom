export default function ({$axios, app}) {
    $axios.onRequest()
    $axios.onResponse((response => {
        if (response.status!==200){
            return Promise.reject(response)
        }
        const respObj=response.data
        if (respObj.success){
            return Promise.resolve(respObj.data)
        }else {
            return Promise.reject(respObj)
        }
    }))
}
