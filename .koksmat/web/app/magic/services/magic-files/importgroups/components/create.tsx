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
    import {ImportGroupsForm} from "./form";
    
    import {ImportGroupsItem} from "../applogic/model";
    export default function CreateImportGroups() {
       
        const importgroups = {} as ImportGroupsItem;
        return (
          <div>{importgroups && 
          <ImportGroupsForm importgroups={importgroups} editmode="create"/>}
         
          </div>
        );
    }
