"use client";

import { JobDashboard } from "@/components/job-dashboard";
import { useRouter } from "next/navigation";

export default function Page() {
  const router = useRouter();
  return (
    <div className="space-x-2 h-[90vh]">
      <JobDashboard />
    </div>
  );
}
