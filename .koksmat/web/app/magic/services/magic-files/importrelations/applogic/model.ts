    
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
// ImportRelations
export interface ImportRelationsItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        tenant : string ;
    name : string ;
    description : string ;
    url : string ;
    user_id : number ;
    type : string ;
    importgroup_id_id : number ;

}


// ImportRelations
export const ImportRelationsSchema = z.object({  
   
        tenant : z.string(), 
    name : z.string(), 
    description : z.string().optional(), 
    url : z.string(), 
    user_id : z.number(), 
    type : z.string(), 
    importgroup_id_id : z.number(), 

});

