/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v2
package importsharedmailboxmodel
import (
	"encoding/json"
	"time"
    // 
)

func UnmarshalImportSharedMailbox(data []byte) (ImportSharedMailbox, error) {
	var r ImportSharedMailbox
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ImportSharedMailbox) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ImportSharedMailbox struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Tenant string `json:"tenant"`
    Name string `json:"name"`
    Description string `json:"description"`
    Url string `json:"url"`
    Email string `json:"email"`
    Owners string `json:"owners"`

}

