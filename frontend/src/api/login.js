import { apiClient } from "./apiClient"

export const login = async(credentials) => {
    const response = await apiClient("/echo/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(credentials)
    })

    return response
}

