"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
import RelationDelete from "@/services/magic-files/endpoints/relation/delete/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestRelationDelete() {
return (
<div>
<VsCodeEdittoolbar
filePath={
  "app/magic/services/magic-files/relation/delete/page.tsx"
}
/>
<RelationDelete />
</div>
);
}
    
