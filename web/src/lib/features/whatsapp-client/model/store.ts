import { writable } from 'svelte/store';

export interface Credentials {
    idInstance: string;
    apiTokenInstance: string;
}

export const credentials = writable<Credentials>({
    idInstance: '',
    apiTokenInstance: ''
});

export const formFields = writable({
    phoneMessage: '',
    messageText: '',
    phoneFile: '',
    fileUrl: ''
});

export const responseLog = writable("");