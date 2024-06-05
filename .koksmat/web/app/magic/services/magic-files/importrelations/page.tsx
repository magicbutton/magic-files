
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
import SearchImportRelations from "@/app/magic/services/magic-files/importrelations/components/search";
import CreateImportRelations from "@/app/magic/services/magic-files/importrelations/components/create";
import {useState} from "react";

export default function ImportRelations() {
    const [isCreating, setisCreating] = useState(false);
return (
<div>
<div>
<Button variant="secondary" onClick={() => setisCreating(true)}>Create</Button>
</div>
<SearchImportRelations />
<Sheet open={isCreating} onOpenChange={setisCreating}>
<SheetContent>
  <SheetHeader>
    <SheetTitle>Create ImportRelations</SheetTitle>
    <SheetDescription>
      <CreateImportRelations  />
    </SheetDescription>
  </SheetHeader>
</SheetContent>
</Sheet>
</div>
);
}
    
