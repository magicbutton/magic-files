/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
"use client";
// piratos
import { useService } from "@/app/koksmat/useservice";
import { useState } from "react";
import { PlatformCategoryForm } from "./form";

import { PlatformCategoryItem } from "../applogic/model";
export default function UpdatePlatformCategory(props: { id: number }) {
  const { id } = props;

  const [transactionId, settransactionId] = useState(0);
  const readResult = useService<PlatformCategoryItem>(
    "magic-files.platformcategory",
    ["read", id?.toString()],
    "",
    6000,
    transactionId.toString()
  );
  const platformcategory = readResult.data;
  return (
    <div>
      {platformcategory && (
        <PlatformCategoryForm
          platformcategory={platformcategory}
          editmode="update"
        />
      )}
    </div>
  );
}
