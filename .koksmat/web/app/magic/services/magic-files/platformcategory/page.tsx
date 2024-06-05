
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
import SearchPlatformCategory from "@/app/magic/services/magic-files/platformcategory/components/search";
import CreatePlatformCategory from "@/app/magic/services/magic-files/platformcategory/components/create";
import {useState} from "react";

export default function PlatformCategory() {
    const [isCreating, setisCreating] = useState(false);
return (
<div>
<div>
<Button variant="secondary" onClick={() => setisCreating(true)}>Create</Button>
</div>
<SearchPlatformCategory />
<Sheet open={isCreating} onOpenChange={setisCreating}>
<SheetContent>
  <SheetHeader>
    <SheetTitle>Create PlatformCategory</SheetTitle>
    <SheetDescription>
      <CreatePlatformCategory  />
    </SheetDescription>
  </SheetHeader>
</SheetContent>
</Sheet>
</div>
);
}
    
