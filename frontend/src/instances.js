import axios from "axios";

const loz_api = axios.create({
    baseURL: "http://127.0.0.1:8080",
    timeout: 5000,
})



export default loz_api