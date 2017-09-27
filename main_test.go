package main


import (
  "testing"
  "os"
)

func TestKey(t *testing.T){
  skey := os.Getenv("OMISE_SKEY")
  pkey := os.Getenv("OMISE_PKEY")
  if pkey == "" {
    t.Error("Test failed")
  }

  if skey ==""{
    t.Error("Test failed")
  }

}
