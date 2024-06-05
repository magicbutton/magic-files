"use client";

import IsDevMode from "../isdevmode";

export default function EditorLeftrail() {
  if (IsDevMode()) return null;
  return <div>DEV</div>;
}
