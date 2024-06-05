    
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
// Platform
export interface PlatformItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        tenant : string ;
    name : string ;
    description : string ;
    url : string ;
    category_id : number ;

}


// Platform
export const PlatformSchema = z.object({  
   
        tenant : z.string(), 
    name : z.string(), 
    description : z.string().optional(), 
    url : z.string(), 
    category_id : z.number(), 

});

