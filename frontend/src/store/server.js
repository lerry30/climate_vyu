import { create } from 'zustand';
//import { devtools, persist } from 'zustand/middleware';
import { persist } from 'zustand/middleware';

export const zServer = create(
    persist((set) => ({
        serverUrl: 'http://localhost:8080',
        setServerUrl: (url) => set({serverUrl: url})
    }), {
        name: 'store-config',
    })
);