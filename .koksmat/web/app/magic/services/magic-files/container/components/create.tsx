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
    import {ContainerForm} from "./form";
    
    import {ContainerItem} from "../applogic/model";
    export default function CreateContainer() {
       
        const container = {} as ContainerItem;
        return (
          <div>{container && 
          <ContainerForm container={container} editmode="create"/>}
         
          </div>
        );
    }
