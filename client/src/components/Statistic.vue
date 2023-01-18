<template>
  <div>
    <div class="col-md-7">
      <table class="table table-striped">
        <thead>
          <tr>
            <th scope="col">Date</th>
            <!-- <th scope="col">Temperature</th>
            <th scope="col">Humidity</th>
            <th scope="col">Dioxide</th>
            <th scope="col">Radiation</th> -->
            <th scope="col">pH</th>
          </tr>
        </thead>
        <tbody v-bind:key="stat.id" v-for="stat in statisticsVals">
          <tr>
            <td>{{ stat.date_reading }}</td>
            <!-- <td>{{ stat.temperature }}</td>
            <td>{{ stat.humidity }}</td> -->
            <td>{{ stat.ph }}</td>
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
          phVals:[]
      }
  },
  created() {
    this.loadStatistics();
    //console.log(this.humidityVals);
  },
  methods: {
    loadStatistics() {
      fetch("http://159.223.205.198:8000/realtime/get-reading")
        .then((res) => res.json())
        .then((data) => {
          this.statisticsVals = data;
          this.statisticsVals.forEach(element => {
            //console.log(element.humidity);
            this.humidityVals.push(element.humidity);
            this.temperatureVals.push(element.temperature);
            this.phVals.push(element.ph);
            this.datesVals.push(element.date_reading);
          });
        });
      setTimeout(this.loadStatistics, 100);

    },
  },
};
</script>
