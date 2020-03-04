<template>
  <v-row align="center" id="gene">
    <v-col class="text-center" cols="12" sm="12">
      <div class="my-2">
        <v-btn @click="generate" text large color="primary">Generate</v-btn>
      </div>
    </v-col>
  </v-row>
</template>

  <script>
  export default {
    name: 'Gene',
    methods: {
      generate() {
        this.axios.get('https://httpbin.org/get')
          .then((response) => {
            alert("ぽけ"+response.data.origin);
          })
          .catch((e) => {
            alert(e);
          });
      }
    }
  }
</script>