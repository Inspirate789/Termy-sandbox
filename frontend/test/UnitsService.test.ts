import {mockAdapter} from "./mocks";
import unitsService from "../src/services/UnitsService";

jest.mock('axios', () => {
    return {
        ...(jest.requireActual('axios') as object),
        create: jest.fn().mockReturnValue(jest.requireActual('axios')),
    };
});

beforeEach(() => {
    mockAdapter.resetHandlers()
    window.sessionStorage.clear()
})


test('saveUnits', async () => {
    // arrange
    let i = 4;
    let j = 4;
    mockAdapter.onPost('/units').reply(200);
    mockAdapter.onPost('/properties').reply(200, {
        "properties_id": [i++],
    });
    mockAdapter.onPost('/models').reply(200, {
        "models_id": [j++],
    });
    mockAdapter.onGet('/properties').reply(200, {
        "properties": [],
    });
    mockAdapter.onGet('/models').reply(200, {
        "models": [
            {"id": 1, "name": "model1"},
            {"id": 2, "name": "model2"},
            {"id": 3, "name": "model3"},
        ],
    });
    // act
    const ok = await unitsService.saveUnits('LAYER', {
        contexts: new Map([
            ["ru", "контекст1, содержащий термин1 и термин2"],
            ["en", "context1 with term1 and term2"],
        ]),
        units: [
            new Map([
                ["ru", {text: "термин1", model: "some_model", properties: ["aboba", "abebe", "abubu"]}],
                ["en", {text: "term1", model: "some_model", properties: ["abubu"]}]
            ]),
            new Map([
                ["ru", {text: "термин2", model: "any_model"}],
                ["en", {text: "term2", model: "any_model", properties: ["abubu"]}]
            ]),
        ],
    });
    // assert
    expect(ok).toStrictEqual(true);
});

test('getUnits', async () => {
    // arrange
    mockAdapter.onGet('/properties').reply(200, {
        "properties": [
            {"id": 1, "name": "property1"},
            {"id": 2, "name": "property2"},
            {"id": 3, "name": "property3"},
        ],
    });
    mockAdapter.onGet('/models').reply(200, {
        "models": [
            {"id": 1, "name": "model1"},
            {"id": 2, "name": "model2"},
            {"id": 3, "name": "model3"},
        ],
    });
    mockAdapter.onGet('/units').reply(200, {
        "units": [
            {
                "ru": {"text": "термин1", "reg_date": "date", "model_id": 1, "properties_id": [4], "contexts_id": [1]},
                "en": {"text": "term1", "reg_date": "date", "model_id": 1, "contexts_id": [1]}
            },
            {
                "ru": {"text": "термин2", "reg_date": "date", "model_id": 3, "contexts_id": [1]},
                "en": {"text": "term2", "reg_date": "date", "model_id": 3, "properties_id": [4], "contexts_id": [1]}
            }
        ],
        "contexts": [
            {"id": 1, "text": "контекст1, содержащий термин1 и термин2"},
            {"id": 2, "text": "context1 with term1 and term2"},
            {"id": 3, "text": "context3"},
        ],
    });
    // act
    const resp = await unitsService.getUnits('LAYER', 'ru', 'термин', 0, 10);
    console.log(resp);
    // assert
    expect(resp).toStrictEqual([
            {"text": "термин1", "model": "model1", "properties": ["aboba", "abebe", "abubu"]},
            {"text": "термин2", "model": "model3", "properties": []},
        ]);
});


test('deleteUnit', async () => {
    // arrange
    mockAdapter.onDelete('/units').reply(200);
    // act
    const ok = await unitsService.deleteUnit('LAYER', {
        lang: 'ru',
        text: 'TERM'
    });
    // assert
    expect(ok).toBe(true);
});