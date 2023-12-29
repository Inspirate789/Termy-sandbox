import api from '../api'
import { Layer } from '../types'

class LayersService {
    private contexts = new Map<Layer, string>()

    async saveLayer(layer: Layer): Promise<boolean> {
        try {
            const response = await api.post("/layers", null, {params: {layer}});
            console.log(response.request);
            if (response.status !== 200) {
                throw new Error(response.statusText);
            }
            return true;
        } catch (error) {
            console.error("Error during saveLayer", error);
            throw error;
        }
    }

    async getLayers(): Promise<Layer[]> {
        try {
            const response = await api.get("/layers");
            if (response.status !== 200) {
                throw new Error(response.statusText);
            }
            if (response?.data?.layers) {
                return response.data.layers;
            } else {
                throw new Error("Invalid API response");
            }
        } catch (error) {
            console.error("Error during getLayers", error);
            throw error;
        }
    }

    switchLayerContext(currentLayer: Layer | null, currentContext: unknown, newLayer: Layer | null): unknown {
        try {
            if (currentLayer !== null) {
                this.contexts.set(currentLayer, JSON.stringify(currentContext));
            }
            if (newLayer !== null) {
                const ctx = this.contexts.get(newLayer);
                return ctx !== undefined ? JSON.parse(ctx) : null;
            }
            return null;
        } catch (error) {
            console.error("Error during switchLayerContext", error);
            throw error;
        }
    }
}

export default new LayersService();
