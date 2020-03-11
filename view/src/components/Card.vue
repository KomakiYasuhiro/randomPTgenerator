<template>
  <v-card class="mx-auto" max-width="344">
    <v-img
      src="https://cdn.vuetifyjs.com/images/cards/sunshine.jpg"
      height="150px"
    ></v-img>
    <v-card-title>
      {{ pokemon.name }}
    </v-card-title>
    <v-card-subtitle> 図鑑No.{{ pokemon.id }} </v-card-subtitle>
  </v-card>
</template>
<script>
export default {
  data: () => ({
    show: true
  }),
  props: {
    pokemon: {
      id: String,
      name: String
    }
  }
};
</script>