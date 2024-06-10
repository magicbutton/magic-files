    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
    "use client";
    import { useService } from "@/koksmat/useservice";
    import { useState } from "react";
    import {BatchForm} from "./form";
    
    import {BatchItem} from "../applogic/model";
    export default function CreateBatch() {
       
        const batch = {} as BatchItem;
        return (
          <div>{batch && 
          <BatchForm batch={batch} editmode="create"/>}
         
          </div>
        );
    }
