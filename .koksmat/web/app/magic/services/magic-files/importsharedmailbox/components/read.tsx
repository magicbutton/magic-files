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
import { ImportSharedMailboxItem } from "../applogic/model";

/* yankiebar */

export default function ReadImportSharedMailbox(props: { id: number }) {
  const { id } = props;
  const [transactionId, settransactionId] = useState(0);
  const readResult = useService<ImportSharedMailboxItem>(
    "magic-files.importsharedmailbox",
    ["read", id?.toString()],
    "",
    6000,
    transactionId.toString()
  );
  const importsharedmailbox = readResult.data;
  return (
    <div>
      {importsharedmailbox && (
        <div>
          <div>
            <div className="font-bold">tenant</div>
            <div>{importsharedmailbox.tenant}</div>
          </div>{" "}
          <div>
            <div className="font-bold">name</div>
            <div>{importsharedmailbox.name}</div>
          </div>{" "}
          <div>
            <div className="font-bold">description</div>
            <div>{importsharedmailbox.description}</div>
          </div>{" "}
          <div>
            <div className="font-bold">url</div>
            <div>{importsharedmailbox.url}</div>
          </div>{" "}
          <div>
            <div className="font-bold">email</div>
            <div>{importsharedmailbox.email}</div>
          </div>{" "}
          <div>
            <div className="font-bold">owners</div>
            <div>{importsharedmailbox.owners}</div>
          </div>
          <div>
            <div>
              <div className="font-bold">id</div>
              <div>{importsharedmailbox.id}</div>
            </div>
            <div>
              <div className="font-bold">created_at</div>
              <div>{importsharedmailbox.created_at}</div>
            </div>
            <div>
              <div className="font-bold">created_by</div>
              <div>{importsharedmailbox.created_by}</div>
            </div>
            <div>
              <div className="font-bold">updated_at</div>
              <div>{importsharedmailbox.updated_at}</div>
            </div>
            <div>
              <div className="font-bold">updated_by</div>
              <div>{importsharedmailbox.updated_by}</div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
