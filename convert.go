package main

import (
    "encoding/base64"
    "fmt"
    "github.com/google/uuid"
)

func ConvertToBase64(uuidStr string) (string, error) {
    uuidParsed, err := uuid.Parse(uuidStr)

    if err != nil {
        return "", fmt.Errorf("failed to convert %q to base64: %v", uuidStr, err)
    }

    return base64.StdEncoding.EncodeToString(uuidParsed[:]), nil
}
