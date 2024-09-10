import gql from "graphql-tag"


  export const REGISTER_USER_MUTATION = gql`
  mutation Register($email:String!,$password:String!,$username:String!){
  registerUser(input:{email:$email,password:$password,username:$username}){
    username
  }
}
`;
export const LOGIN_MUTATION = gql`
 mutation loginUser($email: String!, $password: String!) {
   loginUser(input: {email: $email, password: $password}) {
     token
   }
 }
`
export const GET_USER_BY_ID = gql`
query user($id:uuid!){
  users_by_pk(id:$id){
    id
    username
    events{
      address
      category
      description
      featured_image
      specific_price
      title
      price
      imagestores{
        event_id
      }
    }
  }
}
`;


export const upload_image_action = gql`
  mutation UploadBase64Image($base64_str: String!) {
    uploadBase64Image(base64_str: $base64_str) {
      url
    }
  }
`;


export const insert_event = gql`
  mutation Event(
    $title: String!,
    $description: String!,
    $venue: String!,
    $address: String!,
    $price: String!,
    $specific_price: numeric!,
    $preparation_date: date!,
    $category: String!,
    $featured_image: String!,
    $tags: [String!],
    $user_id: uuid!
  ) {
    insert_events(
      objects: {
        title: $title,
        description: $description,
        venue: $venue,
        address: $address,
        price: $price,
        specific_price: $specific_price,
        preparation_date: $preparation_date,
        category: $category,
        tags: $tags,
        featured_image: $featured_image,
        user_id: $user_id
      }
    ) {
      returning {
        id
      }
    }
  }
`;


export const insert_image_imageTable = gql`
  mutation insertImage($url: String!, $event_id: uuid!) {
    insert_imagestore(objects: { url: $url, event_id: $event_id }) {
      returning {
        url
      }
    }
  }
`;

export const ADD_BOOK_MARK=gql`
mutation Bookmark($user_id: uuid!, $event_id: uuid!) {
  insert_bookmarks(objects: {user_id:$user_id,event_id:$event_id}){
    returning{
      id
      user_id
      event_id
    }
    
  }
}`;
export const REMOVE_BOOKMARK_MUTATION = gql`
  mutation RemoveBookmark($user_id: uuid!, $event_id: uuid!) {
    delete_bookmarks(where: {user_id: {_eq: $user_id}, event_id: {_eq: $event_id}}) {
      returning {
        id
      }
    }
  }
`;

export const Update_Event_ById=gql`
mutation EventUpdate(
  $id: uuid!,
  $title: String,
  $description: String,
  $venue: String,
  $address: String,
  $price: String,
  $specificPrice: numeric,
  $preparationDate: date,
  $category: String,
  $tags: [String!],
 
) {
  update_events(
    where: { id: { _eq: $id } }
    _set: {
      title: $title,
      description: $description,
      venue: $venue,
      address: $address,
      price: $price,
      specific_price: $specificPrice,
      preparation_date: $preparationDate,
      category: $category,
      tags: $tags,
    }
  ) {
    returning {
      id
      title
      description
      venue
      address
      price
      specific_price
      preparation_date
      category
      tags
      featured_image
    }
  }
}
`;

export const GET_USER_BY_HIS_ID=gql`
query userbyId($id:uuid!){
  users_by_pk(id:$id){
    username
    events{
      id
      title
      description
      address
      venue
      price
      specific_price
      preparation_date
      category
      tags
      featured_image
    }
  }
  
}`;


export const DELETE_EVENT_BY_ID=gql`
mutation DeleteEvent($id:uuid!){
  delete_events_by_pk(id:$id){
    id
  }
}`;
export const insert_event_mutation = gql`
mutation Event(
  $title: String!,
  $description: String!,
  $venue: String!,
  $address: String!,
  $price: String!,
  $specific_price: numeric!,
  $preparation_date: date!,
  $category: String!,
  $featured_image: String!,
  $tags: [String!],
  $user_id: uuid!
) {
  insert_events(
    objects: {
      title: $title,
      description: $description,
      venue: $venue,
      address: $address,
      price: $price,
      specific_price: $specific_price,
      preparation_date: $preparation_date,
      category: $category,
      featured_image: $featured_image,
      tags: $tags,
      user_id: $user_id
    }
  ) {
    returning {
      id
    }
  }
}
`;

export const GET_ALL_EVENTS=gql`

query{
  events{
    id
    title
    description
    address
    venue
    price
    specific_price
    featured_image
    category
    tags
    user{
      username
    }
    
    
  }
}
`;