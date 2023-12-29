import api from '../api'
import { AuthRequest } from '../types'

class AuthService {
    private setToken(token: string | null) {
        localStorage.setItem('token', token || "")
    }

    getToken(): string | null {
        return localStorage.getItem('token')
    }

    async login(options?: AuthRequest): Promise<string> {
        try {
            const response = await api.post("/login", options);
            if (response.status !== 200) {
                throw new Error(response.statusText);
            }
            if (response?.data?.token) {
                this.setToken(response.data.token);
                return response.data.token;
            } else {
                throw new Error("Invalid API response");
            }
        } catch (error) {
            console.error("Error during login", error);
            throw error;
        }
    }

    async refresh(): Promise<string> {
        try {
            const response = await api.get("/refresh");
            if (response.status !== 200) {
                throw new Error(response.statusText);
            }
            if (response?.data?.token) {
                this.setToken(response.data.token);
                return response.data.token;
            } else {
                throw new Error("Invalid API response");
            }
        } catch (error) {
            console.error("Error during refresh", error);
            throw error;
        }
    }

    async logout(): Promise<boolean> {
        try {
            const response = await api.delete("/logout");
            if (response.status !== 200) {
                throw new Error(response.statusText);
            }
            this.setToken(null);
            return true;
        } catch (error) {
            console.error("Error during logout", error);
            throw error;
        }
    }
}

export default new AuthService();
