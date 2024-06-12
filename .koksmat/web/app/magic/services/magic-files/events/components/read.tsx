    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
"use client";
import { useService } from "@/app/koksmat/useservice";
import { useState } from "react";
import {EventsItem} from "../applogic/model";


/* yankiebar */

export default function ReadEvents(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<EventsItem>(
      "magic-files.events",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const events = readResult.data;
    return (
      <div>
          
    {events && <div>
        <div>
        <div className="font-bold" >tenant</div>
        <div>{events.tenant}</div>
    </div>    <div>
        <div className="font-bold" >name</div>
        <div>{events.name}</div>
    </div>    <div>
        <div className="font-bold" >description</div>
        <div>{events.description}</div>
    </div>    <div>
        <div className="font-bold" >url</div>
        <div>{events.url}</div>
    </div>    <div>
        <div className="font-bold" >type</div>
        <div>{events.type}</div>
    </div>    <div>
        <div className="font-bold" >eventDate</div>
        <div>{events.eventDate}</div>
    </div>    <div>
        <div className="font-bold" >user</div>
        <div>{events.user_id}</div>
    </div>    <div>
        <div className="font-bold" >operation</div>
        <div>{events.operation}</div>
    </div>    <div>
        <div className="font-bold" >batchname</div>
        <div>{events.batchname}</div>
    </div>    <div>
        <div className="font-bold" >file</div>
        <div>{events.file_id}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{events.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{events.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{events.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{events.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{events.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
