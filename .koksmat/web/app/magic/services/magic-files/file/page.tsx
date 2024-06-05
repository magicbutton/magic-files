
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
import SearchFile from "@/app/magic/services/magic-files/file/components/search";
import CreateFile from "@/app/magic/services/magic-files/file/components/create";
import {useState} from "react";

export default function File() {
    const [isCreating, setisCreating] = useState(false);
return (
<div>
<div>
<Button variant="secondary" onClick={() => setisCreating(true)}>Create</Button>
</div>
<SearchFile />
<Sheet open={isCreating} onOpenChange={setisCreating}>
<SheetContent>
  <SheetHeader>
    <SheetTitle>Create File</SheetTitle>
    <SheetDescription>
      <CreateFile  />
    </SheetDescription>
  </SheetHeader>
</SheetContent>
</Sheet>
</div>
);
}
    
