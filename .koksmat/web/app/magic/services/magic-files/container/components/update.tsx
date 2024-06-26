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
import {ContainerForm} from "./form";

import {ContainerItem} from "../applogic/model";
export default function UpdateContainer(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<ContainerItem>(
      "magic-files.container",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const container = readResult.data;
    return (
      <div>{container && 
      <ContainerForm container={container} editmode="update"/>}
     
      </div>
    );
}
