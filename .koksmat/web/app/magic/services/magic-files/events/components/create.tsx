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
    import {EventsForm} from "./form";
    
    import {EventsItem} from "../applogic/model";
    export default function CreateEvents() {
       
        const events = {} as EventsItem;
        return (
          <div>{events && 
          <EventsForm events={events} editmode="create"/>}
         
          </div>
        );
    }
