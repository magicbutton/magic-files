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
    import {ImportSharePointSiteForm} from "./form";
    
    import {ImportSharePointSiteItem} from "../applogic/model";
    export default function CreateImportSharePointSite() {
       
        const importsharepointsite = {} as ImportSharePointSiteItem;
        return (
          <div>{importsharepointsite && 
          <ImportSharePointSiteForm importsharepointsite={importsharepointsite} editmode="create"/>}
         
          </div>
        );
    }
