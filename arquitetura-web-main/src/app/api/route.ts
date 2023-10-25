import { PrismaClient } from "@prisma/client";
import { NextApiRequest, NextApiResponse } from "next";
import { NextRequest, NextResponse } from "next/server";

const prisma = new PrismaClient();

async function handler(request: NextRequest) {
  const searchParams = request.nextUrl.searchParams
  const cor = searchParams.get('cor')


  const alerts = await prisma.receivedalerts.findMany({
    where: {
      color: cor as string, 
    }
  })

  const serializedAlerts = alerts.map((alert) => ({
    ...alert,
    machinelineartime: alert.machinelineartime?.toString(),
  }));

  return NextResponse.json(serializedAlerts);
}


export {handler as GET, handler as POST}


