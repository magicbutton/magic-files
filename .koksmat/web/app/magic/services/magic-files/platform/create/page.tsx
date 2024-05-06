"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
import PlatformCreate from "@/services/magic-files/endpoints/platform/create/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestPlatformCreate() {
return (
<div>
<VsCodeEdittoolbar
filePath={
  "app/magic/services/magic-files/platform/create/page.tsx"
}
/>
<PlatformCreate />
</div>
);
}
    
