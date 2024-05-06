"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
import PlatformDelete from "@/services/magic-files/endpoints/platform/delete/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestPlatformDelete() {
return (
<div>
<VsCodeEdittoolbar
filePath={
  "app/magic/services/magic-files/platform/delete/page.tsx"
}
/>
<PlatformDelete />
</div>
);
}
    
