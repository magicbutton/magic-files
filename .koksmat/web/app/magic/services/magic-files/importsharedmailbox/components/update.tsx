/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
"use client";
// piratos
import { useService } from "@/koksmat/useservice";
import { useState } from "react";
import {ImportSharedMailboxForm} from "./form";

import {ImportSharedMailboxItem} from "../applogic/model";
export default function UpdateImportSharedMailbox(props: { id: number }) {
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
      <div>{importsharedmailbox && 
      <ImportSharedMailboxForm importsharedmailbox={importsharedmailbox} editmode="update"/>}
     
      </div>
    );
}
