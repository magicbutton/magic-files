/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v1
package containermodel
import (
	"encoding/json"
	"time"
    // 
)

func UnmarshalContainer(data []byte) (Container, error) {
	var r Container
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Container) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Container struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
        Tenant string `json:"tenant"`
    Name string `json:"name"`
    Description string `json:"description"`
    Url string `json:"url"`
    Type string `json:"type"`
    Level1 string `json:"level1"`
    Level2 string `json:"level2"`
    Level3 string `json:"level3"`
    Level4 string `json:"level4"`
    Level5 string `json:"level5"`
    Level6 string `json:"level6"`
    Level7 string `json:"level7"`
    Level8 string `json:"level8"`
    Level9 string `json:"level9"`
    Level10 string `json:"level10"`
    Level11 string `json:"level11"`
    Level12 string `json:"level12"`
    Level13 string `json:"level13"`
    Level14 string `json:"level14"`
    Level15 string `json:"level15"`
    Level16 string `json:"level16"`

}

