<template>
  <div id="map" style="height: 400px; width: 800px;"></div>
</template>

<script setup>
import { Loader } from '@googlemaps/js-api-loader';
import { ref, watch, onMounted } from 'vue';
import {geocodeAddress} from "../../utils/geocode"
const props = defineProps({
  address: {
    type: String,
    default: ''
  }
});

const map = ref(null);
const marker = ref(null);

// Default coordinates for Addis Ababa
const DEFAULT_LAT = 11.5936;
const DEFAULT_LNG = 37.3908;

// Singleton pattern to prevent multiple instances of the loader
let loaderInstance = null;

function getGoogleMapsLoader() {
  if (!loaderInstance) {
    loaderInstance = new Loader({
      apiKey: "AIzaSyDenHT2SGUaSm0MQJ48EqMhbfy0sH379rs",
      version: "weekly",
      libraries: ["places"]
    });
  }
  return loaderInstance;
}

function initializeMap(lat, lng) {
  const mapOptions = {
    center: { lat, lng },
    zoom: 12
  };

  const markerOptions = {
    position: { lat, lng },
    icon: {
      url: "http://maps.google.com/mapfiles/ms/icons/red-dot.png"
    }
  };

  map.value = new google.maps.Map(document.getElementById("map"), mapOptions);
  marker.value = new google.maps.Marker({
    ...markerOptions,
    map: map.value
  });
}

async function updateMapWithAddress(address) {
  console.log('Received address:', address); // Log the incoming address

  try {
    const coordinates = await geocodeAddress(address);
    if (coordinates) {
      initializeMap(coordinates.lat, coordinates.lon);
    } else {
      console.error('Address could not be geocoded.');
      initializeMap(DEFAULT_LAT, DEFAULT_LNG);
    }
  } catch (error) {
    console.error('Error geocoding address:', error);
    initializeMap(DEFAULT_LAT, DEFAULT_LNG);
  }
}

onMounted(() => {
  const loader = getGoogleMapsLoader();
  loader.load()
    .then((google) => {
      updateMapWithAddress(props.address);
    })
    .catch(e => {
      console.error("Error loading Google Maps", e);
    });
});

// Watch for changes in the address prop to update the map
watch(() => props.address, (newAddress) => {
  if (map.value) {
    updateMapWithAddress(newAddress);
  }
});
</script>

<style scoped>
#map {
  height: 100%;
  width: 100%;
}
</style>