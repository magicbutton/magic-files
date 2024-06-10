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
import {ImportRelationsItem} from "../applogic/model";


/* yankiebar */

export default function ReadImportRelations(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<ImportRelationsItem>(
      "magic-files.importrelations",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const importrelations = readResult.data;
    return (
      <div>
          
    {importrelations && <div>
        <div>
        <div className="font-bold" >tenant</div>
        <div>{importrelations.tenant}</div>
    </div>    <div>
        <div className="font-bold" >name</div>
        <div>{importrelations.name}</div>
    </div>    <div>
        <div className="font-bold" >description</div>
        <div>{importrelations.description}</div>
    </div>    <div>
        <div className="font-bold" >url</div>
        <div>{importrelations.url}</div>
    </div>    <div>
        <div className="font-bold" >user</div>
        <div>{importrelations.user_id}</div>
    </div>    <div>
        <div className="font-bold" >type</div>
        <div>{importrelations.type}</div>
    </div>    <div>
        <div className="font-bold" >importgroup_id</div>
        <div>{importrelations.importgroup_id_id}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{importrelations.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{importrelations.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{importrelations.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{importrelations.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{importrelations.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
