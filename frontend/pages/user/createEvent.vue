
<script setup>
import { ref } from 'vue';
import { Form, Field, ErrorMessage } from 'vee-validate';
import { toFieldValidator } from '@vee-validate/zod';
import { z } from 'zod';
import { useMutation,useQuery } from '@vue/apollo-composable';
import {insert_event_mutation} from "../../utils/queries"
import {GET_USER_BY_HIS_ID} from "../../utils/queries"
import {insert_image_imageTable} from "../../utils/queries"
import { upload_image_action } from "../../utils/queries";
import AlertMessage from "../../components/AlertMessage.vue"
import { useRouter } from 'vue-router';
const router=useRouter()
const userid = ref(localStorage.getItem("userId"));
definePageMeta({
    layout: 'user',
    middleware:"auth-log"

});
const alertMessage = ref('');
const alertVisible = ref(false);
const alertType = ref('success');
const loadings = ref(false);
const { result, loading, error ,refetch} = useQuery(GET_USER_BY_HIS_ID);
const {mutate:insertImage}=useMutation(insert_image_imageTable)
const { mutate: uploadBase64Image } = useMutation(upload_image_action);
const onFileChange = (event) => {
  selectedImages.value = Array.from(event.target.files);
};

const formData = ref({
  title: '',
  description: '',
  venue: '',
  address: '',
  price: 'free',
  specificPrice: 0,
  preparationDate: '',
  category: 'Food',
  tags: ''
});


const schema = z.object({
  title: z.string().nonempty('Title is required'),
  description: z.string().nonempty('Description is required'),
  venue: z.string().nonempty('Venue is required'),
  address: z.string().nonempty('Address is required'),
  price: z.enum(['free', 'paid'], 'Invalid price option'),
  specificPrice: z.number().optional(),
  preparationDate: z.string().nonempty('Preparation date is required'),
  category: z.enum(['Food', 'Tech', 'Education', 'Entertainment', 'Sport'], 'Invalid category'),
  tags: z.string().transform(val => val.split(',').map(tag => tag.trim()))
});
const selectedImages = ref([]);
const imageUrls = ref([]);
// Set uploaded image URLs from child component
const setImageUrls = (urls) => {
  imageUrls.value = urls;
};

// Mutation for inserting the event


