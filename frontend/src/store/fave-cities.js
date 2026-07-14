import { create } from 'zustand';
import { persist } from 'zustand/middleware';

export const zFaveCities = create(
    persist((set, get) => ({
        cities: [],
        add: (name) => {
            name = String(name).trim().toLowerCase();
            if(!name) return;
            set(store => {
                if(store.cities?.includes(name)) return store;
                return {cities: [...store.cities, name]};
            })
        },
        remove: (name) => {
            name = String(name).trim().toLowerCase();
            if(!name) return;
            set(store => {
                if(!store.cities.includes(name)) return store;
                return {cities: store.cities.filter(c => c !== name)};
            })
        },
        isSaved: (name) => {
            name = String(name).trim().toLowerCase();
            if(!name) return false;
            return get().cities?.includes(name) || false; // should use ??
        },
    }), {
        name: 'store-fave-cities',
    })
);