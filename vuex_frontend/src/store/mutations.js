export const authenticated = (store, token) => {
    store.token = token
}

export const authguard = (store, auth) => {
    store.auth = auth
}

export const statusLogin = (store, status) => {
    store.status = status
}

export const loading = (store, loading) => {
    store.loading = loading
}
