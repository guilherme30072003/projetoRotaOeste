'use client'

import Image from 'next/image'
import { usePathname } from 'next/navigation'
import { useState } from 'react'
import DESENVOLVE from '@/assets/images/Logo (2).png'
import Link from 'next/link'

export function SideBarResponsive() {
  const [isOpen, setIsOpen] = useState(false)

  const pathname = usePathname()

  return (
    <div className="relative">
      <button
        className="fixed top-15 left-0 z-40 p-3 rounded-lg bg-primary text-white "
        onClick={() => setIsOpen(!isOpen)}
      >
        Filtrar
        {/* <HandPointing className="h-6 w-6 inline" /> */}
      </button>
      <div
        className={`fixed top-32 left-0 h-full w-64 z-30 shadow-lg transform transition-all ease-in-out duration-300 rounded-2xl ${
          isOpen ? 'translate-x-0' : '-translate-x-full'
        }`}
      >
        <div className="w-fit h-full  text-primary-content rounded-3xl">
          <nav className="">
            <div className="h-full bg-neutral p-4 rounded-xl">
              <div className="form-control w-full max-w-xs">
                <label className="label">
                  <span className="label-text">Pesquisar por nome</span>
                </label>
                <input
                  type="text"
                  placeholder="Type here"
                  className="input input-bordered w-full max-w-xs"
                />
              </div>
              <div className="form-control w-full max-w-xs">
                <label className="label">
                  <span className="label-text">Pesquisar por Cliente</span>
                </label>
                <input
                  type="text"
                  placeholder="Type here"
                  className="input input-bordered w-full max-w-xs"
                />
              </div>
              <h2 className="my-4">Gravidade</h2>
              <div className="form-control">
                <label className="label cursor-pointer">
                  <span className="label-text">Sem Filtro</span>
                  <input
                    type="radio"
                    name="radio-10"
                    className="radio checked:bg-white"
                  />
                </label>
              </div>
              <div className="form-control">
                <label className="label cursor-pointer">
                  <span className="label-text">Baixo Risco</span>
                  <input
                    type="radio"
                    name="radio-10"
                    className="radio checked:bg-blue-500"
                  />
                </label>
              </div>

              <div className="form-control">
                <label className="label cursor-pointer">
                  <span className="label-text">Moderado</span>
                  <input
                    type="radio"
                    name="radio-10"
                    className="radio checked:bg-warning"
                  />
                </label>
              </div>
              <div className="form-control">
                <label className="label cursor-pointer">
                  <span className="label-text">Alto risco</span>
                  <input
                    type="radio"
                    name="radio-10"
                    className="radio checked:bg-red-500"
                  />
                </label>
              </div>
            </div>
            )
          </nav>
        </div>
      </div>
    </div>
  )
}
