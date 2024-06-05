/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
"use client";
import { useService } from "@/app/koksmat/useservice";
import { useState } from "react";
import { FileItem } from "../applogic/model";

/* yankiebar */

export default function ReadFile(props: { id: number }) {
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
  return (
    <div>
      {file && (
        <div>
          <div>
            <div className="font-bold">tenant</div>
            <div>{file.tenant}</div>
          </div>{" "}
          <div>
            <div className="font-bold">name</div>
            <div>{file.name}</div>
          </div>{" "}
          <div>
            <div className="font-bold">description</div>
            <div>{file.description}</div>
          </div>{" "}
          <div>
            <div className="font-bold">url</div>
            <div>{file.url}</div>
          </div>{" "}
          <div>
            <div className="font-bold">type</div>
            <div>{file.type}</div>
          </div>{" "}
          <div>
            <div className="font-bold">size</div>
            <div>{file.size}</div>
          </div>{" "}
          <div>
            <div className="font-bold">extension</div>
            <div>{file.extension}</div>
          </div>{" "}
          <div>
            <div className="font-bold">createdby</div>
            <div>{file.createdby}</div>
          </div>{" "}
          <div>
            <div className="font-bold">created</div>
            <div>{file.created}</div>
          </div>{" "}
          <div>
            <div className="font-bold">modified</div>
            <div>{file.modified}</div>
          </div>{" "}
          <div>
            <div className="font-bold">modifiedby</div>
            <div>{file.modifiedby}</div>
          </div>{" "}
          <div>
            <div className="font-bold">version</div>
            <div>{file.version}</div>
          </div>{" "}
          <div>
            <div className="font-bold">published</div>
            <div>{file.published}</div>
          </div>{" "}
          <div>
            <div className="font-bold">level1</div>
            <div>{file.level1}</div>
          </div>{" "}
          <div>
            <div className="font-bold">level2</div>
            <div>{file.level2}</div>
          </div>{" "}
          <div>
            <div className="font-bold">level3</div>
            <div>{file.level3}</div>
          </div>{" "}
          <div>
            <div className="font-bold">level4</div>
            <div>{file.level4}</div>
          </div>{" "}
          <div>
            <div className="font-bold">level5</div>
            <div>{file.level5}</div>
          </div>{" "}
          <div>
            <div className="font-bold">level6</div>
            <div>{file.level6}</div>
          </div>{" "}
          <div>
            <div className="font-bold">level7</div>
            <div>{file.level7}</div>
          </div>{" "}
          <div>
            <div className="font-bold">level8</div>
            <div>{file.level8}</div>
          </div>{" "}
          <div>
            <div className="font-bold">level9</div>
            <div>{file.level9}</div>
          </div>{" "}
          <div>
            <div className="font-bold">level10</div>
            <div>{file.level10}</div>
          </div>{" "}
          <div>
            <div className="font-bold">level11</div>
            <div>{file.level11}</div>
          </div>{" "}
          <div>
            <div className="font-bold">level12</div>
            <div>{file.level12}</div>
          </div>{" "}
          <div>
            <div className="font-bold">level13</div>
            <div>{file.level13}</div>
          </div>{" "}
          <div>
            <div className="font-bold">level14</div>
            <div>{file.level14}</div>
          </div>{" "}
          <div>
            <div className="font-bold">level15</div>
            <div>{file.level15}</div>
          </div>{" "}
          <div>
            <div className="font-bold">level16</div>
            <div>{file.level16}</div>
          </div>
          <div>
            <div>
              <div className="font-bold">id</div>
              <div>{file.id}</div>
            </div>
            <div>
              <div className="font-bold">created_at</div>
              <div>{file.created_at}</div>
            </div>
            <div>
              <div className="font-bold">created_by</div>
              <div>{file.created_by}</div>
            </div>
            <div>
              <div className="font-bold">updated_at</div>
              <div>{file.updated_at}</div>
            </div>
            <div>
              <div className="font-bold">updated_by</div>
              <div>{file.updated_by}</div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
