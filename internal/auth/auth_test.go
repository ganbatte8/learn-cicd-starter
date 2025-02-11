
package auth

import (
  //"errors"
  "net/http"
  //"strings"
  "testing"
)

func TestGetAPIKey(t *testing.T) {

  type test struct {
    input http.Header
    want string
  }
  tests := []test{
    {input: http.Header{
        "Authorization": []string{"ApiKey Youhavetoacceptthefact,basically",},
      },
     want: "Youhavetoacceptthefact,basicall",
    },
  }  
  // var input []string = []string{"Authorization", "You have to accept the fact, basically"} 
  // var headers http.Header   // map[string][]string
  // headers.Add(input[0], input[1])
  

  for i := 0; i < len(tests); i++ {
    result, err := GetAPIKey(tests[i].input)
    if err != nil {
      t.Fatalf("GetAPIKey errored: %s\n", err)
    }
    if result != tests[i].want {
      t.Fatalf("Wrong output: got:%s expected:%s\n", result, tests[i].want)
    }
  }


  
}