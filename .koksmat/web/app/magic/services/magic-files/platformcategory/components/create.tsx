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
import { PlatformCategoryForm } from "./form";

import { PlatformCategoryItem } from "../applogic/model";
export default function CreatePlatformCategory() {
  const platformcategory = {} as PlatformCategoryItem;
  return (
    <div>
      {platformcategory && (
        <PlatformCategoryForm
          platformcategory={platformcategory}
          editmode="create"
        />
      )}
    </div>
  );
}
