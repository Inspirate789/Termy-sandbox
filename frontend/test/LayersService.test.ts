import layersService from '../src/services/LayersService';
import {mockAdapter} from "./mocks";

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

test('saveLayer', async () => {
    // arrange
    mockAdapter.onPost('/layers').reply(200);
    // act
    const ok = await layersService.saveLayer('LAYER');
    // assert
    expect(ok).toBe(true);
});

test('getLayers', async () => {
    // arrange
    mockAdapter.onGet('/layers').reply(200, { layers: ['Layer1', 'Layer2', 'Layer3'] });
    // act
    const layers = await layersService.getLayers();
    // assert
    expect(layers).toStrictEqual(['Layer1', 'Layer2', 'Layer3']);
});

test('switchLayerContext init, switch, switch back', async () => {
    // act
    let ctx = await layersService.switchLayerContext(null, null, 'LAYER');
    // assert
    expect(ctx).toBe(null);

    //arrange
    ctx = {someField: 'NEW_CONTEXT'}
    // act
    ctx = await layersService.switchLayerContext('LAYER', ctx, 'NEW_LAYER');
    // assert
    expect(ctx).toBe(null);

    // act
    ctx = await layersService.switchLayerContext('NEW_LAYER', ctx, 'LAYER');
    // assert
    expect(ctx).toStrictEqual({someField: 'NEW_CONTEXT'})
});
