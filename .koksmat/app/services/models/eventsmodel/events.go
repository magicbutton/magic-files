/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v2
package eventsmodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-files/database/databasetypes"
)

func UnmarshalEvents(data []byte) (Events, error) {
	var r Events
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Events) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Events struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Tenant string `json:"tenant"`
    Name string `json:"name"`
    Description string `json:"description"`
    Url string `json:"url"`
    Type string `json:"type"`
    Eventdate time.Time `json:"eventDate"`
    User_id int `json:"user_id"`
    Operation string `json:"operation"`
    Batchname string `json:"batchname"`
    File_id int `json:"file_id"`

}

