import { FC } from 'react'

export interface IInput {
  placeholder?: string
}
export const Input: FC<IInput> = ({ placeholder }) => {
  return <input className="rounded p-2 text-base focus:outline-none" placeholder={placeholder} />
}
