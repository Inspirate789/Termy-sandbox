import api from '../api'
import {
    Lang,
    Layer,
    UnitLinkages,
    UnitsData,
    UnitShort,
    Units,
    UnitFullSet,
    InputUnits,
    UnitLinkagesSet, InputUnit
} from '../types'

class UnitsService {
    private units: UnitsData = {contexts: [], units: []};
    private properties = new Map<Layer, Map<string, number>>();
    private models = new Map<Layer, Map<string, number>>();
    private unitsActual = new Map<Layer, boolean>();
    private propertiesActual = new Map<Layer, boolean>();
    private modelsActual = new Map<Layer, boolean>();

    private async loadModels(layer: Layer): Promise<boolean> {
        try {
            const response = await api.get("/models", {params: {layer}});
            if (response.status !== 200) {
                throw new Error(response.statusText);
            }
            if (response?.data?.models) {
                if (this.models.get(layer) === undefined) {
                    this.models.set(layer, new Map<string, number>());
                }
                for (const model of response.data.models) {
                    this.models.get(layer)?.set(model.name, model.id);
                }
            } else {
                throw new Error("Invalid API response");
            }

            this.modelsActual.set(layer, true);

            return true;
        } catch (error) {
            console.error("Error during loadModels", error);
            throw error;
        }
    }

    private async loadProperties(layer: Layer): Promise<boolean> {
        try {
            const response = await api.get("/properties", {params: {layer}});
            if (response.status !== 200) {
                throw new Error(response.statusText);
            }
            if (response?.data?.properties) {
                if (this.properties.get(layer) === undefined) {
                    this.properties.set(layer, new Map<string, number>());
                }
                for (const property of response.data.properties) {
                    this.properties.get(layer)?.set(property.name, property.id);
                }
            } else {
                throw new Error("Invalid API response");
            }

            this.propertiesActual.set(layer, true);

            return true;
        } catch (error) {
            console.error("Error during loadProperties", error);
            throw error;
        }
    }

    private async saveModel(layer: Layer, model: string): Promise<number> {
        try {
            let unitModel_id = this.models.get(layer)?.get(model);
            if (unitModel_id === undefined) {
                // console.log("Cannot find unit model _id");
                const response = await api.post("/models", {modelNames: [model]}, {params: {layer}});
                if (response.status !== 200) {
                    throw new Error(response.statusText);
                }
                this.modelsActual.set(layer, false);
                const id: number | undefined = response.data.models_id[0];
                if (id !== undefined) {
                    unitModel_id = id;
                    this.models.get(layer)?.set(model, id);
                    this.modelsActual.set(layer, true);
                } else {
                    throw new Error("Cannot find unit model _id");
                }
            }

            return unitModel_id;
        } catch (error) {
            console.error("Error during saveModel", error);
            throw error;
        }
    }

    private async saveProperties(layer: Layer, properties: string[]): Promise<number[]> {
        try {
            const unitProperties_id: number[] = [];
            for (const property of properties) {
                const id = this.properties.get(layer)?.get(property)
                if (id) {
                    unitProperties_id.push(id);
                } else {
                    // console.log("Cannot find unit property _id");
                    const response = await api.post("/properties", {propertyNames: [property]}, {params: {layer}});
                    if (response.status !== 200) {
                        throw new Error(response.statusText);
                    }
                    this.propertiesActual.set(layer, false);
                    const id: number | undefined = response.data.properties_id[0];
                    if (id !== undefined) {
                        unitProperties_id.push(id);
                        this.properties.get(layer)?.set(property, id);
                        this.propertiesActual.set(layer, true);
                    } else {
                        throw new Error("Cannot find unit model _id");
                    }
                }
            }

            return unitProperties_id;
        } catch (error) {
            console.error("Error during saveModel", error);
            throw error;
        }
    }