const { mutate: insertEvent } = useMutation(insert_event_mutation);
const showAlert = (message, type = 'success') => {
  alertMessage.value = message;
  alertType.value = type;
  alertVisible.value = true;
  setTimeout(() => {
    alertVisible.value = false;
  }, 4000); 
};
// Form submission logic
const onSubmit = async (values) => {
       loading.value=true
  try {
  
    if (selectedImages.value.length === 0) {
    showAlert('Please select images to upload.', 'error');

    return;
  }
  
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
        showAlert('Image upload failed',"error")
      }
    });

    imageUrls.value = await Promise.all(uploadPromises);
    console.log('Uploaded Image URLs:', imageUrls.value);

    const response = await insertEvent({
      title: values.title,
      description: values.description,
      venue: values.venue,
      address: values.address,
      price: values.price,
      specific_price: values.price === 'paid' ? values.specificPrice : 0,
      preparation_date: values.preparationDate,
      category: values.category,
      featured_image: imageUrls.value[0],
      tags: values.tags,
      user_id: String(userid.value)
    });
     const eventId = ref(response.data.insert_events.returning[0].id);
     const imagePromises = imageUrls.value.map(url => 
      insertImage({
        url: url,
        event_id: String(eventId.value)
      })
    );

    await Promise.all(imagePromises);

    showAlert('Event created successfully!', 'success');

    await refetch()
    router.replace("/user")
    console.log('Event ID:', eventId);
  } catch (err) {
    console.error('Error creating event:', err);
    showAlert('Failed to create event.', 'error');

  }
  finally{
    loading.value=false
  }
};
</script>
<template>
  <div class="flex flex-col mt-20">
    <h1 class="text-3xl text-center">Create an Event</h1>
    <p v-if="loading">Loading...</p>
    <div v-else>
      <AlertMessage :message="alertMessage" :type="alertType" :visible="alertVisible" />

    <Form @submit="onSubmit" :validation-schema="toFieldValidator(schema)" class="max-w-2xl mx-auto p-8 bg-white shadow-md rounded-lg">
      <div class="mb-6">
        <label for="title" class="block text-gray-700 text-sm font-bold mb-2">Title</label>
        <Field
          v-model="formData.title"
          name="title"
          type="text"
          placeholder="Event Title"
          class="w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <ErrorMessage name="title" class="text-red-500 text-sm"/>
      </div>

      <div class="mb-6">
        <label for="description" class="block text-gray-700 text-sm font-bold mb-2">Description</label>
        <Field
          v-model="formData.description"
          name="description"
          placeholder="Event Description"
          as="textarea"
          class="w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <ErrorMessage name="description" class="text-red-500 text-sm"/>
      </div>

      <!-- Image Upload Component -->
      <div class="mb-6">
    <label class="block text-gray-700 text-sm font-bold mb-2">Upload Images</label>
    <input
      type="file"
      multiple
      @change="onFileChange"
      class="w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
    />
   
  </div>

      <div class="mb-6">
        <label for="venue" class="block text-gray-700 text-sm font-bold mb-2">Venue and Address</label>
        <Field
          v-model="formData.venue"
          name="venue"
          type="text"
          placeholder="Venue Name"
          class="w-full px-3 py-2 mb-4 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <ErrorMessage name="venue" class="text-red-500 text-sm"/>
        <Field
          v-model="formData.address"
          name="address"
          type="text"
          placeholder="Address"
          class="w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <ErrorMessage name="address" class="text-red-500 text-sm"/>
      </div>

      <div class="mb-6">
        <label for="price" class="block text-gray-700 text-sm font-bold mb-2">Price</label>
        <Field
          as="select"
          v-model="formData.price"
          name="price"
          class="w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
        >
          <option value="free">Free</option>
          <option value="paid">Paid</option>
        </Field>
        <ErrorMessage name="price" class="text-red-500 text-sm"/>
        <Field
          v-if="formData.price === 'paid'"
          v-model="formData.specificPrice"
          name="specificPrice"
          type="number"
          placeholder="Specify Amount"
          class="w-full px-3 py-2 mt-2 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <ErrorMessage name="specificPrice" class="text-red-500 text-sm"/>
      </div>

      <div class="mb-6">
        <label for="preparationDate" class="block text-gray-700 text-sm font-bold mb-2">Preparation Date</label>
        <Field
          v-model="formData.preparationDate"
          name="preparationDate"
          type="date"
          class="w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <ErrorMessage name="preparationDate" class="text-red-500 text-sm"/>
      </div>

      <div class="mb-6">
        <label for="category" class="block text-gray-700 text-sm font-bold mb-2">Category</label>
        <Field
          as="select"
          v-model="formData.category"
          name="category"
          class="w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
        >
          <option value="Food">Food</option>
          <option value="Tech">Tech</option>
          <option value="Education">Education</option>
          <option value="Entertainment">Entertainment</option>
          <option value="Sport">Sport</option>
        </Field>
        <ErrorMessage name="category" class="text-red-500 text-sm"/>
      </div>

      <div class="mb-6">
        <label for="tags" class="block text-gray-700 text-sm font-bold mb-2">Tags</label>
        <Field
          v-model="formData.tags"
          name="tags"
          type="text"
          placeholder="Comma-separated tags"
          class="w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <ErrorMessage name="tags" class="text-red-500 text-sm"/>
      </div>

      <button
        type="submit"
        class="w-full bg-blue-500 text-white py-2 px-4 rounded-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500"
      >
        Create Event
      </button>
    </Form>
    </div>
    
  </div>
</template>
