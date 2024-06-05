export default function IsDevMode() {
  return process.env.NODE_ENV === "production";
}
