import authService from '../src/services/AuthService';
import {mockAdapter} from "./mocks";

jest.mock('axios', () => {
    return {
        ...(jest.requireActual('axios') as object),
        create: jest.fn().mockReturnValue(jest.requireActual('axios')),
    };
});

beforeEach(() => {
    mockAdapter.resetHandlers()
    window.localStorage.clear()
})

test('login', async () => {
    // arrange
    mockAdapter.onPost('/login').reply(200, { token: 'TOKEN' });
    // act
    const token = await authService.login();
    // assert
    expect(token).toBe('TOKEN');
});

test('refresh', async () => {
    // arrange
    mockAdapter.onGet('/refresh').reply(200, { token: 'NEW_TOKEN' });
    // act
    const token = await authService.refresh();
    // assert
    expect(token).toBe('NEW_TOKEN');
});

test('logout', async () => {
    // arrange
    mockAdapter.onDelete('/logout').reply(200);
    // act
    const ok = await authService.logout();
    // assert
    expect(ok).toBe(true);
});
