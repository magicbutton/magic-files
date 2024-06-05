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
        "github.com/magicbutton/magic-files/services/endpoints/importsharedmailbox"
                    "github.com/magicbutton/magic-files/services/models/importsharedmailboxmodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestImportSharedMailboxupdate(t *testing.T) {
                                // transformer v1
            object := importsharedmailboxmodel.ImportSharedMailbox{}
         
            result,err := importsharedmailbox.ImportSharedMailboxUpdate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
