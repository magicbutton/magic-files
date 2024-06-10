/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v2
package importsharepointsitemodel
import (
	"encoding/json"
	"time"
    // 
)

func UnmarshalImportSharePointSite(data []byte) (ImportSharePointSite, error) {
	var r ImportSharePointSite
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ImportSharePointSite) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ImportSharePointSite struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Isteamsconnected bool `json:"isteamsconnected"`
    Localeid string `json:"localeid"`
    Hubsiteid string `json:"hubsiteid"`
    Status string `json:"status"`
    Url string `json:"url"`
    Title string `json:"title"`
    Lastcontentmodified time.Time `json:"lastcontentmodified"`
    Releatedgroupid string `json:"releatedgroupid"`
    Groupid string `json:"groupid"`
    Isteamschannelconnected bool `json:"isteamschannelconnected"`
    Ishubsite bool `json:"ishubsite"`
    Sharingcapability string `json:"sharingcapability"`
    Owner string `json:"owner"`

}

