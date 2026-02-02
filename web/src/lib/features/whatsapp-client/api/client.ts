import { config } from '$lib/shared/config';
import type { Credentials } from '../model/store';

/**
 * Proxies requests to the Green API service.
 *
 * @see {@link https://green-api.com/en/docs/api/account/GetSettings/ | Green API Documentation}
 * @see {@link https://green-api.com/en/docs/api/account/GetStateInstance/ | Green API Documentation}
 * @see {@link https://green-api.com/en/docs/api/sending/SendMessage/ | Green API Documentation}
 * @see {@link https://green-api.com/en/docs/api/sending/SendFileByUrl/ | Green API Documentation}
 *
 * @param method - The Green API method name (e.g., 'sendMessage', 'getSettings').
 * @param credentials - Instance credentials containing ID and Token.
 * @param body - Optional request payload (for POST methods).
 *
 * @returns {Promise<any>} Parsed JSON response from the API.
 * @throws {Error} If the network request fails or the API returns a non-OK status.
 */
export async function fetchGreenApi(method: string, credentials: Credentials, body?: object) {
    const url = `${config.apiBaseUrl}/waInstance${credentials.idInstance}/${method}/${credentials.apiTokenInstance}`;
    console.log('Fetching URL:', url);

    const response = await fetch(url, {
        method: body ? 'POST' : 'GET',
        headers: { 'Content-Type': 'application/json' },
        body: body ? JSON.stringify(body) : null
    });

    console.log('Response:', response);

    if (!response.ok) {
        let errorDetails;
        try {
            errorDetails = await response.json();
        } catch {
            errorDetails = await response.text();
        }
        throw new Error(`HTTP ${response.status} ${response.statusText}\n${JSON.stringify(errorDetails, null, 2)}`);
    }

    return response.json();
}
