<template>
  <div class="flex flex-col min-h-screen bg-gray-100 justify-center items-center">
    <Form @submit="onSubmit" :validation-schema="toFieldValidator(schema)" class="flex flex-col justify-center items-center w-full max-w-md p-6 bg-white rounded-lg ">
      <h1 class="font-bold text-2xl mb-4 text-gray-800">Login Page</h1>

      <div class="flex flex-col w-full mb-4">
        <label class="font-semibold text-gray-600 mb-2">Email</label>
        <Field 
          type="email" 
          name="email"
          v-model="user.email" 
          placeholder="Enter email" 
          class="p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-indigo-500"
        />
        <ErrorMessage name="email" class="text-red-500 text-sm" />
      </div>

      <div class="flex flex-col w-full mb-6">
        <label class="font-semibold text-gray-600 mb-2">Password</label>
        <div class="flex flex-col relative">
          <Field 
            :type="togglePassword ? 'text' : 'password'"
            name="password"
            v-model="user.password" 
            placeholder="Enter password" 
            class="p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-indigo-500"
          />
          <ErrorMessage name="password" class="text-red-500 text-sm" />
          <Eye class="absolute top-3 right-3 cursor-pointer" @click="showPassword" />
        </div>
      </div>

      <div v-if="errorMessage" class="text-red-500 text-sm mb-4">{{ errorMessage }}</div>

      <button 
        type="submit" 
        class="w-full py-2 px-4 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500"
      >
        Login
      </button>

      <p class="font-semibold p-3 text-left">
        Don't have an account?
        <NuxtLink to="/auth/signup" class="text-blue-600">Sign Up</NuxtLink>
      </p>
    </Form>
  </div>
</template>

<script setup>
import { Form, Field, ErrorMessage } from 'vee-validate';
import { z } from 'zod';
import { toFieldValidator } from '@vee-validate/zod';
import { ref } from 'vue';
import gql from 'graphql-tag';
import { useMutation } from '@vue/apollo-composable';
import Eye from "../../assets/icons/Eye.vue";
import { useRouter } from 'vue-router';
import {getUserIdFromToken} from "../../utils/util"
import {useAuthStore} from "../../stores/authstore"

const authStore=useAuthStore()
const router=useRouter()
const user = ref({
  email: "",
  password: ""
});

// Form validation schema
const schema = z.object({
  email: z.string().min(1, "The email must not be empty").email("Invalid email address"),
  password: z.string().min(5, "The password must contain at least 5 characters").max(10, "The password can contain at most 10 characters")
});

// Toggle password visibility
const togglePassword = ref(false);
const showPassword = () => {
  togglePassword.value = !togglePassword.value;
};

// GraphQL login mutation
const LOGIN_MUTATION = gql`
mutation loginUser($email:String!,$password:String!){
  siginUser(email:$email,password:$password){
    token
  }
}
`;

const { mutate: loginUser } = useMutation(LOGIN_MUTATION);

// Error handling
const errorMessage = ref(null);

// Submit function
const onSubmit = async () => {
  try {
    const { data } = await loginUser({
      
        email: user.value.email,
        password: user.value.password,
      
    });

  if (data && data.siginUser && data.siginUser.token) {
          const token = data.siginUser.token;
          // Save the token to localStorage
          localStorage.setItem('token', token);
        
           const userid=getUserIdFromToken(token)
           authStore.setUserId(userid)
       
          // Clear user data
          user.value.email = "";
          user.value.password = "";

          router.push("/user");
        }
  } catch (error) {
    errorMessage.value = "Login failed. Please check your credentials.";
    console.error('Login error:', error);
  }
};
</script>