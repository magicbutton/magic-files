    
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
// Job
export interface JobItem  {
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
    prefix : string ;
    type : string ;
    connection_id : number ;
    data : object ;

}


// Job
export const JobSchema = z.object({  
   
        tenant : z.string(), 
    name : z.string(), 
    description : z.string().optional(), 
    url : z.string(), 
    status : z.string(), 
    prefix : z.string(), 
    type : z.string(), 
    connection_id : z.number(), 
    data : z.object({}).optional(), 

});

