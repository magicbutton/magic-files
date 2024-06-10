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
    import {ImportRelationsForm} from "./form";
    
    import {ImportRelationsItem} from "../applogic/model";
    export default function CreateImportRelations() {
       
        const importrelations = {} as ImportRelationsItem;
        return (
          <div>{importrelations && 
          <ImportRelationsForm importrelations={importrelations} editmode="create"/>}
         
          </div>
        );
    }
