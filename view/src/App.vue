<template>
  <v-app>
    <v-content>
      <Generator/>
    </v-content>
  </v-app>
</template>

<script>
import Generator from './components/Generator';


export default {
  name: 'App',

  components: {
    Generator,
  },

  data: () => ({
    //
  }),
};
</script>
