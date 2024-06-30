import { useQuery } from '@tanstack/react-query'
import axios from 'axios'
import { useState } from 'react'
import { useParams } from 'react-router-dom'
import { IEpisode } from 'src/components/hero'

export const Anime = () => {
  const params = useParams()
  const [current, setEpisode] = useState<IEpisode | null>(null)

  const { data: episodes = [] } = useQuery({
    enabled: !!params.id,
    queryKey: ['animes', params.id],
    queryFn: async ({ signal }) => {
      const { data } = await axios.get<IEpisode[]>(`/api/animes/${params.id}`, {
        signal,
      })

      return data
    },
    onSuccess: (result) => setEpisode(result[0]),
  })

  const {
    data: currentSrc = '',
    isLoading,
    isFetching,
  } = useQuery({
    queryKey: ['src', current?.link],
    queryFn: async ({ signal }) => {
      const { data } = await axios.get<string>(`/api/episode/${current?.id}/src`, {
        signal,
      })

      return data
    },
  })

  return (
    <div className="min-h-screen flex-row flex-wrap bg-gradient-to-b from-stone-700 to-stone-500">
      {isLoading || isFetching ? (
        <div className="mx-auto h-[800px] min-w-full pt-20" />
      ) : (
        <iframe
          onClick={(e) => {
            e.stopPropagation()
          }}
          key={currentSrc}
          className="mx-auto min-w-full pt-20"
          height="800px"
          allowFullScreen
          src={currentSrc}
        />
      )}
      <div
        className={`space-2 flex max-h-96 w-full flex-1 flex-row flex-wrap justify-between gap-2 overflow-y-auto p-2`}
      >
        {episodes.map((episode) => (
          <span
            onClick={(e) => {
              e.preventDefault()
              e.stopPropagation()
              setEpisode(episode)
            }}
            key={episode.link}
            className={`text-md flex h-8 w-12 cursor-pointer flex-wrap justify-center border-2 border-gray-50 p-1 text-center font-semibold text-white`}
          >
            {episode.title}
          </span>
        ))}
      </div>
    </div>
  )
}
