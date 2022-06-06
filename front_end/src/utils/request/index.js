import axios from 'axios'

const api = axios.create({
    baseUrl: 'http://localhost:8080',
    tiemout: 8000,
    header: {}
})

api.interceptors.request.use((req) => {
    //  TODO
    const header = req.headers
    if (header.Authorization) header.Authorization = "nil"
    return req

})

api.interceptors.response.use((req) => {
    const { code, data, msg } = req.data
    if (code == 200) {
        return data
    } else {
        return Promise.reject(msg)
    }
})

function request(options) {
    options.method = options.method || 'get'
    if (options.method.toLowerCase() === 'get') {
        options.params = options.data
    }
    return api(options)
}

export default request