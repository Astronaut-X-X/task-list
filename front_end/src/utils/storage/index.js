
const NAMESPACE = "TL"

export default{
    setItme(key,val){
        let storage = this.getStorage()
        storage[key] = val
        window.localStorage.setItem(NAMESPACE,JSON.stringify(storage))
    },
    getItem(key){
        return this.getStorage()[key]
    },
    getStorage(){
        return JSON.parse(window.localStorage.getItem(NAMESPACE) || "{}")
    },
    clearItem(key){
        let storage = this.getStorage
        delete storage[key]
    },
    clearAll(){
        window.localStorage.clear()
    }
}