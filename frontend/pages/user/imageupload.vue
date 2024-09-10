<template>
  <div class="mb-6">
    <label class="block text-gray-700 text-sm font-bold mb-2">Upload Images</label>
    <input
      type="file"
      multiple
      @change="onFileChange"
      class="w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
    />
    <button
      type="button"
      @click="uploadFiles"
      class="mt-2 w-full bg-blue-500 text-white py-2 px-4 rounded-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500"
    >
      Upload Images
    </button>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useMutation } from '@vue/apollo-composable';
import { upload_image_action } from "../../components/graphql/queries";
const emit = defineEmits(['imagesUploaded']);
// Local state for selected images
const selectedImages = ref([]);
const imageUrls = ref([]);

// Mutation to upload images
const { mutate: uploadBase64Image } = useMutation(upload_image_action);

// File selection handler
const onFileChange = (event) => {
  selectedImages.value = Array.from(event.target.files);
};

// Image upload logic
const uploadFiles = async () => {
  if (selectedImages.value.length === 0) {
    alert('Please select images to upload.');
    return;
  }
  try {
    const uploadPromises = selectedImages.value.map(async (file) => {
      const reader = new FileReader();
      const base64Promise = new Promise((resolve, reject) => {
        reader.onloadend = () => resolve(reader.result.split(',')[1]);
        reader.onerror = reject;
        reader.readAsDataURL(file);
      });

      const base64String = await base64Promise;
      const { data } = await uploadBase64Image({ base64_str: base64String });

      if (data && data.uploadBase64Image && data.uploadBase64Image.url) {
        return data.uploadBase64Image.url;
      } else {
        throw new Error('Image upload failed');
      }
    });

    imageUrls.value = await Promise.all(uploadPromises);
    console.log('Uploaded Image URLs:', imageUrls.value);
    alert('Images uploaded successfully!');

    // Emit the uploaded image URLs to the parent component
    emit('imagesUploaded', imageUrls.value);
  } catch (err) {
    console.error('Error uploading images:', err);
    alert('Failed to upload images.');
  }
};
</script>