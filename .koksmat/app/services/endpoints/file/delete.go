            /*
            File have been automatically created. To prevent the file from getting overwritten
            set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
            ---
            keep: false
            ---
            */
            //generator:  noma3.delete.v2
            package file
            
            import (
                "log"
                "strconv"
                "github.com/magicbutton/magic-files/applogic"
                "github.com/magicbutton/magic-files/database"
                "github.com/magicbutton/magic-files/services/models/filemodel"
            
            )
            
            func FileDelete(arg0 string) ( error) {
                id,_ := strconv.Atoi(arg0)
                log.Println("Calling Filedelete")
            
                return applogic.Delete[database.File, filemodel.File](id)
            
            }
