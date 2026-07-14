import { useState, useEffect } from 'react';
import { zForecast } from '../store/weather-forecast';
import { zWeatherForecastHourIndex } from '../store/wf-hour-index';

import { formattedDateAndTime } from '../utils/datetime';

import TitleFormat from '../utils/titleFormat';
import WeatherForecastChart from './weatherChart';
import Tooltip from './tooltip';
import SaveCity from './saveCity';

const Main = () => {
    // This main dashboard view only based on this zustand store data
    // to display something on screen
    const zForecastCityName = zForecast(state => state.cityName);
    const zForecastCountry = zForecast(state => state.country);
    const zForecastSunrise = zForecast(state => state.sunrise);
    const zForecastSunset = zForecast(state => state.sunset);
    const zForecastList = zForecast(state => state.list);

    const zWFHourIndex = zWeatherForecastHourIndex(state => state.index);
    //const zSetIndex = zWeatherForecastHourIndex(state => state?.setIndex);

    const [ selectedDay, setSelectedDay ] = useState({});

    const repDayList = zForecastList.map(day => day.hour_weather_updates.length);

    useEffect(() => {
        if(zForecastList.length === 0) return;
        let hIndex = zWFHourIndex;
        let dIndex = 0;
        for(let i = 0; i < repDayList.length; i++) {
            const d = hIndex - repDayList[i];
            if(d < 0) {
                dIndex = i;
                break;
            }
            hIndex = d;
        }
        setSelectedDay(zForecastList[dIndex].hour_weather_updates[hIndex]);
    }, [zWFHourIndex, zForecastList]);

    return (
        <main 
            className="w-full lg:w-[80%] 
                h-full bg-gray-700 overflow-y-auto p-4 flex flex-col gap-2
                lg:rounded-bl-none lg:rounded-tr-lg
            "
        >
            <div className="font-sans bg-gray-900 text-white p-4 flex justify-between rounded-xl shadow-sm flex-col-reverse sm:flex-row">
                <span>{zForecastCityName}, {zForecastCountry}</span>
                <div className="w-full flex items-center gap-4 justify-between sm:justify-center sm:w-fit">
                    <span>{formattedDateAndTime(new Date())}</span>
                    <button>
                        <Tooltip text="save city">
                            <SaveCity city={zForecastCityName} />
                        </Tooltip>
                    </button>
                </div>
            </div>

            <div className="bg-gray-900 text-white p-6 flex flex-col rounded-xl shadow-sm">
                <span className="text-sm">Weather Forecast</span>
                <article className="flex justify-between flex-col md:flex-row">
                    <span className="font-sans text-5xl">{TitleFormat(selectedDay?.description)}</span>
                    <div className="md:text-right mt-2 md:mt-0">
                        <div>
                            <span className="text-white/75">Precipitation:&nbsp;</span>
                            <span>{selectedDay?.chance_of_rain_or_snow}</span>
                        </div>
                        <div>
                            <span className="text-white/75">Humidity:&nbsp;</span>
                            <span>{selectedDay?.humidity}</span>
                        </div>
                    </div>
                </article>

                <article className="flex justify-between flex-col-reverse md:flex-row">
                    <div className="flex items-center gap-2 mt-2 md:mt-0">
                        <img src={selectedDay?.icon} alt="" />
                        <p>
                            <span className="mr-2">
                                {formattedDateAndTime(new Date(selectedDay?.local_date_time))}
                            </span>
                            <span>
                                {new Date(selectedDay?.local_date_time).toLocaleDateString('en-US', {weekday: 'short'})}
                            </span>
                        </p>
                    </div>
                    <div>
                        <span className="text-white/75">Wind:&nbsp;</span>
                        <span>{selectedDay?.wind_speed}</span>
                    </div>
                </article>

                <article>
                    <span className="font-sans text-lg text-white/75">
                        Temperature:&nbsp;
                    </span>
                    <span className="font-sans text-2xl">
                        {selectedDay?.temp}
                    </span>
                </article>
                <article>
                    <span className="font-sans text-lg text-white/75">
                        Feels Like:&nbsp;
                    </span>
                    <span className="font-sans text-2xl">
                        {selectedDay?.feels_like}
                    </span>
                </article>
            </div>

            <div>
                <WeatherForecastChart weatherList={zForecastList}  />
            </div>

            <div className="w-full flex gap-2 flex-col md:flex-row">
                {
                    zForecastList.map((wf, i) => {
                        if(i >= 5) return null;

                        const fHForecast = zForecastList[0]?.hour_weather_updates?.length;
                        const fWI = wf?.hour_weather_updates?.length - fHForecast;
                        const updates = wf?.hour_weather_updates[fWI];
                        const dt = new Date(updates?.local_date_time);

                        return (
                            <div 
                                key={i}
                                className="w-full flex flex-col justify-center items-center p-4 cursor-pointer bg-gray-900 text-white rounded-lg shadow-sm"
                                onClick={() => setSelectedDay(updates)}
                            >
                                <span className="font-light font-sans text-2xl mb-2">{wf?.weekday}</span>
                                <img src={updates.icon} alt="Weather Icon" className="mb-2" />
                                <p className="font-sans font-semibold">{dt.toLocaleString('en-US', {hour: 'numeric', hour12: true})}</p>
                            </div>
                        )
                    })
                }
            </div>

            <div className="text-sm text-white flex justify-between mt-auto pt-4 border-t border-gray-200">
                <div>
                    <span className="font-medium">Sunrise:&nbsp;</span>
                    {new Date(zForecastSunrise).toLocaleTimeString([], {hour:'2-digit', minute:'2-digit'})}
                </div>
                <div>
                    <span className="font-medium">Sunset:&nbsp;</span>
                    {new Date(zForecastSunset).toLocaleTimeString([], {hour:'2-digit', minute:'2-digit'})}
                </div>
            </div>
        </main>
    )
}

export default Main;