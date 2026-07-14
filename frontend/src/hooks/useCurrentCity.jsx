import { useEffect, useState } from 'react';
import { getData } from '../utils/send';

const useCurrentCity = () => {
    const [city, setCity] = useState(null); 
    const [status, setStatus] = useState('idle'); // idle | loading | success | error
    const [error, setError] = useState(null);

    useEffect(() => {
        setStatus('loading');

        if(!navigator.geolocation) {
            setError('Geolocation not supported');
            setStatus('error');
            return;
        }

        navigator.geolocation.getCurrentPosition(
            async (pos) => {
                try {
                    const { latitude, longitude } = pos.coords;
                    const res = await fetch(
                        `https://api.bigdatacloud.net/data/reverse-geocode-client?latitude=${latitude}&longitude=${longitude}&localityLanguage=en`
                    );
                    const data = await res.json();
                    setCity(data?.city || data?.locality || 'Unknown');
                    setStatus('success');
                } catch(error) {
                    setError('Failed to resolve city');
                    setStatus('error');
                }
            },
            (error) => {
                setError(error.message);
                setStatus('error');
            }
        );
    }, []);

    return { city, status, error };
}

export {
    useCurrentCity,
};