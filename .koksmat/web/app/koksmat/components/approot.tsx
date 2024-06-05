"use client";

import { ServiceInspector } from "./service-inspector";

export default function AppRoot(props: { children: any }) {
  const { children } = props;
  //const magicbox = useContext(MagicboxContext);

  return (
    <div className="flex">
      <div className="grow">{children}</div>
      <div className="hidden lg:block">x</div>
    </div>
  );
}
