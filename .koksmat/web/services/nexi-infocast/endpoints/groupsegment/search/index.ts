"use server";
/*
Parameters

*/
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
import { run } from "@/koksmat/magicservices";
import { ShowCodeFragment } from "@/services/ShowCodeFragment";
import { Groupsegment } from "@/services/magic-files/models/groupsegment";
/**
 * Search for items
 * 
 * query - The search query

 * @returns
 * 
 * Groupsegment
 */
export default async function call(transactionId: string, query: string) {
  console.log("magic-files.groupsegment", "search");

  return run<Groupsegment>(
    "magic-files.groupsegment",
    ["search", query],
    transactionId,
    600,
    transactionId
  );
}
