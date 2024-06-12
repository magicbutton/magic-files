    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
    "use client";
    import { useService } from "@/app/koksmat/useservice";
    import { useState } from "react";
    import {FileForm} from "./form";
    
    import {FileItem} from "../applogic/model";
    export default function CreateFile() {
       
        const file = {} as FileItem;
        return (
          <div>{file && 
          <FileForm file={file} editmode="create"/>}
         
          </div>
        );
    }
