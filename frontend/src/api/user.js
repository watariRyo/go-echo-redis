import { apiClient } from "./apiClient"

export const getUser = async(token) => {
    const response = await apiClient("/echo/api/", {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": token
        }
    })

    return response
}

