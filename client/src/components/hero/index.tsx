import { useQuery } from '@tanstack/react-query'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'

export interface IEpisode {
  title: string
  image: string
  id: number
  link: string
  src?: string
}

export interface IAnime {
  title: string
  image: string
  id: number
  link: string
  episodes: IEpisode[]
}

export const Hero = () => {
  const { data: animes = [] } = useQuery({
    queryKey: ['animes'],
    queryFn: async ({ signal }) => {
      const { data } = await axios.get<IAnime[]>('/api/animes', {
        params: {
          q: 'one piece',
        },
        signal,
      })

      return data
    },
  })

  const navigate = useNavigate()
  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 to-gray-700 py-8 text-white">
      <div className="container mx-auto px-4">
        <div className="mb-8 flex items-center justify-end">
          <input type="text" placeholder="Search..." value={''} className="rounded-lg px-4 py-2 text-black" />
        </div>
        <div className="grid grid-cols-1 gap-8 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
          {animes.map((anime, index) => (
            <div
              key={index}
              className="transform cursor-pointer overflow-hidden rounded-lg bg-white shadow-lg transition-transform duration-300 hover:scale-105 hover:shadow-2xl"
              onClick={() => navigate(`/anime/${anime.id}`)}
            >
              <img src={anime.image?.replace('100', '900')} alt={anime.title} className="h-48 w-full object-cover" />
              <div className="p-4 text-black">
                <h2 className="text-2xl font-semibold">{anime.title}</h2>
              </div>
            </div>
          ))}
          {animes.length === 0 && <p className="col-span-full text-center">No animes found</p>}
        </div>
      </div>
    </div>
  )
}
