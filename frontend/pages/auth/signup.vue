<script setup>
definePageMeta({
  middleware:"after-log"
})
import { Form, Field, ErrorMessage, useForm } from 'vee-validate';
import { z } from 'zod';
import { toFieldValidator } from '@vee-validate/zod';
import { ref } from 'vue';
import Eye from "../../assets/icons/Eye.vue";
import { useMutation } from '@vue/apollo-composable';
import gql from 'graphql-tag';
import { useRouter } from 'vue-router';
import { REGISTER_USER_MUTATION } from "../../utils/queries";
const router = useRouter();
const { resetForm, setErrors } = useForm();
const alertMessage = ref('');
const alertVisible = ref(false);
const alertType = ref('success');
const showAlert = (message, type = 'success') => {
  alertMessage.value = message;
  alertType.value = type;
  alertVisible.value = true;
  // Extend the timeout based on the alert duration
  setTimeout(() => {
    alertVisible.value = false;
  }, 2000); 
};
const { mutate: registerUser } = useMutation(REGISTER_USER_MUTATION);
const userData = ref({
  username: '',
  email: '',
  password: ''
});
const togglePassword = ref(false);
const successMessage = ref("");
const showPassword = () => {
  togglePassword.value = !togglePassword.value;
};
const signupRegister = async () => {
  console.log('userData values:', userData.value);
 try {
    const { data } = await registerUser({
      email: userData.value.email,
      password: userData.value.password,
      username: userData.value.username,
    });
     console.log('Mutation response:', data);
    if (data && data.signUp) {
      successMessage.value = data.signUp.message;
       showAlert(data.signUp.message, 'success');

      
      // Reset form and validation state
      
      resetForm();
     

      // Wait for 3 seconds, then redirect to login page
      setTimeout(() => {
        
        router.replace("/auth/login");
      }, 3000); // 3000 milliseconds = 3 seconds
    } else {
      console.error('No response from server');
    }
  } catch (error) {
    showAlert("Registration Error", 'error');
  }
};

const schema = z.object({
  username: z.string().min(1, "User name is required"),
  email: z.string().min(1, "The email must not be empty").email("Invalid email address"),
  password: z
    .string()
    .min(5, "The password must contain at least 5 characters")
    .max(10, "The password can contain at most 10 characters")
    .regex(/[!@#$%^&*(),.?":{}|<>]/, "The password must include at least one special character"),
});
</script>
<template>
  <div class="flex flex-col min-h-screen bg-gray-100 justify-center items-center">
        <AlertMessage :message="alertMessage" :type="alertType" :visible="alertVisible" />

    <Form @submit="signupRegister" :validation-schema="toFieldValidator(schema)" class="flex flex-col justify-center items-center w-full max-w-md p-6 bg-white rounded-lg">
      <div class="flex flex-col w-full mb-4">
        <label class="font-semibold text-gray-600 mb-2">User Name</label>
        <Field 
          type="text" 
          name="username"
          v-model="userData.username" 
          placeholder="Enter user name" 
          class="p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-indigo-500"
        />
        <ErrorMessage name="username" class="text-red-500 text-sm" />
      </div>

      <div class="flex flex-col w-full mb-4">
        <label class="font-semibold text-gray-600 mb-2">Email</label>
        <Field 
          type="text" 
          name="email"
          v-model="userData.email" 
          placeholder="Enter email" 
          class="p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-indigo-500"
        />
        <ErrorMessage name="email" class="text-red-500 text-sm" />
      </div>

      <div class="flex flex-col w-full mb-4 relative">
        <label class="font-semibold text-gray-600 mb-2">Password</label>
        <Field 
          :type="togglePassword ? 'text' : 'password'"
          name="password"
          v-model="userData.password" 
          placeholder="Enter password" 
          class="p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-indigo-500"
        />
        <ErrorMessage name="password" class="text-red-500 text-sm" />
        <Eye class="absolute top-10 right-2 cursor-pointer" @click="showPassword">
          {{ togglePassword ? 'Hide' : 'Show' }}
        </Eye>
      </div>

      <button 
        type="submit" 
        class="w-full py-2 px-4 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500"
      >
        Signup
      </button>
      <p class="font-semibold p-3 text-left">Do You have an account?<NuxtLink to="/auth/login" class="text-blue-600">Login</NuxtLink></p>
    </Form>
   
  </div>
</template>
