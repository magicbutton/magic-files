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
import {PlatformItem} from "../applogic/model";


/* yankiebar */

export default function ReadPlatform(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<PlatformItem>(
      "magic-files.platform",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const platform = readResult.data;
    return (
      <div>
          
    {platform && <div>
        <div>
        <div className="font-bold" >tenant</div>
        <div>{platform.tenant}</div>
    </div>    <div>
        <div className="font-bold" >name</div>
        <div>{platform.name}</div>
    </div>    <div>
        <div className="font-bold" >description</div>
        <div>{platform.description}</div>
    </div>    <div>
        <div className="font-bold" >url</div>
        <div>{platform.url}</div>
    </div>    <div>
        <div className="font-bold" >category</div>
        <div>{platform.category_id}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{platform.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{platform.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{platform.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{platform.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{platform.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
