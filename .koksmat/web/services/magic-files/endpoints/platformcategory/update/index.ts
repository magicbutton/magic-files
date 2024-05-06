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
 * Update an existing item
 * 
 * item - The item to update

 * @returns
 * 
 * PlatformCategory
 */
export default async function call(transactionId: string ,item: PlatformCategory) {
  console.log( "magic-files.platformcategory", "update");

  return run<PlatformCategory>(
    "magic-files.platformcategory",
    ["update" , JSON.stringify(item)],
    transactionId,
    600,
    transactionId
  );
}

