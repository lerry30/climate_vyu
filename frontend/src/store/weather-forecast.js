import { create } from 'zustand';
import { persist } from 'zustand/middleware';

import { getData } from '../utils/send';

export const zForecast = create(
    persist((set, get) => ({
        cityName: '',
        country: '',
        sunrise: null,
        sunset: null,
        list: [],

        setWeatherForecast: (cityName, country, sunrise, sunset, list) => {
            if(!cityName || !country) return;
            if(!list || list?.length === 0) return;
            set((state) => ({ cityName, country, sunrise, sunset, list }));
        },

        fetchWeatherForecast: async (baseUrl, city) => {
            city = city.trim().toLowerCase();
            const response = await getData(`${baseUrl}/forecast/${city}`);
            if(response) {
                get().setWeatherForecast(
                    response?.city_name,
                    response?.country,
                    new Date(response?.sunrise),
                    new Date(response?.sunset),
                    response?.list
                );
            }
        },
    }), {
        name: 'store-weather-forecast',
    })
);