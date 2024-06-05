/**
* This code was generated by v0 by Vercel.
* @see https://v0.dev/t/SpLSUZHnHmi
* Documentation: https://v0.dev/docs#integrating-generated-code-into-your-nextjs-app
*/

/** Add fonts into your Next.js project:

import { Inter } from 'next/font/google'

inter({
  subsets: ['latin'],
  display: 'swap',
})

To read more about using these font, please visit the Next.js documentation:
- App Directory: https://nextjs.org/docs/app/building-your-application/optimizing/fonts
- Pages Directory: https://nextjs.org/docs/pages/building-your-application/optimizing/fonts
**/
import { Button } from "@/components/ui/button"
import { TableHead, TableRow, TableHeader, TableBody, Table } from "@/components/ui/table"
import { Label } from "@/components/ui/label"
import { Input } from "@/components/ui/input"
import { Textarea } from "@/components/ui/textarea"

export function JobDashboard() {
  return (
    <div className="flex flex-col h-full">
      <header className="bg-gray-900 text-white p-4 flex items-center justify-between">
        <h1 className="text-2xl font-bold">Job Dashboard</h1>
        <Button>Schedule Job</Button>
      </header>
      <main className="flex-1 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 p-6">
        <section className="bg-white rounded-lg shadow-md p-4">
          <h2 className="text-lg font-bold mb-4">Historical Jobs</h2>
          <div className="overflow-auto max-h-[300px]">
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>Job Name</TableHead>
                  <TableHead>Status</TableHead>
                  <TableHead>Start Time</TableHead>
                  <TableHead>End Time</TableHead>
                  <TableHead>Duration</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody />
            </Table>
          </div>
        </section>
        <section className="bg-white rounded-lg shadow-md p-4">
          <h2 className="text-lg font-bold mb-4">Running Jobs</h2>
          <div className="space-y-4" />
        </section>
        <section className="bg-white rounded-lg shadow-md p-4">
          <h2 className="text-lg font-bold mb-4">Schedule Job</h2>
          <form>
            <div className="space-y-4">
              <div>
                <Label htmlFor="name">Job Name</Label>
                <Input id="name" required />
              </div>
              <div>
                <Label htmlFor="description">Description</Label>
                <Textarea id="description" required />
              </div>
              <div>
                <Label htmlFor="startTime">Start Time</Label>
                <Input id="startTime" required type="datetime-local" />
              </div>
              <Button type="submit">Schedule Job</Button>
            </div>
          </form>
        </section>
      </main>
    </div>
  )
}