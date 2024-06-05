    
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
// ImportSharePointSite
export interface ImportSharePointSiteItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    isteamsconnected : boolean ;
    localeid : string ;
    hubsiteid : string ;
    status : string ;
    url : string ;
    title : string ;
    lastcontentmodified : string ;
    releatedgroupid : string ;
    groupid : string ;
    isteamschannelconnected : boolean ;
    ishubsite : boolean ;
    sharingcapability : string ;
    owner : string ;

}


// ImportSharePointSite
export const ImportSharePointSiteSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    isteamsconnected : z.boolean().optional(), 
    localeid : z.string().optional(), 
    hubsiteid : z.string().optional(), 
    status : z.string().optional(), 
    url : z.string().optional(), 
    title : z.string().optional(), 
    lastcontentmodified : z.string().optional(), 
    releatedgroupid : z.string().optional(), 
    groupid : z.string().optional(), 
    isteamschannelconnected : z.boolean().optional(), 
    ishubsite : z.boolean().optional(), 
    sharingcapability : z.string().optional(), 
    owner : z.string().optional(), 

});

