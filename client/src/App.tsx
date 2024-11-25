import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'


import { NextResponse } from 'next/server';
import { get } from '@vercel/edge-config';

export const config: { matcher: string } = { matcher: '/welcome' };

export async function middleware(): Promise<NextResponse> {
  const greeting = await get('greeting'); // greeting is type string | undefined

  if (!greeting) {
    return NextResponse.json({ error: 'Greeting not found' }, { status: 404 });
  }

  return NextResponse.json({ greeting }); // Safe to use greeting here as it's definitely a string
}

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App