    async saveUnits(layer: Layer, units: InputUnits): Promise<boolean> {
        try {
            const outputUnits: Units = {units: [], contexts: units.contexts};

            const promises = [];
            if (!this.modelsActual.get(layer)) {
                promises.push(this.loadModels(layer));
            }
            if (!this.propertiesActual.get(layer)) {
                promises.push(this.loadProperties(layer));
            }
            await Promise.all(promises);

            for (const unitSet of units.units) {
                const outputUnitSet: UnitLinkagesSet = new Map<Lang, UnitLinkages>();

                for (const [lang, unit] of unitSet) {
                    const [unitModel_id, unitProperties_id] = await Promise.all([
                        this.saveModel(layer, unit.model),
                        unit.properties ?
                        this.saveProperties(layer, unit.properties) : undefined,
                    ])

                    outputUnitSet.set(lang, {
                        text: unit.text,
                        model_id: unitModel_id,
                        properties_id: unitProperties_id,
                    });
                }

                outputUnits.units.push(outputUnitSet);
            }

            const response = await api.post("/units", outputUnits, {params: {layer}});
            if (response.status !== 200) {
                throw new Error(response.statusText);
            }

            this.unitsActual.set(layer, false);

            return true;
        } catch (error) {
            console.error("Error during saveUnits", error);
            throw error;
        }
    }

    static getByValue(map: Map<string, number>, searchValue: number) {
        for (const [key, value] of map) {
            if (value === searchValue)
                return key;
        }
    }

    static getByValues(map: Map<string, number>, searchValues: number[]) {
        const res: string[] = [];
        // console.log(searchValues);
        for (const [key, value] of map) {
            if (searchValues.find((v) => v === value))
                res.push(key);
        }
        return res;
    }

    async getUnits(layer: Layer, lang: Lang, pattern: string, offset: number, limit: number): Promise<InputUnit[]> {
        try {
            let units: UnitsData;
            if (!this.unitsActual.get(layer)) {
                const response = await api.get("/units", {params: {layer}});
                if (response.status !== 200) {
                    throw new Error(response.statusText);
                }
                const unitSets: UnitFullSet[] = []
                for (const unitSet of response.data.units) {
                    unitSets.push(new Map(Object.entries(unitSet)));
                }
                this.units = units = {
                    contexts: response.data.contexts,
                    units: unitSets,
                };
                this.unitsActual.set(layer, true);
            } else {
                units = this.units;
            }

            const promises = [];
            // console.log(this.modelsActual.get(layer), this.propertiesActual.get(layer));
            if (!this.modelsActual.get(layer)) {
                promises.push(this.loadModels(layer));
            }
            if (!this.propertiesActual.get(layer)) {
                promises.push(this.loadProperties(layer));
            }
            await Promise.all(promises);
            const models = this.models.get(layer);
            const properties = this.properties.get(layer);
            // console.log(models || new Map(), properties || new Map());

            const res: InputUnit[] = [];
            for (const unitsSet of units.units) {
                // console.log(unitsSet);
                const unit = unitsSet.get(lang);
                if (unit && unit?.text.includes(pattern)) {
                    if (offset === 0) {
                        if (limit === 0) {
                            break;
                        }
                        // console.log(unit);
                        // console.log(models);
                        // console.log(properties);
                        const unitInfo = {
                            text: unit.text,
                            model: UnitsService.getByValue(models || new Map(), unit.model_id) || "",
                            properties: UnitsService.getByValues(properties || new Map(), unit.properties_id || []),
                        };
                        // console.log(unitInfo);
                        res.push(unitInfo);
                        limit--;
                    } else  {
                        offset--;
                    }
                }
            }

            return res;
        } catch (error) {
            console.error("Error during getUnits", error);
            throw error;
        }
    }

    async deleteUnit(layer: Layer, unit: UnitShort): Promise<boolean> {
        try {
            const response = await api.delete("/units", {params: {layer}, data: unit});
            if (response.status !== 200) {
                throw new Error(response.statusText);
            }
            this.unitsActual.set(layer, false);
            return true;
        } catch (error) {
            console.error("Error during deleteUnit", error);
            throw error;
        }
    }
}

export default new UnitsService();
