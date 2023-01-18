
<template>
  <div>
    <canvas id="measurements-chart"></canvas>
  </div>
</template>

<script>
import Chart from 'chart.js'
import MeasurementsChartData from '../services/measurementsChart.js'

export default {
  name: "MeasurementsChart",
  data() {
    return {
      MeasurementsChartData: MeasurementsChartData
    }
  },
  mounted() {
    const ctx = document.getElementById('measurements-chart');
    new Chart(ctx, this.MeasurementsChartData);
  },
  created() {
    this.loadData();
  },
  methods: {
    loadData() {
      //console.log("working!!");
      setTimeout(this.loadData, 100);
    },

  }
}
</script>
