import gql from "graphql-tag"



export const REGISTER_USER_MUTATION = gql`
  mutation Register($email: String!, $password: String!, $username: String!) {
    signUp(email: $email, password: $password, username: $username) {
      message
    }
  }
`;
export const LOGIN_MUTATION = gql`
mutation loginUser($email:String!,$password:String!){
  siginUser(email:$email,password:$password){
    token
  }
}
`;

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
  $tags: [String],
 
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

query getEvent($search:String!,$limit:Int!,$offset:Int!){
  events(where:{_or:[
    {title:{_ilike:$search}},
    {address:{_ilike:$search}},
    {venue:{_ilike:$search}},
     {description:{_ilike:$search}},
     {category:{_ilike:$search}},
  ]},
    limit:$limit,
    offset:$offset
  )
  {
    id
    title
    description
    address
    venue
    category
    price
    specific_price
    featured_image
    preparation_date
    tags
  }
}
`;
export const GET_ALL_EVENTS_WTHOUT_FILTER=gql`

query{
  events{
    id
    title
    description
    address
    venue
    price
    specific_price
    preparation_date
    featured_image
    category
    tags
    user{
      username
    }
    
    
  }
}
`;
export const GET_EVENT_BY_CATEGORY=gql`
query getEvent($category:String!){
  events(where:{category:{_eq:$category}}){
    id
    title
    description
    price
    address
    venue
    specific_price
    featured_image
    category
    tags
    preparation_date
  }
}`;

export const GET_EVENT_BY_ID=gql`
query getEventById($id:uuid!){
  events_by_pk(id:$id){
      id
    title
    description
    address
    venue
    price
    specific_price
    featured_image
    imagestores{
      id
      url
      event_id
    }
    category
    preparation_date
    tags
    
  }
}`;

export const USER_MAKE_BOOK_MARK=gql`
mutation bookMark($user_id: uuid!, $event_id: uuid!,$isbookMarked:Boolean!) {
  insert_bookmarks_one(object: {user_id:$user_id,event_id:$event_id,isbookMarked:$isbookMarked}){
    id
    
    
  }
}

`;

export const GET_BOOK_MARK_BY_USER_ID=gql`
query getBookMark($user_id:uuid!){
  bookmarks(where:{user_id:{_eq:$user_id}}){
    id
    event{
      id
      title
      description
      address
      venue
      specific_price
      tags
      category
      featured_image
      price
      preparation_date
      
    }
    
  }
}`;

export const SEARCH_TERMS=gql`
query searchItems($search: String, $category: String!, $limit: Int) {
  events_aggregate(where: {
    _and: [
      { category: { _eq: $category } },
      {
        _or: [
          { title: { _ilike: $search } },
          { description: { _ilike: $search } },
          { address: { _ilike: $search } },
          { venue: { _ilike: $search } }
        ]
      }
    ]
  }) {
    aggregate {
      count
    }
  }
  
  events(where: {
    _and: [
      { category: { _eq: $category } },
      {
        _or: [
          { title: { _ilike: $search } },
          { description: { _ilike: $search } },
          { address: { _ilike: $search } },
          { venue: { _ilike: $search } }
        ]
      }
    ]
  }, limit: $limit) {
    id
    title
    category
    description
    address
    venue
    tags
    preparation_date
    price
    specific_price
    featured_image

  }
}
`;

export const GET_BOOK_MARKED_EVENT=gql`
query getBookmark($user_id:uuid!,$event_id:uuid!){
  bookmarks(where:{user_id:{_eq:$user_id},event_id:{_eq:$event_id}}){
    isbookMarked
  }
    
  }
`;

export const CATCH_TICKET=gql`
mutation insertTicket($user_id:uuid!,$event_id:uuid!,$quantity:Int!,$catchedTicket:Boolean!){
  insert_tickets_one(object:{user_id:$user_id,event_id:$event_id,quantity:$quantity,catchedTicket:$catchedTicket}){
     id
  
  }
}
`;

export const GET_TICKET_USER=gql`
query getTicket($user_id:uuid!){
  tickets(where:{user_id:{_eq:$user_id}}){
    id
    catchedTicket
    user_id
    quantity
    event{
      title
      preparation_date
      
      
    }
  }
  
}
`;

export const REMOVE_BOOKMARK_FROM_FAVOURITE=gql`
mutation deleteBookmark($id: uuid!) {
  delete_bookmarks_by_pk(id: $id){
    id
  }
}


`;