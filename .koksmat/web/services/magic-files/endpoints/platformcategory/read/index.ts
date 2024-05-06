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
import { PlatformCategory } from "@/services/magic-files/models/platformcategory";
/**
 * Read a single item
 * 
 * id - The id of the item

 * @returns
 * 
 * PlatformCategory
 */
export default async function call(transactionId: string, id: int) {
  console.log("magic-files.platformcategory", "read");

  return run<PlatformCategory>(
    "magic-files.platformcategory",
    ["read", id],
    transactionId,
    600,
    transactionId
  );
}