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
import { ImportGroupsItem } from "../applogic/model";

/* yankiebar */

export default function ReadImportGroups(props: { id: number }) {
  const { id } = props;
  const [transactionId, settransactionId] = useState(0);
  const readResult = useService<ImportGroupsItem>(
    "magic-files.importgroups",
    ["read", id?.toString()],
    "",
    6000,
    transactionId.toString()
  );
  const importgroups = readResult.data;
  return (
    <div>
      {importgroups && (
        <div>
          <div>
            <div className="font-bold">tenant</div>
            <div>{importgroups.tenant}</div>
          </div>{" "}
          <div>
            <div className="font-bold">name</div>
            <div>{importgroups.name}</div>
          </div>{" "}
          <div>
            <div className="font-bold">description</div>
            <div>{importgroups.description}</div>
          </div>{" "}
          <div>
            <div className="font-bold">url</div>
            <div>{importgroups.url}</div>
          </div>{" "}
          <div>
            <div className="font-bold">groupid</div>
            <div>{importgroups.groupid}</div>
          </div>{" "}
          <div>
            <div className="font-bold">type</div>
            <div>{importgroups.type}</div>
          </div>
          <div>
            <div>
              <div className="font-bold">id</div>
              <div>{importgroups.id}</div>
            </div>
            <div>
              <div className="font-bold">created_at</div>
              <div>{importgroups.created_at}</div>
            </div>
            <div>
              <div className="font-bold">created_by</div>
              <div>{importgroups.created_by}</div>
            </div>
            <div>
              <div className="font-bold">updated_at</div>
              <div>{importgroups.updated_at}</div>
            </div>
            <div>
              <div className="font-bold">updated_by</div>
              <div>{importgroups.updated_by}</div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
