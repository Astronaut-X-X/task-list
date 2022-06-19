import axios from 'axios'
import storage from '../storage'

const api = axios.create({
    baseUrl: 'http://localhost:8080',
    tiemout: 8000,
    header: {}
})

api.interceptors.request.use((req) => {
    let token = storage.getItem('token');
    req.headers = {
        'Content-Type': 'application/json',
        'Authorization': token,
    }
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

function request (options) {
    options.method = options.method || 'get'
    if (options.method.toLowerCase() === 'get') {
        options.params = options.data
    }
    return api(options)
}

export default request