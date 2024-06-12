    
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/       
"use client";
import { z } from "zod";
// spunk
// Events
export interface EventsItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        tenant : string ;
    name : string ;
    description : string ;
    url : string ;
    type : string ;
    eventDate : string ;
    user_id : number ;
    operation : string ;
    batchname : string ;
    file_id : number ;

}


// Events
export const EventsSchema = z.object({  
   
        tenant : z.string(), 
    name : z.string(), 
    description : z.string().optional(), 
    url : z.string(), 
    type : z.string(), 
    eventDate : z.string(), 
    user_id : z.number(), 
    operation : z.string(), 
    batchname : z.string().optional(), 
    file_id : z.number(), 

});

