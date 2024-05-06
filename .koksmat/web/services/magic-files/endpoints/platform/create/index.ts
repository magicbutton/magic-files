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
 * Create a new item
 * 
 * item - The item to create

 * @returns
 * 
 * Platform
 */
export default async function call(transactionId: string ,item: Platform) {
  console.log( "magic-files.platform", "create");

  return run<Platform>(
    "magic-files.platform",
    ["create" , JSON.stringify(item)],
    transactionId,
    600,
    transactionId
  );
}

