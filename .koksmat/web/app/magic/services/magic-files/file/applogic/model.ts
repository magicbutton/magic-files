    
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
// File
export interface FileItem  {
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
    size : number ;
    extension : string ;
    createdby : string ;
    created : string ;
    modified : string ;
    modifiedby : string ;
    version : string ;
    published : boolean ;
    level1 : string ;
    level2 : string ;
    level3 : string ;
    level4 : string ;
    level5 : string ;
    level6 : string ;
    level7 : string ;
    level8 : string ;
    level9 : string ;
    level10 : string ;
    level11 : string ;
    level12 : string ;
    level13 : string ;
    level14 : string ;
    level15 : string ;
    level16 : string ;

}


// File
export const FileSchema = z.object({  
   
        tenant : z.string(), 
    name : z.string(), 
    description : z.string().optional(), 
    url : z.string(), 
    type : z.string(), 
    size : z.number(), 
    extension : z.string(), 
    createdby : z.string().optional(), 
    created : z.string().optional(), 
    modified : z.string().optional(), 
    modifiedby : z.string().optional(), 
    version : z.string().optional(), 
    published : z.boolean().optional(), 
    level1 : z.string().optional(), 
    level2 : z.string().optional(), 
    level3 : z.string().optional(), 
    level4 : z.string().optional(), 
    level5 : z.string().optional(), 
    level6 : z.string().optional(), 
    level7 : z.string().optional(), 
    level8 : z.string().optional(), 
    level9 : z.string().optional(), 
    level10 : z.string().optional(), 
    level11 : z.string().optional(), 
    level12 : z.string().optional(), 
    level13 : z.string().optional(), 
    level14 : z.string().optional(), 
    level15 : z.string().optional(), 
    level16 : z.string().optional(), 

});

