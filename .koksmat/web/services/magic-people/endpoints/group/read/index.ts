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
import { Group } from "@/services/magic-files/models/group";
/**
 * Read a single item
 * 
 * id - The id of the item

 * @returns
 * 
 * Group
 */
export default async function call(transactionId: string, id: int) {
  console.log("magic-files.group", "read");

  return run<Group>(
    "magic-files.group",
    ["read", id],
    transactionId,
    600,
    transactionId
  );
}
