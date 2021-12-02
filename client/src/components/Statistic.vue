<template>
    <div>
        <div class="col-md-7">
          <table class="table table-striped">
            <thead>
              <tr>
                <th scope="col">Fecha</th>
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
    </div>
</template>


<script>


export default ({
    name: "Statistic",
    props:["statistics"],
    // data(){
    //     return {
    //         statisticsVals:[]
    //     }
    // },
    created(){
        this.loadStatistics();
    },
    methods:{
        loadStatistics(){
            fetch('http://localhost:8000/realtime/get-reading')
                .then(res => res.json())
                .then(data =>{
                    this.statisticsVals = data;
                    console.log(this.statisticsVals)
                })
        }
    }
})
</script>
