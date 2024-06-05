/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v1
package importrelationsmodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-files/database/databasetypes"
)

func UnmarshalImportRelations(data []byte) (ImportRelations, error) {
	var r ImportRelations
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ImportRelations) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ImportRelations struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Tenant string `json:"tenant"`
    Name string `json:"name"`
    Description string `json:"description"`
    Url string `json:"url"`
    User_id int `json:"user_id"`
    Type string `json:"type"`
    Importgroup_Id_id int `json:"importgroup_id_id"`

}

