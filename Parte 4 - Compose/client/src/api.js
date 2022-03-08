import http from './http-common';

const getRoot = () => {
    return http.get('/');
}

export default {
    getRoot
}