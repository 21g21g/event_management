import axios from 'axios';
export const uploadImageToCloudinary=async(file)=>{

    const CLOUDINARY_NAME="dnfl1adwg";
    const UPLOAD_PREIST="upload_profile";
    const url=`https://api.cloudinary.com/v1_1/${CLOUDINARY_NAME}/image/upload`;

    const formData = new FormData();
  formData.append('file', file);
  formData.append('upload_preset', UPLOAD_PREIST);

  try {
    const response = await axios.post(url, formData);
    return response.data.secure_url; 
  } catch (error) {
    console.error('Error uploading image:', error);
  }
};

