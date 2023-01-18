let dataFetched = [];
let datesVals = [];
let humidityVals = [];
let temperatureVals = [];
let dioxideVals = [];
let radiationVals = [];

fetch("http://159.223.205.198:8000/realtime/get-reading")
  .then((res) => res.json())
  .then((data) => {
    dataFetched = data;
    dataFetched.forEach((element) => {
      humidityVals.push(element.humidity);
      temperatureVals.push(element.temperature);
      dioxideVals.push(element.dioxide);
      radiationVals.push(element.radiation);
      datesVals.push(element.id);
    });
  });

export const MeasurementsChartData = {
  type: "line",
  data: {
    labels: datesVals,
    datasets: [
      {
        label: "Temperature",
        data: temperatureVals,
        backgroundColor: "rgba(54,73,93,.5)",
        borderColor: "#36495d",
        borderWidth: 3,
      },
      {
        label: "Humidity",
        data: humidityVals,
        backgroundColor: "rgba(71, 183,132,.5)",
        borderColor: "#47b784",
        borderWidth: 3,
      },
      {
        label: "Dioxide",
        data: dioxideVals,
        backgroundColor: "rgba(153, 20, 60, 0.8)",
        borderColor: "#47b784",
        borderWidth: 3,
      },
      {
        label: "Radiation",
        data: radiationVals,
        backgroundColor: "rgba(20, 102, 153, 0.8)",
        borderColor: "#47b784",
        borderWidth: 3,
      },
    ],
  },
  options: {
    responsive: true,
    lineTension: 1,
    scales: {
      yAxes: [
        {
          ticks: {
            beginAtZero: true,
            padding: 25,
          },
        },
      ],
    },
  },
};

export default MeasurementsChartData;
