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
import {BatchItem} from "../applogic/model";


/* yankiebar */

export default function ReadBatch(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<BatchItem>(
      "magic-files.batch",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const batch = readResult.data;
    return (
      <div>
          
    {batch && <div>
        <div>
        <div className="font-bold" >tenant</div>
        <div>{batch.tenant}</div>
    </div>    <div>
        <div className="font-bold" >name</div>
        <div>{batch.name}</div>
    </div>    <div>
        <div className="font-bold" >description</div>
        <div>{batch.description}</div>
    </div>    <div>
        <div className="font-bold" >url</div>
        <div>{batch.url}</div>
    </div>    <div>
        <div className="font-bold" >status</div>
        <div>{batch.status}</div>
    </div>    <div>
        <div className="font-bold" >scheduledstart</div>
        <div>{batch.scheduledstart}</div>
    </div>    <div>
        <div className="font-bold" >type</div>
        <div>{batch.type}</div>
    </div>    <div>
        <div className="font-bold" >job</div>
        <div>{batch.job_id}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{batch.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{batch.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{batch.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{batch.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{batch.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
