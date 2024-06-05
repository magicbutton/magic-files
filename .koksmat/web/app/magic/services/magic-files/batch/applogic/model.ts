    
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
// Batch
export interface BatchItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        tenant : string ;
    name : string ;
    description : string ;
    url : string ;
    status : string ;
    scheduledstart : string ;
    type : string ;
    job_id : number ;

}


// Batch
export const BatchSchema = z.object({  
   
        tenant : z.string(), 
    name : z.string(), 
    description : z.string().optional(), 
    url : z.string(), 
    status : z.string(), 
    scheduledstart : z.string(), 
    type : z.string(), 
    job_id : z.number(), 

});

