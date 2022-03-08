import axios from 'axios';

export default axios.create({
    baseURL: "http://localhost:3020/",
    headers: {
        "Content-type": "application/json"
    }
})