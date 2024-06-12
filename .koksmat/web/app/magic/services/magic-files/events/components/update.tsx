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
import {EventsForm} from "./form";

import {EventsItem} from "../applogic/model";
export default function UpdateEvents(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<EventsItem>(
      "magic-files.events",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const events = readResult.data;
    return (
      <div>{events && 
      <EventsForm events={events} editmode="update"/>}
     
      </div>
    );
}
