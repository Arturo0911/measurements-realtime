<template>
  <div>
    <div class="col-md-7">
      <table class="table table-striped">
        <thead>
          <tr>
            <th scope="col">Temperature</th>
            <th scope="col">Humidity</th>
            <th scope="col">Dioxide</th>
            <th scope="col">Radiation</th>
          </tr>
        </thead>
        <tbody v-bind:key="stat.id" v-for="stat in statisticsVals">
          <tr>
            <td>{{ stat.temperature }}</td>
            <td>{{ stat.humidity }}</td>
            <td>{{ stat.dioxide }}</td>
            <td>{{ stat.radiation }}</td>
          </tr>
        </tbody>
      </table>
    </div>
    <canvas id="measurements-chart"></canvas>
  </div>
</template>

<script>




export default {
  name: "Statistic",
  props: ["statistic"],
  data(){
      return {
          statisticsVals:[],
          datesVals:[],
          humidityVals: [],
          temperatureVals:[],
          dioxideVals:[],
          radiationVals:[]
      }
  },
  created() {
    this.loadStatistics();
    //console.log(this.humidityVals);
  },
  methods: {
    loadStatistics() {
      fetch("http://localhost:8000/realtime/get-reading")
        .then((res) => res.json())
        .then((data) => {
          this.statisticsVals = data;
          this.statisticsVals.forEach(element => {
            //console.log(element.humidity);
            this.humidityVals.push(element.humidity);
            this.temperatureVals.push(element.temperature);
            this.dioxideVals.push(element.dioxide);
            this.radiationVals.push(element.radiation);
            this.datesVals.push(element.date_reading);
          });
        });
      setTimeout(this.loadStatistics, 100);

    },
  },
};
</script>
