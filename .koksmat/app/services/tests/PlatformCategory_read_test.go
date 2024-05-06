    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
    //generator:  noma4.1
    package tests
    import (
        "testing"
        "github.com/magicbutton/magic-files/services/endpoints/platformcategory"
        
        "github.com/stretchr/testify/assert"
    )
    
    func TestPlatformCategoryread(t *testing.T) {
                    
            result,err := platformcategory.PlatformCategoryRead(-1)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    