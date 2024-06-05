
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
import SearchBatch from "@/app/magic/services/magic-files/batch/components/search";
import CreateBatch from "@/app/magic/services/magic-files/batch/components/create";
import {useState} from "react";

export default function Batch() {
    const [isCreating, setisCreating] = useState(false);
return (
<div>
<div>
<Button variant="secondary" onClick={() => setisCreating(true)}>Create</Button>
</div>
<SearchBatch />
<Sheet open={isCreating} onOpenChange={setisCreating}>
<SheetContent>
  <SheetHeader>
    <SheetTitle>Create Batch</SheetTitle>
    <SheetDescription>
      <CreateBatch  />
    </SheetDescription>
  </SheetHeader>
</SheetContent>
</Sheet>
</div>
);
}
    
