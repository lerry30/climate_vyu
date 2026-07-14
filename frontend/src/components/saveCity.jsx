import { useState, useEffect } from 'react';
import { Heart } from 'lucide-react';
import { zFaveCities } from '../store/fave-cities';
import { zForecast } from '../store/weather-forecast';

const SaveCity = ({city}) => {
    const add = zFaveCities(fave => fave.add);
    const remove = zFaveCities(fave => fave.remove);
    const isSaved = zFaveCities(fave => fave.isSaved);
    // helps to update the heart button
    const cities = zFaveCities(fave => fave.cities);
    
    const [ saved, setSaved ] = useState(false);

    // I've added it here since I need to know if the main dashboard
    // has changed to make sure that this component also gets an update
    const zForecastCityName = zForecast(state => state.cityName);

    const save = () => {
        if(!String(city).trim()) return;

        if(!saved) {
            add(city);
        } else {
            remove(city);
        }
    }

    useEffect(() => {
        if(!String(city).trim()) return;
        setSaved(isSaved(city));
    }, [zForecastCityName, cities]);

    return (
        <Heart 
            fill={`${saved ? '#b11' : 'none'}`}
            stroke={`${saved ? 'none' : 'white'}`}
            onClick={save} 
        />
    );
}

export default SaveCity;