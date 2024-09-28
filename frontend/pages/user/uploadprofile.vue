<script setup>
definePageMeta({
  layout: 'user',
    middleware:"auth-log"
});
import { useMutation, useQuery } from "@vue/apollo-composable";
import { watchEffect, ref } from "vue";
import { uploadImageToCloudinary } from "../../utils/cloudinary";
import { GET_USER_FOR_UPLOAD_PROFILE, UPLOAD_PROFILE } from "../../utils/queries";
import { useRouter } from 'vue-router';
import AlertMessage from "../../components/AlertMessage.vue"

const router=useRouter()
const user_id = ref(localStorage.getItem("userId"));
const { result,loading,refetch } = useQuery(GET_USER_FOR_UPLOAD_PROFILE);
const { mutate: upload } = useMutation(UPLOAD_PROFILE);
const initialDataSet = ref(false);
const loadings = ref(false);
const error = ref(null);
const alertMessage = ref('');
const alertVisible = ref(false);
const alertType = ref('success');
const success = ref(false);
const uploading = ref(false); 
const previewImage=ref(null)
const user = ref({
  username: '',
  email: '',
  image: null
});

watchEffect(() => {
  if (!loading.value && result.value?.users) {

    const userData = result.value.users[0]; 
    if (!initialDataSet.value){ 
    user.value.username = userData.username;
    user.value.email = userData.email;
    user.value.image = userData.image;

    if (user.value.image) {
      previewImage.value = user.value.image; 
    }
     initialDataSet.value = true;
    }
  
  }
});
const showAlert = (message, type = 'success') => {
  alertMessage.value = message;
  alertType.value = type;
  alertVisible.value = true;
  setTimeout(() => {
    alertVisible.value = false;
  }, 2000); 
};

const handleFileUpload = async (event) => {
  const file = event.target.files[0];
  if (file) {
    previewImage.value = URL.createObjectURL(file); 
    uploading.value = true;

    try {
      const cloudinaryResponse = await uploadImageToCloudinary(file); 
      console.log("Cloudinary upload response:", cloudinaryResponse); 

      if (cloudinaryResponse) {
        user.value.image = cloudinaryResponse; 
        console.log("Updated user image URL:", user.value.image); 
      }
    } catch (err) {
      showAlert('Image upload failed. Please try again.','error')
      
    } finally {
      uploading.value = false; 
    }
  }
};
// Function to handle form submission
const handleSubmit = async () => {
  loadings.value = true;
  error.value = null;
  success.value = false;
  if (uploading.value) {
    error.value = "Please wait for the image upload to complete.";
    loadings.value = false;
    return;
  }
  try {
    const response = await upload({
      id: String(user_id.value),
      username: user.value.username,
      email: user.value.email,
      image: user.value.image,
    });

    if (response && response.data) {
      const { data } = response;
      if (data.update_users) {
        loadings.value = false;
        error.value = null;
        success.value = true;
        showAlert("Profile Updated succesfully",'success')
        await refetch(); 
        setTimeout(()=>{
            router.push("/user")
        },3000)
       
        
      }
    }
  } catch (err) {
    error.value = 'Failed to update profile. Please try again.';
  } finally {
    loadings.value = false;
  }
};

</script>

<template>
  <div class="flex justify-center items-center min-h-screen bg-gray-100 mt-20">
    <div class="bg-white p-8 rounded-lg shadow-lg w-full max-w-lg">
      <h1 class="text-2xl font-semibold mb-6 text-center text-gray-800">Update Profile</h1>
        <div v-if="loading" class="text-center text-gray-500">Loading user data...</div>
        <AlertMessage :message="alertMessage" :type="alertType" :visible="alertVisible" />

      <form @submit.prevent="handleSubmit" class="space-y-6">
        <div>
          <label class="block text-gray-700 font-medium mb-2">Username</label>
          <input 
            v-model="user.username" 
            type="text" 
            placeholder="Enter your username"
            class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500" 
            
          />
        </div>

        <div>
          <label class="block text-gray-700 font-medium mb-2">Email</label>
          <input 
            v-model="user.email" 
            type="email" 
            placeholder="Enter your email"
            class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500" 
            
          />
        </div>

     
        <div>
          <label class="block text-gray-700 font-medium mb-2">Profile Image</label>
          <input 
            type="file" 
            @change="handleFileUpload"
            class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500" 
            
          />
          <div v-if="previewImage" class="mt-4">
            <img :src="previewImage" alt="Image Preview" class="w-32 h-32 object-cover rounded-full mx-auto" />
          </div>
        </div>
          <p v-if="uploading" class="self-center w-full">uploading image....</p>
        <button 
        v-else
          type="submit"
          class="w-full bg-indigo-600 text-white py-2 rounded-lg hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500">
          Update Profile
        </button>
      </form>
    </div>
  </div>
</template>