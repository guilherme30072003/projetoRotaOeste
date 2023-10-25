'use client'

import Aviso from '@/components/Aviso'
import Navbar from '@/components/Navbar'
import { SideBarResponsive } from '@/components/SideBarResponsive'
import SideMenu from '@/components/SideMenu'
import { url } from 'inspector'
import Image from 'next/image'

//import { PrismaClient, receivedalerts } from '@prisma/client'
import { Key, useEffect, useState } from 'react'

import { prisma } from '@/lib/prisma'


const getBaseData = async () => {
    const res = await fetch("http://localhost:8080/dados") 
    return res.json()
}



export default async function Home() {
  // const [filtro, setFiltro] = useState<receivedalerts[]>([])

  // const [cor, setCor] = useState('BLUE')

  const alertas = await getBaseData()


  // useEffect(() => {
  //   async function getData() {
  //     const data = await (await fetch(`/api?cor=${cor}`)).json()
  //     setFiltro(data)
  //     console.log(data)
  //   }
  //   getData()
  // }, [cor])

  console.log(alertas)

  return (
    <main className="">
      <Navbar></Navbar>
      <div className="flex flex-row">
        <div className="min-h-screen w-1/5 h-screen mr-4 hidden lg:block">
        <SideMenu></SideMenu>
        </div>
        <div className="block lg:hidden rounded-xl">
          <SideBarResponsive></SideBarResponsive>
        </div>
        <div className=" lg:p-0 p-12">
          {alertas.map((item: any) => (
            <Aviso
              title={item.occurrences}
              description={item.definitionRel}
              gravidade={item.acknowledgementStatus}
              id={item.id}
            ></Aviso>
          ))}
        </div>
      </div>
    </main>
  )
}