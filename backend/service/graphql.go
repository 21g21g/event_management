package service

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/machinebox/graphql"
	"golang.org/x/crypto/bcrypt"
)

var client = graphql.NewClient("http://graphql-engine:8080/v1/graphql")

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")

	}
	adminSecret := os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET")
	log.Printf("Admin secret loaded: %s\n", adminSecret)

}

func RunGraphQLQuery(ctx context.Context, userQuery *graphql.Request, userResp interface{}) error {

	userQuery.Header.Set("X-Hasura-Admin-Secret", os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET"))

	err := client.Run(ctx, userQuery, userResp)
	if err != nil {
		log.Println("GraphQL query failed:", err)
		return err
	}
	log.Printf("GraphQL query succeeded. Response: %+v\n", userResp)

	return nil
}

func UserExists(email string) (bool, error) {
	req := graphql.NewRequest(`
	query user($email:String!){
		users(where:{email:{_eq:$email}}){
	id
		}
	  }
    `)

	req.Var("email", email)

	var respData struct {
		Data struct {
			Users []struct {
				ID string `json:"id"`
			} `json:"users"`
		} `json:"data"`
	}

	ctx := context.Background()
	if err := RunGraphQLQuery(ctx, req, &respData); err != nil {
		return false, err
	}
	return len(respData.Data.Users) > 0, nil
}

func RegisterUser(username, email, password string) error {
	exists, err := UserExists(email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user already exists")

	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req := graphql.NewRequest(`
	mutation registerUser($email:String!,$username:String!,$password:String!){
		insert_users(objects:{email:$email,username:$username,password:$password}){
		  returning{
			id
		  }
		}
	  }
    `)
	req.Var("username", username)
	req.Var("email", email)
	req.Var("password", string(hashedPassword))

	var respData struct {
		Data struct {
			RegisterUser struct {
				Username string `json:"username"`
			} `json:"registerUser"`
		} `json:"data"`
	}

	ctx := context.Background()
	if err := RunGraphQLQuery(ctx, req, &respData); err != nil {
		return err
	}
	return nil
}
