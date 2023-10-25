'use client'

import { useEffect } from 'react'

interface SideMenuProps {
  setCor: Function
}

export default function SideMenu({ setCor }: SideMenuProps) {
  useEffect(() => {
    console.log('oi')
  }, [])
  return (
    <div className="h-full bg-neutral p-4">
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
            className="radio bg-white"
            onChange={() => setCor('')}
          />
        </label>
      </div>
      <div className="form-control">
        <label className="label cursor-pointer">
          <span className="label-text">Baixo Risco</span>
          <input
            type="radio"
            name="radio-10"
            className="radio bg-blue-500"
            onChange={() => setCor('BLUE')}
          />
        </label>
      </div>

      <div className="form-control">
        <label className="label cursor-pointer">
          <span className="label-text">Moderado</span>
          <input
            type="radio"
            name="radio-10"
            className="radio bg-warning"
            onChange={() => setCor('yellow')}
          />
        </label>
      </div>
      <div className="form-control">
        <label className="label cursor-pointer">
          <span className="label-text">Alto risco</span>
          <input
            type="radio"
            name="radio-10"
            className="radio bg-red-500"
            onChange={() => setCor('RED')}
          />
        </label>
      </div>
    </div>
  )
}
