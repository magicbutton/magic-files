    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
"use client";
import { useService } from "@/koksmat/useservice";
import { useState } from "react";
import {ContainerItem} from "../applogic/model";


/* yankiebar */

export default function ReadContainer(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<ContainerItem>(
      "magic-files.container",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const container = readResult.data;
    return (
      <div>
          
    {container && <div>
        <div>
        <div className="font-bold" >tenant</div>
        <div>{container.tenant}</div>
    </div>    <div>
        <div className="font-bold" >name</div>
        <div>{container.name}</div>
    </div>    <div>
        <div className="font-bold" >description</div>
        <div>{container.description}</div>
    </div>    <div>
        <div className="font-bold" >url</div>
        <div>{container.url}</div>
    </div>    <div>
        <div className="font-bold" >type</div>
        <div>{container.type}</div>
    </div>    <div>
        <div className="font-bold" >level1</div>
        <div>{container.level1}</div>
    </div>    <div>
        <div className="font-bold" >level2</div>
        <div>{container.level2}</div>
    </div>    <div>
        <div className="font-bold" >level3</div>
        <div>{container.level3}</div>
    </div>    <div>
        <div className="font-bold" >level4</div>
        <div>{container.level4}</div>
    </div>    <div>
        <div className="font-bold" >level5</div>
        <div>{container.level5}</div>
    </div>    <div>
        <div className="font-bold" >level6</div>
        <div>{container.level6}</div>
    </div>    <div>
        <div className="font-bold" >level7</div>
        <div>{container.level7}</div>
    </div>    <div>
        <div className="font-bold" >level8</div>
        <div>{container.level8}</div>
    </div>    <div>
        <div className="font-bold" >level9</div>
        <div>{container.level9}</div>
    </div>    <div>
        <div className="font-bold" >level10</div>
        <div>{container.level10}</div>
    </div>    <div>
        <div className="font-bold" >level11</div>
        <div>{container.level11}</div>
    </div>    <div>
        <div className="font-bold" >level12</div>
        <div>{container.level12}</div>
    </div>    <div>
        <div className="font-bold" >level13</div>
        <div>{container.level13}</div>
    </div>    <div>
        <div className="font-bold" >level14</div>
        <div>{container.level14}</div>
    </div>    <div>
        <div className="font-bold" >level15</div>
        <div>{container.level15}</div>
    </div>    <div>
        <div className="font-bold" >level16</div>
        <div>{container.level16}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{container.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{container.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{container.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{container.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{container.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
