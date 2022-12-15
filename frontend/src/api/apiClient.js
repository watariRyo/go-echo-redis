export const apiClient = async(endPoint, config) => {
    const response = await fetch(BACKEND_URL + `${endPoint}`, {
        mode: 'cors',
        ...config,
        headers: {
            ...config.headers
        },
    }).catch(() => {
        return Promise.reject(new ApiError())
    })

    if (!response.ok) {
        return Promise.reject(new ApiError(response, await response.json()))
    } else {
        return await response.json()
    }
}

export class ApiError extends Error{
    name
    url
    status
    statusText
    serverErrorContent
    constructor(response, serverErrorContent) {
        super(response.statusText || 'network error')
        this.name = "API Error"
        this.status = response.status
        this.statusText = response.statusText
        this.url = response.url
        this.serverErrorContent = serverErrorContent
    }
    serialize() {
        return Object.assign({}, this)
    }
}