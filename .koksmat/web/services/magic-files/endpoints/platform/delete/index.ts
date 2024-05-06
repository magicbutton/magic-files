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
import { Platform } from "@/services/magic-files/models/platform";
/**
 * Delete an existing item
 * 
 * id - The id of the item

 * @returns
 * 
 * Platform
 */
export default async function call(transactionId: string, id: int) {
  console.log("magic-files.platform", "delete");

  return run<Platform>(
    "magic-files.platform",
    ["delete", id],
    transactionId,
    600,
    transactionId
  );
}
