import {
    LineChart,
    Line,
    XAxis,
    YAxis,
    CartesianGrid,
    Tooltip,
    Legend,
    ResponsiveContainer,
} from 'recharts';

import { zWeatherForecastHourIndex } from '../store/wf-hour-index';

const prepareChartData = (weatherList) => {
    const data = [];

    weatherList.forEach((day) => {
        day.hour_weather_updates.forEach((hour) => {
            const temp = parseFloat(hour.temp.replace('°C', ''));
            const feelsLike = parseFloat(hour.feels_like.replace('°C', ''));

            // Short format: Mon 18:00
            const timeLabel = `${day.weekday.slice(0, 3)} ${new Date(hour.local_date_time).getHours().toString().padStart(2, '0')}:00`;

            data.push({
                time: timeLabel,
                temp: temp,
                feelsLike: feelsLike,
                fullDateTime: hour.local_date_time,
                description: hour.description,
                icon: hour.icon,
                humidity: hour.humidity,
                wind: hour.wind_speed,
            });
        });
    });

    return data;
};

const CustomDot = (props) => {
    //const { cx, cy, payload, dataKey, onDotClick, index } = props;
    const { cx, cy, dataKey, onDotClick, index } = props;

    return (
        <circle
            cx={cx}
            cy={cy}
            r={8}
            fill={dataKey === 'temp' ? '#f97316' : '#3b82f6'}
            stroke={dataKey === 'temp' ? '#c2410c' : '#2563eb'}
            strokeWidth={2}
            onClick={(data) => {
                onDotClick && onDotClick(index);
            }}
            style={{ cursor: 'pointer' }}
        />
    );
};

const WeatherForecastChart = ({ weatherList }) => {
    const zSetIndex = zWeatherForecastHourIndex(state => state?.setIndex);
    const chartData = prepareChartData(weatherList);

    const handleDotClick = (index) => {
        zSetIndex(index);
    };

    return (
        <div className="w-full h-[500px] p-4 bg-gray-900 text-white rounded-xl shadow-2xl">
            <ResponsiveContainer width="100%" height="100%">
                <LineChart
                    data={chartData}
                    margin={{ top: 20, right: 30, left: 20, bottom: 40 }}
                >
                    <CartesianGrid strokeDasharray="3 3" stroke="#374151" />
                    <XAxis
                        dataKey="time"
                        angle={-35}
                        textAnchor="end"
                        height={70}
                        tick={{ fill: '#d1d5db', fontSize: 12 }}
                        stroke="#6b7280"
                    />
                    <YAxis
                        yAxisId="temp"
                        domain={['dataMin - 5', 'dataMax + 5']}
                        tick={{ fill: '#d1d5db' }}
                        stroke="#60a5fa"
                        label={{
                            value: 'Temperature °C',
                            angle: -90,
                            position: 'insideLeft',
                            fill: '#93c5fd',
                            fontSize: 14,
                        }}
                    />
                    <Tooltip
                        contentStyle={{
                            backgroundColor: '#1f2937',
                            border: '1px solid #4b5563',
                            borderRadius: '8px',
                            color: '#f3f4f6',
                        }}
                        formatter={(value, name) => [`${value} °C`, name]}
                        labelStyle={{ color: '#fbbf24' }}
                    />
                    <Legend wrapperStyle={{ paddingTop: 10 }} />

                    <Line
                        yAxisId="temp"
                        type="monotone"
                        dataKey="temp"
                        name="Temperature"
                        stroke="#f97316"
                        strokeWidth={2.5}
                        dot={{ r: 5, fill: '#fb923c', stroke: '#c2410c' }}
                        activeDot={<CustomDot onDotClick={handleDotClick} />}
                    />

                    <Line
                        yAxisId="temp"
                        type="monotone"
                        dataKey="feelsLike"
                        name="Feels Like"
                        stroke="#3b82f6"
                        strokeWidth={2}
                        strokeDasharray="5 5"
                        dot={{ r: 5, fill: '#60a5fa' }}
                        activeDot={<CustomDot onDotClick={handleDotClick} />}
                    />
                </LineChart>
            </ResponsiveContainer>
        </div>
    );
};

export default WeatherForecastChart;