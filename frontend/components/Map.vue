<template>
  <div id="map" class="h-96 w-full"></div>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import L from "leaflet"
import {geocodeAddress} from "../utils/geocode"

const props = defineProps({
  address: String,
});

const coordinates = ref({ lat: null, lon: null });

onMounted(async () => {
  if (props.address) {
    const { lat, lon } = await geocodeAddress(props.address);
    coordinates.value = { lat, lon };

    if (lat && lon) {
      const map = L.map('map').setView([lat, lon], 13);

      L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
      }).addTo(map);

      L.marker([lat, lon])
        .addTo(map)
        .bindPopup(`<b>${props.address}</b>`)
        .openPopup();
    }
  }
});
</script>

<style scoped>
#map {
  height: 100%;
  width: 100%;
}
</style>