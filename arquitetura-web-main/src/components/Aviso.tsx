interface AvisoProps {
  title: string
  description: string
  date?: string
  avatar?: string
  gravidade?: string
  id?:string
}

export default function Aviso({
  title,
  description,
  date,
  avatar,
  gravidade,
  id
}: AvisoProps) {
  return (
    <div className="bg-base-100 rounded p-6 max-w-7xl mt-10">
      <div className="flex flex-row gap-4 items-center">
        <div className="avatar">
          <div className="w-10 rounded-full">
            <img src="https://images.unsplash.com/photo-1614977645540-7abd88ba8e56?auto=format&fit=crop&q=80&w=1546&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D" />
          </div>
        </div>
        <div className="flex flex-col">
          <p className="font-bold">{title}</p>
          <h4 className="text-xs">{id}</h4>
        </div>
      </div>
      <div className="mt-8 mb-8">
        <p>{description}</p>
      </div>
      <button className="btn btn-outline">ACTION</button>
      <p className="mt-6 flex justify-end">23/10/2023</p>
    </div>
  )
}
