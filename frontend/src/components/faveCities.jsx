import { zFaveCities } from '../store/fave-cities';
import { zServer } from '../store/server';
import { zForecast } from '../store/weather-forecast';
import { X } from 'lucide-react';

import TitleFormat from '../utils/titleFormat';

const FaveCities = () => {
    const zCities = zFaveCities(state => state.cities);
    const zRemoveCity = zFaveCities(state => state.remove);
    const zServerUrl = zServer(state => state.serverUrl);
    const zFetchWeatherForecast = zForecast(state => state.fetchWeatherForecast);

    function cleanCityName(locationName) {
        return locationName
            .replace(/\b(city\s+of|municipality\s+of)\s+/gi, '')
            .replace(/\s+(city|municipality)\b/gi, '')
            .trim();
    }

    const getFaveCity = async (city) => {
        try {
            city = cleanCityName(city);
            await zFetchWeatherForecast(zServerUrl, city);
        } catch(error) {
            console.log(error);
        }
    }

    return (
        <div className="w-full h-full max-h-96 flex flex-col bg-gray-900 rounded-xl p-4">
            <span className="text-white text-2xl mb-4">Favorites</span>
            {
                zCities.map((city, i) => {
                    return (
                        <button 
                            key={i} 
                            className="flex justify-between items-center p-1 border-t border-white/25 hover:bg-gray-700/75 cursor-pointer"
                            onClick={() => getFaveCity(city)}
                        >
                            <span className="text-white">
                                {TitleFormat(city)}
                            </span>
                            <X color="white" size={16} onClick={(ev) => {
                                    ev.stopPropagation();
                                    zRemoveCity(city)
                                }}
                            />
                        </button>
                    )
                })
            }
        </div>
    );
}

export default FaveCities;