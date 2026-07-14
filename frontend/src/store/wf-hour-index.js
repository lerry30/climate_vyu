import { create } from 'zustand';
import { persist } from 'zustand/middleware';

export const zWeatherForecastHourIndex = create(
    persist((set) => ({
        index: 0,
        setIndex: (index) => set({index}),
    }), {
        name: 'store-weather-forecast-hour-index',
    })
);