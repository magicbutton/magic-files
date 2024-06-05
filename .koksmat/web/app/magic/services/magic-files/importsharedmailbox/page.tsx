
"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
/* dumle */
import {
    Sheet,
    SheetContent,
    SheetDescription,
    SheetHeader,
    SheetTitle,
    SheetTrigger,
  } from "@/components/ui/sheet";
import { Button } from "@/components/ui/button";
import SearchImportSharedMailbox from "@/app/magic/services/magic-files/importsharedmailbox/components/search";
import CreateImportSharedMailbox from "@/app/magic/services/magic-files/importsharedmailbox/components/create";
import {useState} from "react";

export default function ImportSharedMailbox() {
    const [isCreating, setisCreating] = useState(false);
return (
<div>
<div>
<Button variant="secondary" onClick={() => setisCreating(true)}>Create</Button>
</div>
<SearchImportSharedMailbox />
<Sheet open={isCreating} onOpenChange={setisCreating}>
<SheetContent>
  <SheetHeader>
    <SheetTitle>Create ImportSharedMailbox</SheetTitle>
    <SheetDescription>
      <CreateImportSharedMailbox  />
    </SheetDescription>
  </SheetHeader>
</SheetContent>
</Sheet>
</div>
);
}
    
