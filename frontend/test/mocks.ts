import axios from 'axios';
import MockAdapter from 'axios-mock-adapter';

export const mockAdapter = new MockAdapter(axios);

const localStorageMock = (() => {
    let store: any = {};

    return {
        getItem(key: string) {
            return store[key] || null;
        },
        setItem(key: string, value: string) {
            store[key] = value.toString();
        },
        removeItem(key: string) {
            delete store[key];
        },
        clear() {
            store = {};
        }
    };
})();

Object.defineProperty(window, 'localStorage', {
    value: localStorageMock
});
