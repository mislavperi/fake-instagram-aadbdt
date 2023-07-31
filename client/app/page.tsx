'use client';
import Image from 'next/image'
import { Button } from '@/components/ui/button'

export default function Home() {
  const dummyImage: { title: string, description: string, url: string }[] = [
    {
      title: "image",
      description: "image",
      url: "https://images.unsplash.com/photo-1638486071992-536e48c8fa3e?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=2072&q=80"
    }
  ] 

  return (
    <main>
      {dummyImage.map(image => {
        return (
          <div>
            <Image src={image.url} alt="description" width={500} height={500} />
          </div>
        )
      })}
    </main>
  )
}
