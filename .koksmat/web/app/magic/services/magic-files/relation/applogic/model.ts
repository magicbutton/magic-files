    
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
// Relation
export interface RelationItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    user_id : number ;
    date : string ;
    type : string ;
    batchname : string ;

}


// Relation
export const RelationSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    user_id : z.number(), 
    date : z.string(), 
    type : z.string(), 
    batchname : z.string().optional(), 

});

