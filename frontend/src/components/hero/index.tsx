import { useQuery } from '@tanstack/react-query'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'

export const Hero = () => {
  const { data: animes = [] } = useQuery({
    queryKey: ['animes'],
    queryFn: async ({ signal }) => {
      const { data } = await axios.get('/api/anime', {
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
    <div className="flex min-h-screen bg-gradient-to-b from-stone-700 to-stone-500">
      <section className="flex w-full py-32 md:py-48">
        <div className={`flex w-full flex-1 flex-row justify-center space-x-4 p-10`}>
          {animes.map((anime: any) => (
            <div
              key={anime.id}
              style={{
                backgroundImage: `url("${anime.Image.replace('100', '4000')}")`,
              }}
              className={`flex-column relative flex max-h-72 w-56 cursor-pointer  flex-wrap justify-center space-x-2 bg-cover  bg-no-repeat align-bottom`}
              onClick={() => navigate(`/anime/${anime.Id}`)}
            >
              <div className="absolute h-full w-full bg-gradient-to-b from-transparent to-stone-500" />
              <h4 className="z-10 mx-auto mt-auto text-lg font-semibold text-white">{anime.Title}</h4>
            </div>
          ))}
        </div>
      </section>
    </div>
  )
}
