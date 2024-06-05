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
import { FileForm } from "./form";

import { FileItem } from "../applogic/model";
export default function UpdateFile(props: { id: number }) {
  const { id } = props;

  const [transactionId, settransactionId] = useState(0);
  const readResult = useService<FileItem>(
    "magic-files.file",
    ["read", id?.toString()],
    "",
    6000,
    transactionId.toString()
  );
  const file = readResult.data;
  return <div>{file && <FileForm file={file} editmode="update" />}</div>;
}