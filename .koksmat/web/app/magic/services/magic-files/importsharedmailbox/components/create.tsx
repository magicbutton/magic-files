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
import { ImportSharedMailboxForm } from "./form";

import { ImportSharedMailboxItem } from "../applogic/model";
export default function CreateImportSharedMailbox() {
  const importsharedmailbox = {} as ImportSharedMailboxItem;
  return (
    <div>
      {importsharedmailbox && (
        <ImportSharedMailboxForm
          importsharedmailbox={importsharedmailbox}
          editmode="create"
        />
      )}
    </div>
  );
}
