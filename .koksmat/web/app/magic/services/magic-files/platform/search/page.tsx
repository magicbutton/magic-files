"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
import PlatformSearch from "@/services/magic-files/endpoints/platform/search/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestPlatformSearch() {
return (
<div>
<VsCodeEdittoolbar
filePath={
  "app/magic/services/magic-files/platform/search/page.tsx"
}
/>
<PlatformSearch />
</div>
);
}
    
