import {defineStore} from "pinia"

export const usecounterStore=defineStore("counter",{
    state:()=>({
        count:0,
    }),
    actions:{
        increment(){
            this.count++
        }
    }

})