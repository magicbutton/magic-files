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
import {BatchForm} from "./form";

import {BatchItem} from "../applogic/model";
export default function UpdateBatch(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<BatchItem>(
      "magic-files.batch",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const batch = readResult.data;
    return (
      <div>{batch && 
      <BatchForm batch={batch} editmode="update"/>}
     
      </div>
    );
}
