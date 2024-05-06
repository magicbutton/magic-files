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
import { Relation } from "@/services/magic-files/models/relation";
/**
 * Read a single item
 * 
 * id - The id of the item

 * @returns
 * 
 * Relation
 */
export default async function call(transactionId: string, id: int) {
  console.log("magic-files.relation", "read");

  return run<Relation>(
    "magic-files.relation",
    ["read", id],
    transactionId,
    600,
    transactionId
  );
}
