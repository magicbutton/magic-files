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
import {ImportSharePointSiteItem} from "../applogic/model";


/* yankiebar */

export default function ReadImportSharePointSite(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<ImportSharePointSiteItem>(
      "magic-files.importsharepointsite",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const importsharepointsite = readResult.data;
    return (
      <div>
          
    {importsharepointsite && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{importsharepointsite.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{importsharepointsite.description}</div>
    </div>    <div>
        <div className="font-bold" >isteamsconnected</div>
        <div>{importsharepointsite.isteamsconnected}</div>
    </div>    <div>
        <div className="font-bold" >localeid</div>
        <div>{importsharepointsite.localeid}</div>
    </div>    <div>
        <div className="font-bold" >hubsiteid</div>
        <div>{importsharepointsite.hubsiteid}</div>
    </div>    <div>
        <div className="font-bold" >status</div>
        <div>{importsharepointsite.status}</div>
    </div>    <div>
        <div className="font-bold" >url</div>
        <div>{importsharepointsite.url}</div>
    </div>    <div>
        <div className="font-bold" >title</div>
        <div>{importsharepointsite.title}</div>
    </div>    <div>
        <div className="font-bold" >lastcontentmodified</div>
        <div>{importsharepointsite.lastcontentmodified}</div>
    </div>    <div>
        <div className="font-bold" >releatedgroupid</div>
        <div>{importsharepointsite.releatedgroupid}</div>
    </div>    <div>
        <div className="font-bold" >groupid</div>
        <div>{importsharepointsite.groupid}</div>
    </div>    <div>
        <div className="font-bold" >isteamschannelconnected</div>
        <div>{importsharepointsite.isteamschannelconnected}</div>
    </div>    <div>
        <div className="font-bold" >ishubsite</div>
        <div>{importsharepointsite.ishubsite}</div>
    </div>    <div>
        <div className="font-bold" >sharingcapability</div>
        <div>{importsharepointsite.sharingcapability}</div>
    </div>    <div>
        <div className="font-bold" >owner</div>
        <div>{importsharepointsite.owner}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{importsharepointsite.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{importsharepointsite.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{importsharepointsite.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{importsharepointsite.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{importsharepointsite.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
