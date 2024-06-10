/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
"use client";
// piratos
import { useService } from "@/koksmat/useservice";
import { useState } from "react";
import {PlatformForm} from "./form";

import {PlatformItem} from "../applogic/model";
export default function UpdatePlatform(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<PlatformItem>(
      "magic-files.platform",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const platform = readResult.data;
    return (
      <div>{platform && 
      <PlatformForm platform={platform} editmode="update"/>}
     
      </div>
    );
}
