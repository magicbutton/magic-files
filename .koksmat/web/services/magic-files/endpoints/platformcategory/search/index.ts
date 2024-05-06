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
 * Search for items
 * 
 * query - The search query

 * @returns
 * 
 * PlatformCategory
 */
export default async function call(transactionId: string ,query: string) {
  console.log( "magic-files.platformcategory", "search");

  return run<PlatformCategory>(
    "magic-files.platformcategory",
    ["search" , query],
    transactionId,
    600,
    transactionId
  );
}
