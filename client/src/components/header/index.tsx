import { Input } from '../ui/input'

export function Header() {
  return (
    <div className="fixed left-0 top-0 flex w-full items-center justify-between bg-stone-700 bg-opacity-70 px-4 py-4 md:px-12">
      <Input placeholder="Search..." />
    </div>
  )
}
