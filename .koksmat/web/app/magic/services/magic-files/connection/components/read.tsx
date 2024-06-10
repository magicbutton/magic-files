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
import {ConnectionItem} from "../applogic/model";


/* yankiebar */

export default function ReadConnection(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<ConnectionItem>(
      "magic-files.connection",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const connection = readResult.data;
    return (
      <div>
          
    {connection && <div>
        <div>
        <div className="font-bold" >tenant</div>
        <div>{connection.tenant}</div>
    </div>    <div>
        <div className="font-bold" >name</div>
        <div>{connection.name}</div>
    </div>    <div>
        <div className="font-bold" >description</div>
        <div>{connection.description}</div>
    </div>    <div>
        <div className="font-bold" >url</div>
        <div>{connection.url}</div>
    </div>                <div>
                    <div className="font-bold" >data</div>
                    <div>{JSON.stringify(connection.data,null,2)}</div>
                </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{connection.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{connection.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{connection.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{connection.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{connection.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
