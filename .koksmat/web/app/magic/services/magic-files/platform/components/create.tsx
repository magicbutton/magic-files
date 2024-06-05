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
import { PlatformForm } from "./form";

import { PlatformItem } from "../applogic/model";
export default function CreatePlatform() {
  const platform = {} as PlatformItem;
  return (
    <div>
      {platform && <PlatformForm platform={platform} editmode="create" />}
    </div>
  );
}