package main

import (
    "fmt"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
    // Initialize a session in us-west-2 that the SDK will use to load
    // credentials from the shared credentials file ~/.aws/credentials.
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-east-1")},
    )

    // Create DynamoDB client
    svc := dynamodb.New(sess)

    // Get the list of tables
    result, err := svc.ListTables(&dynamodb.ListTablesInput{})

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println("Tables:")
    fmt.Println("")

    for _, n := range result.TableNames {
        fmt.Println(*n)
    }

    fmt.Println("")
}
