export interface Config {
    api_url: string
    // token: string | null
    units_per_page: number
    // months: string[]
}

export interface AuthRequest {
    username: string
    password: string
}

export type Layer = string;

export type Lang = string;

export interface UnitShort {
    lang: Lang
    text: string
}

export interface UnitLinkages {
    text: string
    model_id: number
    properties_id?: number[]
}

export interface UnitFull {
    model_id: number
    regDate: string
    text: string
    properties_id: number[]
    contexts_id: number[]
}

export type UnitLinkagesSet = Map<Lang, UnitLinkages>;

export type UnitFullSet = Map<Lang, UnitFull>;

export interface Context {
    id: number
    text: string
}

export interface Units {
    contexts: Map<Lang, string>
    units: UnitLinkagesSet[]
}

export interface UnitsData {
    contexts: Context[]
    units: UnitFullSet[]
}

export interface InputUnit {
    text: string
    model: string
    properties?: string[]
}

export interface InputUnits {
    contexts: Map<Lang, string>
    units: Map<Lang, InputUnit>[]
}

export interface Property {
    id: number
    name: string
}

export type ModelElement = Property;

export type Model = Property;
