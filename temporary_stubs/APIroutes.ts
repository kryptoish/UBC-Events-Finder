//This will be put in the appropriate directory and config later

// pages/api/events.ts
import { NextApiRequest, NextApiResponse } from 'next';

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
    // TODO: Forward request to Go backend to retrieve events
    res.status(200).json([]);
}

