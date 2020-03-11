<template>
    <v-container>
        <v-row align="center" id="gene">
            <v-col class="text-center" cols="12" sm="12">
                <div class="my-2">
                    <v-btn @click="generate" text large color="primary">Generate</v-btn>
                </div>
            </v-col>
        </v-row>
        <v-row>
            <v-col
            align="center"
            cols="4"
            v-for="(pokemon,key) in pokemons"
            :key="key"
            >
            <Card :pokemon="pokemon"/>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
    import Card from "./Card";
    export default{
        name:"Gene",
        components:{
            Card
        },
        data(){
            return{
                pokemons:[]
            };
        },
        methods:{
            generate(){
                console.log("aa")
                this.axios
                .get("https://jsonplaceholder.typicode.com/users")
                .then(response=>{
                    this.pokemons=response.data;
                    console.log(this.pokemons);
                })
                .catch(e=>{
                    alert(e);
                });
            }
        }
    };
</script>
